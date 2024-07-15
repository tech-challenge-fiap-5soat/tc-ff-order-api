package gateway

import (
	"fmt"

	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

type orderGateway struct {
	datasource     interfaces.DatabaseSource
	kitchenService interfaces.KitchenService
}

func NewOrderGateway(datasource interfaces.DatabaseSource,
	kitchenService interfaces.KitchenService) interfaces.OrderGateway {
	return &orderGateway{datasource: datasource, kitchenService: kitchenService}
}

func (og *orderGateway) FindAll() ([]entity.Order, error) {
	orders, err := og.datasource.FindAll("", "")

	if err != nil {
		return nil, err
	}

	foundOrders := []entity.Order{}

	for _, order := range orders {
		foundOrders = append(foundOrders, order.(entity.Order))
	}

	return foundOrders, nil
}

func (og *orderGateway) FindById(id string) (*entity.Order, error) {
	order, err := og.datasource.FindOne("_id", id)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, nil
	}

	foundOrder := order.(*entity.Order)
	return foundOrder, nil
}

func (og *orderGateway) FindAllByStatus(status valueobject.OrderStatus) ([]entity.Order, error) {
	orders, err := og.datasource.FindAll("orderStatus", string(status))

	if err != nil {
		return nil, err
	}

	foundOrders := []entity.Order{}

	for _, order := range orders {
		foundOrders = append(foundOrders, order.(entity.Order))
	}

	return foundOrders, nil
}

func (og *orderGateway) Save(order *entity.Order) (string, error) {
	insertResult, err := og.datasource.Save(
		dto.OrderEntityToSaveRecordDTO(order),
	)

	if err != nil {
		return "", err
	}

	fmt.Println(insertResult)

	orderInserted := insertResult.(string)
	return orderInserted, nil
}

func (og *orderGateway) Update(order *entity.Order) error {
	_, err := og.datasource.Update(
		order.ID,
		dto.OrderEntityToUpdateRecordDTO(order),
	)

	if err != nil {
		return err
	}
	return nil
}

func (og *orderGateway) RequetOrderPreparation(order *entity.Order) error {
	err := og.kitchenService.RequetOrderPreparation(*order)
	if err != nil {
		return err
	}
	return nil
}
