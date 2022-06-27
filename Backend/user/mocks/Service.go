// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	user "diary/user"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetUserByID provides a mock function with given fields: ID
func (_m *Service) GetUserByID(ID int) (user.User, error) {
	ret := _m.Called(ID)

	var r0 user.User
	if rf, ok := ret.Get(0).(func(int) user.User); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsEmailAvailable provides a mock function with given fields: input
func (_m *Service) IsEmailAvailable(input user.CheckEmailInput) (bool, error) {
	ret := _m.Called(input)

	var r0 bool
	if rf, ok := ret.Get(0).(func(user.CheckEmailInput) bool); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.CheckEmailInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: input
func (_m *Service) Login(input user.LoginInput) (user.User, error) {
	ret := _m.Called(input)

	var r0 user.User
	if rf, ok := ret.Get(0).(func(user.LoginInput) user.User); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.LoginInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: input
func (_m *Service) RegisterUser(input user.RegisterUserInput) (user.User, error) {
	ret := _m.Called(input)

	var r0 user.User
	if rf, ok := ret.Get(0).(func(user.RegisterUserInput) user.User); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.RegisterUserInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAvatar provides a mock function with given fields: ID, fileLocation
func (_m *Service) SaveAvatar(ID int, fileLocation string) (user.User, error) {
	ret := _m.Called(ID, fileLocation)

	var r0 user.User
	if rf, ok := ret.Get(0).(func(int, string) user.User); ok {
		r0 = rf(ID, fileLocation)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(ID, fileLocation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}