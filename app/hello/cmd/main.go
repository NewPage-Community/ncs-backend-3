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
	api.InitServer("tcp", "0.0.0.0:2333", &service.Hello{})
	api.InitClient("0.0.0.0:2333")

	log.Info("Hello app started!")

	go func() {
		for {
			res, err := api.CallSay(context.Background(), "world")
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
		api.StopServer()
		api.CloseClient()
	})
}
