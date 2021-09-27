package main

import (
	pb "backend/app/bot/kaiheila/api/grpc/v1"
	"backend/app/bot/kaiheila/conf"
	"backend/app/bot/kaiheila/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/rpc"
	"backend/pkg/tracer"

	"google.golang.org/grpc"
)

// 这里修改服务名字
const serviceName = "bot-kaiheila"

func main() {
	// Init 初始化模块
	config := conf.Init()
	log.Init(config.Log)
	tracer.Init(serviceName)
	srv := service.Init(config, serviceName)

	// rpc 服务注册
	server := rpc.NewServer(nil)
	server.Grpc(func(s *grpc.Server) {
		pb.RegisterKaiheilaServer(s, srv)
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
