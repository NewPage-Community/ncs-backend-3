package main

import (
	api "backend/app/service/user/sign/api/grpc"
	"backend/app/service/user/sign/conf"
	"backend/app/service/user/sign/service"
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

	log.Info("Sign service started!")

	// cmd
	cmd.Run("Sign", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}
