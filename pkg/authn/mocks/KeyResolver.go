// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	authn "github.com/sourcenetwork/orbis-go/pkg/authn"

	mock "github.com/stretchr/testify/mock"
)

// KeyResolver is an autogenerated mock type for the KeyResolver type
type KeyResolver struct {
	mock.Mock
}

type KeyResolver_Expecter struct {
	mock *mock.Mock
}

func (_m *KeyResolver) EXPECT() *KeyResolver_Expecter {
	return &KeyResolver_Expecter{mock: &_m.Mock}
}

// Resolve provides a mock function with given fields: _a0, _a1
func (_m *KeyResolver) Resolve(_a0 context.Context, _a1 string) (authn.SubjectInfo, error) {
	ret := _m.Called(_a0, _a1)

	var r0 authn.SubjectInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (authn.SubjectInfo, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) authn.SubjectInfo); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(authn.SubjectInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeyResolver_Resolve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resolve'
type KeyResolver_Resolve_Call struct {
	*mock.Call
}

// Resolve is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *KeyResolver_Expecter) Resolve(_a0 interface{}, _a1 interface{}) *KeyResolver_Resolve_Call {
	return &KeyResolver_Resolve_Call{Call: _e.mock.On("Resolve", _a0, _a1)}
}

func (_c *KeyResolver_Resolve_Call) Run(run func(_a0 context.Context, _a1 string)) *KeyResolver_Resolve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *KeyResolver_Resolve_Call) Return(_a0 authn.SubjectInfo, _a1 error) *KeyResolver_Resolve_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KeyResolver_Resolve_Call) RunAndReturn(run func(context.Context, string) (authn.SubjectInfo, error)) *KeyResolver_Resolve_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewKeyResolver interface {
	mock.TestingT
	Cleanup(func())
}

// NewKeyResolver creates a new instance of KeyResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKeyResolver(t mockConstructorTestingTNewKeyResolver) *KeyResolver {
	mock := &KeyResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
