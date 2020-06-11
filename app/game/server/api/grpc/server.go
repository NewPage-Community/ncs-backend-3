package grpc

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

func InitServer(network string, address string, srv ServerServer) (s *rpc.Server) {
	s = rpc.NewServer(&rpc.ServerConfig{
		Network: network,
		Addr:    address,
		RegFunc: func(s *grpc.Server) {
			RegisterServerServer(s, srv)
		},
	})
	return
}
