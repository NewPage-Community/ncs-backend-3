package service

import (
	chatGW "backend/app/game/chat/api/grpc/v1"
	cookieGW "backend/app/game/cookie/api/grpc/v1"
	cvarGW "backend/app/game/cvar/api/grpc/v1"
	passGW "backend/app/game/pass/api/grpc/v1"
	serverGW "backend/app/game/server/api/grpc/v1"
	skinGW "backend/app/game/skin/api/grpc/v1"
	statsGW "backend/app/game/stats/api/grpc/v1"
	storeGW "backend/app/game/store/api/grpc/v1"
	userGW "backend/app/game/user/api/grpc/v1"
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
	gws.AddGateway(
		cookieGW.RegisterCookieHandlerFromEndpoint,
		cookieGW.ServiceAddr,
	)
	gws.AddGateway(
		chatGW.RegisterChatHandlerFromEndpoint,
		chatGW.ServiceAddr,
	)
	gws.AddGateway(
		cvarGW.RegisterCVarHandlerFromEndpoint,
		cvarGW.ServiceAddr,
	)
}
