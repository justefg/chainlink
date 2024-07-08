// Code generated by mockery v2.43.2. DO NOT EDIT.

package client

import (
	context "context"

	types "github.com/smartcontractkit/chainlink/v2/common/types"
	mock "github.com/stretchr/testify/mock"
)

// mockNode is an autogenerated mock type for the Node type
type mockNode[CHAIN_ID types.ID, RPC interface{}] struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, RPC]) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConfiguredChainID provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, RPC]) ConfiguredChainID() CHAIN_ID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ConfiguredChainID")
	}

	var r0 CHAIN_ID
	if rf, ok := ret.Get(0).(func() CHAIN_ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(CHAIN_ID)
	}

	return r0
}

// HighestUserObservations provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, RPC]) HighestUserObservations() ChainInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HighestUserObservations")
	}

	var r0 ChainInfo
	if rf, ok := ret.Get(0).(func() ChainInfo); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ChainInfo)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, RPC]) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Order provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, RPC]) Order() int32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Order")
	}

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// RPC provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, RPC]) RPC() RPC {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RPC")
	}

	var r0 RPC
	if rf, ok := ret.Get(0).(func() RPC); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(RPC)
	}

	return r0
}

// SetPoolChainInfoProvider provides a mock function with given fields: _a0
func (_m *mockNode[CHAIN_ID, RPC]) SetPoolChainInfoProvider(_a0 PoolChainInfoProvider) {
	_m.Called(_a0)
}

// Start provides a mock function with given fields: _a0
func (_m *mockNode[CHAIN_ID, RPC]) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// State provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, RPC]) State() NodeState {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for State")
	}

	var r0 NodeState
	if rf, ok := ret.Get(0).(func() NodeState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(NodeState)
	}

	return r0
}

// StateAndLatest provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, RPC]) StateAndLatest() (NodeState, ChainInfo) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StateAndLatest")
	}

	var r0 NodeState
	var r1 ChainInfo
	if rf, ok := ret.Get(0).(func() (NodeState, ChainInfo)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() NodeState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(NodeState)
	}

	if rf, ok := ret.Get(1).(func() ChainInfo); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(ChainInfo)
	}

	return r0, r1
}

// String provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, RPC]) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UnsubscribeAllExceptAliveLoop provides a mock function with given fields:
func (_m *mockNode[CHAIN_ID, RPC]) UnsubscribeAllExceptAliveLoop() {
	_m.Called()
}

// newMockNode creates a new instance of mockNode. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockNode[CHAIN_ID types.ID, RPC interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *mockNode[CHAIN_ID, RPC] {
	mock := &mockNode[CHAIN_ID, RPC]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
