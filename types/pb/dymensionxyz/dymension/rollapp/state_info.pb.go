// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/dymensionxyz/dymension/rollapp/state_info.proto

package rollapp

import (
	fmt "fmt"
	common "github.com/dymensionxyz/dymint/types/pb/dymensionxyz/dymension/common"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// StateInfoIndex is the data used for indexing and retrieving a StateInfo
// it updated and saved with every UpdateState in StateInfo.
// We use the this structure also for:
// 1. LatestStateInfoIndex which defines the rollapps' current (latest) index of the last UpdateState
// 2. LatestFinalizedStateIndex which defines the rollapps' current (latest) index of the latest StateInfo that was finalized
type StateInfoIndex struct {
	// rollappId is the rollapp that the sequencer belongs to and asking to update
	// it used to identify the what rollapp a StateInfo belongs
	// The rollappId follows the same standard as cosmos chain_id
	RollappId string `protobuf:"bytes,1,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// index is a sequential increasing number, updating on each
	// state update used for indexing to a specific state info, the first index is 1
	Index uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *StateInfoIndex) Reset()         { *m = StateInfoIndex{} }
func (m *StateInfoIndex) String() string { return proto.CompactTextString(m) }
func (*StateInfoIndex) ProtoMessage()    {}
func (*StateInfoIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_deebb9ffbcdd017e, []int{0}
}
func (m *StateInfoIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateInfoIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateInfoIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateInfoIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateInfoIndex.Merge(m, src)
}
func (m *StateInfoIndex) XXX_Size() int {
	return m.Size()
}
func (m *StateInfoIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_StateInfoIndex.DiscardUnknown(m)
}

var xxx_messageInfo_StateInfoIndex proto.InternalMessageInfo

func (m *StateInfoIndex) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *StateInfoIndex) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// StateInfo defines a rollapps' state.
type StateInfo struct {
	// stateInfoIndex defines what rollapp the state belongs to
	// and in which index it can be referenced
	StateInfoIndex StateInfoIndex `protobuf:"bytes,1,opt,name=stateInfoIndex,proto3" json:"stateInfoIndex"`
	// sequencer is the bech32-encoded address of the sequencer sent the update
	Sequencer string `protobuf:"bytes,2,opt,name=sequencer,proto3" json:"sequencer,omitempty"`
	// startHeight is the block height of the first block in the batch
	StartHeight uint64 `protobuf:"varint,3,opt,name=startHeight,proto3" json:"startHeight,omitempty"`
	// numBlocks is the number of blocks included in this batch update
	NumBlocks uint64 `protobuf:"varint,4,opt,name=numBlocks,proto3" json:"numBlocks,omitempty"`
	// DAPath is the description of the location on the DA layer
	DAPath string `protobuf:"bytes,5,opt,name=DAPath,proto3" json:"DAPath,omitempty"`
	// creationHeight is the height at which the UpdateState took place
	CreationHeight uint64 `protobuf:"varint,7,opt,name=creationHeight,proto3" json:"creationHeight,omitempty"`
	// status is the status of the state update
	Status common.Status `protobuf:"varint,8,opt,name=status,proto3,enum=dymensionxyz.dymension.common.Status" json:"status,omitempty"`
	// BDs is a list of block description objects (one per block)
	// the list must be ordered by height, starting from startHeight to startHeight+numBlocks-1
	BDs BlockDescriptors `protobuf:"bytes,9,opt,name=BDs,proto3" json:"BDs"`
	// created_at is the timestamp at which the StateInfo was created
	CreatedAt time.Time `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at" yaml:"created_at"`
}

func (m *StateInfo) Reset()         { *m = StateInfo{} }
func (m *StateInfo) String() string { return proto.CompactTextString(m) }
func (*StateInfo) ProtoMessage()    {}
func (*StateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_deebb9ffbcdd017e, []int{1}
}
func (m *StateInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateInfo.Merge(m, src)
}
func (m *StateInfo) XXX_Size() int {
	return m.Size()
}
func (m *StateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StateInfo proto.InternalMessageInfo

func (m *StateInfo) GetStateInfoIndex() StateInfoIndex {
	if m != nil {
		return m.StateInfoIndex
	}
	return StateInfoIndex{}
}

func (m *StateInfo) GetSequencer() string {
	if m != nil {
		return m.Sequencer
	}
	return ""
}

func (m *StateInfo) GetStartHeight() uint64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *StateInfo) GetNumBlocks() uint64 {
	if m != nil {
		return m.NumBlocks
	}
	return 0
}

func (m *StateInfo) GetDAPath() string {
	if m != nil {
		return m.DAPath
	}
	return ""
}

func (m *StateInfo) GetCreationHeight() uint64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

func (m *StateInfo) GetStatus() common.Status {
	if m != nil {
		return m.Status
	}
	return common.Status_PENDING
}

func (m *StateInfo) GetBDs() BlockDescriptors {
	if m != nil {
		return m.BDs
	}
	return BlockDescriptors{}
}

func (m *StateInfo) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

// StateInfoSummary is a compact representation of StateInfo
type StateInfoSummary struct {
	// stateInfoIndex defines what rollapp the state belongs to
	// and in which index it can be referenced
	StateInfoIndex StateInfoIndex `protobuf:"bytes,1,opt,name=stateInfoIndex,proto3" json:"stateInfoIndex"`
	// status is the status of the state update
	Status common.Status `protobuf:"varint,2,opt,name=status,proto3,enum=dymensionxyz.dymension.common.Status" json:"status,omitempty"`
	// creationHeight is the height at which the UpdateState took place
	CreationHeight uint64 `protobuf:"varint,3,opt,name=creationHeight,proto3" json:"creationHeight,omitempty"`
}

func (m *StateInfoSummary) Reset()         { *m = StateInfoSummary{} }
func (m *StateInfoSummary) String() string { return proto.CompactTextString(m) }
func (*StateInfoSummary) ProtoMessage()    {}
func (*StateInfoSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_deebb9ffbcdd017e, []int{2}
}
func (m *StateInfoSummary) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateInfoSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateInfoSummary.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateInfoSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateInfoSummary.Merge(m, src)
}
func (m *StateInfoSummary) XXX_Size() int {
	return m.Size()
}
func (m *StateInfoSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_StateInfoSummary.DiscardUnknown(m)
}

