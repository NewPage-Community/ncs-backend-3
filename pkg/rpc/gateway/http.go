package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
)

func Redirect(ctx context.Context, url string) {
	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "302"))
	_ = grpc.SetHeader(ctx, metadata.Pairs("x-location", url))
}

func SetCookie(ctx context.Context, cookie *http.Cookie) {
	_ = grpc.SetHeader(ctx, metadata.Pairs("x-set-cookie", cookie.String()))
}

func HttpResponseModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// Set http status code
	if vals := md.HeaderMD.Get("x-http-code"); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, "x-http-code")
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
		w.WriteHeader(code)
	}

	// Set http location
	if vals := md.HeaderMD.Get("x-location"); len(vals) > 0 {
		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, "x-location")
		delete(w.Header(), "Grpc-Metadata-X-Location")
		w.Header().Set("Location", vals[0])
	}

	// Set http location
	if vals := md.HeaderMD.Get("x-set-cookie"); len(vals) > 0 {
		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, "x-set-cookie")
		delete(w.Header(), "Grpc-Metadata-X-Set-Cookie")
		w.Header().Set("Set-Cookie", vals[0])
	}

	return nil
}
