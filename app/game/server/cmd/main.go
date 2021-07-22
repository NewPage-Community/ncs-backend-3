package main

import (
	api "backend/app/game/server/api/grpc/v1"
	"backend/app/game/server/conf"
	"backend/app/game/server/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/tracer"
)

const serviceName = "game-server"

func main() {
	// Init
	config := conf.Init()
	tracer.Init(serviceName)
	log.Init()
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
