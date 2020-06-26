package grpc

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

func InitServer(srv SkinServer, health func() bool) (s *rpc.Server) {
	s = rpc.NewServer(nil)
	s.Grpc(func(s *grpc.Server) {
		RegisterSkinServer(s, srv)
	})
	s.HealthCheck(health)
	return
}
