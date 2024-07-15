package interfaces

import "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"

type KitchenService interface {
	RequetOrderPreparation(order entity.Order) error
}
