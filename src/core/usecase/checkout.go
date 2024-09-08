package usecase

import (
	"fmt"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	coreErrors "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/errors"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	orderStatus "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

type checkoutUseCase struct {
	orderUseCase   interfaces.OrderUseCase
	paymentGateway interfaces.PaymentGateway
}

func NewCheckoutUseCase(
	orderUseCase interfaces.OrderUseCase,
	paymentGateway interfaces.PaymentGateway,
) interfaces.CheckoutUseCase {
	return &checkoutUseCase{
		orderUseCase:   orderUseCase,
		paymentGateway: paymentGateway,
	}
}

func (uc *checkoutUseCase) CreateCheckout(orderId string) (*dto.CreateCheckout, error) {
	order, err := uc.orderUseCase.FindById(orderId)
	nextStatus := orderStatus.ORDER_PAYMENT_PENDING

	if err != nil {
		return nil, err
	}

	if !order.OrderStatus.IsValidNextStatus(nextStatus.String()) {
		return &dto.CreateCheckout{
			CheckoutURL: "",
			Message:     coreErrors.ErrCheckoutOrderAlreadyCompleted.Error(),
		}, nil
	}

	requestedPayment, err := uc.paymentGateway.RequestAssyncronousPayment(*order)
	if err != nil {
		fmt.Println("error SQS: ", err)
		return nil, fmt.Errorf("error on request payment to orderId %s", order.ID)
	}

	return &dto.CreateCheckout{
		CheckoutURL: requestedPayment.CheckoutURL,
		Message:     "Payment request successfully created",
	}, nil
}

func (uc *checkoutUseCase) UpdateCheckout(orderId string, status orderStatus.OrderStatus) error {
	order, err := uc.orderUseCase.FindById(orderId)

	if err != nil {
		return err
	}

	if !order.OrderStatus.IsValidNextStatus(status.String()) {
		return coreErrors.ErrCheckoutOrderAlreadyCompleted
	}

	err = uc.orderUseCase.UpdateOrderStatus(orderId, status)

	if err != nil {
		return fmt.Errorf("error updating order status %s to %s", order.OrderStatus.String(), status.String())
	}

	return nil
}
