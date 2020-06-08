package rpc

import (
	"context"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"net/http"
)

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

type regHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type Gateways struct {
	mux    *runtime.ServeMux
	opts   []grpc.DialOption
	cancel []context.CancelFunc
}

func (gws *Gateways) AddGateway(handler regHandler, endpoint string) {
	ctx, cancel := context.WithCancel(context.Background())
	err := handler(ctx, gws.mux, endpoint, gws.opts)
	if err != nil {
		glog.Fatal(err)
	}
	gws.cancel = append(gws.cancel, cancel)
}

func (gws *Gateways) Close() {
	for _, cancel := range gws.cancel {
		if cancel != nil {
			cancel()
		}
	}
}

func (gws *Gateways) ServerMux() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}

func NewGateway(conf *ClientConfig) *Gateways {
	return &Gateways{
		mux:    runtime.NewServeMux(),
		opts:   getDialOption(setCliConf(conf)),
		cancel: make([]context.CancelFunc, 0),
	}
}
