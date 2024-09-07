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
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
)

type paymentGateway struct {
	httpClient         http.Client
	checkoutServiceURL string
	sqsClient          *sqs.Client
	queueURL           string
}

type PaymentGatewayConfig struct {
	Timeout            time.Duration
	CheckoutServiceURL string
	SQSEndpoint        string
	SQSQueueURL        string
	AWSRegion          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
}

func NewPaymentGateway(gtConfig PaymentGatewayConfig) interfaces.PaymentGateway {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	var cfg aws.Config
	var err error

	if gtConfig.SQSEndpoint != "" {
		// Usando o LocalStack com um endpoint personalizado
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(gtConfig.AWSRegion),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					gtConfig.AWSAccessKeyID, gtConfig.AWSSecretAccessKey, "")),
			config.WithEndpointResolver(
				aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:           gtConfig.SQSEndpoint,
						SigningRegion: region,
					}, nil
				}),
			),
		)
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(gtConfig.AWSRegion),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					gtConfig.AWSAccessKeyID, gtConfig.AWSSecretAccessKey, "")),
		)
	}

	if err != nil {
		fmt.Printf("unable to load SDK config, %v", err)
	}

	sqsClient := sqs.NewFromConfig(cfg)

	return &paymentGateway{
		httpClient:         *client,
		checkoutServiceURL: gtConfig.CheckoutServiceURL,
		sqsClient:          sqsClient,
		queueURL:           gtConfig.SQSQueueURL,
	}
}

func (pg *paymentGateway) RequestSyncronousPayment(order entity.Order) (dto.CreateCheckout, error) {

	jsonData, err := json.Marshal(order)
	if err != nil {
		return dto.CreateCheckout{}, fmt.Errorf("error occurred while encoding order data: %s", err.Error())
	}

	url := fmt.Sprintf("%s/payments", pg.checkoutServiceURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return dto.CreateCheckout{}, fmt.Errorf("error occurred while creating request: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := pg.httpClient.Do(req)
	if err != nil {
		return dto.CreateCheckout{}, fmt.Errorf("error occurred while sending post request: %s", err.Error())
	}
	defer resp.Body.Close()

	var checkoutResponse dto.CreateCheckout
	err = json.NewDecoder(resp.Body).Decode(&checkoutResponse)
	if err != nil {
		return dto.CreateCheckout{}, fmt.Errorf("error occurred while decoding response data: %s", err.Error())
	}
	return checkoutResponse, nil
}

func (pg *paymentGateway) RequestAssyncronousPayment(order entity.Order) (dto.CreateCheckout, error) {

	event := entity.OrderEvent{
		Id:          uuid.New().String(),
		EventType:   "CreatedCheckout",
		OrderStatus: string(order.OrderStatus),
		Order:       &order,
	}

	jsonData, err := json.Marshal(event)
	if err != nil {
		return dto.CreateCheckout{}, fmt.Errorf("error occurred while encoding order data: %s", err.Error())
	}

	_, err = pg.sqsClient.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &pg.queueURL,
		MessageBody: aws.String(string(jsonData)),
	})
	if err != nil {
		return dto.CreateCheckout{}, fmt.Errorf("error occurred while sending message to SQS: %s", err.Error())
	}
	return dto.CreateCheckout{}, nil
}
