package repo

//ProductModel ...
type ProductModel struct {
	Name   string  `faker:"name"`
	Price  float64 `faker:"oneof: 4.95, 9.99, 31997.97"`
	Type   string  `faker:"oneof: simple, virtual, group, configurable, bundle"`
	Status uint8   `faker:"oneof: 1, 2, 3, 4"`
}

//ProductRepoInterface ...
type ProductRepoInterface interface {
	GetProduct() *ProductModel
}
