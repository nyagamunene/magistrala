// Code generated by mockery v2.42.3. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RouteMapRepository is an autogenerated mock type for the RouteMapRepository type
type RouteMapRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *RouteMapRepository) Get(_a0 context.Context, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: _a0, _a1
func (_m *RouteMapRepository) Remove(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: _a0, _a1, _a2
func (_m *RouteMapRepository) Save(_a0 context.Context, _a1 string, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRouteMapRepository creates a new instance of RouteMapRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRouteMapRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *RouteMapRepository {
	mock := &RouteMapRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}