// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	entity "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockCustomerUseCase is an autogenerated mock type for the CustomerUseCase type
type MockCustomerUseCase struct {
	mock.Mock
}

type MockCustomerUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCustomerUseCase) EXPECT() *MockCustomerUseCase_Expecter {
	return &MockCustomerUseCase_Expecter{mock: &_m.Mock}
}

// CreateCustomer provides a mock function with given fields: _a0, _a1
func (_m *MockCustomerUseCase) CreateCustomer(_a0 context.Context, _a1 dto.CustomerCreateDTO) (*entity.Customer, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomer")
	}

	var r0 *entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.CustomerCreateDTO) (*entity.Customer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.CustomerCreateDTO) *entity.Customer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.CustomerCreateDTO) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCustomerUseCase_CreateCustomer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomer'
type MockCustomerUseCase_CreateCustomer_Call struct {
	*mock.Call
}

// CreateCustomer is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 dto.CustomerCreateDTO
func (_e *MockCustomerUseCase_Expecter) CreateCustomer(_a0 interface{}, _a1 interface{}) *MockCustomerUseCase_CreateCustomer_Call {
	return &MockCustomerUseCase_CreateCustomer_Call{Call: _e.mock.On("CreateCustomer", _a0, _a1)}
}

func (_c *MockCustomerUseCase_CreateCustomer_Call) Run(run func(_a0 context.Context, _a1 dto.CustomerCreateDTO)) *MockCustomerUseCase_CreateCustomer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dto.CustomerCreateDTO))
	})
	return _c
}

func (_c *MockCustomerUseCase_CreateCustomer_Call) Return(_a0 *entity.Customer, _a1 error) *MockCustomerUseCase_CreateCustomer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCustomerUseCase_CreateCustomer_Call) RunAndReturn(run func(context.Context, dto.CustomerCreateDTO) (*entity.Customer, error)) *MockCustomerUseCase_CreateCustomer_Call {
	_c.Call.Return(run)
	return _c
}

// DisableCustomer provides a mock function with given fields: ctx, id
func (_m *MockCustomerUseCase) DisableCustomer(ctx context.Context, id string) (*entity.Customer, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DisableCustomer")
	}

	var r0 *entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Customer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Customer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCustomerUseCase_DisableCustomer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisableCustomer'
type MockCustomerUseCase_DisableCustomer_Call struct {
	*mock.Call
}

// DisableCustomer is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockCustomerUseCase_Expecter) DisableCustomer(ctx interface{}, id interface{}) *MockCustomerUseCase_DisableCustomer_Call {
	return &MockCustomerUseCase_DisableCustomer_Call{Call: _e.mock.On("DisableCustomer", ctx, id)}
}

func (_c *MockCustomerUseCase_DisableCustomer_Call) Run(run func(ctx context.Context, id string)) *MockCustomerUseCase_DisableCustomer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCustomerUseCase_DisableCustomer_Call) Return(_a0 *entity.Customer, _a1 error) *MockCustomerUseCase_DisableCustomer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCustomerUseCase_DisableCustomer_Call) RunAndReturn(run func(context.Context, string) (*entity.Customer, error)) *MockCustomerUseCase_DisableCustomer_Call {
	_c.Call.Return(run)
	return _c
}

// GetCustomer provides a mock function with given fields: ctx, params
func (_m *MockCustomerUseCase) GetCustomer(ctx context.Context, params map[string]string) (*entity.Customer, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomer")
	}

	var r0 *entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string) (*entity.Customer, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string) *entity.Customer); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]string) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCustomerUseCase_GetCustomer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCustomer'
type MockCustomerUseCase_GetCustomer_Call struct {
	*mock.Call
}

// GetCustomer is a helper method to define mock.On call
//   - ctx context.Context
//   - params map[string]string
func (_e *MockCustomerUseCase_Expecter) GetCustomer(ctx interface{}, params interface{}) *MockCustomerUseCase_GetCustomer_Call {
	return &MockCustomerUseCase_GetCustomer_Call{Call: _e.mock.On("GetCustomer", ctx, params)}
}

func (_c *MockCustomerUseCase_GetCustomer_Call) Run(run func(ctx context.Context, params map[string]string)) *MockCustomerUseCase_GetCustomer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]string))
	})
	return _c
}

func (_c *MockCustomerUseCase_GetCustomer_Call) Return(_a0 *entity.Customer, _a1 error) *MockCustomerUseCase_GetCustomer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCustomerUseCase_GetCustomer_Call) RunAndReturn(run func(context.Context, map[string]string) (*entity.Customer, error)) *MockCustomerUseCase_GetCustomer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCustomerUseCase creates a new instance of MockCustomerUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCustomerUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCustomerUseCase {
	mock := &MockCustomerUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
