// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	models "github.com/smartcontractkit/chainlink/core/store/models"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddJob provides a mock function with given fields: _a0
func (_m *Service) AddJob(_a0 models.JobSpec) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.JobSpec) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveJob provides a mock function with given fields: _a0
func (_m *Service) RemoveJob(_a0 *models.ID) {
	_m.Called(_a0)
}

// Start provides a mock function with given fields:
func (_m *Service) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Service) Stop() {
	_m.Called()
}
