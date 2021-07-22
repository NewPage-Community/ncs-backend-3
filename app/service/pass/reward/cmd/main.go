package main

import (
	api "backend/app/service/pass/reward/api/grpc/v1"
	"backend/app/service/pass/reward/conf"
	"backend/app/service/pass/reward/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/tracer"
)

const serviceName = "pass-reward"

func main() {
	// Init
	config := conf.Init()
	log.Init()
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
