// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	generated "github.com/smartcontractkit/chainlink/core/gethwrappers/generated"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// AbigenContract is an autogenerated mock type for the AbigenContract type
type AbigenContract struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *AbigenContract) Address() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// ParseLog provides a mock function with given fields: _a0
func (_m *AbigenContract) ParseLog(_a0 types.Log) (generated.AbigenLog, error) {
	ret := _m.Called(_a0)

	var r0 generated.AbigenLog
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (generated.AbigenLog, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(types.Log) generated.AbigenLog); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.AbigenLog)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAbigenContract interface {
	mock.TestingT
	Cleanup(func())
}

// NewAbigenContract creates a new instance of AbigenContract. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAbigenContract(t mockConstructorTestingTNewAbigenContract) *AbigenContract {
	mock := &AbigenContract{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
