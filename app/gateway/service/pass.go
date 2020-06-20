package service

import (
	userGW "backend/app/service/pass/user/api/grpc"
	"backend/pkg/rpc"
)

func regPassService(gws *rpc.Gateways) {
	gws.AddGateway(
		userGW.RegisterUserHandlerFromEndpoint,
		userGW.ServiceAddr,
	)
}
