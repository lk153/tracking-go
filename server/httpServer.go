package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"

	services_pb "github.com/lk153/proto-tracking-gen/go/services"

	"factory/exam/handler"
	"factory/exam/utils/gateway"
	"factory/exam/utils/shutdown"
)

var _ shutdown.ServerAbstract = &HTTPServer{}

const (
	httpLatencyName         = "http.server.requests.duration.seconds"
	httpRequestNumberName   = "http.server.requests.number"
	externalRequestsTotal   = "http.external.requests.total"
	externalRequestsLatency = "http.external.requests.duration.seconds"
)

const (
	httpCodeLabel   = label.Key("code")
	httpMethodLabel = label.Key("method")
	httpPathLabel   = label.Key("path")
	httpHostLabel   = label.Key("host")
)

//HTTPServer ...
type HTTPServer struct {
	server *http.Server
}

//Builder ...
type Builder struct {
	*chi.Mux
	name string
}

func (b *Builder) Named(name string) *Builder {
	return b
}

func (b *Builder) Build() http.Handler {
	return otelhttp.NewHandler(b.Mux, b.name,
		otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents),
		otelhttp.WithFilter(shouldLogRequest),
	)
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

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome changes"))
	})

	router.Get("/products", productHandler.Get)
	router.Route("/", func(r chi.Router) {
		r.Mount("/v1", gateway)
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router.Build(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &HTTPServer{
		server: server,
	}, nil
}

// httpMetricsObserver is a handler that exposes otel metrics for the number of requests,
// the latency and the response size, partitioned by status code, method and HTTP path.
type httpMetricsObserver struct {
	latency       metric.Float64ValueRecorder
	meter         metric.Meter
	total_request metric.Int64Counter
	reqFilter     otelhttp.Filter
}

// NewHTTPMetricsObserver returns a new otel HTTPMiddleware handler.
func NewHTTPMetricsObserver(name string, reqFilter otelhttp.Filter) func(next http.Handler) http.Handler {
	var m httpMetricsObserver
	m.meter = global.Meter(name)
	m.latency = metric.Must(m.meter).NewFloat64ValueRecorder(
		httpLatencyName,
		metric.WithDescription("How long it took to process the request, partitioned by status code, method and HTTP path."),
	)

	m.total_request = metric.Must(m.meter).NewInt64Counter(
		httpRequestNumberName,
		metric.WithDescription("How many request throughput, partitioned by status code, method and HTTP path."),
	)
	m.reqFilter = reqFilter
	return m.handler
}

func (c httpMetricsObserver) handler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if c.reqFilter != nil && !c.reqFilter(r) {
			next.ServeHTTP(w, r)
			return
		}

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()
		next.ServeHTTP(ww, r)
		c.latency.Record(
			r.Context(),
			time.Since(start).Seconds(),
			httpCodeLabel.Int(ww.Status()),
			httpMethodLabel.String(r.Method),
			httpPathLabel.String(r.URL.Path),
		)

		c.total_request.Add(r.Context(),
			1,
			httpCodeLabel.Int(ww.Status()),
			httpMethodLabel.String(r.Method),
			httpPathLabel.String(r.URL.Path),
		)
	}
	return http.HandlerFunc(fn)
}

var (
	excludeHTTPPath = map[string]struct{}{
		"/":            {},
		"/health":      {},
		"/favicon.ico": {},
	}
)

func shouldLogRequest(r *http.Request) bool {
	_, ok := excludeHTTPPath[r.URL.Path]
	return !ok
}

func AllowTrackingService(_ *http.Request, origin string) bool {
	return false
}

//NewHTTPBuilder ...
func NewHTTPBuilder() *Builder {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  AllowTrackingService,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Use(NewHTTPMetricsObserver("tracking", shouldLogRequest))

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
