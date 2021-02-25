package services

import (
	"factory/exam/repo"
)

//ProductServiceInterface ...
type ProductServiceInterface interface {
	GetProducts(limit uint32) []*repo.ProductModel
}
