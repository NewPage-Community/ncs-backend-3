package v1

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

func InitServer(srv ChatServer, health func() bool) (s *rpc.Server) {
	s = rpc.NewServer(nil)
	s.Grpc(func(s *grpc.Server) {
		RegisterChatServer(s, srv)
	})
	s.HealthCheck(health)
	return
}
