package services

import (
	"context"

	entities_pb "github.com/lk153/proto-tracking-gen/go/entities"

	"factory/exam/repo"
)

//ProductServiceInterface ...
type ProductServiceInterface interface {
	GetProducts(ctx context.Context, limit int) []*repo.ProductModel
	GetProduct(ctx context.Context, id int) *repo.ProductModel
	CreateProduct(ctx context.Context, data *entities_pb.ProductInfo) *repo.ProductModel
	Transform(input []*repo.ProductModel) []*entities_pb.ProductInfo
	TransformSingle(prod *repo.ProductModel) *entities_pb.ProductInfo
}
