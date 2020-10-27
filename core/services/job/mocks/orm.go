// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	pipeline "github.com/smartcontractkit/chainlink/core/services/pipeline"

	postgres "github.com/smartcontractkit/chainlink/core/services/postgres"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// ClaimUnclaimedJobs provides a mock function with given fields: ctx
func (_m *ORM) ClaimUnclaimedJobs(ctx context.Context) ([]models.JobSpecV2, error) {
	ret := _m.Called(ctx)

	var r0 []models.JobSpecV2
	if rf, ok := ret.Get(0).(func(context.Context) []models.JobSpecV2); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.JobSpecV2)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *ORM) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateJob provides a mock function with given fields: ctx, jobSpec, taskDAG
func (_m *ORM) CreateJob(ctx context.Context, jobSpec *models.JobSpecV2, taskDAG pipeline.TaskDAG) error {
	ret := _m.Called(ctx, jobSpec, taskDAG)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.JobSpecV2, pipeline.TaskDAG) error); ok {
		r0 = rf(ctx, jobSpec, taskDAG)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteJob provides a mock function with given fields: ctx, id
func (_m *ORM) DeleteJob(ctx context.Context, id int32) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListenForNewJobs provides a mock function with given fields:
func (_m *ORM) ListenForNewJobs() (postgres.Subscription, error) {
	ret := _m.Called()

	var r0 postgres.Subscription
	if rf, ok := ret.Get(0).(func() postgres.Subscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(postgres.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
