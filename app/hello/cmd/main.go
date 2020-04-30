package main

import (
	"backend/pkg/cmd"
	"backend/pkg/grpc"
	"backend/pkg/log"
)

func main() {
	// Init
	log.Init(true)
	defer log.Close()

	log.Info("Hello world!")
	s := grpc.NewServer()

	// cmd
	cmd.Run("Hello", func(){})
}