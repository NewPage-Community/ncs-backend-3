package v1

import (
	"backend/pkg/rpc"

	"google.golang.org/grpc"
)

func InitServer(srv interface{}, health func() bool) (s *rpc.Server) {
	s = rpc.NewServer(nil)
	s.Grpc(func(s *grpc.Server) {
		RegisterServerServer(s, srv.(ServerServer))
		RegisterServerPublicServer(s, srv.(ServerPublicServer))
	})
	s.HealthCheck(health)
	return
}
