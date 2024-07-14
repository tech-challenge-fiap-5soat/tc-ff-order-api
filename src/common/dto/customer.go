package dto

import "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"

type CustomerCreateDTO struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
	Cpf   string `json:"cpf" binding:"required,IsCpfValid" `
}

func CustomerEntityToSaveRecordDTO(customer *entity.Customer) map[string]interface{} {
	return map[string]interface{}{
		"_id":   customer.CPF,
		"name":  customer.Name,
		"email": customer.Email,
	}
}
