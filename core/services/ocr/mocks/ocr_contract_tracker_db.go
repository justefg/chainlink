// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	ocr "github.com/smartcontractkit/chainlink/v2/core/services/ocr"
	mock "github.com/stretchr/testify/mock"

	offchainaggregator "github.com/justefg/libocr/gethwrappers/offchainaggregator"

	sqlutil "github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
)

// OCRContractTrackerDB is an autogenerated mock type for the OCRContractTrackerDB type
type OCRContractTrackerDB struct {
	mock.Mock
}

type OCRContractTrackerDB_Expecter struct {
	mock *mock.Mock
}

func (_m *OCRContractTrackerDB) EXPECT() *OCRContractTrackerDB_Expecter {
	return &OCRContractTrackerDB_Expecter{mock: &_m.Mock}
}

// LoadLatestRoundRequested provides a mock function with given fields: ctx
func (_m *OCRContractTrackerDB) LoadLatestRoundRequested(ctx context.Context) (offchainaggregator.OffchainAggregatorRoundRequested, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LoadLatestRoundRequested")
	}

	var r0 offchainaggregator.OffchainAggregatorRoundRequested
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (offchainaggregator.OffchainAggregatorRoundRequested, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) offchainaggregator.OffchainAggregatorRoundRequested); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(offchainaggregator.OffchainAggregatorRoundRequested)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OCRContractTrackerDB_LoadLatestRoundRequested_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadLatestRoundRequested'
type OCRContractTrackerDB_LoadLatestRoundRequested_Call struct {
	*mock.Call
}

// LoadLatestRoundRequested is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OCRContractTrackerDB_Expecter) LoadLatestRoundRequested(ctx interface{}) *OCRContractTrackerDB_LoadLatestRoundRequested_Call {
	return &OCRContractTrackerDB_LoadLatestRoundRequested_Call{Call: _e.mock.On("LoadLatestRoundRequested", ctx)}
}

func (_c *OCRContractTrackerDB_LoadLatestRoundRequested_Call) Run(run func(ctx context.Context)) *OCRContractTrackerDB_LoadLatestRoundRequested_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OCRContractTrackerDB_LoadLatestRoundRequested_Call) Return(rr offchainaggregator.OffchainAggregatorRoundRequested, err error) *OCRContractTrackerDB_LoadLatestRoundRequested_Call {
	_c.Call.Return(rr, err)
	return _c
}

func (_c *OCRContractTrackerDB_LoadLatestRoundRequested_Call) RunAndReturn(run func(context.Context) (offchainaggregator.OffchainAggregatorRoundRequested, error)) *OCRContractTrackerDB_LoadLatestRoundRequested_Call {
	_c.Call.Return(run)
	return _c
}

// SaveLatestRoundRequested provides a mock function with given fields: ctx, rr
func (_m *OCRContractTrackerDB) SaveLatestRoundRequested(ctx context.Context, rr offchainaggregator.OffchainAggregatorRoundRequested) error {
	ret := _m.Called(ctx, rr)

	if len(ret) == 0 {
		panic("no return value specified for SaveLatestRoundRequested")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, offchainaggregator.OffchainAggregatorRoundRequested) error); ok {
		r0 = rf(ctx, rr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OCRContractTrackerDB_SaveLatestRoundRequested_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveLatestRoundRequested'
type OCRContractTrackerDB_SaveLatestRoundRequested_Call struct {
	*mock.Call
}

// SaveLatestRoundRequested is a helper method to define mock.On call
//   - ctx context.Context
//   - rr offchainaggregator.OffchainAggregatorRoundRequested
func (_e *OCRContractTrackerDB_Expecter) SaveLatestRoundRequested(ctx interface{}, rr interface{}) *OCRContractTrackerDB_SaveLatestRoundRequested_Call {
	return &OCRContractTrackerDB_SaveLatestRoundRequested_Call{Call: _e.mock.On("SaveLatestRoundRequested", ctx, rr)}
}

func (_c *OCRContractTrackerDB_SaveLatestRoundRequested_Call) Run(run func(ctx context.Context, rr offchainaggregator.OffchainAggregatorRoundRequested)) *OCRContractTrackerDB_SaveLatestRoundRequested_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(offchainaggregator.OffchainAggregatorRoundRequested))
	})
	return _c
}

func (_c *OCRContractTrackerDB_SaveLatestRoundRequested_Call) Return(_a0 error) *OCRContractTrackerDB_SaveLatestRoundRequested_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OCRContractTrackerDB_SaveLatestRoundRequested_Call) RunAndReturn(run func(context.Context, offchainaggregator.OffchainAggregatorRoundRequested) error) *OCRContractTrackerDB_SaveLatestRoundRequested_Call {
	_c.Call.Return(run)
	return _c
}

// WithDataSource provides a mock function with given fields: _a0
func (_m *OCRContractTrackerDB) WithDataSource(_a0 sqlutil.DataSource) ocr.OCRContractTrackerDB {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for WithDataSource")
	}

	var r0 ocr.OCRContractTrackerDB
	if rf, ok := ret.Get(0).(func(sqlutil.DataSource) ocr.OCRContractTrackerDB); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ocr.OCRContractTrackerDB)
		}
	}

	return r0
}

// OCRContractTrackerDB_WithDataSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithDataSource'
type OCRContractTrackerDB_WithDataSource_Call struct {
	*mock.Call
}

// WithDataSource is a helper method to define mock.On call
//   - _a0 sqlutil.DataSource
func (_e *OCRContractTrackerDB_Expecter) WithDataSource(_a0 interface{}) *OCRContractTrackerDB_WithDataSource_Call {
	return &OCRContractTrackerDB_WithDataSource_Call{Call: _e.mock.On("WithDataSource", _a0)}
}

func (_c *OCRContractTrackerDB_WithDataSource_Call) Run(run func(_a0 sqlutil.DataSource)) *OCRContractTrackerDB_WithDataSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sqlutil.DataSource))
	})
	return _c
}

func (_c *OCRContractTrackerDB_WithDataSource_Call) Return(_a0 ocr.OCRContractTrackerDB) *OCRContractTrackerDB_WithDataSource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OCRContractTrackerDB_WithDataSource_Call) RunAndReturn(run func(sqlutil.DataSource) ocr.OCRContractTrackerDB) *OCRContractTrackerDB_WithDataSource_Call {
	_c.Call.Return(run)
	return _c
}

// NewOCRContractTrackerDB creates a new instance of OCRContractTrackerDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOCRContractTrackerDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *OCRContractTrackerDB {
	mock := &OCRContractTrackerDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
