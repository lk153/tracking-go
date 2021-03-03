package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	client_prometheus "github.com/prometheus/client_golang/prometheus"
	otel_instrumentation_host "go.opentelemetry.io/contrib/instrumentation/host"
	otel_instrumentation_runtime "go.opentelemetry.io/contrib/instrumentation/runtime"
	otel_exporters_prometheus "go.opentelemetry.io/otel/exporters/metric/prometheus"
	otel_global "go.opentelemetry.io/otel/metric/global"
	otel_metric_controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	otel_resource "go.opentelemetry.io/otel/sdk/resource"

	"factory/exam/config"
)

//MetricServer ...
type MetricServer struct {
	server *http.Server
}

//NewMetricServer ...
func NewMetricServer(
	metricPort config.MetricPort,
) (*MetricServer, error) {
	resource, err := otel_resource.New(
		context.Background(),
	)
	if err != nil {
		return nil, err
	}
	exporter, err := otel_exporters_prometheus.InstallNewPipeline(
		otel_exporters_prometheus.Config{
			Registry:                   client_prometheus.NewRegistry(),
			DefaultHistogramBoundaries: []float64{.0005, 0.0075, 0.001, 0.002, 0.003, 0.004, .005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		},
		otel_metric_controller.WithResource(resource),
		otel_metric_controller.WithCollectPeriod(10*time.Second),
	)

	if err != nil {
		return nil, fmt.Errorf("Init metrics: %w", err)
	}

	if err := WithRuntime(); err != nil {
		return nil, fmt.Errorf("Init runtimemetrics: %w", err)
	}

	otel_global.SetMeterProvider(exporter.MeterProvider())

	mux := chi.NewRouter()
	mux.Get("/metrics", exporter.ServeHTTP)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", metricPort),
		Handler: mux,
	}
	return &MetricServer{server: server}, nil
}

//Start ...
func (s *MetricServer) Start() error {
	return s.server.ListenAndServe()
}

//Close ...
func (s *MetricServer) Close() error {
	s.server.Shutdown(context.Background())
	return nil
}

//WithRuntime ...
func WithRuntime() error {
	err := otel_instrumentation_runtime.Start(
		otel_instrumentation_runtime.WithMinimumReadMemStatsInterval(time.Second),
	)
	if err != nil {
		return err
	}
	err = otel_instrumentation_host.Start()
	return err
}
