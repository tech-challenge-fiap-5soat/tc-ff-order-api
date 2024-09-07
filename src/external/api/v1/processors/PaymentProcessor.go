package processors

import (
	"fmt"

	"github.com/inaciogu/go-sqs/consumer"
	"github.com/inaciogu/go-sqs/consumer/message"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

type PaymentProcessorConfig struct {
	Endpoint        string
	QueueName       string
	Region          string
	WaitTimeSeconds int
}

type PaymentHandler struct {
	OrderUseCase interfaces.OrderUseCase
}

func PaymentProcessor(config PaymentProcessorConfig, orderUseCase interfaces.OrderUseCase) {

	handler := PaymentHandler{
		OrderUseCase: orderUseCase,
	}

	clientOptions := consumer.SQSClientOptions{
		Region:          config.Region,
		QueueName:       config.QueueName,
		Handle:          handler.paymentEventHandler,
		WaitTimeSeconds: 30,
	}
	if config.Endpoint != "" {
		clientOptions.Endpoint = config.Endpoint
	}
	go consumer.New(nil, clientOptions).Start()
}

func (pg PaymentHandler) paymentEventHandler(message *message.Message) bool {
	paymentEvent := entity.PaymentEvent{}

	err := message.Unmarshal(&paymentEvent)
	if err != nil {
		fmt.Println(err)
		return false
	}

	pg.OrderUseCase.UpdateOrderStatus(paymentEvent.Order.Id, valueobject.OrderStatus(paymentEvent.EventType))

	fmt.Print("Payment Event Received: ", paymentEvent)
	fmt.Print("Payment Event: ", paymentEvent.EventType)

	return true
}
