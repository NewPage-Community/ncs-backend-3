package service

import (
	cookieGW "backend/app/game/cookie/api/grpc/v1"
	serverGW "backend/app/game/server/api/grpc/v1"
	statsGW "backend/app/game/stats/api/grpc/v1"
	"backend/pkg/rpc"
)

func regGameService(gws *rpc.Gateways) {
	gws.AddGateway(
		serverGW.RegisterServerPublicHandlerFromEndpoint,
		serverGW.ServiceAddr,
	)
	gws.AddGateway(
		cookieGW.RegisterCookieHandlerFromEndpoint,
		cookieGW.ServiceAddr,
	)
	gws.AddGateway(
		statsGW.RegisterStatsPublicHandlerFromEndpoint,
		statsGW.ServiceAddr,
	)
}
