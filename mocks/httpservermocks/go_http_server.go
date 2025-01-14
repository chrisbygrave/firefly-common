// Code generated by mockery v2.38.0. DO NOT EDIT.

package httpservermocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	net "net"
)

// GoHTTPServer is an autogenerated mock type for the GoHTTPServer type
type GoHTTPServer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *GoHTTPServer) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Serve provides a mock function with given fields: l
func (_m *GoHTTPServer) Serve(l net.Listener) error {
	ret := _m.Called(l)

	if len(ret) == 0 {
		panic("no return value specified for Serve")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(net.Listener) error); ok {
		r0 = rf(l)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServeTLS provides a mock function with given fields: l, certFile, keyFile
func (_m *GoHTTPServer) ServeTLS(l net.Listener, certFile string, keyFile string) error {
	ret := _m.Called(l, certFile, keyFile)

	if len(ret) == 0 {
		panic("no return value specified for ServeTLS")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(net.Listener, string, string) error); ok {
		r0 = rf(l, certFile, keyFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields: ctx
func (_m *GoHTTPServer) Shutdown(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Shutdown")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewGoHTTPServer creates a new instance of GoHTTPServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGoHTTPServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *GoHTTPServer {
	mock := &GoHTTPServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
