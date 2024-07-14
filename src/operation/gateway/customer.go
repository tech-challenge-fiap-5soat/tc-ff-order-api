package gateway

import (
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

type customerGateway struct {
	datasource interfaces.DatabaseSource
}

func NewCustomerGateway(datasource interfaces.DatabaseSource) interfaces.CustomerGateway {
	return &customerGateway{datasource: datasource}
}

func (cg *customerGateway) Find(cpf valueobject.CPF) (*entity.Customer, error) {
	customer, err := cg.datasource.FindOne("_id", string(cpf))

	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, nil
	}

	found := customer.(*entity.Customer)
	composed := entity.Customer{
		CPF:   cpf,
		Name:  found.Name,
		Email: found.Email,
	}
	return &composed, nil
}

func (cg *customerGateway) Save(customer *entity.Customer) error {
	_, err := cg.datasource.Save(
		dto.CustomerEntityToSaveRecordDTO(customer),
	)

	if err != nil {
		return err
	}

	return nil
}
