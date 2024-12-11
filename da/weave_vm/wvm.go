package weave_vm

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/celestiaorg/celestia-openrpc/types/blob"
	"github.com/celestiaorg/celestia-openrpc/types/header"
	"github.com/dymensionxyz/dymint/da"
	"github.com/dymensionxyz/dymint/da/weave_vm/rpc"
	"github.com/dymensionxyz/dymint/da/weave_vm/signer"
	weaveVMtypes "github.com/dymensionxyz/dymint/da/weave_vm/types"
	"github.com/dymensionxyz/dymint/store"
	"github.com/dymensionxyz/dymint/types"
	pb "github.com/dymensionxyz/dymint/types/pb/dymint"
	uretry "github.com/dymensionxyz/dymint/utils/retry"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/pubsub"
)

type WeaveVM interface {
	SendTransaction(ctx context.Context, to string, data []byte) (string, error)
	GetTransactionReceipt(ctx context.Context, txHash string) (*ethtypes.Receipt, error)
}

// TODO: adjust
const (
	defaultRpcRetryDelay    = 3 * time.Second
	namespaceVersion        = 0
	DefaultGasPrices        = 0.1
	defaultRpcRetryAttempts = 5
)

var defaultSubmitBackoff = uretry.NewBackoffConfig(
	uretry.WithInitialDelay(time.Second*6),
	uretry.WithMaxDelay(time.Second*6),
)

// DataAvailabilityLayerClient use celestia-node public API.
type DataAvailabilityLayerClient struct {
	client       WeaveVM
	pubsubServer *pubsub.Server
	config       *weaveVMtypes.Config
	logger       types.Logger
	ctx          context.Context
	cancel       context.CancelFunc
	synced       chan struct{}
}

var (
	_ da.DataAvailabilityLayerClient = &DataAvailabilityLayerClient{}
	_ da.BatchRetriever              = &DataAvailabilityLayerClient{}
)

// WithRPCClient sets rpc client.
func WithRPCClient(rpc WeaveVM) da.Option {
	return func(daLayerClient da.DataAvailabilityLayerClient) {
		daLayerClient.(*DataAvailabilityLayerClient).client = rpc
	}
}

// WithRPCRetryDelay sets failed rpc calls retry delay.
func WithRPCRetryDelay(delay time.Duration) da.Option {
	return func(daLayerClient da.DataAvailabilityLayerClient) {
		daLayerClient.(*DataAvailabilityLayerClient).config.RetryDelay = delay
	}
}

// WithRPCAttempts sets failed rpc calls retry attempts.
func WithRPCAttempts(attempts int) da.Option {
	return func(daLayerClient da.DataAvailabilityLayerClient) {
		daLayerClient.(*DataAvailabilityLayerClient).config.RetryAttempts = &attempts
	}
}

// WithSubmitBackoff sets submit retry delay config.
func WithSubmitBackoff(c uretry.BackoffConfig) da.Option {
	return func(daLayerClient da.DataAvailabilityLayerClient) {
		daLayerClient.(*DataAvailabilityLayerClient).config.Backoff = c
	}
}

