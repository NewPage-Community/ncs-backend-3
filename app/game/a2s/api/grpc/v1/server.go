package v1

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

func InitServer(srv A2SServer, health func() bool) (s *rpc.Server) {
	s = rpc.NewServer()
	s.Grpc(func(s *grpc.Server) {
		RegisterA2SServer(s, srv)
	})
	s.HealthCheck(health)
	return
}
