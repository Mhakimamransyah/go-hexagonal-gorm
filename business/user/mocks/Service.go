// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	user "go-hexagonal/business/user"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// FindAllUserWithPagination provides a mock function with given fields: skip, rowPerPage
func (_m *Service) FindAllUserWithPagination(skip int, rowPerPage int) ([]user.User, error) {
	ret := _m.Called(skip, rowPerPage)

	var r0 []user.User
	if rf, ok := ret.Get(0).(func(int, int) []user.User); ok {
		r0 = rf(skip, rowPerPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(skip, rowPerPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByID provides a mock function with given fields: id
func (_m *Service) FindUserByID(id string) (*user.User, error) {
	ret := _m.Called(id)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: insertUserSpec, createdBy
func (_m *Service) InsertUser(insertUserSpec user.InsertUserSpec, createdBy string) error {
	ret := _m.Called(insertUserSpec, createdBy)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.InsertUserSpec, string) error); ok {
		r0 = rf(insertUserSpec, createdBy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: id, name, modifiedBy, currentVersion
func (_m *Service) UpdateUser(id string, name string, modifiedBy string, currentVersion int) error {
	ret := _m.Called(id, name, modifiedBy, currentVersion)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, int) error); ok {
		r0 = rf(id, name, modifiedBy, currentVersion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}