// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	dto "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	entity "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockPaymentGateway is an autogenerated mock type for the PaymentGateway type
type MockPaymentGateway struct {
	mock.Mock
}

type MockPaymentGateway_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPaymentGateway) EXPECT() *MockPaymentGateway_Expecter {
	return &MockPaymentGateway_Expecter{mock: &_m.Mock}
}

// RequestPayment provides a mock function with given fields: order
func (_m *MockPaymentGateway) RequestPayment(order entity.Order) (dto.CreateCheckout, error) {
	ret := _m.Called(order)

	if len(ret) == 0 {
		panic("no return value specified for RequestPayment")
	}

	var r0 dto.CreateCheckout
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Order) (dto.CreateCheckout, error)); ok {
		return rf(order)
	}
	if rf, ok := ret.Get(0).(func(entity.Order) dto.CreateCheckout); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Get(0).(dto.CreateCheckout)
	}

	if rf, ok := ret.Get(1).(func(entity.Order) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPaymentGateway_RequestPayment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestPayment'
type MockPaymentGateway_RequestPayment_Call struct {
	*mock.Call
}

// RequestPayment is a helper method to define mock.On call
//   - order entity.Order
func (_e *MockPaymentGateway_Expecter) RequestPayment(order interface{}) *MockPaymentGateway_RequestPayment_Call {
	return &MockPaymentGateway_RequestPayment_Call{Call: _e.mock.On("RequestPayment", order)}
}

func (_c *MockPaymentGateway_RequestPayment_Call) Run(run func(order entity.Order)) *MockPaymentGateway_RequestPayment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entity.Order))
	})
	return _c
}

func (_c *MockPaymentGateway_RequestPayment_Call) Return(_a0 dto.CreateCheckout, _a1 error) *MockPaymentGateway_RequestPayment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaymentGateway_RequestPayment_Call) RunAndReturn(run func(entity.Order) (dto.CreateCheckout, error)) *MockPaymentGateway_RequestPayment_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPaymentGateway creates a new instance of MockPaymentGateway. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPaymentGateway(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPaymentGateway {
	mock := &MockPaymentGateway{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
