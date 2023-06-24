package gateway

import (
	"context"
	"taxes-balancer-wrapper/pkg/business/models"
)

type ProductRepository interface {
	Save(ctx context.Context, product models.Product) error
	GetById(ctx context.Context, id string) (*models.Product, error)
	GetAll(ctx context.Context) ([]models.Product, error)
}
