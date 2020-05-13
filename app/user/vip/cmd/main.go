package main

import (
	api "backend/app/user/vip/api/grpc"
	"backend/app/user/vip/conf"
	"backend/app/user/vip/service"
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

	log.Info("VIP app started!")

	// cmd
	cmd.Run("VIP", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}
