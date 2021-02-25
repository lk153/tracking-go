package repo

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

var _ ProductRepoInterface = &ProductRepoImp{}

//ProductRepoImp ...
type ProductRepoImp struct {
}

//ProductRepoProvider ...
func ProductRepoProvider() *ProductRepoImp {
	return &ProductRepoImp{}
}

//GetProduct ...
func (p *ProductRepoImp) GetProduct() *ProductModel {
	product := &ProductModel{}
	err := faker.FakeData(&product)
	if err != nil {
		fmt.Println(err)
	}

	return product
}
