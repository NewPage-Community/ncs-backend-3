package rpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

const (
	prefixTracerState  = "x-b3-"
	zipkinTraceID      = prefixTracerState + "traceid"
	zipkinSpanID       = prefixTracerState + "spanid"
	zipkinParentSpanID = prefixTracerState + "parentspanid"
	zipkinSampled      = prefixTracerState + "sampled"
	zipkinFlags        = prefixTracerState + "flags"
)

var otHeaders = []string{
	zipkinTraceID,
	zipkinSpanID,
	zipkinParentSpanID,
	zipkinSampled,
	zipkinFlags,
}

type annotator func(context.Context, *http.Request) metadata.MD

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

func NewGateway(conf *ClientConfig) (mux *runtime.ServeMux, opts []grpc.DialOption) {
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	annotators := []annotator{injectHeadersIntoMetadata}
	mux = runtime.NewServeMux(
		runtime.WithMetadata(chainGrpcAnnotators(annotators...)),
	)
	opts = getDialOption(setCliConf(conf))
	return
}
