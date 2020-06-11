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
	log.Init(&log.Config{Debug: true})
	gateways := rpc.NewGateway()
	service.RegService(gateways)
	defer gateways.Close()

	srv := &http.Server{
		Addr:           ":8081",
		Handler:        gateways.HTTPHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error(err)
		}
	}()

	log.Info("Gateway started!")

	cmd.Run("Gateway", func() {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*15)
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Error(err)
		}
	})
}
