package usecase

import (
	"context"
	"testing"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	. "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"

	"github.com/stretchr/testify/assert"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/usecase"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/test/mocks"
)

var customerGatewayMock *mocks.MockCustomerGateway

func TestCustomerUseCase(t *testing.T) {
	t.Parallel()

	t.Run("Should find a customer by CPF", func(t *testing.T) {
		cpf := CPF("12345678900")

		expected := entity.Customer{
			Name:  "John Doe",
			Email: "john@email.com",
			CPF:   cpf,
		}

		customerGatewayMock = mocks.NewMockCustomerGateway(t)
		customerGatewayMock.On("Find", cpf).Return(&expected, nil)

		useCase := usecase.NewCustomerUseCase(customerGatewayMock)

		params := map[string]string{
			"cpf": string(cpf),
		}
		customer, err := useCase.GetCustomer(context.TODO(), params)

		assert.Nil(t, err)
		assert.NotNil(t, customer)
		assert.Equal(t, customer.CPF, cpf)
	})

	t.Run("Should return error when search params was invalid", func(t *testing.T) {

		customerGatewayMock = mocks.NewMockCustomerGateway(t)
		useCase := usecase.NewCustomerUseCase(customerGatewayMock)

		params := map[string]string{
			"nome": "john",
		}
		_, err := useCase.GetCustomer(context.TODO(), params)

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, usecase.ErrCustomerSearchParams)
	})

	t.Run("Should return error when a customer has invalid attributes", func(t *testing.T) {
		createRequest := dto.CustomerCreateDTO{
			Name:  "John Doe",
			Email: "email.com",
			Cpf:   "111",
		}

		customerGatewayMock = mocks.NewMockCustomerGateway(t)

		useCase := usecase.NewCustomerUseCase(customerGatewayMock)

		_, err := useCase.CreateCustomer(context.TODO(), createRequest)

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, usecase.ErrCustomerInvalid)
	})

	t.Run("Should create customer successfully when has valid attributes", func(t *testing.T) {
		createRequest := dto.CustomerCreateDTO{
			Name:  "John Doe",
			Email: "john@email.com",
			Cpf:   "35679254077",
		}

		customerArg := entity.Customer{
			Name:  "John Doe",
			Email: "john@email.com",
			CPF:   CPF("35679254077"),
		}

		customerGatewayMock = mocks.NewMockCustomerGateway(t)
		customerGatewayMock.On("Save", &customerArg).Return(nil)

		useCase := usecase.NewCustomerUseCase(customerGatewayMock)
		result, err := useCase.CreateCustomer(context.TODO(), createRequest)

		assert.Nil(t, err)
		assert.Equal(t, result, &entity.Customer{
			Name:  "John Doe",
			Email: "john@email.com",
			CPF:   CPF("35679254077"),
		})
	})
}
