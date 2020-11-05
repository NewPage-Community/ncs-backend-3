// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: app/hello/api/grpc/v1/hello.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_Hello_Hello_0(ctx context.Context, marshaler runtime.Marshaler, client HelloClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq HelloReq
	var metadata runtime.ServerMetadata

	msg, err := client.Hello(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Hello_Hello_0(ctx context.Context, marshaler runtime.Marshaler, server HelloServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq HelloReq
	var metadata runtime.ServerMetadata

	msg, err := server.Hello(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterHelloHandlerServer registers the http handlers for service Hello to "mux".
// UnaryRPC     :call HelloServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterHelloHandlerFromEndpoint instead.
func RegisterHelloHandlerServer(ctx context.Context, mux *runtime.ServeMux, server HelloServer) error {

	mux.Handle("GET", pattern_Hello_Hello_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ncs.hello.v1.Hello/Hello")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Hello_Hello_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Hello_Hello_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterHelloHandlerFromEndpoint is same as RegisterHelloHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterHelloHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterHelloHandler(ctx, mux, conn)
}

// RegisterHelloHandler registers the http handlers for service Hello to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterHelloHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterHelloHandlerClient(ctx, mux, NewHelloClient(conn))
}

// RegisterHelloHandlerClient registers the http handlers for service Hello
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "HelloClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "HelloClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "HelloClient" to call the correct interceptors.
func RegisterHelloHandlerClient(ctx context.Context, mux *runtime.ServeMux, client HelloClient) error {

	mux.Handle("GET", pattern_Hello_Hello_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ncs.hello.v1.Hello/Hello")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Hello_Hello_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Hello_Hello_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Hello_Hello_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"hello"}, ""))
)

var (
	forward_Hello_Hello_0 = runtime.ForwardResponseMessage
)
