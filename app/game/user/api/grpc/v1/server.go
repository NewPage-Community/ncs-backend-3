package v1

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

func InitServer(srv GameServer, health func() bool) (s *rpc.Server) {
	s = rpc.NewServer()
	s.Grpc(func(s *grpc.Server) {
		RegisterGameServer(s, srv)
	})
	s.HealthCheck(health)
	return
}
