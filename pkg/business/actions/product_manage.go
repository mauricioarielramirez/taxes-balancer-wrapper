package actions

import (
	"context"
	"taxes-balancer-wrapper/pkg/business/models"
	"taxes-balancer-wrapper/pkg/gateway"
	"taxes-balancer-wrapper/pkg/infraestructure/repository"
)

type ProductManager interface {
	Storage(ctx context.Context, product models.Product) error
	GetAll(ctx context.Context) ([]models.Product, error)
}

type ProductManagerAction struct {
	ProductRepository gateway.ProductRepository
}

func NewProductManager() ProductManagerAction {
	return NewProductManagerWithDep(repository.NewInMemoryProductRepository())
}

func NewProductManagerWithDep(repository gateway.ProductRepository) ProductManagerAction {
	return ProductManagerAction{
		ProductRepository: repository,
	}
}

func (pm ProductManagerAction) Storage(ctx context.Context, product models.Product) error {

	err := pm.ProductRepository.Save(ctx, product)

	if err != nil {
		return err
	}

	return nil
}

func (pm ProductManagerAction) GetAll(ctx context.Context) ([]models.Product, error) {

	allProducts, err := pm.ProductRepository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return allProducts, nil
}
