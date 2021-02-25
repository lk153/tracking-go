package handler

import (
	"net/http"
	"strconv"

	"github.com/getsentry/sentry-go"

	"factory/exam/services"
)

//ProductHandler ...
type ProductHandler struct {
	services services.ProductServiceInterface
}

//ProductHandlerProvider ...
func ProductHandlerProvider(
	services services.ProductServiceInterface,
) *ProductHandler {
	return &ProductHandler{
		services: services,
	}
}

//Get ...
func (h *ProductHandler) Get(resp http.ResponseWriter, req *http.Request) {
	limitString := req.URL.Query().Get("limit")
	limit, err := strconv.ParseUint(limitString, 10, 32)
	if err != nil {
		sentry.CaptureException(err)
		responseError(resp, req, "Parse error")
		return
	}

	products := h.services.GetProducts(uint32(limit))
	responseSuccess(resp, req, products)
}
