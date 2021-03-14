package repo

import "context"

//ProductModel ...
type ProductModel struct {
	ID     uint64  `json:"id,omitempty" faker:"id"`
	Name   string  `json:"name,omitempty" faker:"name"`
	Price  float64 `json:"price,omitempty" faker:"oneof: 4.95, 9.99, 31997.97"`
	Type   string  `json:"type,omitempty" faker:"oneof: simple, virtual, group, configurable, bundle"`
	Status uint8   `json:"status,omitempty" faker:"oneof: 1, 2, 3, 4"`
}

//TableName ...
func (p *ProductModel) TableName() string {
	return "products"
}

//ProductRepoInterface ...
type ProductRepoInterface interface {
	GetProduct(context context.Context) (productDAO *ProductModel, err error)
}
