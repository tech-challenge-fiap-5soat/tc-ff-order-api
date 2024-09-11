package config

import (
	"os"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/constants"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/usecase"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/external/api/v1/processors"
	mongodb "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/external/datasource"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/operation/gateway"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitProcessors(dbClient mongo.Client) {

	productDbAdapter := mongodb.NewMongoAdapter[entity.Product](
		dbClient,
		GetMongoCfg().Database,
		constants.ProductCollection,
	)
	productGateway := gateway.NewProductGateway(productDbAdapter)
	productUseCase := usecase.NewProductUseCase(productGateway)

	customerDbAdapter := mongodb.NewMongoAdapter[entity.Customer](
		dbClient,
		GetMongoCfg().Database,
		constants.CustomerCollection,
	)
	customerGateway := gateway.NewCustomerGateway(customerDbAdapter)
	customerUseCase := usecase.NewCustomerUseCase(customerGateway)

	orderDbAdapter := mongodb.NewMongoAdapter[entity.Order](
		dbClient,
		GetMongoCfg().Database,
		constants.OrderCollection,
	)
	kitchenService := gateway.NewKitchenService(gateway.KitchenServiceConfig{
		Timeout:               5,
		KitchenServiceBaseUrl: GetApiCfg().KitchenServiceURL,
		SQSEndpoint:           GetQueueProcessorsCfg().OrderPreparationEventsQueueEndpoint,
		SQSQueueURL:           GetQueueProcessorsCfg().OrderPreparationEventsQueue,
		AWSRegion:             GetQueueProcessorsCfg().OrderPreparationEventsQueueRegion,
		AWSAccessKeyID:        os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey:    os.Getenv("AWS_SECRET_ACCESS_KEY"),
	})

	orderGateway := gateway.NewOrderGateway(orderDbAdapter, kitchenService)
	orderUseCase := usecase.NewOrderUseCase(orderGateway, productUseCase, customerUseCase)

	paymentProcessorConfig := processors.PaymentProcessorConfig{
		Endpoint:        GetQueueProcessorsCfg().CheckoutEventsQueueEndpoint,
		Region:          GetQueueProcessorsCfg().CheckoutEventsQueueRegion,
		QueueName:       GetQueueProcessorsCfg().CheckoutEventsQueue,
		WaitTimeSeconds: 30,
	}

	kitchenProcessorConfig := processors.KitchenProcessorConfig{
		Endpoint:        GetQueueProcessorsCfg().KitchenEventsQueueEndpoint,
		Region:          GetQueueProcessorsCfg().KitchenEventsQueueRegion,
		QueueName:       GetQueueProcessorsCfg().KitchenEventsQueue,
		WaitTimeSeconds: 30,
	}

	processors.PaymentProcessor(paymentProcessorConfig, orderUseCase)
	processors.KitchenProcessor(kitchenProcessorConfig, orderUseCase)
}
