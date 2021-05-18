package handler

import (
	"context"
	"factory/exam/services"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	services_pb "github.com/lk153/proto-tracking-gen/go/services"
)

//ProductPBHandler ...
type ProductPBHandler struct {
	services_pb.UnimplementedProductServiceServer
	productService services.ProductServiceInterface
	tracer         trace.Tracer
}

//NewProductPBHandler ...
func NewProductPBHandler(
	productService services.ProductServiceInterface,
) *ProductPBHandler {
	tracer := otel.Tracer("ProductHandlerProvider")
	return &ProductPBHandler{
		tracer:         tracer,
		productService: productService,
	}
}

//Get ...
func (p *ProductPBHandler) Get(ctx context.Context,
	req *services_pb.ProductRequest) (*services_pb.ProductResponse, error) {
	ctx, span := p.tracer.Start(ctx, "Get")
	defer span.End()

	limit := req.GetLimit()
	products := p.productService.GetProducts(ctx, int(limit))
	data := p.productService.Transform(products)

	return &services_pb.ProductResponse{
		Data: data,
	}, nil
}

//GetSingle ...
func (p *ProductPBHandler) GetSingle(ctx context.Context,
	req *services_pb.SingleProductRequest) (*services_pb.SingleProductResponse, error) {
	ctx, span := p.tracer.Start(ctx, "GetSingle")
	defer span.End()

	ID := req.GetId()
	product := p.productService.GetProduct(ctx, int(ID))
	data := p.productService.TransformSingle(product)

	return &services_pb.SingleProductResponse{
		Data: data,
	}, nil
}
