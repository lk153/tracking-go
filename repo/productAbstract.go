package repo

import "context"

//ProductRepoInterface ...
type ProductRepoInterface interface {
	GetProduct(context context.Context, limit int) (productDAO []*ProductModel, err error)
}
