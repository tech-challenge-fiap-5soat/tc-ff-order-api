// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	entity "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"

	mock "github.com/stretchr/testify/mock"

	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

// MockProductUseCase is an autogenerated mock type for the ProductUseCase type
type MockProductUseCase struct {
	mock.Mock
}

type MockProductUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProductUseCase) EXPECT() *MockProductUseCase_Expecter {
	return &MockProductUseCase_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: product
func (_m *MockProductUseCase) Create(product *entity.Product) error {
	ret := _m.Called(product)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProductUseCase_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockProductUseCase_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - product *entity.Product
func (_e *MockProductUseCase_Expecter) Create(product interface{}) *MockProductUseCase_Create_Call {
	return &MockProductUseCase_Create_Call{Call: _e.mock.On("Create", product)}
}

func (_c *MockProductUseCase_Create_Call) Run(run func(product *entity.Product)) *MockProductUseCase_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Product))
	})
	return _c
}

func (_c *MockProductUseCase_Create_Call) Return(_a0 error) *MockProductUseCase_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProductUseCase_Create_Call) RunAndReturn(run func(*entity.Product) error) *MockProductUseCase_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: productId
func (_m *MockProductUseCase) Delete(productId string) error {
	ret := _m.Called(productId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(productId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProductUseCase_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockProductUseCase_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - productId string
func (_e *MockProductUseCase_Expecter) Delete(productId interface{}) *MockProductUseCase_Delete_Call {
	return &MockProductUseCase_Delete_Call{Call: _e.mock.On("Delete", productId)}
}

func (_c *MockProductUseCase_Delete_Call) Run(run func(productId string)) *MockProductUseCase_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockProductUseCase_Delete_Call) Return(_a0 error) *MockProductUseCase_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProductUseCase_Delete_Call) RunAndReturn(run func(string) error) *MockProductUseCase_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: id
func (_m *MockProductUseCase) FindById(id string) (*entity.Product, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Product, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProductUseCase_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type MockProductUseCase_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - id string
func (_e *MockProductUseCase_Expecter) FindById(id interface{}) *MockProductUseCase_FindById_Call {
	return &MockProductUseCase_FindById_Call{Call: _e.mock.On("FindById", id)}
}

func (_c *MockProductUseCase_FindById_Call) Run(run func(id string)) *MockProductUseCase_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockProductUseCase_FindById_Call) Return(_a0 *entity.Product, _a1 error) *MockProductUseCase_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProductUseCase_FindById_Call) RunAndReturn(run func(string) (*entity.Product, error)) *MockProductUseCase_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *MockProductUseCase) GetAll() ([]entity.Product, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.Product, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProductUseCase_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockProductUseCase_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *MockProductUseCase_Expecter) GetAll() *MockProductUseCase_GetAll_Call {
	return &MockProductUseCase_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *MockProductUseCase_GetAll_Call) Run(run func()) *MockProductUseCase_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProductUseCase_GetAll_Call) Return(_a0 []entity.Product, _a1 error) *MockProductUseCase_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProductUseCase_GetAll_Call) RunAndReturn(run func() ([]entity.Product, error)) *MockProductUseCase_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByCategory provides a mock function with given fields: category
func (_m *MockProductUseCase) GetByCategory(category valueobject.Category) ([]entity.Product, error) {
	ret := _m.Called(category)

	if len(ret) == 0 {
		panic("no return value specified for GetByCategory")
	}

	var r0 []entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(valueobject.Category) ([]entity.Product, error)); ok {
		return rf(category)
	}
	if rf, ok := ret.Get(0).(func(valueobject.Category) []entity.Product); ok {
		r0 = rf(category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(valueobject.Category) error); ok {
		r1 = rf(category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProductUseCase_GetByCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByCategory'
type MockProductUseCase_GetByCategory_Call struct {
	*mock.Call
}

// GetByCategory is a helper method to define mock.On call
//   - category valueobject.Category
func (_e *MockProductUseCase_Expecter) GetByCategory(category interface{}) *MockProductUseCase_GetByCategory_Call {
	return &MockProductUseCase_GetByCategory_Call{Call: _e.mock.On("GetByCategory", category)}
}

func (_c *MockProductUseCase_GetByCategory_Call) Run(run func(category valueobject.Category)) *MockProductUseCase_GetByCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(valueobject.Category))
	})
	return _c
}

func (_c *MockProductUseCase_GetByCategory_Call) Return(_a0 []entity.Product, _a1 error) *MockProductUseCase_GetByCategory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProductUseCase_GetByCategory_Call) RunAndReturn(run func(valueobject.Category) ([]entity.Product, error)) *MockProductUseCase_GetByCategory_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: productId, product
func (_m *MockProductUseCase) Update(productId string, product *entity.Product) error {
	ret := _m.Called(productId, product)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *entity.Product) error); ok {
		r0 = rf(productId, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockProductUseCase_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockProductUseCase_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - productId string
//   - product *entity.Product
func (_e *MockProductUseCase_Expecter) Update(productId interface{}, product interface{}) *MockProductUseCase_Update_Call {
	return &MockProductUseCase_Update_Call{Call: _e.mock.On("Update", productId, product)}
}

func (_c *MockProductUseCase_Update_Call) Run(run func(productId string, product *entity.Product)) *MockProductUseCase_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*entity.Product))
	})
	return _c
}

func (_c *MockProductUseCase_Update_Call) Return(_a0 error) *MockProductUseCase_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProductUseCase_Update_Call) RunAndReturn(run func(string, *entity.Product) error) *MockProductUseCase_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProductUseCase creates a new instance of MockProductUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProductUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProductUseCase {
	mock := &MockProductUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
