package grpc

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
}

func InitServer(network string, address string, srv SignServer) (s *Server) {
	s = &Server{}
	s.server = rpc.NewServer(&rpc.ServerConfig{
		Network: network,
		Addr:    address,
		RegFunc: func(s *grpc.Server) {
			RegisterSignServer(s, srv)
		},
	})
	return
}

func (s *Server) Stop() {
	if s.server != nil {
		s.server.Stop()
	}
}