package usecase

import (
	"errors"
	"testing"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/usecase"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	coreErrors "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/errors"
	orderStatus "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"

	"github.com/stretchr/testify/assert"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/test/mocks"
)

func TestCheckoutUseCase(t *testing.T) {
	t.Parallel()

	productGatewayMock := mocks.NewMockProductGateway(t)
	productUseCase := usecase.NewProductUseCase(productGatewayMock)

	customerGatewayMock := mocks.NewMockCustomerGateway(t)
	customerUseCase := usecase.NewCustomerUseCase(customerGatewayMock)

	paymentGatewayMock := mocks.NewMockPaymentGateway(t)

	t.Run("should create checkout with success", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_STARTED}
		updatedOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}

		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock, productUseCase, customerUseCase)

		orderGatewayMock.On("FindById", orderId).Return(existentOrder, nil)
		orderGatewayMock.On("Update", updatedOrder).Return(nil)

		spyOrder := entity.Order{
			ID:          "order-123",
			Customer:    entity.Customer{Name: "", Email: "", CPF: ""},
			OrderStatus: "STARTED", OrderItems: []entity.OrderItem(nil),
			CreatedAt: orderStatus.CustomTime{},
			UpdatedAt: orderStatus.CustomTime{},
			Amount:    0,
		}

		paymentGatewayMock.On("RequestPayment", spyOrder).Return(
			dto.CreateCheckout{CheckoutURL: "https://fake-checkout.com/payment/" + orderId}, nil,
		)

		useCase := usecase.NewCheckoutUseCase(orderUseCase, paymentGatewayMock)

		createdCheckout, err := useCase.CreateCheckout(orderId)

		assert.Nil(t, err)
		assert.Equal(t, createdCheckout.Message, "Payment request successfully created")
		assert.Equal(t, createdCheckout.CheckoutURL, "https://fake-checkout.com/payment/"+orderId)
	})

	t.Run("should not create checkout When order status is diff of STARTED", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}

		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock, productUseCase, customerUseCase)
		orderGatewayMock.On("FindById", orderId).Return(existentOrder, nil)

		useCase := usecase.NewCheckoutUseCase(orderUseCase, paymentGatewayMock)

		createdCheckout, err := useCase.CreateCheckout(orderId)

		assert.Nil(t, err)
		assert.Equal(t, createdCheckout.Message, coreErrors.ErrCheckoutOrderAlreadyCompleted.Error())
		assert.Equal(t, createdCheckout.CheckoutURL, "")
	})

	t.Run("should not create checkout When some error occurs during update operation", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_STARTED}
		updatedOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}

		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock, productUseCase, customerUseCase)

		orderGatewayMock.On("FindById", orderId).Return(existentOrder, nil)
		orderGatewayMock.On("Update", updatedOrder).Return(errors.New("error updating order status"))
		spyOrder := entity.Order{
			ID:          "order-123",
			Customer:    entity.Customer{Name: "", Email: "", CPF: ""},
			OrderStatus: "STARTED", OrderItems: []entity.OrderItem(nil),
			CreatedAt: orderStatus.CustomTime{},
			UpdatedAt: orderStatus.CustomTime{},
			Amount:    0,
		}

		paymentGatewayMock.On("RequestPayment", spyOrder).Return(
			dto.CreateCheckout{CheckoutURL: "https://fake-checkout.com/payment/" + orderId}, nil,
		)

		useCase := usecase.NewCheckoutUseCase(orderUseCase, paymentGatewayMock)

		_, err := useCase.CreateCheckout(orderId)

		assert.NotNil(t, err)
	})

	t.Run("should update order status to PAYMENT_APPROVED with success", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}
		updatedOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_APPROVED}
		prepareOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_BEING_PREPARED}

		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock, productUseCase, customerUseCase)

		orderGatewayMock.On("FindById", orderId).Return(existentOrder, nil)
		orderGatewayMock.On("Update", updatedOrder).Return(nil)
		orderGatewayMock.On("RequetOrderPreparation", existentOrder).Return(nil)
		orderGatewayMock.On("Update", prepareOrder).Return(nil)

		useCase := usecase.NewCheckoutUseCase(orderUseCase, paymentGatewayMock)

		err := useCase.UpdateCheckout(orderId, orderStatus.ORDER_PAYMENT_APPROVED)

		assert.Nil(t, err)
	})

	t.Run("should update order status to PAYMENT_REFUSED with success", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}
		updatedOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_REFUSED}

		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock, productUseCase, customerUseCase)

		orderGatewayMock.On("FindById", orderId).Return(existentOrder, nil)
		orderGatewayMock.On("Update", updatedOrder).Return(nil)
		useCase := usecase.NewCheckoutUseCase(orderUseCase, paymentGatewayMock)

		err := useCase.UpdateCheckout(orderId, orderStatus.ORDER_PAYMENT_REFUSED)

		assert.Nil(t, err)
	})

	t.Run("should not update order status When new status is equal to the current", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_APPROVED}

		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock, productUseCase, customerUseCase)

		orderGatewayMock.On("FindById", orderId).Return(existentOrder, nil)
		useCase := usecase.NewCheckoutUseCase(orderUseCase, paymentGatewayMock)

		err := useCase.UpdateCheckout(orderId, orderStatus.ORDER_PAYMENT_APPROVED)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), coreErrors.ErrCheckoutOrderAlreadyCompleted.Error())
	})

	t.Run("should not update order status When new status is not valid next status", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_APPROVED}

		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock, productUseCase, customerUseCase)

		orderGatewayMock.On("FindById", orderId).Return(existentOrder, nil)
		useCase := usecase.NewCheckoutUseCase(orderUseCase, paymentGatewayMock)

		err := useCase.UpdateCheckout(orderId, orderStatus.ORDER_COMPLETED)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), coreErrors.ErrCheckoutOrderAlreadyCompleted.Error())
	})

	t.Run("should not update order status When some error occurrs during update operation", func(t *testing.T) {
		orderId := "order-123"
		existentOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_PENDING}
		updatedOrder := &entity.Order{ID: orderId, OrderStatus: orderStatus.ORDER_PAYMENT_APPROVED}

		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock, productUseCase, customerUseCase)

		orderGatewayMock.On("FindById", orderId).Return(existentOrder, nil)
		orderGatewayMock.On("Update", updatedOrder).Return(errors.New("error updating order status"))
		useCase := usecase.NewCheckoutUseCase(orderUseCase, paymentGatewayMock)

		err := useCase.UpdateCheckout(orderId, orderStatus.ORDER_PAYMENT_APPROVED)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error updating order status PAYMENT_PENDING to PAYMENT_APPROVED")
	})
}
