package dto

import (
	"github.com/google/uuid"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
)

func ProductEntityToSaveRecordDTO(product *entity.Product) map[string]interface{} {
	return map[string]interface{}{
		"_id":      uuid.New().String(),
		"name":     product.Name,
		"price":    product.Price,
		"category": product.Category,
	}
}

func ProductEntityToUpdateRecordDTO(product *entity.Product) map[string]interface{} {
	return map[string]interface{}{
		"name":     product.Name,
		"price":    product.Price,
		"category": product.Category,
	}
}
