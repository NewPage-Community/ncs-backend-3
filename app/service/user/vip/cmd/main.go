package main

import (
	api "backend/app/service/user/vip/api/grpc"
	"backend/app/service/user/vip/conf"
	"backend/app/service/user/vip/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/tracer"
)

const serviceName = "Gateway"

func main() {
	// Init
	config := conf.Init()
	log.Init(config.Log)
	tracer.Init(serviceName)
	srv := service.Init(config)

	// rpc
	// TODO: Health check
	server := api.InitServer(srv, func() bool {
		return true
	})

	log.Info(serviceName, "service started!")

	// cmd
	cmd.Run(serviceName, func() {
		server.Stop()
		srv.Close()
		tracer.Close()
		log.Close()
	})
}
