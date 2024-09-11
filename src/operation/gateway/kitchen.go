package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/uuid"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
)

type kitchenService struct {
	httpClient            http.Client
	kitchenServiceBaseUrl string
	sqsClient             *sqs.Client
	queueURL              string
}

type KitchenServiceConfig struct {
	Timeout               time.Duration
	KitchenServiceBaseUrl string
	SQSEndpoint           string
	SQSQueueURL           string
	AWSRegion             string
	AWSAccessKeyID        string
	AWSSecretAccessKey    string
}

func NewKitchenService(ksConfig KitchenServiceConfig) interfaces.KitchenService {

	var cfg aws.Config
	var err error

	if ksConfig.SQSEndpoint != "" {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(ksConfig.AWSRegion),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					ksConfig.AWSAccessKeyID, ksConfig.AWSSecretAccessKey, "")),
			config.WithEndpointResolver(
				aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:           ksConfig.SQSEndpoint,
						SigningRegion: region,
					}, nil
				}),
			),
		)
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(ksConfig.AWSRegion),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					ksConfig.AWSAccessKeyID, ksConfig.AWSSecretAccessKey, "")),
		)
	}

	if err != nil {
		fmt.Printf("unable to load SDK config, %v", err)
	}

	sqsClient := sqs.NewFromConfig(cfg)

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return &kitchenService{
		httpClient:            *client,
		kitchenServiceBaseUrl: ksConfig.KitchenServiceBaseUrl,
		sqsClient:             sqsClient,
		queueURL:              ksConfig.SQSQueueURL,
	}
}

func (ks *kitchenService) AssyncRequestOrderPreparation(order entity.Order) error {
	event := entity.OrderEvent{
		Id:          uuid.New().String(),
		EventType:   "RequestOrderPreparation",
		OrderStatus: string(order.OrderStatus),
		Order:       &order,
	}

	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("error occurred while encoding order data: %s", err.Error())
	}

	_, err = ks.sqsClient.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:     &ks.queueURL,
		MessageBody:  aws.String(string(jsonData)),
		DelaySeconds: 10,
	})

	if err != nil {
		return fmt.Errorf("error occurred while sending message to SQS: %s", err.Error())
	}
	return nil
}

func (ks *kitchenService) RequetOrderPreparation(order entity.Order) error {
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}
	requestBody := bytes.NewBuffer(orderJSON)
	request, err := http.NewRequest("POST", ks.kitchenServiceBaseUrl+"/prepare-order", requestBody)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	response, err := ks.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to request order preparation: %s", response.Status)
	}

	return nil
}
