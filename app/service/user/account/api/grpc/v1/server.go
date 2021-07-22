package v1

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

func InitServer(srv interface{}, health func() bool) (s *rpc.Server) {
	s = rpc.NewServer()
	s.Grpc(func(s *grpc.Server) {
		RegisterAccountServer(s, srv.(AccountServer))
		RegisterWebServer(s, srv.(WebServer))
	})
	s.HealthCheck(health)
	return
}