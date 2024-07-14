package controller

import (
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/usecase"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

type CheckoutController struct {
	useCase interfaces.CheckoutUseCase
}

func NewCheckoutController(orderUseCase interfaces.OrderUseCase,
	paymentGateway interfaces.PaymentGateway) interfaces.CheckoutController {

	return &CheckoutController{
		useCase: usecase.NewCheckoutUseCase(orderUseCase, paymentGateway),
	}
}

func (cc *CheckoutController) CreateCheckout(orderId string) (*dto.CreateCheckout, error) {
	return cc.useCase.CreateCheckout(orderId)
}

func (cc *CheckoutController) UpdateCheckout(orderId string, status valueobject.OrderStatus) error {
	return cc.useCase.UpdateCheckout(orderId, status)
}
