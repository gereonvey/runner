// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package runner

import mock "github.com/stretchr/testify/mock"

// MockRunnerService is an autogenerated mock type for the RunnerService type
type MockRunnerService struct {
	mock.Mock
}

// BuildCatalog provides a mock function with given fields:
func (_m *MockRunnerService) BuildCatalog() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Execute provides a mock function with given fields: e
func (_m *MockRunnerService) Execute(e *ExecutionEvent) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ExecutionEvent) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCatalog provides a mock function with given fields:
func (_m *MockRunnerService) GetCatalog() *Catalog {
	ret := _m.Called()

	var r0 *Catalog
	if rf, ok := ret.Get(0).(func() *Catalog); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Catalog)
		}
	}

	return r0
}

// GetChannel provides a mock function with given fields:
func (_m *MockRunnerService) GetChannel() chan *ExecutionEvent {
	ret := _m.Called()

	var r0 chan *ExecutionEvent
	if rf, ok := ret.Get(0).(func() chan *ExecutionEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *ExecutionEvent)
		}
	}

	return r0
}

// IsCatalogReady provides a mock function with given fields:
func (_m *MockRunnerService) IsCatalogReady() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ScheduleExecution provides a mock function with given fields: e
func (_m *MockRunnerService) ScheduleExecution(e *ExecutionEvent) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ExecutionEvent) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
