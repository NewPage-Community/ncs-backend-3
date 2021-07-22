package main

import (
	api "backend/app/game/skin/api/grpc/v1"
	"backend/app/game/skin/conf"
	"backend/app/game/skin/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/tracer"
)

const serviceName = "game-skin"

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
