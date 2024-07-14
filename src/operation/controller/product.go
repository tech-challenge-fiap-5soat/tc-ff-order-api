package controller

import (
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/entity"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/usecase"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/core/valueObject"
	"github.com/tech-challenge-fiap-5soat/tc-ff-order-api/src/operation/gateway"
)

type ProductController struct {
	useCase interfaces.ProductUseCase
}

func NewProductController(datasource interfaces.DatabaseSource) interfaces.ProductController {
	gateway := gateway.NewProductGateway(datasource)
	return &ProductController{
		useCase: usecase.NewProductUseCase(gateway),
	}
}

func (pc *ProductController) GetAll() ([]entity.Product, error) {
	return pc.useCase.GetAll()
}

func (pc *ProductController) GetByCategory(category valueobject.Category) ([]entity.Product, error) {
	return pc.useCase.GetByCategory(category)
}

func (pc *ProductController) Create(product *entity.Product) error {
	return pc.useCase.Create(product)
}

func (pc *ProductController) Update(productId string, product *entity.Product) error {
	return pc.useCase.Update(productId, product)
}

func (pc *ProductController) Delete(productId string) error {
	return pc.useCase.Delete(productId)
}

func (pc *ProductController) FindById(id string) (*entity.Product, error) {
	return pc.useCase.FindById(id)
}
