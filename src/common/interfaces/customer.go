package interfaces

import (
	"context"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

type CustomerUseCase interface {
	CreateCustomer(context.Context, dto.CustomerCreateDTO) (*entity.Customer, error)
	GetCustomer(ctx context.Context, params map[string]string) (*entity.Customer, error)
	DisableCustomer(ctx context.Context, id string) (*entity.Customer, error)
}

type CustomerGateway interface {
	Find(cpf valueobject.CPF) (*entity.Customer, error)
	Save(customer *entity.Customer) error
}

type CustomerController interface {
	CreateCustomer(ctx context.Context,
		customerRequest dto.CustomerCreateDTO) (*entity.Customer, error)
	GetCustomer(ctx context.Context, params map[string]string) (*entity.Customer, error)
	DisableCustomer(ctx context.Context, id string) (*entity.Customer, error)
}
