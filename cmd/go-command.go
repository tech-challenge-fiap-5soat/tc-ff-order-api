package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/docs"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/external/api"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/external/api/infra/config"
)

// @title Fast Food API
// @version 0.1.0
// @description Fast Food Order API for FIAP Tech course

// @host localhost:8080
// @BasePath /
func main() {

	mongoClient, err := config.InitMongoDbConfiguration(context.TODO())
	if err != nil {
		log.Fatal("error on create mongoConnection")
	}

	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	gServer := gin.New()
	api.Run(gServer, mongoClient)
}
