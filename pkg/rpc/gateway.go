package rpc

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/opentracing/opentracing-go"
	ot "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"net/http"
)

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

type regHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type Gateways struct {
	mux    *runtime.ServeMux
	opts   []grpc.DialOption
	ctx    context.Context
	cancel context.CancelFunc
}

func (gws *Gateways) AddGateway(handler regHandler, endpoint string) {
	err := handler(gws.ctx, gws.mux, endpoint, gws.opts)
	if err != nil {
		glog.Fatal(err)
	}
}

func (gws *Gateways) Close() {
	gws.cancel()
}

func (gws *Gateways) HTTPHandler() func(w http.ResponseWriter, r *http.Request) {
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
		gws.mux.ServeHTTP(w, r)
	}
}

func NewGateway() *Gateways {
	ctx, cancel := context.WithCancel(context.Background())
	return &Gateways{
		mux: runtime.NewServeMux(),
		opts: []grpc.DialOption{
			grpc.WithUnaryInterceptor(
				grpc_opentracing.UnaryClientInterceptor(
					grpc_opentracing.WithTracer(ot.GlobalTracer()))),
			grpc.WithStreamInterceptor(
				grpc_opentracing.StreamClientInterceptor(
					grpc_opentracing.WithTracer(ot.GlobalTracer()))),
			grpc.WithInsecure(),
		},
		ctx:    ctx,
		cancel: cancel,
	}
}

func HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "Healthy")
	}
}
