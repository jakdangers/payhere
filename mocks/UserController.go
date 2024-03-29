// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// UserController is an autogenerated mock type for the UserController type
type UserController struct {
	mock.Mock
}

type UserController_Expecter struct {
	mock *mock.Mock
}

func (_m *UserController) EXPECT() *UserController_Expecter {
	return &UserController_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: c
func (_m *UserController) CreateUser(c *gin.Context) {
	_m.Called(c)
}

// UserController_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type UserController_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - c *gin.Context
func (_e *UserController_Expecter) CreateUser(c interface{}) *UserController_CreateUser_Call {
	return &UserController_CreateUser_Call{Call: _e.mock.On("CreateUser", c)}
}

func (_c *UserController_CreateUser_Call) Run(run func(c *gin.Context)) *UserController_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *UserController_CreateUser_Call) Return() *UserController_CreateUser_Call {
	_c.Call.Return()
	return _c
}

func (_c *UserController_CreateUser_Call) RunAndReturn(run func(*gin.Context)) *UserController_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// LoginUser provides a mock function with given fields: c
func (_m *UserController) LoginUser(c *gin.Context) {
	_m.Called(c)
}

// UserController_LoginUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoginUser'
type UserController_LoginUser_Call struct {
	*mock.Call
}

// LoginUser is a helper method to define mock.On call
//   - c *gin.Context
func (_e *UserController_Expecter) LoginUser(c interface{}) *UserController_LoginUser_Call {
	return &UserController_LoginUser_Call{Call: _e.mock.On("LoginUser", c)}
}

func (_c *UserController_LoginUser_Call) Run(run func(c *gin.Context)) *UserController_LoginUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *UserController_LoginUser_Call) Return() *UserController_LoginUser_Call {
	_c.Call.Return()
	return _c
}

func (_c *UserController_LoginUser_Call) RunAndReturn(run func(*gin.Context)) *UserController_LoginUser_Call {
	_c.Call.Return(run)
	return _c
}

// LogoutUser provides a mock function with given fields: c
func (_m *UserController) LogoutUser(c *gin.Context) {
	_m.Called(c)
}

// UserController_LogoutUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogoutUser'
type UserController_LogoutUser_Call struct {
	*mock.Call
}

// LogoutUser is a helper method to define mock.On call
//   - c *gin.Context
func (_e *UserController_Expecter) LogoutUser(c interface{}) *UserController_LogoutUser_Call {
	return &UserController_LogoutUser_Call{Call: _e.mock.On("LogoutUser", c)}
}

func (_c *UserController_LogoutUser_Call) Run(run func(c *gin.Context)) *UserController_LogoutUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *UserController_LogoutUser_Call) Return() *UserController_LogoutUser_Call {
	_c.Call.Return()
	return _c
}

func (_c *UserController_LogoutUser_Call) RunAndReturn(run func(*gin.Context)) *UserController_LogoutUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserController creates a new instance of UserController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserController(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserController {
	mock := &UserController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
