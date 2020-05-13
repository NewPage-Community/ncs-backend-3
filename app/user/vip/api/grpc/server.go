package grpc

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
}

func InitServer(network string, address string, srv VIPServer) (s *Server) {
	s = &Server{}
	s.server = rpc.NewServer(&rpc.ServerConfig{
		Network: network,
		Addr:    address,
		RegFunc: func(s *grpc.Server) {
			RegisterVIPServer(s, srv)
		},
	})
	return
}

func (s *Server) Stop() {
	if s.server != nil {
		s.server.Stop()
	}
}
