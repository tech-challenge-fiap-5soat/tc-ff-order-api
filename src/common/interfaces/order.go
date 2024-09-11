package interfaces

import (
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
)

type OrderUseCase interface {
	FindAll() ([]entity.Order, error)
	FindById(id string) (*entity.Order, error)
	GetAllByStatus(status valueobject.OrderStatus) ([]entity.Order, error)
	CreateOrder(order dto.OrderCreateDTO) (string, error)
	UpdateOrder(orderId string, order dto.OrderUpdateDTO) error
	UpdateOrderStatus(orderId string, status valueobject.OrderStatus) error
	RequetOrderPreparation(order *entity.Order) error
	RequestOrderCancellation(order *entity.Order) error
}

type OrderGateway interface {
	FindAll() ([]entity.Order, error)
	FindById(id string) (*entity.Order, error)
	FindAllByStatus(status valueobject.OrderStatus) ([]entity.Order, error)
	Save(order *entity.Order) (string, error)
	Update(order *entity.Order) error
	RequetOrderPreparation(order *entity.Order) error
}

type OrderController interface {
	FindAll() ([]entity.Order, error)
	FindById(id string) (*entity.Order, error)
	GetAllByStatus(status valueobject.OrderStatus) ([]entity.Order, error)
	CreateOrder(order dto.OrderCreateDTO) (string, error)
	UpdateOrder(orderId string, order dto.OrderUpdateDTO) error
	UpdateOrderStatus(orderId string, status valueobject.OrderStatus) error
}
