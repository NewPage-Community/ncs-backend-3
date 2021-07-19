package service

import (
	serverGW "backend/app/game/server/api/grpc"
	"backend/pkg/rpc"
)

func regGameService(gws *rpc.Gateways) {
	gws.AddGateway(
		serverGW.RegisterServerHandlerFromEndpoint,
		serverGW.ServiceAddr,
	)
}
