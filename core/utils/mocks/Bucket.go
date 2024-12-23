// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	bytes "bytes"

	mock "github.com/stretchr/testify/mock"
)

// Bucket is an autogenerated mock type for the Bucket type
type Bucket struct {
	mock.Mock
}

// GeneratePreSignedURL provides a mock function with given fields: fileName
func (_m *Bucket) GeneratePreSignedURL(fileName string) (string, error) {
	ret := _m.Called(fileName)

	if len(ret) == 0 {
		panic("no return value specified for GeneratePreSignedURL")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(fileName)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(fileName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadFileToBucket provides a mock function with given fields: fileContent, fileName
func (_m *Bucket) UploadFileToBucket(fileContent *bytes.Reader, fileName string) (string, error) {
	ret := _m.Called(fileContent, fileName)

	if len(ret) == 0 {
		panic("no return value specified for UploadFileToBucket")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*bytes.Reader, string) (string, error)); ok {
		return rf(fileContent, fileName)
	}
	if rf, ok := ret.Get(0).(func(*bytes.Reader, string) string); ok {
		r0 = rf(fileContent, fileName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*bytes.Reader, string) error); ok {
		r1 = rf(fileContent, fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBucket creates a new instance of Bucket. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBucket(t interface {
	mock.TestingT
	Cleanup(func())
}) *Bucket {
	mock := &Bucket{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
