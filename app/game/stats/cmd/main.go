package main

import (
	api "backend/app/game/stats/api/grpc/v1"
	"backend/app/game/stats/conf"
	"backend/app/game/stats/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/tracer"
)

const serviceName = "game-stats"

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
