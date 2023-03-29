// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	keystore "github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	mock "github.com/stretchr/testify/mock"
)

// Master is an autogenerated mock type for the Master type
type Master struct {
	mock.Mock
}

// CSA provides a mock function with given fields:
func (_m *Master) CSA() keystore.CSA {
	ret := _m.Called()

	var r0 keystore.CSA
	if rf, ok := ret.Get(0).(func() keystore.CSA); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.CSA)
		}
	}

	return r0
}

// Cosmos provides a mock function with given fields:
func (_m *Master) Cosmos() keystore.Cosmos {
	ret := _m.Called()

	var r0 keystore.Cosmos
	if rf, ok := ret.Get(0).(func() keystore.Cosmos); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Cosmos)
		}
	}

	return r0
}

// DKGEncrypt provides a mock function with given fields:
func (_m *Master) DKGEncrypt() keystore.DKGEncrypt {
	ret := _m.Called()

	var r0 keystore.DKGEncrypt
	if rf, ok := ret.Get(0).(func() keystore.DKGEncrypt); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.DKGEncrypt)
		}
	}

	return r0
}

// DKGSign provides a mock function with given fields:
func (_m *Master) DKGSign() keystore.DKGSign {
	ret := _m.Called()

	var r0 keystore.DKGSign
	if rf, ok := ret.Get(0).(func() keystore.DKGSign); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.DKGSign)
		}
	}

	return r0
}

// Eth provides a mock function with given fields:
func (_m *Master) Eth() keystore.Eth {
	ret := _m.Called()

	var r0 keystore.Eth
	if rf, ok := ret.Get(0).(func() keystore.Eth); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Eth)
		}
	}

	return r0
}

// IsEmpty provides a mock function with given fields:
func (_m *Master) IsEmpty() (bool, error) {
	ret := _m.Called()

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrate provides a mock function with given fields: vrfPassword, f
func (_m *Master) Migrate(vrfPassword string, f keystore.DefaultEVMChainIDFunc) error {
	ret := _m.Called(vrfPassword, f)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, keystore.DefaultEVMChainIDFunc) error); ok {
		r0 = rf(vrfPassword, f)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OCR provides a mock function with given fields:
func (_m *Master) OCR() keystore.OCR {
	ret := _m.Called()

	var r0 keystore.OCR
	if rf, ok := ret.Get(0).(func() keystore.OCR); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.OCR)
		}
	}

	return r0
}

// OCR2 provides a mock function with given fields:
func (_m *Master) OCR2() keystore.OCR2 {
	ret := _m.Called()

	var r0 keystore.OCR2
	if rf, ok := ret.Get(0).(func() keystore.OCR2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.OCR2)
		}
	}

	return r0
}

// P2P provides a mock function with given fields:
func (_m *Master) P2P() keystore.P2P {
	ret := _m.Called()

	var r0 keystore.P2P
	if rf, ok := ret.Get(0).(func() keystore.P2P); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.P2P)
		}
	}

	return r0
}

// Solana provides a mock function with given fields:
func (_m *Master) Solana() keystore.Solana {
	ret := _m.Called()

	var r0 keystore.Solana
	if rf, ok := ret.Get(0).(func() keystore.Solana); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Solana)
		}
	}

	return r0
}

// Starknet provides a mock function with given fields:
func (_m *Master) Starknet() keystore.Starknet {
	ret := _m.Called()

	var r0 keystore.Starknet
	if rf, ok := ret.Get(0).(func() keystore.Starknet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Starknet)
		}
	}

	return r0
}

// Unlock provides a mock function with given fields: password
func (_m *Master) Unlock(password string) error {
	ret := _m.Called(password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VRF provides a mock function with given fields:
func (_m *Master) VRF() keystore.VRF {
	ret := _m.Called()

	var r0 keystore.VRF
	if rf, ok := ret.Get(0).(func() keystore.VRF); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.VRF)
		}
	}

	return r0
}

type mockConstructorTestingTNewMaster interface {
	mock.TestingT
	Cleanup(func())
}

// NewMaster creates a new instance of Master. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMaster(t mockConstructorTestingTNewMaster) *Master {
	mock := &Master{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
