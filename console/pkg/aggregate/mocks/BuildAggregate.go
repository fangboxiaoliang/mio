// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "hidevops.io/mio/pkg/apis/mio/v1alpha1"

// BuildAggregate is an autogenerated mock type for the BuildAggregate type
type BuildAggregate struct {
	mock.Mock
}

func (_m *BuildAggregate) ImagePush(build *v1alpha1.Build) error {
	panic("implement me")
}

// Compile provides a mock function with given fields: build
func (_m *BuildAggregate) Compile(build *v1alpha1.Build) error {
	ret := _m.Called(build)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Build) error); ok {
		r0 = rf(build)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: buildConfig, pipelineName, version
func (_m *BuildAggregate) Create(buildConfig *v1alpha1.BuildConfig, pipelineName string, version string) (*v1alpha1.Build, error) {
	ret := _m.Called(buildConfig, pipelineName, version)

	var r0 *v1alpha1.Build
	if rf, ok := ret.Get(0).(func(*v1alpha1.BuildConfig, string, string) *v1alpha1.Build); ok {
		r0 = rf(buildConfig, pipelineName, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1alpha1.BuildConfig, string, string) error); ok {
		r1 = rf(buildConfig, pipelineName, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateService provides a mock function with given fields: build
func (_m *BuildAggregate) CreateService(build *v1alpha1.Build) error {
	ret := _m.Called(build)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Build) error); ok {
		r0 = rf(build)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNode provides a mock function with given fields: build
func (_m *BuildAggregate) DeleteNode(build *v1alpha1.Build) error {
	ret := _m.Called(build)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Build) error); ok {
		r0 = rf(build)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployNode provides a mock function with given fields: build
func (_m *BuildAggregate) DeployNode(build *v1alpha1.Build) error {
	ret := _m.Called(build)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Build) error); ok {
		r0 = rf(build)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImageBuild provides a mock function with given fields: build
func (_m *BuildAggregate) ImageBuild(build *v1alpha1.Build) error {
	ret := _m.Called(build)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Build) error); ok {
		r0 = rf(build)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Selector provides a mock function with given fields: build
func (_m *BuildAggregate) Selector(build *v1alpha1.Build) error {
	ret := _m.Called(build)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Build) error); ok {
		r0 = rf(build)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SourceCodePull provides a mock function with given fields: build
func (_m *BuildAggregate) SourceCodePull(build *v1alpha1.Build) error {
	ret := _m.Called(build)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Build) error); ok {
		r0 = rf(build)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: build, event, phase
func (_m *BuildAggregate) Update(build *v1alpha1.Build, event string, phase string) error {
	ret := _m.Called(build, event, phase)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Build, string, string) error); ok {
		r0 = rf(build, event, phase)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Watch provides a mock function with given fields: name, namespace
func (_m *BuildAggregate) Watch(name string, namespace string) (*v1alpha1.Build, error) {
	ret := _m.Called(name, namespace)

	var r0 *v1alpha1.Build
	if rf, ok := ret.Get(0).(func(string, string) *v1alpha1.Build); ok {
		r0 = rf(name, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchPod provides a mock function with given fields: name, namespace
func (_m *BuildAggregate) WatchPod(name string, namespace string) error {
	ret := _m.Called(name, namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
