// Code generated by mockery v2.43.2. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	certs "github.com/absmach/magistrala/certs"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Remove provides a mock function with given fields: ctx, ownerID, thingID
func (_m *Repository) Remove(ctx context.Context, ownerID string, thingID string) error {
	ret := _m.Called(ctx, ownerID, thingID)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, ownerID, thingID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveAll provides a mock function with given fields: ctx, ownerID, offset, limit
func (_m *Repository) RetrieveAll(ctx context.Context, ownerID string, offset uint64, limit uint64) (certs.Page, error) {
	ret := _m.Called(ctx, ownerID, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAll")
	}

	var r0 certs.Page
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, uint64) (certs.Page, error)); ok {
		return rf(ctx, ownerID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, uint64) certs.Page); ok {
		r0 = rf(ctx, ownerID, offset, limit)
	} else {
		r0 = ret.Get(0).(certs.Page)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uint64, uint64) error); ok {
		r1 = rf(ctx, ownerID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveBySerial provides a mock function with given fields: ctx, ownerID, serialID
func (_m *Repository) RetrieveBySerial(ctx context.Context, ownerID string, serialID string) (certs.Cert, error) {
	ret := _m.Called(ctx, ownerID, serialID)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBySerial")
	}

	var r0 certs.Cert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (certs.Cert, error)); ok {
		return rf(ctx, ownerID, serialID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) certs.Cert); ok {
		r0 = rf(ctx, ownerID, serialID)
	} else {
		r0 = ret.Get(0).(certs.Cert)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, ownerID, serialID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveByThing provides a mock function with given fields: ctx, ownerID, thingID, offset, limit
func (_m *Repository) RetrieveByThing(ctx context.Context, ownerID string, thingID string, offset uint64, limit uint64) (certs.Page, error) {
	ret := _m.Called(ctx, ownerID, thingID, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByThing")
	}

	var r0 certs.Page
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, uint64, uint64) (certs.Page, error)); ok {
		return rf(ctx, ownerID, thingID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, uint64, uint64) certs.Page); ok {
		r0 = rf(ctx, ownerID, thingID, offset, limit)
	} else {
		r0 = ret.Get(0).(certs.Page)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, uint64, uint64) error); ok {
		r1 = rf(ctx, ownerID, thingID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, cert
func (_m *Repository) Save(ctx context.Context, cert certs.Cert) (string, error) {
	ret := _m.Called(ctx, cert)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, certs.Cert) (string, error)); ok {
		return rf(ctx, cert)
	}
	if rf, ok := ret.Get(0).(func(context.Context, certs.Cert) string); ok {
		r0 = rf(ctx, cert)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, certs.Cert) error); ok {
		r1 = rf(ctx, cert)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
