package grpc

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

func InitServer(srv RewardServer, health func() bool) (s *rpc.Server) {
	s = rpc.NewServer(nil)
	s.Grpc(func(s *grpc.Server) {
		RegisterRewardServer(s, srv)
	})
	s.HealthCheck(health)
	return
}
