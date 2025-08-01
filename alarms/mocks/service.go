// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify
// Copyright (c) Abstract Machines

// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"context"

	"github.com/absmach/magistrala/alarms"
	"github.com/absmach/supermq/pkg/authn"
	mock "github.com/stretchr/testify/mock"
)

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// CreateAlarm provides a mock function for the type Service
func (_mock *Service) CreateAlarm(ctx context.Context, alarm alarms.Alarm) error {
	ret := _mock.Called(ctx, alarm)

	if len(ret) == 0 {
		panic("no return value specified for CreateAlarm")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, alarms.Alarm) error); ok {
		r0 = returnFunc(ctx, alarm)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// Service_CreateAlarm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAlarm'
type Service_CreateAlarm_Call struct {
	*mock.Call
}

// CreateAlarm is a helper method to define mock.On call
//   - ctx context.Context
//   - alarm alarms.Alarm
func (_e *Service_Expecter) CreateAlarm(ctx interface{}, alarm interface{}) *Service_CreateAlarm_Call {
	return &Service_CreateAlarm_Call{Call: _e.mock.On("CreateAlarm", ctx, alarm)}
}

func (_c *Service_CreateAlarm_Call) Run(run func(ctx context.Context, alarm alarms.Alarm)) *Service_CreateAlarm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 alarms.Alarm
		if args[1] != nil {
			arg1 = args[1].(alarms.Alarm)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *Service_CreateAlarm_Call) Return(err error) *Service_CreateAlarm_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *Service_CreateAlarm_Call) RunAndReturn(run func(ctx context.Context, alarm alarms.Alarm) error) *Service_CreateAlarm_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAlarm provides a mock function for the type Service
