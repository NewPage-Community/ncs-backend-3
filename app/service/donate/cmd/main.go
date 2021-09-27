package main

import (
	api "backend/app/service/donate/api/grpc/v1"
	"backend/app/service/donate/conf"
	"backend/app/service/donate/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/tracer"
)

const serviceName = "donate-main"

func main() {
	// Init
	config := conf.Init()
	log.Init(config.Log)
	tracer.Init(serviceName)
	srv := service.Init(config, serviceName)

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
