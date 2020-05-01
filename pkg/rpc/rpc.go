package rpc

import (
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	ot "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
)

func NewServer(network string, address string, RegFunc func(s *grpc.Server)) *grpc.Server {
	// Options
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(ot.GlobalTracer())),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(ot.GlobalTracer())),
	}

	// Network
	lis, err := net.Listen(network, address)
	if err != nil {
		panic(err)
	}

	// Initialize
	s := grpc.NewServer(opts...)

	// Register
	RegFunc(s)

	// Serve
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return s
}

func Dial(target string) *grpc.ClientConn {
	// Options
	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(ot.GlobalTracer())),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(ot.GlobalTracer())),
		grpc.WithInsecure(),
	}

	// Initialize
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		panic(err)
	}
	return conn
}