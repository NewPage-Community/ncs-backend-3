package service

import (
	cookieGW "backend/app/game/cookie/api/grpc/v1"
	serverGW "backend/app/game/server/api/grpc/v1"
	"backend/pkg/rpc"
)

func regGameService(gws *rpc.Gateways) {
	gws.AddGateway(
		serverGW.RegisterServerHandlerFromEndpoint,
		serverGW.ServiceAddr,
	)
	gws.AddGateway(
		cookieGW.RegisterCookieHandlerFromEndpoint,
		cookieGW.ServiceAddr,
	)
}
