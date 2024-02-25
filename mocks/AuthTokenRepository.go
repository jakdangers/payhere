// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "payhere/domain"

	mock "github.com/stretchr/testify/mock"
)

// AuthTokenRepository is an autogenerated mock type for the AuthTokenRepository type
type AuthTokenRepository struct {
	mock.Mock
}

type AuthTokenRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthTokenRepository) EXPECT() *AuthTokenRepository_Expecter {
	return &AuthTokenRepository_Expecter{mock: &_m.Mock}
}

// CreateAuthToken provides a mock function with given fields: ctx, token
func (_m *AuthTokenRepository) CreateAuthToken(ctx context.Context, token domain.AuthToken) (int, error) {
	ret := _m.Called(ctx, token)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.AuthToken) (int, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.AuthToken) int); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.AuthToken) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthTokenRepository_CreateAuthToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAuthToken'
type AuthTokenRepository_CreateAuthToken_Call struct {
	*mock.Call
}

// CreateAuthToken is a helper method to define mock.On call
//   - ctx context.Context
//   - token domain.AuthToken
func (_e *AuthTokenRepository_Expecter) CreateAuthToken(ctx interface{}, token interface{}) *AuthTokenRepository_CreateAuthToken_Call {
	return &AuthTokenRepository_CreateAuthToken_Call{Call: _e.mock.On("CreateAuthToken", ctx, token)}
}

func (_c *AuthTokenRepository_CreateAuthToken_Call) Run(run func(ctx context.Context, token domain.AuthToken)) *AuthTokenRepository_CreateAuthToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.AuthToken))
	})
	return _c
}

func (_c *AuthTokenRepository_CreateAuthToken_Call) Return(_a0 int, _a1 error) *AuthTokenRepository_CreateAuthToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthTokenRepository_CreateAuthToken_Call) RunAndReturn(run func(context.Context, domain.AuthToken) (int, error)) *AuthTokenRepository_CreateAuthToken_Call {
	_c.Call.Return(run)
	return _c
}

// DeactivateAuthToken provides a mock function with given fields: ctx, params
func (_m *AuthTokenRepository) DeactivateAuthToken(ctx context.Context, params domain.DeactivateAuthTokenParams) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.DeactivateAuthTokenParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthTokenRepository_DeactivateAuthToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeactivateAuthToken'
type AuthTokenRepository_DeactivateAuthToken_Call struct {
	*mock.Call
}

// DeactivateAuthToken is a helper method to define mock.On call
//   - ctx context.Context
//   - params domain.DeactivateAuthTokenParams
func (_e *AuthTokenRepository_Expecter) DeactivateAuthToken(ctx interface{}, params interface{}) *AuthTokenRepository_DeactivateAuthToken_Call {
	return &AuthTokenRepository_DeactivateAuthToken_Call{Call: _e.mock.On("DeactivateAuthToken", ctx, params)}
}

func (_c *AuthTokenRepository_DeactivateAuthToken_Call) Run(run func(ctx context.Context, params domain.DeactivateAuthTokenParams)) *AuthTokenRepository_DeactivateAuthToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.DeactivateAuthTokenParams))
	})
	return _c
}

func (_c *AuthTokenRepository_DeactivateAuthToken_Call) Return(_a0 error) *AuthTokenRepository_DeactivateAuthToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthTokenRepository_DeactivateAuthToken_Call) RunAndReturn(run func(context.Context, domain.DeactivateAuthTokenParams) error) *AuthTokenRepository_DeactivateAuthToken_Call {
	_c.Call.Return(run)
	return _c
}

// FindAuthTokenByUserIDAndJwtToken provides a mock function with given fields: ctx, params
func (_m *AuthTokenRepository) FindAuthTokenByUserIDAndJwtToken(ctx context.Context, params domain.FindByUserIDAndJwtTokenParams) (domain.AuthToken, error) {
	ret := _m.Called(ctx, params)

	var r0 domain.AuthToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.FindByUserIDAndJwtTokenParams) (domain.AuthToken, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.FindByUserIDAndJwtTokenParams) domain.AuthToken); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(domain.AuthToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.FindByUserIDAndJwtTokenParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthTokenRepository_FindAuthTokenByUserIDAndJwtToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAuthTokenByUserIDAndJwtToken'
type AuthTokenRepository_FindAuthTokenByUserIDAndJwtToken_Call struct {
	*mock.Call
}

// FindAuthTokenByUserIDAndJwtToken is a helper method to define mock.On call
//   - ctx context.Context
//   - params domain.FindByUserIDAndJwtTokenParams
func (_e *AuthTokenRepository_Expecter) FindAuthTokenByUserIDAndJwtToken(ctx interface{}, params interface{}) *AuthTokenRepository_FindAuthTokenByUserIDAndJwtToken_Call {
	return &AuthTokenRepository_FindAuthTokenByUserIDAndJwtToken_Call{Call: _e.mock.On("FindAuthTokenByUserIDAndJwtToken", ctx, params)}
}

func (_c *AuthTokenRepository_FindAuthTokenByUserIDAndJwtToken_Call) Run(run func(ctx context.Context, params domain.FindByUserIDAndJwtTokenParams)) *AuthTokenRepository_FindAuthTokenByUserIDAndJwtToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.FindByUserIDAndJwtTokenParams))
	})
	return _c
}

func (_c *AuthTokenRepository_FindAuthTokenByUserIDAndJwtToken_Call) Return(_a0 domain.AuthToken, _a1 error) *AuthTokenRepository_FindAuthTokenByUserIDAndJwtToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthTokenRepository_FindAuthTokenByUserIDAndJwtToken_Call) RunAndReturn(run func(context.Context, domain.FindByUserIDAndJwtTokenParams) (domain.AuthToken, error)) *AuthTokenRepository_FindAuthTokenByUserIDAndJwtToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthTokenRepository creates a new instance of AuthTokenRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthTokenRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthTokenRepository {
	mock := &AuthTokenRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}