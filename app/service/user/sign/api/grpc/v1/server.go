package v1

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

func InitServer(srv SignServer, health func() bool) (s *rpc.Server) {
	s = rpc.NewServer()
	s.Grpc(func(s *grpc.Server) {
		RegisterSignServer(s, srv)
	})
	s.HealthCheck(health)
	return
}
