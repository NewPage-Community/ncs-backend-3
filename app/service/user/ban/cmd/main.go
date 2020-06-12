package main

import (
	api "backend/app/service/user/ban/api/grpc"
	"backend/app/service/user/ban/conf"
	"backend/app/service/user/ban/service"
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

	log.Info("Ban app started!")

	// cmd
	cmd.Run("Ban", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}
