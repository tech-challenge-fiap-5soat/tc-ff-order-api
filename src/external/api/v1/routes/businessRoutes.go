package routes

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/constants"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/usecase"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/external/api/infra/config"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/external/api/v1/handlers"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/operation/controller"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/operation/gateway"

	mongodb "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/external/datasource"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterBusinessRoutes(gServer *gin.RouterGroup, dbClient mongo.Client) {
	groupServer := gServer.Group("/v1")

	registerCustomerHandler(groupServer, dbClient)
	registerProductHandler(groupServer, dbClient)
	registerOrderHandler(groupServer, dbClient)
	registerCheckoutHandler(groupServer, dbClient)
}

func registerCustomerHandler(groupServer *gin.RouterGroup, dbClient mongo.Client) {
	mongoAdapter := mongodb.NewMongoAdapter[entity.Customer](
		dbClient,
		config.GetMongoCfg().Database,
		constants.CustomerCollection,
	)

	customerInteractor := controller.NewCustomerController(mongoAdapter)
	handlers.NewCustomerHandler(groupServer, customerInteractor)
}

func registerProductHandler(groupServer *gin.RouterGroup, dbClient mongo.Client) {
	mongoAdapter := mongodb.NewMongoAdapter[entity.Product](
		dbClient,
		config.GetMongoCfg().Database,
		constants.ProductCollection,
	)

	productInteractor := controller.NewProductController(mongoAdapter)
	handlers.NewProductHandler(groupServer, productInteractor)
}

func registerOrderHandler(groupServer *gin.RouterGroup, dbClient mongo.Client) {
	orderDbAdapter := mongodb.NewMongoAdapter[entity.Order](
		dbClient,
		config.GetMongoCfg().Database,
		constants.OrderCollection,
	)

	productDbAdapter := mongodb.NewMongoAdapter[entity.Product](
		dbClient,
		config.GetMongoCfg().Database,
		constants.ProductCollection,
	)
	productGateway := gateway.NewProductGateway(productDbAdapter)
	productUseCase := usecase.NewProductUseCase(productGateway)

	customerDbAdapter := mongodb.NewMongoAdapter[entity.Customer](
		dbClient,
		config.GetMongoCfg().Database,
		constants.CustomerCollection,
	)
	customerGateway := gateway.NewCustomerGateway(customerDbAdapter)
	customerUseCase := usecase.NewCustomerUseCase(customerGateway)

	kitchenService := gateway.NewKitchenService(gateway.KitchenServiceConfig{
		Timeout:               5,
		KitchenServiceBaseUrl: config.GetApiCfg().KitchenServiceURL,
		SQSEndpoint:           config.GetQueueProcessorsCfg().OrderPreparationEventsQueueEndpoint,
		SQSQueueURL:           config.GetQueueProcessorsCfg().OrderPreparationEventsQueue,
		AWSRegion:             config.GetQueueProcessorsCfg().OrderPreparationEventsQueueRegion,
		AWSAccessKeyID:        os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey:    os.Getenv("AWS_SECRET_ACCESS_KEY"),
	})
	orderInteractor := controller.NewOrderController(orderDbAdapter, productUseCase, customerUseCase, kitchenService)

	handlers.NewOrderHandler(groupServer, orderInteractor)
}

func registerCheckoutHandler(groupServer *gin.RouterGroup, dbClient mongo.Client) {

	productDbAdapter := mongodb.NewMongoAdapter[entity.Product](
		dbClient,
		config.GetMongoCfg().Database,
		constants.ProductCollection,
	)
	productGateway := gateway.NewProductGateway(productDbAdapter)
	productUseCase := usecase.NewProductUseCase(productGateway)

	customerDbAdapter := mongodb.NewMongoAdapter[entity.Customer](
		dbClient,
		config.GetMongoCfg().Database,
		constants.CustomerCollection,
	)
	customerGateway := gateway.NewCustomerGateway(customerDbAdapter)
	customerUseCase := usecase.NewCustomerUseCase(customerGateway)

	orderDbAdapter := mongodb.NewMongoAdapter[entity.Order](
		dbClient,
		config.GetMongoCfg().Database,
		constants.OrderCollection,
	)

	kitchenService := gateway.NewKitchenService(gateway.KitchenServiceConfig{
		Timeout:               5,
		KitchenServiceBaseUrl: config.GetApiCfg().KitchenServiceURL,
		SQSEndpoint:           config.GetQueueProcessorsCfg().OrderPreparationEventsQueueEndpoint,
		SQSQueueURL:           config.GetQueueProcessorsCfg().OrderPreparationEventsQueue,
		AWSRegion:             config.GetQueueProcessorsCfg().OrderPreparationEventsQueueRegion,
		AWSAccessKeyID:        os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey:    os.Getenv("AWS_SECRET_ACCESS_KEY"),
	})

	orderGateway := gateway.NewOrderGateway(orderDbAdapter, kitchenService)
	orderUseCase := usecase.NewOrderUseCase(orderGateway, productUseCase, customerUseCase)

	paymentGateway := gateway.NewPaymentGateway(gateway.PaymentGatewayConfig{
		Timeout:            5,
		CheckoutServiceURL: config.GetApiCfg().CheckoutServiceURL,
		SQSEndpoint:        config.GetQueueProcessorsCfg().OrderEventsQueueEndpoint,
		SQSQueueURL:        config.GetQueueProcessorsCfg().OrderEventsQueue,
		AWSRegion:          config.GetQueueProcessorsCfg().OrderEventsQueueRegion,
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	})
	checkoutInteractor := controller.NewCheckoutController(orderUseCase, paymentGateway)

	handlers.NewCheckoutHandler(groupServer, checkoutInteractor)
}
