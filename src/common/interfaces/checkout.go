package interfaces

import (
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

type CheckoutUseCase interface {
	CreateCheckout(orderId string) (*dto.CreateCheckout, error)
	UpdateCheckout(orderId string, status valueobject.OrderStatus) error
}

type CheckoutController interface {
	CreateCheckout(orderId string) (*dto.CreateCheckout, error)
	UpdateCheckout(orderId string, status valueobject.OrderStatus) error
}
