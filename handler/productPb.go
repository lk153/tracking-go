package handler

import (
	"context"
	"factory/exam/services"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	servicesPb "github.com/lk153/proto-tracking-gen/go/services"
)

//ProductPBHandler
type ProductPBHandler struct {
	servicesPb.UnimplementedProductServiceServer
	productService services.ProductServiceInterface
	tracer         trace.Tracer
}

//NewProductPBHandler
func NewProductPBHandler(
	productService services.ProductServiceInterface,
) *ProductPBHandler {
	tracer := otel.Tracer("NewProductPBHandler")
	return &ProductPBHandler{
		tracer:         tracer,
		productService: productService,
	}
}

//Get
func (p *ProductPBHandler) Get(ctx context.Context,
	req *servicesPb.ProductRequest) (*servicesPb.ProductResponse, error) {
	ctx, span := p.tracer.Start(ctx, "Get")
	defer span.End()

	limit := req.GetLimit()
	page := req.GetPage()
	ids := req.GetIds()
	products := p.productService.GetProducts(ctx, int(limit), int(page), ids)
	data := p.productService.Transform(products)

	return &servicesPb.ProductResponse{
		Data: data,
	}, nil
}

//GetSingle
func (p *ProductPBHandler) GetSingle(ctx context.Context,
	req *servicesPb.SingleProductRequest) (*servicesPb.SingleProductResponse, error) {
	ctx, span := p.tracer.Start(ctx, "GetSingle")
	defer span.End()

	ID := req.GetId()
	product := p.productService.GetProduct(ctx, int(ID))
	data := p.productService.TransformSingle(product)

	return &servicesPb.SingleProductResponse{
		Data: data,
	}, nil
}

//Create
func (p *ProductPBHandler) Create(ctx context.Context,
	req *servicesPb.ProductCreateRequest) (*servicesPb.ProductCreateResponse, error) {
	ctx, span := p.tracer.Start(ctx, "Create")
	defer span.End()

	data := req.GetData()
	product := p.productService.CreateProduct(ctx, data)
	data = p.productService.TransformSingle(product)

	return &servicesPb.ProductCreateResponse{
		Data: data,
	}, nil
}
