// Code generated by mockery v2.38.0. DO NOT EDIT.

package tokendata

import (
	context "context"

	internal "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

// MockReader is an autogenerated mock type for the Reader type
type MockReader struct {
	mock.Mock
}

// Close provides a mock function with given fields: qopts
func (_m *MockReader) Close(qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...pg.QOpt) error); ok {
		r0 = rf(qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadTokenData provides a mock function with given fields: ctx, msg
func (_m *MockReader) ReadTokenData(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) ([]byte, error) {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for ReadTokenData")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) ([]byte, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) []byte); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterFilters provides a mock function with given fields: qopts
func (_m *MockReader) RegisterFilters(qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterFilters")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...pg.QOpt) error); ok {
		r0 = rf(qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockReader creates a new instance of MockReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReader {
	mock := &MockReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
