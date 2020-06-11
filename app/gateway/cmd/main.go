package main

import (
	"backend/app/gateway/service"
	"backend/pkg/cmd"
	"backend/pkg/log"
	"backend/pkg/rpc"
	"context"
	"net/http"
	"time"
)

func main() {
	healthy := true

	log.Init(&log.Config{Debug: true})
	gateways := rpc.NewGateway()
	service.RegService(gateways)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", rpc.HealthCheck(&healthy))
	mux.HandleFunc("/", gateways.HTTPHandler())
	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Panic(err)
		}
	}()

	log.Info("Gateway started!")

	cmd.Run("Gateway", func() {
		// Waiting k8s deal with terminating
		healthy = false
		time.Sleep(time.Minute)
		// Close http -> grpc
		ctx, _ := context.WithTimeout(context.Background(), time.Second*15)
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Error(err)
		}
		gateways.Close()
		log.Close()
	})
}
