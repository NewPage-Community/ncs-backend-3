package v1

import (
	"backend/pkg/rpc"
	"google.golang.org/grpc"
)

// InitServer 初始化服务端
func InitServer(srv HelloServer, health func() bool) (s *rpc.Server) {
	s = rpc.NewServer(nil)
	s.Grpc(func(s *grpc.Server) {
		// 这里修改 grpc 接口
		RegisterHelloServer(s, srv)
	})
	s.HealthCheck(health)
	return
}
