package main

import (
	api "backend/app/user/admin/api/grpc"
	"backend/app/user/admin/conf"
	"backend/app/user/admin/service"
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

	log.Info("Admin service started!")

	// cmd
	cmd.Run("Admin", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}
