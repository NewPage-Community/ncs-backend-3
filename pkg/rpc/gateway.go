package rpc

import (
	"backend/pkg/log"
	"context"
	"github.com/golang/glog"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"net/http"
)

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

type regHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type Gateways struct {
	server *http.Server
	mux    *runtime.ServeMux
	opts   []grpc.DialOption
	ctx    context.Context
	cancel context.CancelFunc
}

func tracingWrapper(h http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		parentSpanContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err == nil || err == opentracing.ErrSpanContextNotFound {
			serverSpan := opentracing.GlobalTracer().StartSpan(
				"ServeHTTP",
				// this is magical, it attaches the new span to the parent parentSpanContext, and creates an unparented one if empty.
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		h.ServeHTTP(w, r)
	}
}

func (gws *Gateways) AddGateway(handler regHandler, endpoint string) {
	err := handler(gws.ctx, gws.mux, endpoint, gws.opts)
	if err != nil {
		glog.Fatal(err)
	}
}

func (gws *Gateways) Close() {
	if gws.server != nil {
		if err := gws.server.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
	}
	gws.cancel()
}

func NewGateway() *Gateways {
	ctx, cancel := context.WithCancel(context.Background())
	tracer := grpc_opentracing.WithTracer(opentracing.GlobalTracer())
	return &Gateways{
		mux: runtime.NewServeMux(runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{OrigName: true, EmitDefaults: true}),
		),
		opts: []grpc.DialOption{
			grpc.WithUnaryInterceptor(
				grpc_middleware.ChainUnaryClient(
					grpc_opentracing.UnaryClientInterceptor(tracer),
					grpc_retry.UnaryClientInterceptor(
						grpc_retry.WithMax(_defaultCliConf.MaxRetry),
						grpc_retry.WithCodes(_defaultCliConf.RetryCode...),
						grpc_retry.WithPerRetryTimeout(_defaultCliConf.Timeout),
					),
				),
			),
			grpc.WithInsecure(),
		},
		ctx:    ctx,
		cancel: cancel,
	}
}

func (gws *Gateways) Gateway(health func() bool) {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", HttpHealthHandler(health))
	mux.HandleFunc("/", tracingWrapper(gws.mux))
	gws.server = &http.Server{
		Addr:    ":23333",
		Handler: mux,
	}

	go func() {
		err := gws.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}
