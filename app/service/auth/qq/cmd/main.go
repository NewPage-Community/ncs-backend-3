package main

import (
	pb "backend/app/service/auth/qq/api/grpc/v1"
	"backend/app/service/auth/qq/conf"
	"backend/app/service/auth/qq/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/rpc"
	"backend/pkg/tracer"
	"google.golang.org/grpc"
)

// 服务名
const (
	serviceName = pb.ServiceName
)

func main() {
	// Init 初始化模块
	config := conf.Init()
	log.Init()
	tracer.Init(serviceName)
	srv := service.Init(config)

	// rpc 服务注册
	server := rpc.NewServer()
	server.Grpc(func(s *grpc.Server) {
		pb.RegisterWebServer(s, srv)
	})
	server.HealthCheck(srv.Healthy)

	log.Info(serviceName, "service started!")

	// cmd 控制服务优雅关闭
	cmd.Run(serviceName, func() {
		server.Stop()
		srv.Close()
		tracer.Close()
		log.Close()
	})
}
