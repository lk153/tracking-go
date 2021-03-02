package handler

import (
	"net/http"
	"strconv"

	"github.com/getsentry/sentry-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"factory/exam/services"
)

//ProductHandler ...
type ProductHandler struct {
	services services.ProductServiceInterface
	tracer   trace.Tracer
}

//ProductHandlerProvider ...
func ProductHandlerProvider(
	services services.ProductServiceInterface,
) *ProductHandler {
	tracer := otel.Tracer("ProductHandlerProvider")
	return &ProductHandler{
		tracer:   tracer,
		services: services,
	}
}

//Get ...
func (h *ProductHandler) Get(resp http.ResponseWriter, req *http.Request) {
	ctx, span := h.tracer.Start(req.Context(), "Get")
	defer span.End()

	span.AddEvent("req.Parse.Limit")
	limitString := req.URL.Query().Get("limit")
	limit, err := strconv.ParseUint(limitString, 10, 32)
	if err != nil {
		sentry.CaptureException(err)
		responseError(resp, req, "Parse error")
		return
	}
	span.AddEvent("end req.Parse.Limit")

	products := h.services.GetProducts(ctx, uint32(limit))
	responseSuccess(resp, req, products)
}
