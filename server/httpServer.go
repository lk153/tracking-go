package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	services_pb "github.com/lk153/proto-tracking-gen/go/services"

	"factory/exam/handler"
	"factory/exam/utils/gateway"
	"factory/exam/utils/shutdown"
)

var _ shutdown.ServerAbstract = &HTTPServer{}

//HTTPServer ...
type HTTPServer struct {
	server *http.Server
}

//Builder ...
type Builder struct {
	*chi.Mux
	name string
}

//HTTPProvider ...
func HTTPProvider(
	ctx context.Context,
	productHandler *handler.ProductHandler,
	prodPBHandler services_pb.ProductServiceServer,
) (*HTTPServer, error) {
	router := NewHTTPBuilder()

	gateway := runtime.NewServeMux(gateway.DefaultGateMuxOpts()...)
	err := services_pb.RegisterProductServiceHandlerServer(ctx, gateway, prodPBHandler)
	if err != nil {
		return nil, err
	}

	router.Get("/products", productHandler.Get)
	router.Route("/", func(r chi.Router) {
		r.Mount("/v1", gateway)
	})

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &HTTPServer{
		server: server,
	}, nil
}

//NewHTTPBuilder ...
func NewHTTPBuilder() *Builder {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome changes"))
	})

	return &Builder{mux, "tracking"}
}

//Start ...
func (s *HTTPServer) Start() error {
	return s.server.ListenAndServe()
}

//Close ...
func (s *HTTPServer) Close() error {
	return s.server.Shutdown(context.Background())
}
