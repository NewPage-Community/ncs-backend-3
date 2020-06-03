package main

import (
	api "backend/app/user/sign/api/grpc"
	"backend/app/user/sign/conf"
	"backend/app/user/sign/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
)

func main() {
	// Init
	config := conf.Init()
	log.Init(config.Log)
	srv := service.Init(config)

	// rpc
	server := api.InitServer("tcp", "0.0.0.0:2333", srv)

	log.Info("Sign service started!")

	// cmd
	cmd.Run("Sign", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}
