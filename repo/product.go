package repo

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"

	"github.com/bxcodec/faker/v3"
)

var _ ProductRepoInterface = &ProductRepoImp{}

//ProductRepoImp ...
type ProductRepoImp struct {
	tracer trace.Tracer
}

//ProductRepoProvider ...
func ProductRepoProvider() *ProductRepoImp {
	tracer := otel.Tracer("ProductRepoProvider")
	return &ProductRepoImp{
		tracer: tracer,
	}
}

//GetProduct ...
func (p *ProductRepoImp) GetProduct(ctx context.Context, limit int) (productDAO []*ProductModel, err error) {
	_, span := p.tracer.Start(ctx, "GetProduct")
	defer span.End()

	for i := 0; i < limit; i++ {
		product := &ProductModel{}
		span.SetAttributes(label.KeyValue{
			Key:   label.Key("name"),
			Value: label.StringValue("Viet Nguyen"),
		})

		span.AddEvent("faker.FakeData")
		err = faker.FakeData(&product)
		span.AddEvent("end faker.FakeData")

		if err != nil {
			fmt.Println(err)
		}

		productDAO = append(productDAO, product)
	}

	return productDAO, err
}

//GetProduct ...
func (p *ProductRepoImp) Find(ctx context.Context, id int) (productDAO *ProductModel, err error) {
	_, span := p.tracer.Start(ctx, "Find")
	defer span.End()

	product := &ProductModel{}
	span.SetAttributes(label.KeyValue{
		Key:   label.Key("name"),
		Value: label.StringValue("Viet Nguyen"),
	})

	span.AddEvent("faker.FakeData")
	err = faker.FakeData(&product)
	span.AddEvent("end faker.FakeData")

	if err != nil {
		fmt.Println(err)
	}

	return product, err
}
