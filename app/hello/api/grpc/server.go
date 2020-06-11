package grpc

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

func InitServer(network string, address string, srv HelloServer) (s *rpc.Server) {
	s = rpc.NewServer(&rpc.ServerConfig{
		Network: network,
		Addr:    address,
		RegFunc: func(s *grpc.Server) {
			RegisterHelloServer(s, srv)
		},
	})
	return
}