var xxx_messageInfo_StateInfoSummary proto.InternalMessageInfo

func (m *StateInfoSummary) GetStateInfoIndex() StateInfoIndex {
	if m != nil {
		return m.StateInfoIndex
	}
	return StateInfoIndex{}
}

func (m *StateInfoSummary) GetStatus() common.Status {
	if m != nil {
		return m.Status
	}
	return common.Status_PENDING
}

func (m *StateInfoSummary) GetCreationHeight() uint64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

// BlockHeightToFinalizationQueue defines a map from block height to list of states to finalized
type BlockHeightToFinalizationQueue struct {
	// creationHeight is the block height that the state should be finalized
	CreationHeight uint64 `protobuf:"varint,1,opt,name=creationHeight,proto3" json:"creationHeight,omitempty"`
	// finalizationQueue is a list of states that are waiting to be finalized
	// when the block height becomes creationHeight
	FinalizationQueue []StateInfoIndex `protobuf:"bytes,2,rep,name=finalizationQueue,proto3" json:"finalizationQueue"`
}

func (m *BlockHeightToFinalizationQueue) Reset()         { *m = BlockHeightToFinalizationQueue{} }
func (m *BlockHeightToFinalizationQueue) String() string { return proto.CompactTextString(m) }
func (*BlockHeightToFinalizationQueue) ProtoMessage()    {}
func (*BlockHeightToFinalizationQueue) Descriptor() ([]byte, []int) {
	return fileDescriptor_deebb9ffbcdd017e, []int{3}
}
func (m *BlockHeightToFinalizationQueue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeightToFinalizationQueue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeightToFinalizationQueue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeightToFinalizationQueue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeightToFinalizationQueue.Merge(m, src)
}
func (m *BlockHeightToFinalizationQueue) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeightToFinalizationQueue) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeightToFinalizationQueue.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeightToFinalizationQueue proto.InternalMessageInfo

