package main

import (
	api "backend/app/service/user/admin/api/grpc"
	"backend/app/service/user/admin/conf"
	"backend/app/service/user/admin/service"
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

	log.Info("Admin service started!")

	// cmd
	cmd.Run("Admin", func() {
		server.Stop()
		srv.Close()
		log.Close()
	})
}
