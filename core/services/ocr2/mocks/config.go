// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	time "time"

	models "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/models"
	mock "github.com/stretchr/testify/mock"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// MercuryCredentials provides a mock function with given fields: credName
func (_m *Config) MercuryCredentials(credName string) *models.MercuryCredentials {
	ret := _m.Called(credName)

	var r0 *models.MercuryCredentials
	if rf, ok := ret.Get(0).(func(string) *models.MercuryCredentials); ok {
		r0 = rf(credName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MercuryCredentials)
		}
	}

	return r0
}

// OCR2BlockchainTimeout provides a mock function with given fields:
func (_m *Config) OCR2BlockchainTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCR2CaptureEATelemetry provides a mock function with given fields:
func (_m *Config) OCR2CaptureEATelemetry() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OCR2ContractConfirmations provides a mock function with given fields:
func (_m *Config) OCR2ContractConfirmations() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// OCR2ContractPollInterval provides a mock function with given fields:
func (_m *Config) OCR2ContractPollInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCR2ContractSubscribeInterval provides a mock function with given fields:
func (_m *Config) OCR2ContractSubscribeInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCR2ContractTransmitterTransmitTimeout provides a mock function with given fields:
func (_m *Config) OCR2ContractTransmitterTransmitTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCR2DatabaseTimeout provides a mock function with given fields:
func (_m *Config) OCR2DatabaseTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCR2DefaultTransactionQueueDepth provides a mock function with given fields:
func (_m *Config) OCR2DefaultTransactionQueueDepth() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// OCR2KeyBundleID provides a mock function with given fields:
func (_m *Config) OCR2KeyBundleID() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OCR2SimulateTransactions provides a mock function with given fields:
func (_m *Config) OCR2SimulateTransactions() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OCR2TraceLogging provides a mock function with given fields:
func (_m *Config) OCR2TraceLogging() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OCRDevelopmentMode provides a mock function with given fields:
func (_m *Config) OCRDevelopmentMode() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ThresholdKeyShare provides a mock function with given fields:
func (_m *Config) ThresholdKeyShare() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfig creates a new instance of Config. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfig(t mockConstructorTestingTNewConfig) *Config {
	mock := &Config{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