func (m *BlockHeightToFinalizationQueue) GetCreationHeight() uint64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

func (m *BlockHeightToFinalizationQueue) GetFinalizationQueue() []StateInfoIndex {
	if m != nil {
		return m.FinalizationQueue
	}
	return nil
}

func init() {
	proto.RegisterType((*StateInfoIndex)(nil), "dymensionxyz.dymension.rollapp.StateInfoIndex")
	proto.RegisterType((*StateInfo)(nil), "dymensionxyz.dymension.rollapp.StateInfo")
	proto.RegisterType((*StateInfoSummary)(nil), "dymensionxyz.dymension.rollapp.StateInfoSummary")
	proto.RegisterType((*BlockHeightToFinalizationQueue)(nil), "dymensionxyz.dymension.rollapp.BlockHeightToFinalizationQueue")
}

func init() {
	proto.RegisterFile("types/dymensionxyz/dymension/rollapp/state_info.proto", fileDescriptor_deebb9ffbcdd017e)
}

var fileDescriptor_deebb9ffbcdd017e = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x41, 0x6a, 0xdb, 0x40,
	0x14, 0xf5, 0xc4, 0x8e, 0x13, 0x4d, 0xc0, 0x24, 0x43, 0x28, 0xc2, 0xb4, 0xb2, 0x11, 0xb4, 0x78,
	0x25, 0x85, 0x94, 0x6e, 0x5a, 0xba, 0x88, 0x31, 0x21, 0xee, 0xaa, 0x55, 0xb2, 0x28, 0xa5, 0xe0,
	0x8e, 0xa4, 0xb1, 0x3c, 0x54, 0x9a, 0x51, 0x35, 0x23, 0x88, 0x72, 0x8a, 0x1c, 0xa4, 0x07, 0xc9,
	0x32, 0xbb, 0x76, 0x95, 0x16, 0xfb, 0x02, 0xa5, 0x27, 0x28, 0x1a, 0x29, 0x56, 0x62, 0xbb, 0x49,
	0x09, 0x74, 0xe7, 0xff, 0xfd, 0xdf, 0xd3, 0xfb, 0xef, 0x3f, 0x06, 0xbe, 0x90, 0x59, 0x4c, 0x84,
	0xed, 0x67, 0x11, 0x61, 0x82, 0x72, 0x76, 0x9a, 0x9d, 0x55, 0x85, 0x9d, 0xf0, 0x30, 0xc4, 0x71,
	0x6c, 0x0b, 0x89, 0x25, 0x19, 0x51, 0x36, 0xe6, 0x56, 0x9c, 0x70, 0xc9, 0x91, 0x71, 0x13, 0x60,
	0xcd, 0x0b, 0xab, 0x04, 0xb4, 0x77, 0x03, 0x1e, 0x70, 0x35, 0x6a, 0xe7, 0xbf, 0x0a, 0x54, 0xbb,
	0x13, 0x70, 0x1e, 0x84, 0xc4, 0x56, 0x95, 0x9b, 0x8e, 0x6d, 0x49, 0x23, 0x22, 0x24, 0x8e, 0xe2,
	0x72, 0xe0, 0xd5, 0x3f, 0xa9, 0x71, 0x43, 0xee, 0x7d, 0x1e, 0xf9, 0x44, 0x78, 0x09, 0x8d, 0x25,
	0x4f, 0x4a, 0xf0, 0xde, 0x9d, 0x60, 0x8f, 0x47, 0x11, 0x67, 0x6a, 0x93, 0x54, 0x14, 0x08, 0x73,
	0x00, 0x5b, 0xc7, 0xf9, 0x66, 0x43, 0x36, 0xe6, 0x43, 0xe6, 0x93, 0x53, 0xf4, 0x18, 0x6a, 0xe5,
	0x57, 0x86, 0xbe, 0x0e, 0xba, 0xa0, 0xa7, 0x39, 0x55, 0x03, 0xed, 0xc2, 0x75, 0x9a, 0x8f, 0xe9,
	0x6b, 0x5d, 0xd0, 0x6b, 0x38, 0x45, 0x61, 0xfe, 0xaa, 0x43, 0x6d, 0x4e, 0x83, 0x3e, 0xc2, 0x96,
	0xb8, 0xc5, 0xa9, 0x68, 0xb6, 0xf6, 0x2d, 0xeb, 0x6e, 0xcb, 0xac, 0xdb, 0x4a, 0xfa, 0x8d, 0x8b,
	0xab, 0x4e, 0xcd, 0x59, 0xe0, 0xca, 0xf5, 0x09, 0xf2, 0x25, 0x25, 0xcc, 0x23, 0x89, 0x52, 0xa1,
	0x39, 0x55, 0x03, 0x75, 0xe1, 0x96, 0x90, 0x38, 0x91, 0x47, 0x84, 0x06, 0x13, 0xa9, 0xd7, 0x95,
	0xca, 0x9b, 0xad, 0x1c, 0xcf, 0xd2, 0xa8, 0x9f, 0x1b, 0x28, 0xf4, 0x86, 0xfa, 0xbf, 0x6a, 0xa0,
	0x47, 0xb0, 0x39, 0x38, 0x78, 0x8b, 0xe5, 0x44, 0x5f, 0x57, 0xd4, 0x65, 0x85, 0x9e, 0xc1, 0x96,
	0x97, 0x10, 0x2c, 0x29, 0x67, 0x25, 0xf5, 0x86, 0x82, 0x2e, 0x74, 0xd1, 0x6b, 0xd8, 0x2c, 0xfc,
	0xd5, 0x37, 0xbb, 0xa0, 0xd7, 0xda, 0x7f, 0xfa, 0xb7, 0x9d, 0x8b, 0x63, 0xa8, 0x95, 0x53, 0xe1,
	0x94, 0x20, 0x74, 0x04, 0xeb, 0xfd, 0x81, 0xd0, 0x35, 0xe5, 0xd7, 0xde, 0x7d, 0x7e, 0x29, 0xcd,
	0x83, 0x79, 0x08, 0x44, 0xe9, 0x58, 0x4e, 0x81, 0xde, 0x43, 0xa8, 0xa4, 0x11, 0x7f, 0x84, 0xa5,
	0x0e, 0x15, 0x61, 0xdb, 0x2a, 0xd2, 0x67, 0x5d, 0xa7, 0xcf, 0x3a, 0xb9, 0x4e, 0x5f, 0xff, 0x49,
	0x0e, 0xfd, 0x7d, 0xd5, 0xd9, 0xc9, 0x70, 0x14, 0xbe, 0x34, 0x2b, 0xac, 0x79, 0xfe, 0xa3, 0x03,
	0x1c, 0xad, 0x6c, 0x1c, 0xc8, 0x37, 0x8d, 0xcd, 0xe6, 0xf6, 0x86, 0xf9, 0x0d, 0xc0, 0xed, 0xf9,
	0xbd, 0x8e, 0xd3, 0x28, 0xc2, 0x49, 0xf6, 0x9f, 0x2f, 0x5f, 0x79, 0xbb, 0xf6, 0x10, 0x6f, 0x97,
	0x4f, 0x58, 0x5f, 0x75, 0x42, 0xf3, 0x2b, 0x80, 0x86, 0x72, 0xb6, 0xa8, 0x4f, 0xf8, 0x21, 0x65,
	0x38, 0xa4, 0x67, 0x6a, 0xe6, 0x5d, 0x4a, 0x52, 0xb2, 0x82, 0x0a, 0xac, 0x4c, 0x83, 0x0b, 0x77,
	0xc6, 0x8b, 0x60, 0x7d, 0xad, 0x5b, 0x7f, 0xb0, 0x25, 0xcb, 0x74, 0xfd, 0x4f, 0x17, 0x53, 0x03,
	0x5c, 0x4e, 0x0d, 0xf0, 0x73, 0x6a, 0x80, 0xf3, 0x99, 0x51, 0xbb, 0x9c, 0x19, 0xb5, 0xef, 0x33,
	0xa3, 0xf6, 0xe1, 0x30, 0xa0, 0x72, 0x92, 0xba, 0xb9, 0x1d, 0x4b, 0x4f, 0x02, 0x65, 0xd2, 0x2e,
	0x1e, 0x8b, 0xd8, 0xbd, 0xe7, 0xb1, 0x71, 0x9b, 0x2a, 0x2e, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xd2, 0xe2, 0x5a, 0x02, 0x29, 0x05, 0x00, 0x00,
}

