package service

import (
	accountGW "backend/app/service/user/account/api/grpc/v1"
	banGW "backend/app/service/user/ban/api/grpc/v1"
	"backend/pkg/rpc"
)

func regUserService(gws *rpc.Gateways) {
	gws.AddGateway(
		accountGW.RegisterAccountPublicHandlerFromEndpoint,
		accountGW.ServiceAddr,
	)
	gws.AddGateway(
		banGW.RegisterWebHandlerFromEndpoint,
		banGW.ServiceAddr,
	)
}
