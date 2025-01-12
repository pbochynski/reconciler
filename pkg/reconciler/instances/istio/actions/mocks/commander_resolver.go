// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	istioctl "github.com/kyma-incubator/reconciler/pkg/reconciler/instances/istio/istioctl"
	mock "github.com/stretchr/testify/mock"
)

// CommanderResolver is an autogenerated mock type for the CommanderResolver type
type CommanderResolver struct {
	mock.Mock
}

// GetCommander provides a mock function with given fields: version
func (_m *CommanderResolver) GetCommander(version istioctl.Version) (istioctl.Commander, error) {
	ret := _m.Called(version)

	var r0 istioctl.Commander
	if rf, ok := ret.Get(0).(func(istioctl.Version) istioctl.Commander); ok {
		r0 = rf(version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(istioctl.Commander)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(istioctl.Version) error); ok {
		r1 = rf(version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommanderResolver interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommanderResolver creates a new instance of CommanderResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommanderResolver(t mockConstructorTestingTNewCommanderResolver) *CommanderResolver {
	mock := &CommanderResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
