package main

import (
	api "backend/app/service/user/ban/api/grpc"
	"backend/app/service/user/ban/conf"
	"backend/app/service/user/ban/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/tracer"
)

const serviceName = "user-ban"

func main() {
	// Init
	config := conf.Init()
	log.Init(config.Log)
	tracer.Init(serviceName)
	srv := service.Init(config)

	// rpc
	server := api.InitServer(srv, func() bool {
		return srv.Healthy()
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
