package v1

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

func InitServer(srv TitleServer, health func() bool) (s *rpc.Server) {
	s = rpc.NewServer(nil)
	s.Grpc(func(s *grpc.Server) {
		RegisterTitleServer(s, srv)
	})
	s.HealthCheck(health)
	return
}
