package grpc

import (
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	ot "github.com/opentracing/opentracing-go"
)

func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	opts = append(opts,
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(ot.GlobalTracer())),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(ot.GlobalTracer())),
		)

	// Initialize the gRPC server.
	return grpc.NewServer(opts...)
}

func Dial(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts,
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(ot.GlobalTracer())),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(ot.GlobalTracer())),
	)

	// Initialize the gRPC client.
	return grpc.Dial(target, opts...)
}