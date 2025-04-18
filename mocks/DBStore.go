// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DBStore is an autogenerated mock type for the DBStore type
type DBStore struct {
	mock.Mock
}

// CleanupExpiredKeys provides a mock function with given fields:
func (_m *DBStore) CleanupExpiredKeys() {
	_m.Called()
}

// Delete provides a mock function with given fields: key
func (_m *DBStore) Delete(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *DBStore) Get(key string) (interface{}, error) {
	ret := _m.Called(key)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (interface{}, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *DBStore) GetAll() map[string]interface{} {
	ret := _m.Called()

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// GetTTL provides a mock function with given fields: key
func (_m *DBStore) GetTTL(key string) (int64, error) {
	ret := _m.Called(key)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: key, value, ttlSeconds
func (_m *DBStore) Set(key string, value interface{}, ttlSeconds int64) {
	_m.Called(key, value, ttlSeconds)
}

// SetTTL provides a mock function with given fields: key, ttlSeconds
func (_m *DBStore) SetTTL(key string, ttlSeconds int64) bool {
	ret := _m.Called(key, ttlSeconds)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, int64) bool); ok {
		r0 = rf(key, ttlSeconds)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewDBStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewDBStore creates a new instance of DBStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDBStore(t mockConstructorTestingTNewDBStore) *DBStore {
	mock := &DBStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
