package rpc

import (
	"backend/pkg/log"
	"backend/pkg/rpc/gateway"
	"context"
	"net/http"

	"github.com/golang/glog"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

type regHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type Gateways struct {
	server *http.Server
	mux    *runtime.ServeMux
	opts   []grpc.DialOption
	ctx    context.Context
	cancel context.CancelFunc
	id     string
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

func metadataModify(id string) func(ctx context.Context, req *http.Request) metadata.MD {
	return func(ctx context.Context, req *http.Request) metadata.MD {
		return gateway.InjectID(ctx, id)
	}
}

func (gws *Gateways) AddGateway(handler regHandler, endpoint string) {
	dnsEndpoint := "dns:///" + endpoint
	err := handler(gws.ctx, gws.mux, dnsEndpoint, gws.opts)
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

func NewGateway(name string) *Gateways {
	ctx, cancel := context.WithCancel(context.Background())
	tracer := grpc_opentracing.WithTracer(opentracing.GlobalTracer())
	return &Gateways{
		mux: runtime.NewServeMux(runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			}),
			runtime.WithMetadata(metadataModify(name)),
			runtime.WithForwardResponseOption(gateway.HttpResponseModifier),
		),
		opts: []grpc.DialOption{
			grpc.WithUnaryInterceptor(
				grpc_middleware.ChainUnaryClient(
					grpc_opentracing.UnaryClientInterceptor(tracer),
					grpc_retry.UnaryClientInterceptor(
						grpc_retry.WithMax(_defaultCliConf.RetryMaxTime),
						grpc_retry.WithCodes(_defaultCliConf.RetryCode...),
					),
				),
			),
			grpc.WithInsecure(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                _defaultCliConf.KeepaliveInterval,
				Timeout:             _defaultCliConf.KeepaliveTimeout,
				PermitWithoutStream: _defaultCliConf.PermitWithoutStream,
			}),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		},
		ctx:    ctx,
		cancel: cancel,
		id:     name,
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
