// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	log "github.com/smartcontractkit/chainlink/core/chains/evm/log"
	mock "github.com/stretchr/testify/mock"
)

// Listener is an autogenerated mock type for the Listener type
type Listener struct {
	mock.Mock
}

// HandleLog provides a mock function with given fields: b
func (_m *Listener) HandleLog(b log.Broadcast) {
	_m.Called(b)
}

// JobID provides a mock function with given fields:
func (_m *Listener) JobID() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

type NewListenerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewListener creates a new instance of Listener. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewListener(t NewListenerT) *Listener {
	mock := &Listener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
