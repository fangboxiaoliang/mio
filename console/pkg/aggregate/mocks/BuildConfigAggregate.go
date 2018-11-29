// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import command "hidevops.io/mio/console/pkg/command"
import mock "github.com/stretchr/testify/mock"
import v1alpha1 "hidevops.io/mio/pkg/apis/mio/v1alpha1"

// BuildConfigAggregate is an autogenerated mock type for the BuildConfigAggregate type
type BuildConfigAggregate struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, pipelineName, namespace, sourceType, version
func (_m *BuildConfigAggregate) Create(name string, pipelineName string, namespace string, sourceType string, version string) (*v1alpha1.BuildConfig, error) {
	ret := _m.Called(name, pipelineName, namespace, sourceType, version)

	var r0 *v1alpha1.BuildConfig
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) *v1alpha1.BuildConfig); ok {
		r0 = rf(name, pipelineName, namespace, sourceType, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.BuildConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, string) error); ok {
		r1 = rf(name, pipelineName, namespace, sourceType, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name, namespace
func (_m *BuildConfigAggregate) Delete(name string, namespace string) error {
	ret := _m.Called(name, namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Template provides a mock function with given fields: buildConfigTemplate
func (_m *BuildConfigAggregate) Template(buildConfigTemplate *command.BuildConfig) (*v1alpha1.BuildConfig, error) {
	ret := _m.Called(buildConfigTemplate)

	var r0 *v1alpha1.BuildConfig
	if rf, ok := ret.Get(0).(func(*command.BuildConfig) *v1alpha1.BuildConfig); ok {
		r0 = rf(buildConfigTemplate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.BuildConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*command.BuildConfig) error); ok {
		r1 = rf(buildConfigTemplate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}