// Init initializes the WeaveVM DA client
func (c *DataAvailabilityLayerClient) Init(config []byte, pubsubServer *pubsub.Server, kvStore store.KV, logger types.Logger, options ...da.Option) error {
	var cfg weaveVMtypes.Config
	if len(config) > 0 {
		err := json.Unmarshal(config, &c.config)
		if err != nil {
			return fmt.Errorf("failed to unmarshal config: %w", err)
		}
	}

	// Set defaults
	c.pubsubServer = pubsubServer
	c.logger = logger
	c.config = &cfg

	// Validate and set defaults
	if c.config.RetryDelay == 0 {
		c.config.RetryDelay = defaultRpcRetryDelay
	}
	if c.config.Backoff == (uretry.BackoffConfig{}) {
		c.config.Backoff = defaultSubmitBackoff
	}
	if c.config.RetryAttempts == nil {
		attempts := defaultRpcRetryAttempts
		c.config.RetryAttempts = &attempts
	}

	// Apply options
	for _, apply := range options {
		apply(c)
	}
	types.RollappConsecutiveFailedDASubmission.Set(0)

	c.ctx, c.cancel = context.WithCancel(context.Background())

	if c.client != nil {
		if cfg.Web3SignerEndpoint != "" {
			web3signer, err := signer.NewWeb3SignerClient(&cfg, logger)
			if err != nil {
				return fmt.Errorf("failed to initialize web3signer client: %w", err)
			}
			client, err := rpc.NewWvmRPCClient(logger, &cfg, web3signer)
			if err != nil {
				return fmt.Errorf("failed to initialize rpc client for weaveVM chain: %w", err)
			}
			c.client = client
			return nil
		}

		// Use PrivateKey signer
		if cfg.PrivateKeyHex == "" {
			return fmt.Errorf("weaveVM private key is empty and weaveVM web3 signer is empty")
		}
		privateKeySigner := signer.NewPrivateKeySigner(cfg.PrivateKeyHex, logger, cfg.ChainID)
		client, err := rpc.NewWvmRPCClient(logger, &cfg, privateKeySigner)
		if err != nil {
			return fmt.Errorf("failed to initialize rpc client for weaveVM chain: %w", err)
		}

		c.client = client
	}

	return nil
}

// Start starts DataAvailabilityLayerClient instance.
func (c *DataAvailabilityLayerClient) Start() error {
	c.synced <- struct{}{}
	return nil
}

// Stop stops DataAvailabilityLayerClient instance.
func (c *DataAvailabilityLayerClient) Stop() error {
	c.cancel()
	close(c.synced)
	return nil
}

// WaitForSyncing is used to check when the DA light client finished syncing
func (m *DataAvailabilityLayerClient) WaitForSyncing() {
	<-m.synced
}

// GetClientType returns client type.
func (c *DataAvailabilityLayerClient) GetClientType() da.Client {
	return da.WeaveVM
}

// SubmitBatch submits a batch to the DA layer.
func (c *DataAvailabilityLayerClient) SubmitBatch(batch *types.Batch) da.ResultSubmitBatch {
	data, err := batch.MarshalBinary()
	if err != nil {
		return da.ResultSubmitBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: err.Error(),
				Error:   err,
			},
		}
	}

	if len(data) > weaveVMtypes.WeaveVMMaxTransactionSize {
		return da.ResultSubmitBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: fmt.Sprintf("size bigger than maximum blob size: max n bytes: %d", weaveVMtypes.WeaveVMMaxTransactionSize),
				Error:   errors.New("blob size too big"),
			},
		}
	}

	backoff := c.config.Backoff.Backoff()

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Debug("Context cancelled.")
			return da.ResultSubmitBatch{}
		default:
			height, commitment, err := c.submit(data)
			if errors.Is(err, gerrc.ErrInternal) {
				err = fmt.Errorf("submit: %w", err)
				return da.ResultSubmitBatch{
					BaseResult: da.BaseResult{
						Code:    da.StatusError,
						Message: err.Error(),
						Error:   err,
					},
				}
			}

			if err != nil {
				c.logger.Error("Submit blob.", "error", err)
				types.RollappConsecutiveFailedDASubmission.Inc()
				backoff.Sleep()
				continue
			}

			daMetaData := &da.DASubmitMetaData{
				Client:     da.WeaveVM,
				Height:     height,
				Commitment: commitment,
				Namespace:  nil,
			}

			c.logger.Debug("Submitted blob to DA successfully.")

			types.RollappConsecutiveFailedDASubmission.Set(0)
			return da.ResultSubmitBatch{
				BaseResult: da.BaseResult{
					Code:    da.StatusSuccess,
					Message: "Submission successful",
				},
				SubmitMetaData: daMetaData,
			}
		}
	}
}

