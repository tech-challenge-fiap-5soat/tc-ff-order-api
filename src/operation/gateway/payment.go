package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
)

type paymentGateway struct {
	httpClient         http.Client
	checkoutServiceURL string
}

type PaymentGatewayConfig struct {
	Timeout            time.Duration
	CheckoutServiceURL string
}

func NewPaymentGateway(config PaymentGatewayConfig) interfaces.PaymentGateway {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return &paymentGateway{
		httpClient:         *client,
		checkoutServiceURL: config.CheckoutServiceURL,
	}
}

func (pg *paymentGateway) RequestPayment(order entity.Order) (dto.CreateCheckout, error) {

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
