package main

import (
	"backend/app/gateway/game/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/rpc"
	"backend/pkg/tracer"
)

const serviceName = "gateway-game"

func main() {
	// Init
	log.Init(&log.Config{Debug: true})
	tracer.Init(serviceName)
	gateways := rpc.NewGateway(serviceName)

	// Gateway
	service.RegService(gateways)
	gateways.Gateway(func() bool {
		return true
	})

	log.Info(serviceName, "started!")

	// cmd
	cmd.Run(serviceName, func() {
		// Waiting k8s deal with terminating
		// Close http -> grpc
		gateways.Close()
		tracer.Close()
		log.Close()
	})
}
