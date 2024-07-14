package usecase

import (
	"context"
	"errors"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

var ErrCustomerInvalid = errors.New("customer has invalid attributes")
var ErrCustomerSearchParams = errors.New("invalid params to search customer")

type customerUseCase struct {
	gateway interfaces.CustomerGateway
}

func NewCustomerUseCase(gateway interfaces.CustomerGateway) interfaces.CustomerUseCase {
	return &customerUseCase{
		gateway: gateway,
	}
}

func (interactor *customerUseCase) CreateCustomer(ctx context.Context,
	customerRequest dto.CustomerCreateDTO) (*entity.Customer, error) {

	customerToCreate := entity.Customer{
		Name:  customerRequest.Name,
		Email: valueobject.Email(customerRequest.Email),
		CPF:   valueobject.CPF(customerRequest.Cpf),
	}

	if !customerToCreate.IsValid() {
		return nil, ErrCustomerInvalid
	}

	err := interactor.gateway.Save(&customerToCreate)
	return &customerToCreate, err
}

func (interactor *customerUseCase) GetCustomer(ctx context.Context, params map[string]string) (*entity.Customer, error) {
	param, exists := params["cpf"]
	if !exists {
		return &entity.Customer{}, ErrCustomerSearchParams
	}
	return interactor.gateway.Find(valueobject.CPF(param))
}
