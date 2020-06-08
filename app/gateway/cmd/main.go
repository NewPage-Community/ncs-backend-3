package main

import (
	"backend/app/gateway/service"
	"backend/pkg/rpc"
	"flag"
	"net/http"

	"github.com/golang/glog"
)

func run() error {
	mux, opts := rpc.NewGateway(nil)
	srv := service.NewService(mux, opts)
	srv.Register()
	defer srv.Close()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
