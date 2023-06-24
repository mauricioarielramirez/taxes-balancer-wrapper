package repository

import (
	"context"
	"taxes-balancer-wrapper/pkg/business/models"
)

var inMemoryProductStorage []models.Product

type InMemoryProductRepository struct{}

func NewInMemoryProductRepository() InMemoryProductRepository {
	return InMemoryProductRepository{}
}

func (pr InMemoryProductRepository) Save(ctx context.Context, product models.Product) error {
	inMemoryProductStorage = append(inMemoryProductStorage, product)
	return nil
}

func (pr InMemoryProductRepository) GetById(ctx context.Context, id string) (*models.Product, error) {
	return nil, nil
}

func (pr InMemoryProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	return inMemoryProductStorage, nil
}
