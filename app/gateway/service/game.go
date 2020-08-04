package service

import (
	passGW "backend/app/game/pass/api/grpc"
	serverGW "backend/app/game/server/api/grpc"
	skinGW "backend/app/game/skin/api/grpc"
	statsGW "backend/app/game/stats/api/grpc"
	storeGW "backend/app/game/store/api/grpc"
	userGW "backend/app/game/user/api/grpc"
	"backend/pkg/rpc"
)

func regGameService(gws *rpc.Gateways) {
	gws.AddGateway(
		serverGW.RegisterServerHandlerFromEndpoint,
		serverGW.ServiceAddr,
	)
	gws.AddGateway(
		userGW.RegisterGameHandlerFromEndpoint,
		userGW.ServiceAddr,
	)
	gws.AddGateway(
		skinGW.RegisterSkinHandlerFromEndpoint,
		skinGW.ServiceAddr,
	)
	gws.AddGateway(
		storeGW.RegisterStoreHandlerFromEndpoint,
		storeGW.ServiceAddr,
	)
	gws.AddGateway(
		statsGW.RegisterStatsHandlerFromEndpoint,
		statsGW.ServiceAddr,
	)
	gws.AddGateway(
		passGW.RegisterPassHandlerFromEndpoint,
		passGW.ServiceAddr,
	)
}
