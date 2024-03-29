// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "payhere/domain"

	mock "github.com/stretchr/testify/mock"
)

// ProductService is an autogenerated mock type for the ProductService type
type ProductService struct {
	mock.Mock
}

type ProductService_Expecter struct {
	mock *mock.Mock
}

func (_m *ProductService) EXPECT() *ProductService_Expecter {
	return &ProductService_Expecter{mock: &_m.Mock}
}

// CreateProduct provides a mock function with given fields: ctx, req
func (_m *ProductService) CreateProduct(ctx context.Context, req domain.CreateProductRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateProductRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProductService_CreateProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProduct'
type ProductService_CreateProduct_Call struct {
	*mock.Call
}

// CreateProduct is a helper method to define mock.On call
//   - ctx context.Context
//   - req domain.CreateProductRequest
func (_e *ProductService_Expecter) CreateProduct(ctx interface{}, req interface{}) *ProductService_CreateProduct_Call {
	return &ProductService_CreateProduct_Call{Call: _e.mock.On("CreateProduct", ctx, req)}
}

func (_c *ProductService_CreateProduct_Call) Run(run func(ctx context.Context, req domain.CreateProductRequest)) *ProductService_CreateProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.CreateProductRequest))
	})
	return _c
}

func (_c *ProductService_CreateProduct_Call) Return(_a0 error) *ProductService_CreateProduct_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProductService_CreateProduct_Call) RunAndReturn(run func(context.Context, domain.CreateProductRequest) error) *ProductService_CreateProduct_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProduct provides a mock function with given fields: ctx, req
func (_m *ProductService) DeleteProduct(ctx context.Context, req domain.DeleteProductRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.DeleteProductRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProductService_DeleteProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProduct'
type ProductService_DeleteProduct_Call struct {
	*mock.Call
}

// DeleteProduct is a helper method to define mock.On call
//   - ctx context.Context
//   - req domain.DeleteProductRequest
func (_e *ProductService_Expecter) DeleteProduct(ctx interface{}, req interface{}) *ProductService_DeleteProduct_Call {
	return &ProductService_DeleteProduct_Call{Call: _e.mock.On("DeleteProduct", ctx, req)}
}

func (_c *ProductService_DeleteProduct_Call) Run(run func(ctx context.Context, req domain.DeleteProductRequest)) *ProductService_DeleteProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.DeleteProductRequest))
	})
	return _c
}

func (_c *ProductService_DeleteProduct_Call) Return(_a0 error) *ProductService_DeleteProduct_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProductService_DeleteProduct_Call) RunAndReturn(run func(context.Context, domain.DeleteProductRequest) error) *ProductService_DeleteProduct_Call {
	_c.Call.Return(run)
	return _c
}

// GetProduct provides a mock function with given fields: ctx, req
func (_m *ProductService) GetProduct(ctx context.Context, req domain.GetProductRequest) (domain.GetProductResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.GetProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetProductRequest) (domain.GetProductResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.GetProductRequest) domain.GetProductResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.GetProductResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.GetProductRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductService_GetProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProduct'
type ProductService_GetProduct_Call struct {
	*mock.Call
}

// GetProduct is a helper method to define mock.On call
//   - ctx context.Context
//   - req domain.GetProductRequest
func (_e *ProductService_Expecter) GetProduct(ctx interface{}, req interface{}) *ProductService_GetProduct_Call {
	return &ProductService_GetProduct_Call{Call: _e.mock.On("GetProduct", ctx, req)}
}

func (_c *ProductService_GetProduct_Call) Run(run func(ctx context.Context, req domain.GetProductRequest)) *ProductService_GetProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.GetProductRequest))
	})
	return _c
}

func (_c *ProductService_GetProduct_Call) Return(_a0 domain.GetProductResponse, _a1 error) *ProductService_GetProduct_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProductService_GetProduct_Call) RunAndReturn(run func(context.Context, domain.GetProductRequest) (domain.GetProductResponse, error)) *ProductService_GetProduct_Call {
	_c.Call.Return(run)
	return _c
}

// ListProducts provides a mock function with given fields: ctx, req
func (_m *ProductService) ListProducts(ctx context.Context, req domain.ListProductsRequest) (domain.ListProductsResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.ListProductsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ListProductsRequest) (domain.ListProductsResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.ListProductsRequest) domain.ListProductsResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.ListProductsResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.ListProductsRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProductService_ListProducts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProducts'
type ProductService_ListProducts_Call struct {
	*mock.Call
}

// ListProducts is a helper method to define mock.On call
//   - ctx context.Context
//   - req domain.ListProductsRequest
func (_e *ProductService_Expecter) ListProducts(ctx interface{}, req interface{}) *ProductService_ListProducts_Call {
	return &ProductService_ListProducts_Call{Call: _e.mock.On("ListProducts", ctx, req)}
}

func (_c *ProductService_ListProducts_Call) Run(run func(ctx context.Context, req domain.ListProductsRequest)) *ProductService_ListProducts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.ListProductsRequest))
	})
	return _c
}

func (_c *ProductService_ListProducts_Call) Return(_a0 domain.ListProductsResponse, _a1 error) *ProductService_ListProducts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProductService_ListProducts_Call) RunAndReturn(run func(context.Context, domain.ListProductsRequest) (domain.ListProductsResponse, error)) *ProductService_ListProducts_Call {
	_c.Call.Return(run)
	return _c
}

// PatchProduct provides a mock function with given fields: ctx, req
func (_m *ProductService) PatchProduct(ctx context.Context, req domain.PatchProductRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PatchProductRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProductService_PatchProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchProduct'
type ProductService_PatchProduct_Call struct {
	*mock.Call
}

// PatchProduct is a helper method to define mock.On call
//   - ctx context.Context
//   - req domain.PatchProductRequest
func (_e *ProductService_Expecter) PatchProduct(ctx interface{}, req interface{}) *ProductService_PatchProduct_Call {
	return &ProductService_PatchProduct_Call{Call: _e.mock.On("PatchProduct", ctx, req)}
}

func (_c *ProductService_PatchProduct_Call) Run(run func(ctx context.Context, req domain.PatchProductRequest)) *ProductService_PatchProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.PatchProductRequest))
	})
	return _c
}

func (_c *ProductService_PatchProduct_Call) Return(_a0 error) *ProductService_PatchProduct_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProductService_PatchProduct_Call) RunAndReturn(run func(context.Context, domain.PatchProductRequest) error) *ProductService_PatchProduct_Call {
	_c.Call.Return(run)
	return _c
}

// NewProductService creates a new instance of ProductService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductService {
	mock := &ProductService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
