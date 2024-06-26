// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	automation "github.com/smartcontractkit/chainlink-common/pkg/types/automation"

	mock "github.com/stretchr/testify/mock"
)

// UpkeepStateReader is an autogenerated mock type for the UpkeepStateReader type
type UpkeepStateReader struct {
	mock.Mock
}

// SelectByWorkIDs provides a mock function with given fields: ctx, workIDs
func (_m *UpkeepStateReader) SelectByWorkIDs(ctx context.Context, workIDs ...string) ([]automation.UpkeepState, error) {
	_va := make([]interface{}, len(workIDs))
	for _i := range workIDs {
		_va[_i] = workIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SelectByWorkIDs")
	}

	var r0 []automation.UpkeepState
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) ([]automation.UpkeepState, error)); ok {
		return rf(ctx, workIDs...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...string) []automation.UpkeepState); ok {
		r0 = rf(ctx, workIDs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]automation.UpkeepState)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(ctx, workIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUpkeepStateReader creates a new instance of UpkeepStateReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUpkeepStateReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *UpkeepStateReader {
	mock := &UpkeepStateReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
