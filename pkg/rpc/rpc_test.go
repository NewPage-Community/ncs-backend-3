package rpc

import (
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestAll(t *testing.T) {
	server := NewServer(&ServerConfig{RegFunc: func(s *grpc.Server) {
		t.Log("RegFunc()")
	}})
	defer server.Stop()

	client := Dial(context.Background(), "0.0.0.0:2333", &ClientConfig{})
	defer client.Close()
}
