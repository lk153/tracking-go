package tracer

import (
	"log"

	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

//InitTracer ...
func InitTracer() func() {
	// Create and install Jaeger export pipeline.
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "trace.tracking.service",
			Tags:        nil,
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.TraceIDRatioBased(0.1)}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return flush
}
