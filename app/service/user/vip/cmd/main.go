package main

import (
	api "backend/app/service/user/vip/api/grpc"
	"backend/app/service/user/vip/conf"
	"backend/app/service/user/vip/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
)

func main() {
	// Init
	config := conf.Init()
	log.Init(config.Log)
	srv := service.Init(config)

	// rpc
	// TODO: Health check
	server := api.InitServer(srv, func() bool {
		return true
	})

	log.Info("VIP app started!")

	// cmd
	cmd.Run("VIP", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}
