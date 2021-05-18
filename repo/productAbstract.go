package repo

import "context"

//ProductRepoInterface ...
type ProductRepoInterface interface {
	GetProduct(context context.Context, limit int) (productDAO []*ProductModel, err error)
	Find(context context.Context, id int) (productDAO *ProductModel, err error)
}
