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
func (p *ProductRepoImp) GetProduct(ctx context.Context) *ProductModel {
	ctx, span := p.tracer.Start(ctx, "GetProduct")
	defer span.End()

	product := &ProductModel{}

	span.SetAttributes(label.KeyValue{label.Key("name"), label.StringValue("Viet Nguyen")})
	span.AddEvent("faker.FakeData")
	err := faker.FakeData(&product)
	span.AddEvent("end faker.FakeData")

	if err != nil {
		fmt.Println(err)
	}

	return product
}
