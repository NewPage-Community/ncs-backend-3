package main

import (
	"backend/app/gateway/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/rpc"
)

func main() {
	log.Init(&log.Config{Debug: true})
	gateways := rpc.NewGateway()
	service.RegService(gateways)
	gateways.Gateway(func() bool {
		return true
	})

	log.Info("Gateway started!")

	cmd.Run("Gateway", func() {
		// Waiting k8s deal with terminating
		// Close http -> grpc
		gateways.Close()
		log.Close()
	})
}
