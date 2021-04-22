package services

import (
	"context"

	entities_pb "github.com/lk153/proto-tracking-gen/go/entities"

	"factory/exam/repo"
)

//ProductServiceInterface ...
type ProductServiceInterface interface {
	GetProducts(ctx context.Context, limit int) []*repo.ProductModel
	Transform(input []*repo.ProductModel) []*entities_pb.ProductInfo
}
