package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"factory/exam/handler"
	"factory/exam/utils/shutdown"
)

var _ shutdown.ServerAbstract = &HTTPServer{}

//HTTPServer ...
type HTTPServer struct {
	server *http.Server
}

//HTTPProvider ...
func HTTPProvider(
	productHandler *handler.ProductHandler,
) *HTTPServer {
	router := NewHTTPRouter()

	router.Get("/products", productHandler.Get)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &HTTPServer{
		server: server,
	}
}

//NewHTTPRouter ...
func NewHTTPRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome changes"))
	})

	return r
}

//Start ...
func (s *HTTPServer) Start() error {
	return s.server.ListenAndServe()
}

//Close ...
func (s *HTTPServer) Close() error {
	return s.server.Shutdown(context.Background())
}
