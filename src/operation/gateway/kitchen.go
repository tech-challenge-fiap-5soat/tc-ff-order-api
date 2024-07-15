package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
)

type kitchenService struct {
	httpClient            http.Client
	kitchenServiceBaseUrl string
}

type KitchenServiceConfig struct {
	Timeout               time.Duration
	KitchenServiceBaseUrl string
}

func NewKitchenService(config KitchenServiceConfig) interfaces.KitchenService {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return &kitchenService{
		httpClient:            *client,
		kitchenServiceBaseUrl: config.KitchenServiceBaseUrl,
	}
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
