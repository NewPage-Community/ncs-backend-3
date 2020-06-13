package main

import (
	api "backend/app/game/user/api/grpc"
	"backend/app/game/user/conf"
	"backend/app/game/user/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
)

func main() {
	// Init
	config := conf.Init()
	log.Init(config.Log)
	srv := service.Init()

	// rpc
	// TODO: health check
	server := api.InitServer(srv, func() bool {
		return true
	})

	log.Info("Game user service started!")

	// cmd
	cmd.Run("Game user", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}
