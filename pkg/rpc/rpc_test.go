package rpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestAll(t *testing.T) {
	server := NewServer(&ServerConfig{RegFunc: func(s *grpc.Server) {
		t.Log("RegFunc()")
	}})
	defer server.Stop()

	client := Dial(context.Background(), "0.0.0.0:2333", &ClientConfig{})
	defer client.Close()

	err := Errorf(codes.Internal, "OK!")
	if GetError(err).Code != codes.Internal {
		t.Errorf("Error code: v = %d, want = 13", GetError(err).Code)
	}
	if GetError(nil) != nil {
		t.Errorf("GetError(): v = %v, want = nil", GetError(err))
	}
}