func (c *DataAvailabilityLayerClient) RetrieveBatches(daMetaData *da.DASubmitMetaData) da.ResultRetrieveBatch {
	for {
		select {
		case <-c.ctx.Done():
			c.logger.Debug("Context cancelled.")
			return da.ResultRetrieveBatch{}
		default:
			var resultRetrieveBatch da.ResultRetrieveBatch
			err := retry.Do(
				func() error {
					resultRetrieveBatch = c.retrieveBatches(daMetaData)

					if errors.Is(resultRetrieveBatch.Error, da.ErrRetrieval) {
						c.logger.Error("Retrieve batch.", "error", resultRetrieveBatch.Error)
						return resultRetrieveBatch.Error
					}

					return nil
				},
				retry.Attempts(uint(*c.config.RetryAttempts)), //nolint:gosec // RetryAttempts should be always positive
				retry.DelayType(retry.FixedDelay),
				retry.Delay(c.config.RetryDelay),
			)
			if err != nil {
				c.logger.Error("RetrieveBatches process failed.", "error", err)
			}
			return resultRetrieveBatch

		}
	}
}

func (c *DataAvailabilityLayerClient) retrieveBatches(daMetaData *da.DASubmitMetaData) da.ResultRetrieveBatch {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.Timeout)
	defer cancel()

	c.logger.Debug("Getting blob from DA.", "height", daMetaData.Height, "namespace", hex.EncodeToString(daMetaData.Namespace), "commitment", hex.EncodeToString(daMetaData.Commitment))
	var batches []*types.Batch

	// get wvm tx hash

	txHash := common.BytesToHash(daMetaData.Commitment)

	blob, err := c.getFromGateway(ctx, txHash.Hex())
	if err != nil {
		return da.ResultRetrieveBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: err.Error(),
				Error:   da.ErrRetrieval,
			},
		}
	}
	if blob == nil {
		return da.ResultRetrieveBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: "Blob not found",
				Error:   da.ErrBlobNotFound,
			},
		}
	}

	var batch pb.Batch
	err = proto.Unmarshal(blob, &batch)
	if err != nil {
		c.logger.Error("Unmarshal blob.", "daHeight", daMetaData.Height, "error", err)
		return da.ResultRetrieveBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: err.Error(),
				Error:   da.ErrBlobNotParsed,
			},
		}
	}

	c.logger.Debug("Blob retrieved successfully from DA.", "DA height", daMetaData.Height, "lastBlockHeight", batch.EndHeight)

	parsedBatch := new(types.Batch)
	err = parsedBatch.FromProto(&batch)
	if err != nil {
		return da.ResultRetrieveBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: err.Error(),
				Error:   da.ErrBlobNotParsed,
			},
		}
	}
	batches = append(batches, parsedBatch)
	return da.ResultRetrieveBatch{
		BaseResult: da.BaseResult{
			Code:    da.StatusSuccess,
			Message: "Batch retrieval successful",
		},
		Batches: batches,
	}
}

const weaveVMGatewayURL = "https://gateway.wvm.dev/v1/calldata/%s"

type WvmDymintBlob struct {
	ArweaveBlockHash string
	WvmBlockHash     string
	WvmTxHash        string
	Blob             []byte
}

