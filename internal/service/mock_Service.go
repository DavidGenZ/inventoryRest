// Code generated by mockery v2.16.0. DO NOT EDIT.

package service

import (
	context "context"

	models "github.com/DavidGenZ/inventory-project/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// AddUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockService) AddUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoginUser provides a mock function with given fields: ctx, email, password
func (_m *MockService) LoginUser(ctx context.Context, email string, password string) (*models.User, error) {
	ret := _m.Called(ctx, email, password)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.User); ok {
		r0 = rf(ctx, email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, email, name, password
func (_m *MockService) RegisterUser(ctx context.Context, email string, name string, password string) error {
	ret := _m.Called(ctx, email, name, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, email, name, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveUserRole provides a mock function with given fields: ctx, userID, roleID
func (_m *MockService) RemoveUserRole(ctx context.Context, userID int64, roleID int64) error {
	ret := _m.Called(ctx, userID, roleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, userID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockService(t mockConstructorTestingTNewMockService) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
