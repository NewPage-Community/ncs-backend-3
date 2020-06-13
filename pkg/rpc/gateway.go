package rpc

import (
	"context"
	"github.com/golang/glog"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

const (
	prefixTracerState  = "x-b3-"
	zipkinTraceID      = prefixTracerState + "traceid"
	zipkinSpanID       = prefixTracerState + "spanid"
	zipkinParentSpanID = prefixTracerState + "parentspanid"
	zipkinSampled      = prefixTracerState + "sampled"
	zipkinFlags        = prefixTracerState + "flags"
)

type regHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type annotator func(context.Context, *http.Request) metadata.MD

type Gateways struct {
	server *http.Server
	mux    *runtime.ServeMux
	opts   []grpc.DialOption
	ctx    context.Context
	cancel context.CancelFunc
}

var otHeaders = []string{
	zipkinTraceID,
	zipkinSpanID,
	zipkinParentSpanID,
	zipkinSampled,
	zipkinFlags,
}

func injectHeadersIntoMetadata(ctx context.Context, req *http.Request) metadata.MD {
	var pairs []string
	for _, h := range otHeaders {
		if v := req.Header.Get(h); len(v) > 0 {
			pairs = append(pairs, h, v)
		}
	}
	return metadata.Pairs(pairs...)
}

func chainGrpcAnnotators(annotators ...annotator) annotator {
	return func(c context.Context, r *http.Request) metadata.MD {
		var mds []metadata.MD
		for _, a := range annotators {
			mds = append(mds, a(c, r))
		}
		return metadata.Join(mds...)
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
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		gws.server.Shutdown(ctx)
	}
	gws.cancel()
}

func NewGateway() *Gateways {
	ctx, cancel := context.WithCancel(context.Background())
	return &Gateways{
		mux: runtime.NewServeMux(
			runtime.WithMetadata(chainGrpcAnnotators([]annotator{
				injectHeadersIntoMetadata,
			}...))),
		opts: []grpc.DialOption{
			grpc.WithUnaryInterceptor(
				grpc_retry.UnaryClientInterceptor(
					grpc_retry.WithMax(_defaultCliConf.MaxRetry),
					grpc_retry.WithCodes(_defaultCliConf.RetryCode...),
					grpc_retry.WithPerRetryTimeout(_defaultCliConf.Timeout),
				)),
			grpc.WithInsecure(),
		},
		ctx:    ctx,
		cancel: cancel,
	}
}

func (gws *Gateways) Gateway(health func() bool) {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", HealthCheck(health))
	mux.HandleFunc("/", gws.mux.ServeHTTP)
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