// Modified get function with improved error handling
func (c *DataAvailabilityLayerClient) getFromGateway(ctx context.Context, weaveVMTxHash string) (*WvmDymintBlob, error) {
	type WvmRetrieverResponse struct {
		ArweaveBlockHash   string `json:"arweave_block_hash"`
		Calldata           string `json:"calldata"`
		WarDecodedCalldata string `json:"war_decoded_calldata"`
		WvmBlockHash       string `json:"weavevm_block_hash"`
	}
	r, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf(weaveVMGatewayURL,
			weaveVMTxHash), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	r.Header.Set("Accept", "application/json")
	client := &http.Client{
		Timeout: c.config.Timeout,
	}

	c.logger.Debug("sending request to WeaveVM data retriever",
		"url", r.URL.String(),
		"headers", r.Header)

	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to call weaveVM-data-retriever: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := validateResponse(resp, body); err != nil {
		c.logger.Error("invalid response from WeaveVM data retriever",
			"status", resp.Status,
			"content_type", resp.Header.Get("Content-Type"),
			"body", string(body))
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	var weaveVMData WvmRetrieverResponse
	if err = json.Unmarshal(body, &weaveVMData); err != nil {
		c.logger.Error("failed to unmarshal response",
			"error", err,
			"body", string(body),
			"content_type", resp.Header.Get("Content-Type"))
		return nil, fmt.Errorf("failed to unmarshal response: %w, body: %s", err, string(body))
	}

	c.logger.Info("weaveVM backend: get data from weaveVM",
		"arweave_block_hash", weaveVMData.ArweaveBlockHash,
		"weavevm_block_hash", weaveVMData.WvmBlockHash,
		"calldata_length", len(weaveVMData.Calldata))

	blob, err := hexutil.Decode(weaveVMData.Calldata)
	if err != nil {
		return nil, fmt.Errorf("failed to decode calldata: %w", err)
	}

	if len(blob) == 0 {
		return nil, fmt.Errorf("decoded blob has length zero")
	}

	return &WvmDymintBlob{ArweaveBlockHash: weaveVMData.ArweaveBlockHash, WvmBlockHash: weaveVMData.WvmBlockHash, WvmTxHash: weaveVMTxHash, Blob: blob}, nil
}

func validateResponse(resp *http.Response, body []byte) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return fmt.Errorf("unexpected content type: %s, body: %s", contentType, string(body))
	}

	return nil
}

func (c *DataAvailabilityLayerClient) CheckBatchAvailability(daMetaData *da.DASubmitMetaData) da.ResultCheckBatch {
	var availabilityResult da.ResultCheckBatch
	for {
		select {
		case <-c.ctx.Done():
			c.logger.Debug("Context cancelled")
			return da.ResultCheckBatch{}
		default:
			err := retry.Do(
				func() error {
					result := c.checkBatchAvailability(daMetaData)
					availabilityResult = result

					if result.Code != da.StatusSuccess {
						c.logger.Error("Blob submitted not found in DA. Retrying availability check.")
						return da.ErrBlobNotFound
					}

					return nil
				},
				retry.Attempts(uint(*c.config.RetryAttempts)), //nolint:gosec // RetryAttempts should be always positive
				retry.DelayType(retry.FixedDelay),
				retry.Delay(c.config.RetryDelay),
			)
			if err != nil {
				c.logger.Error("CheckAvailability process failed.", "error", err)
			}
			return availabilityResult
		}
	}
}

func (c *DataAvailabilityLayerClient) checkBatchAvailability(daMetaData *da.DASubmitMetaData) da.ResultCheckBatch {
	DACheckMetaData := &da.DACheckMetaData{
		Client:     daMetaData.Client,
		Height:     daMetaData.Height,
		Commitment: daMetaData.Commitment,
	}

	txHash := common.BytesToHash(daMetaData.Commitment)
	res, err := c.getFromGateway(context.Background(), txHash.Hex())
	if err != nil {
		return da.ResultCheckBatch{
			BaseResult: da.BaseResult{
				Code:    da.StatusError,
				Message: err.Error(),
				Error:   da.ErrBlobNotFound,
			},
			CheckMetaData: DACheckMetaData,
		}
	}

	return da.ResultCheckBatch{
		BaseResult: da.BaseResult{
			Code:    da.StatusSuccess,
			Message: "batch available",
		},
		CheckMetaData: DACheckMetaData,
	}
}

// Submit submits the Blobs to Data Availability layer.
func (c *DataAvailabilityLayerClient) submit(daBlob da.Blob) (uint64, da.Commitment, error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.ctx, c.config.Timeout)
	defer cancel()

	// Submit transaction to WeaveVM with the blob data
	txHash, err := c.client.SendTransaction(ctx, weaveVMtypes.ArchivePoolAddress, daBlob)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to send transaction: %w", gerrc.ErrInternal)
	}

	c.logger.Info("wvm tx hash", "hash", txHash)

	// Wait for receipt
	receipt, err := c.waitForTxReceipt(ctx, txHash)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get tx receipt: %w", err)
	}

	return receipt.BlockNumber.Uint64(), receipt.TxHash.Bytes(), nil
}

