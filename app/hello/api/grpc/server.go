package grpc

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

var server *grpc.Server

func InitServer(network string, address string, srv HelloServer) {
	server = rpc.NewServer(&rpc.ServerConfig{
		Network: network,
		Addr:    address,
		RegFunc: func(s *grpc.Server) {
			RegisterHelloServer(s, srv)
		},
	})
}

func StopServer() {
	if server != nil {
		server.Stop()
	}
}