func (m *StateInfoIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateInfoIndex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateInfoIndex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintStateInfo(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StateInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintStateInfo(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	{
		size, err := m.BDs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStateInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.Status != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if m.CreationHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x38
	}
	if len(m.DAPath) > 0 {
		i -= len(m.DAPath)
		copy(dAtA[i:], m.DAPath)
		i = encodeVarintStateInfo(dAtA, i, uint64(len(m.DAPath)))
		i--
		dAtA[i] = 0x2a
	}
	if m.NumBlocks != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.NumBlocks))
		i--
		dAtA[i] = 0x20
	}
	if m.StartHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Sequencer) > 0 {
		i -= len(m.Sequencer)
		copy(dAtA[i:], m.Sequencer)
		i = encodeVarintStateInfo(dAtA, i, uint64(len(m.Sequencer)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.StateInfoIndex.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStateInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *StateInfoSummary) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateInfoSummary) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateInfoSummary) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreationHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.Status != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.StateInfoIndex.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStateInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BlockHeightToFinalizationQueue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeightToFinalizationQueue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeightToFinalizationQueue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FinalizationQueue) > 0 {
		for iNdEx := len(m.FinalizationQueue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FinalizationQueue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStateInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CreationHeight != 0 {
		i = encodeVarintStateInfo(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStateInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovStateInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StateInfoIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovStateInfo(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovStateInfo(uint64(m.Index))
	}
	return n
}

func (m *StateInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.StateInfoIndex.Size()
	n += 1 + l + sovStateInfo(uint64(l))
	l = len(m.Sequencer)
	if l > 0 {
		n += 1 + l + sovStateInfo(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.StartHeight))
	}
	if m.NumBlocks != 0 {
		n += 1 + sovStateInfo(uint64(m.NumBlocks))
	}
	l = len(m.DAPath)
	if l > 0 {
		n += 1 + l + sovStateInfo(uint64(l))
	}
	if m.CreationHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.CreationHeight))
	}
	if m.Status != 0 {
		n += 1 + sovStateInfo(uint64(m.Status))
	}
	l = m.BDs.Size()
	n += 1 + l + sovStateInfo(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovStateInfo(uint64(l))
	return n
}

