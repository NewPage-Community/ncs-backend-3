package main

import (
	"backend/app/gateway/service"
	"backend/pkg/rpc"
	"flag"
	"net/http"

	"github.com/golang/glog"
)

func run() error {
	gateways := rpc.NewGateway(nil)
	service.RegService(gateways)
	defer gateways.Close()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", gateways.ServerMux())
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
