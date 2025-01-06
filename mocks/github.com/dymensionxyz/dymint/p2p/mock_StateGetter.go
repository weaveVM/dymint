// Code generated by mockery v2.50.2. DO NOT EDIT.

package p2p

import (
	mock "github.com/stretchr/testify/mock"
	crypto "github.com/tendermint/tendermint/crypto"
)

// MockStateGetter is an autogenerated mock type for the StateGetter type
type MockStateGetter struct {
	mock.Mock
}

type MockStateGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStateGetter) EXPECT() *MockStateGetter_Expecter {
	return &MockStateGetter_Expecter{mock: &_m.Mock}
}

// GetProposerPubKey provides a mock function with no fields
func (_m *MockStateGetter) GetProposerPubKey() crypto.PubKey {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProposerPubKey")
	}

	var r0 crypto.PubKey
	if rf, ok := ret.Get(0).(func() crypto.PubKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PubKey)
		}
	}

	return r0
}

// MockStateGetter_GetProposerPubKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProposerPubKey'
type MockStateGetter_GetProposerPubKey_Call struct {
	*mock.Call
}

// GetProposerPubKey is a helper method to define mock.On call
func (_e *MockStateGetter_Expecter) GetProposerPubKey() *MockStateGetter_GetProposerPubKey_Call {
	return &MockStateGetter_GetProposerPubKey_Call{Call: _e.mock.On("GetProposerPubKey")}
}

func (_c *MockStateGetter_GetProposerPubKey_Call) Run(run func()) *MockStateGetter_GetProposerPubKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateGetter_GetProposerPubKey_Call) Return(_a0 crypto.PubKey) *MockStateGetter_GetProposerPubKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateGetter_GetProposerPubKey_Call) RunAndReturn(run func() crypto.PubKey) *MockStateGetter_GetProposerPubKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetRevision provides a mock function with no fields
func (_m *MockStateGetter) GetRevision() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRevision")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// MockStateGetter_GetRevision_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRevision'
type MockStateGetter_GetRevision_Call struct {
	*mock.Call
}

// GetRevision is a helper method to define mock.On call
func (_e *MockStateGetter_Expecter) GetRevision() *MockStateGetter_GetRevision_Call {
	return &MockStateGetter_GetRevision_Call{Call: _e.mock.On("GetRevision")}
}

func (_c *MockStateGetter_GetRevision_Call) Run(run func()) *MockStateGetter_GetRevision_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStateGetter_GetRevision_Call) Return(_a0 uint64) *MockStateGetter_GetRevision_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateGetter_GetRevision_Call) RunAndReturn(run func() uint64) *MockStateGetter_GetRevision_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStateGetter creates a new instance of MockStateGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStateGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStateGetter {
	mock := &MockStateGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
