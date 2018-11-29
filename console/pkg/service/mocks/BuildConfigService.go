// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import command "hidevops.io/mio/console/pkg/command"
import mock "github.com/stretchr/testify/mock"

// BuildConfigService is an autogenerated mock type for the BuildConfigService type
type BuildConfigService struct {
	mock.Mock
}

// Compile provides a mock function with given fields: host, port, cmd
func (_m *BuildConfigService) Compile(host string, port string, cmd *command.CompileCommand) error {
	ret := _m.Called(host, port, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *command.CompileCommand) error); ok {
		r0 = rf(host, port, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImageBuild provides a mock function with given fields: host, port, cmd
func (_m *BuildConfigService) ImageBuild(host string, port string, cmd *command.ImageBuildCommand) error {
	ret := _m.Called(host, port, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *command.ImageBuildCommand) error); ok {
		r0 = rf(host, port, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImagePush provides a mock function with given fields: host, port, cmd
func (_m *BuildConfigService) ImagePush(host string, port string, cmd *command.ImagePushCommand) error {
	ret := _m.Called(host, port, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *command.ImagePushCommand) error); ok {
		r0 = rf(host, port, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SourceCodePull provides a mock function with given fields: host, port, _a2
func (_m *BuildConfigService) SourceCodePull(host string, port string, _a2 *command.SourceCodePullCommand) error {
	ret := _m.Called(host, port, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *command.SourceCodePullCommand) error); ok {
		r0 = rf(host, port, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
