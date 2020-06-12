package main

import (
	api "backend/app/game/server/api/grpc"
	"backend/app/game/server/conf"
	"backend/app/game/server/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
)

func main() {
	// Init
	config := conf.Init()
	log.Init(config.Log)
	srv := service.Init(config)

	// rpc
	server := api.InitServer(srv, func() bool {
		return srv.Healthy()
	})

	log.Info("Server service started!")

	// cmd
	cmd.Run("Server", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}