func (_mock *Service) DeleteAlarm(ctx context.Context, session authn.Session, id string) error {
	ret := _mock.Called(ctx, session, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAlarm")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, authn.Session, string) error); ok {
		r0 = returnFunc(ctx, session, id)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// Service_DeleteAlarm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAlarm'
type Service_DeleteAlarm_Call struct {
	*mock.Call
}

// DeleteAlarm is a helper method to define mock.On call
//   - ctx context.Context
//   - session authn.Session
//   - id string
func (_e *Service_Expecter) DeleteAlarm(ctx interface{}, session interface{}, id interface{}) *Service_DeleteAlarm_Call {
	return &Service_DeleteAlarm_Call{Call: _e.mock.On("DeleteAlarm", ctx, session, id)}
}

func (_c *Service_DeleteAlarm_Call) Run(run func(ctx context.Context, session authn.Session, id string)) *Service_DeleteAlarm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 authn.Session
		if args[1] != nil {
			arg1 = args[1].(authn.Session)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *Service_DeleteAlarm_Call) Return(err error) *Service_DeleteAlarm_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *Service_DeleteAlarm_Call) RunAndReturn(run func(ctx context.Context, session authn.Session, id string) error) *Service_DeleteAlarm_Call {
	_c.Call.Return(run)
	return _c
}

// ListAlarms provides a mock function for the type Service
func (_mock *Service) ListAlarms(ctx context.Context, session authn.Session, pm alarms.PageMetadata) (alarms.AlarmsPage, error) {
	ret := _mock.Called(ctx, session, pm)

	if len(ret) == 0 {
		panic("no return value specified for ListAlarms")
	}

	var r0 alarms.AlarmsPage
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, authn.Session, alarms.PageMetadata) (alarms.AlarmsPage, error)); ok {
		return returnFunc(ctx, session, pm)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, authn.Session, alarms.PageMetadata) alarms.AlarmsPage); ok {
		r0 = returnFunc(ctx, session, pm)
	} else {
		r0 = ret.Get(0).(alarms.AlarmsPage)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, authn.Session, alarms.PageMetadata) error); ok {
		r1 = returnFunc(ctx, session, pm)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Service_ListAlarms_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAlarms'
type Service_ListAlarms_Call struct {
	*mock.Call
}

// ListAlarms is a helper method to define mock.On call
//   - ctx context.Context
//   - session authn.Session
//   - pm alarms.PageMetadata
func (_e *Service_Expecter) ListAlarms(ctx interface{}, session interface{}, pm interface{}) *Service_ListAlarms_Call {
	return &Service_ListAlarms_Call{Call: _e.mock.On("ListAlarms", ctx, session, pm)}
}

func (_c *Service_ListAlarms_Call) Run(run func(ctx context.Context, session authn.Session, pm alarms.PageMetadata)) *Service_ListAlarms_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 authn.Session
		if args[1] != nil {
			arg1 = args[1].(authn.Session)
		}
		var arg2 alarms.PageMetadata
		if args[2] != nil {
			arg2 = args[2].(alarms.PageMetadata)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *Service_ListAlarms_Call) Return(alarmsPage alarms.AlarmsPage, err error) *Service_ListAlarms_Call {
	_c.Call.Return(alarmsPage, err)
	return _c
}

func (_c *Service_ListAlarms_Call) RunAndReturn(run func(ctx context.Context, session authn.Session, pm alarms.PageMetadata) (alarms.AlarmsPage, error)) *Service_ListAlarms_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAlarm provides a mock function for the type Service
func (_mock *Service) UpdateAlarm(ctx context.Context, session authn.Session, alarm alarms.Alarm) (alarms.Alarm, error) {
	ret := _mock.Called(ctx, session, alarm)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAlarm")
	}

	var r0 alarms.Alarm
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, authn.Session, alarms.Alarm) (alarms.Alarm, error)); ok {
		return returnFunc(ctx, session, alarm)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, authn.Session, alarms.Alarm) alarms.Alarm); ok {
		r0 = returnFunc(ctx, session, alarm)
	} else {
		r0 = ret.Get(0).(alarms.Alarm)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, authn.Session, alarms.Alarm) error); ok {
		r1 = returnFunc(ctx, session, alarm)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Service_UpdateAlarm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAlarm'
type Service_UpdateAlarm_Call struct {
	*mock.Call
}

// UpdateAlarm is a helper method to define mock.On call
//   - ctx context.Context
//   - session authn.Session
//   - alarm alarms.Alarm
func (_e *Service_Expecter) UpdateAlarm(ctx interface{}, session interface{}, alarm interface{}) *Service_UpdateAlarm_Call {
	return &Service_UpdateAlarm_Call{Call: _e.mock.On("UpdateAlarm", ctx, session, alarm)}
}

func (_c *Service_UpdateAlarm_Call) Run(run func(ctx context.Context, session authn.Session, alarm alarms.Alarm)) *Service_UpdateAlarm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 authn.Session
		if args[1] != nil {
			arg1 = args[1].(authn.Session)
		}
		var arg2 alarms.Alarm
		if args[2] != nil {
			arg2 = args[2].(alarms.Alarm)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *Service_UpdateAlarm_Call) Return(alarm1 alarms.Alarm, err error) *Service_UpdateAlarm_Call {
	_c.Call.Return(alarm1, err)
	return _c
}

func (_c *Service_UpdateAlarm_Call) RunAndReturn(run func(ctx context.Context, session authn.Session, alarm alarms.Alarm) (alarms.Alarm, error)) *Service_UpdateAlarm_Call {
	_c.Call.Return(run)
	return _c
}

// ViewAlarm provides a mock function for the type Service
func (_mock *Service) ViewAlarm(ctx context.Context, session authn.Session, id string) (alarms.Alarm, error) {
	ret := _mock.Called(ctx, session, id)

	if len(ret) == 0 {
		panic("no return value specified for ViewAlarm")
	}

	var r0 alarms.Alarm
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, authn.Session, string) (alarms.Alarm, error)); ok {
		return returnFunc(ctx, session, id)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, authn.Session, string) alarms.Alarm); ok {
		r0 = returnFunc(ctx, session, id)
	} else {
		r0 = ret.Get(0).(alarms.Alarm)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, authn.Session, string) error); ok {
		r1 = returnFunc(ctx, session, id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Service_ViewAlarm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ViewAlarm'
type Service_ViewAlarm_Call struct {
	*mock.Call
}

// ViewAlarm is a helper method to define mock.On call
//   - ctx context.Context
//   - session authn.Session
//   - id string
func (_e *Service_Expecter) ViewAlarm(ctx interface{}, session interface{}, id interface{}) *Service_ViewAlarm_Call {
	return &Service_ViewAlarm_Call{Call: _e.mock.On("ViewAlarm", ctx, session, id)}
}

func (_c *Service_ViewAlarm_Call) Run(run func(ctx context.Context, session authn.Session, id string)) *Service_ViewAlarm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 authn.Session
		if args[1] != nil {
			arg1 = args[1].(authn.Session)
		}
		var arg2 string
		if args[2] != nil {
			arg2 = args[2].(string)
		}
		run(
			arg0,
			arg1,
			arg2,
		)
	})
	return _c
}

func (_c *Service_ViewAlarm_Call) Return(alarm alarms.Alarm, err error) *Service_ViewAlarm_Call {
	_c.Call.Return(alarm, err)
	return _c
}

func (_c *Service_ViewAlarm_Call) RunAndReturn(run func(ctx context.Context, session authn.Session, id string) (alarms.Alarm, error)) *Service_ViewAlarm_Call {
	_c.Call.Return(run)
	return _c
}
