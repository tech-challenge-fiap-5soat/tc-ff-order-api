package controller

import (
	"context"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/usecase"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/operation/gateway"
)

type CustomerController struct {
	useCase interfaces.CustomerUseCase
}

func NewCustomerController(datasource interfaces.DatabaseSource) interfaces.CustomerController {
	gateway := gateway.NewCustomerGateway(datasource)
	return &CustomerController{
		useCase: usecase.NewCustomerUseCase(gateway),
	}
}

func (cc *CustomerController) CreateCustomer(ctx context.Context,
	customerRequest dto.CustomerCreateDTO) (*entity.Customer, error) {
	return cc.useCase.CreateCustomer(ctx, customerRequest)
}

func (cc *CustomerController) GetCustomer(ctx context.Context, params map[string]string) (*entity.Customer, error) {
	return cc.useCase.GetCustomer(ctx, params)
}
