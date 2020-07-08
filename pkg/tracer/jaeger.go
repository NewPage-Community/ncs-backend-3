package tracer

import (
	"backend/pkg/log"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-client-go/zipkin"
	"github.com/uber/jaeger-lib/metrics"
	"io"
	"time"
)

var tracerCloser io.Closer

func Init(serviceName string) {
	// Recommended configuration for production.
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Reporter: &jaegercfg.ReporterConfig{
			BufferFlushInterval: time.Second,
			CollectorEndpoint:   "http://jaeger-collector.istio-system:14268/api/traces",
		},
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Zipkin shares span ID between client and server spans; it must be enabled via the following option.
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()

	// Create tracer and then initialize global tracer
	closer, err := cfg.InitGlobalTracer(
		serviceName,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
		jaegercfg.Injector(opentracing.HTTPHeaders, zipkinPropagator),
		jaegercfg.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
		jaegercfg.ZipkinSharedRPCSpan(true),
	)

	if err != nil {
		log.Error("Could not initialize jaeger tracer:", err.Error())
		return
	}
	tracerCloser = closer
}

func Close() {
	if tracerCloser != nil {
		if err := tracerCloser.Close(); err != nil {
			log.Error(err)
		}
	}
}
