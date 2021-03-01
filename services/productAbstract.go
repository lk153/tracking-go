package services

import (
	"context"
	"factory/exam/repo"
)

//ProductServiceInterface ...
type ProductServiceInterface interface {
	GetProducts(ctx context.Context, limit uint32) []*repo.ProductModel
}