func (c *DataAvailabilityLayerClient) waitForTxReceipt(ctx context.Context, txHash string) (*ethtypes.Receipt, error) {
	var receipt *ethtypes.Receipt
	err := retry.Do(
		func() error {
			var err error
			receipt, err = c.client.GetTransactionReceipt(ctx, txHash)
			if err != nil {
				return fmt.Errorf("get receipt failed: %w", err)
			}
			if receipt == nil {
				return errors.New("receipt not found")
			}
			if receipt.BlockNumber == big.NewInt(0) {
				return errors.New("no block number in receipt")
			}
			return nil
		},
		retry.Context(ctx),
		retry.Attempts(uint(*c.config.RetryAttempts)),
		retry.Delay(c.config.RetryDelay),
		retry.OnRetry(func(n uint, err error) {
			c.logger.Debug("waiting for receipt",
				"txHash", txHash,
				"attempt", n,
				"error", err)
		}),
	)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func (c *DataAvailabilityLayerClient) getProof(daMetaData *da.DASubmitMetaData) (*blob.Proof, error) {
	return nil, nil
	/*
		c.logger.Debug("Getting proof via RPC call.", "height", daMetaData.Height, "namespace", daMetaData.Namespace, "commitment", daMetaData.Commitment)
		ctx, cancel := context.WithTimeout(c.ctx, c.config.Timeout)
		defer cancel()

		proof, err := c.rpc.GetProof(ctx, daMetaData.Height, daMetaData.Namespace, daMetaData.Commitment)
		if err != nil {
			return nil, err
		}

		return proof, nil
	*/
}

func (c *DataAvailabilityLayerClient) blobsAndCommitments(daBlob da.Blob) ([]*blob.Blob, []da.Commitment, error) {
	return nil, nil, nil
	/*
		var blobs []*blob.Blob
		var commitments []da.Commitment
		b, err := blob.NewBlobV0(c.config.NamespaceID.Bytes(), daBlob)
		if err != nil {
			return nil, nil, err
		}
		blobs = append(blobs, b)

		commitments = append(commitments, b.Commitment)
		return blobs, commitments, nil
	*/
}

func (c *DataAvailabilityLayerClient) validateProof(daMetaData *da.DASubmitMetaData, proof *blob.Proof) (bool, error) {
	return false, nil
	/*
		c.logger.Debug("Validating proof via RPC call.", "height", daMetaData.Height, "namespace", daMetaData.Namespace, "commitment", daMetaData.Commitment)
		ctx, cancel := context.WithTimeout(c.ctx, c.config.Timeout)
		defer cancel()

		return c.rpc.Included(ctx, daMetaData.Height, daMetaData.Namespace, proof, daMetaData.Commitment)
	*/
}

func (c *DataAvailabilityLayerClient) getDataAvailabilityHeaders(height uint64) (*header.DataAvailabilityHeader, error) {
	return nil, nil
	/*
		c.logger.Debug("Getting extended headers via RPC call.", "height", height)
		ctx, cancel := context.WithTimeout(c.ctx, c.config.Timeout)
		defer cancel()

		headers, err := c.rpc.GetByHeight(ctx, height)
		if err != nil {
			return nil, err
		}

		return headers.DAH, nil
	*/
}

// GetMaxBlobSizeBytes returns the maximum allowed blob size in the DA, used to check the max batch size configured
func (d *DataAvailabilityLayerClient) GetMaxBlobSizeBytes() uint32 {
	// return maxBlobSizeBytes
	return 0
}

// GetSignerBalance returns the balance for a specific address
func (d *DataAvailabilityLayerClient) GetSignerBalance() (da.Balance, error) {
	return da.Balance{}, nil
	/*
		ctx, cancel := context.WithTimeout(d.ctx, d.config.Timeout)
		defer cancel()

		balance, err := d.rpc.GetSignerBalance(ctx)
		if err != nil {
			return da.Balance{}, fmt.Errorf("get balance: %w", err)
		}

		daBalance := da.Balance{
			Amount: balance.Amount,
			Denom:  balance.Denom,
		}

		return daBalance, nil
	*/
}