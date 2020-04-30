package main

import (
	"backend/pkg/cmd"
	"backend/pkg/log"
)

func main() {
	log.Init(true)
	defer log.Close()

	log.Info("Hello world!")

	cmd.Run(&cmd.App{"Hello", func(){}})
}