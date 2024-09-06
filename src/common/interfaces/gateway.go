package interfaces

import (
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
)

type PaymentGateway interface {
	RequestSyncronousPayment(order entity.Order) (dto.CreateCheckout, error)
	RequestAssyncronousPayment(order entity.Order) (dto.CreateCheckout, error)
}
