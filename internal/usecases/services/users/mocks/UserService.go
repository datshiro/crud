// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	models "github.com/datshiro/crud/internal/infras/models"
	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, email
func (_m *UserService) Create(name string, email string) (*models.User, error) {
	ret := _m.Called(name, email)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string, string) *models.User); ok {
		r0 = rf(name, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteById provides a mock function with given fields: id
func (_m *UserService) DeleteById(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: id
func (_m *UserService) FindById(id int) (*models.User, error) {
	ret := _m.Called(id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(int) *models.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByIdForUpdate provides a mock function with given fields: id
func (_m *UserService) FindByIdForUpdate(id int) (*models.User, error) {
	ret := _m.Called(id)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(int) *models.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPagination provides a mock function with given fields: page, limit
func (_m *UserService) GetPagination(page int, limit int) (models.UserSlice, error) {
	ret := _m.Called(page, limit)

	var r0 models.UserSlice
	if rf, ok := ret.Get(0).(func(int, int) models.UserSlice); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.UserSlice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: user
func (_m *UserService) Update(user *models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
