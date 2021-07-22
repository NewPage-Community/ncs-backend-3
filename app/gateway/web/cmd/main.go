package main

import (
	"backend/app/gateway/web/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/rpc"
	"backend/pkg/tracer"
)

const serviceName = "gateway-web"

func main() {
	// Init
	log.Init()
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
