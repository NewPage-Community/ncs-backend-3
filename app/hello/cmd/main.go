package main

import (
	api "backend/app/hello/api/grpc"
	"backend/app/hello/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"context"
	"time"
)

func main() {
	// Init
	log.Init(true)
	defer log.Close()

	// rpc
	server := api.InitServer("tcp", "0.0.0.0:2333", &service.Hello{})
	client := api.InitClient("0.0.0.0:2333")

	log.Info("Hello app started!")

	go func() {
		for {
			res, err := client.Say(context.Background(), "world")
			if err != nil {
				log.Error(err.Error())
			} else {
				log.Info(res)
			}
			time.Sleep(time.Second)
		}
	}()

	// cmd
	cmd.Run("Hello", func() {
		client.Close()
		server.Stop()
	})
}