func (m *StateInfoSummary) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.StateInfoIndex.Size()
	n += 1 + l + sovStateInfo(uint64(l))
	if m.Status != 0 {
		n += 1 + sovStateInfo(uint64(m.Status))
	}
	if m.CreationHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.CreationHeight))
	}
	return n
}

func (m *BlockHeightToFinalizationQueue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreationHeight != 0 {
		n += 1 + sovStateInfo(uint64(m.CreationHeight))
	}
	if len(m.FinalizationQueue) > 0 {
		for _, e := range m.FinalizationQueue {
			l = e.Size()
			n += 1 + l + sovStateInfo(uint64(l))
		}
	}
	return n
}

func sovStateInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStateInfo(x uint64) (n int) {
	return sovStateInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StateInfoIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StateInfoIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateInfoIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StateInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StateInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateInfoIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StateInfoIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequencer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sequencer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlocks", wireType)
			}
			m.NumBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DAPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DAPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= common.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BDs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StateInfoSummary) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StateInfoSummary: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateInfoSummary: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateInfoIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StateInfoIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= common.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockHeightToFinalizationQueue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockHeightToFinalizationQueue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeightToFinalizationQueue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizationQueue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStateInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalizationQueue = append(m.FinalizationQueue, StateInfoIndex{})
			if err := m.FinalizationQueue[len(m.FinalizationQueue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStateInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStateInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStateInfo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStateInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStateInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStateInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStateInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStateInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStateInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStateInfo = fmt.Errorf("proto: unexpected end of group")
)