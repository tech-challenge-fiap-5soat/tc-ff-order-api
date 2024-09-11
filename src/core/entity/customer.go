package entity

import (
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

type Customer struct {
	Name    string            `json:"name"`
	Email   valueobject.Email `json:"email"`
	CPF     valueobject.CPF   `json:"cpf"`
	Enabled bool
}

func (c *Customer) IsValid() bool {
	return len(c.Name) > 0 && c.Email.IsValid() && c.CPF.IsValid()
}
