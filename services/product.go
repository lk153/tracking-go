package services

import (
	"factory/exam/repo"
)

var _ ProductServiceInterface = &ProductService{}

//ProductProvider ...
func ProductProvider(
	productRepo repo.ProductRepoInterface,
) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

//ProductService ...
type ProductService struct {
	productRepo repo.ProductRepoInterface
}

//GetProducts ...
func (ps *ProductService) GetProducts(limit uint32) []*repo.ProductModel {
	var products []*repo.ProductModel
	for i := uint32(0); i < limit; i++ {
		prod := ps.productRepo.GetProduct()
		products = append(products, prod)
	}

	return products
}
