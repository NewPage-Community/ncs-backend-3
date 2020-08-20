package service

import (
	accountGW "backend/app/service/user/account/api/grpc"
	"backend/pkg/rpc"
)

func regUserService(gws *rpc.Gateways) {
	gws.AddGateway(
		accountGW.RegisterWebHandlerFromEndpoint,
		accountGW.ServiceAddr,
	)
}
