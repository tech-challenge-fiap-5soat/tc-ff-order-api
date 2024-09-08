package processors

import (
	"fmt"

	"github.com/inaciogu/go-sqs/consumer"
	"github.com/inaciogu/go-sqs/consumer/message"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

type KitchenProcessorConfig struct {
	Endpoint        string
	QueueName       string
	Region          string
	WaitTimeSeconds int
}

type KitchenHandler struct {
	OrderUseCase interfaces.OrderUseCase
}

func KitchenProcessor(config KitchenProcessorConfig, orderUseCase interfaces.OrderUseCase) {

	handler := KitchenHandler{
		OrderUseCase: orderUseCase,
	}

	clientOptions := consumer.SQSClientOptions{
		Region:          config.Region,
		QueueName:       config.QueueName,
		Handle:          handler.kitchenEventHandler,
		WaitTimeSeconds: 10,
	}
	if config.Endpoint != "" {
		clientOptions.Endpoint = config.Endpoint
	}
	go consumer.New(nil, clientOptions).Start()
}

func (pg KitchenHandler) kitchenEventHandler(message *message.Message) bool {
	orderEvent := entity.OrderEvent{}

	err := message.Unmarshal(&orderEvent)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if orderEvent.EventType == "READY_TO_TAKEOUT" {
		pg.OrderUseCase.UpdateOrderStatus(orderEvent.Order.ID, valueobject.OrderStatus("READY"))
		return true
	}

	if orderEvent.EventType == "COMPLETED" {
		pg.OrderUseCase.UpdateOrderStatus(orderEvent.Order.ID, valueobject.OrderStatus(orderEvent.EventType))
		return true
	}

	return true
}
