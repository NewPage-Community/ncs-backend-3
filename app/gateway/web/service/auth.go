package service

import (
	qqGW "backend/app/service/auth/qq/api/grpc/v1"
	steamGW "backend/app/service/auth/steam/api/grpc/v1"
	"backend/pkg/rpc"
)

func regAuthService(gws *rpc.Gateways) {
	gws.AddGateway(
		qqGW.RegisterWebHandlerFromEndpoint,
		qqGW.ServiceAddr,
	)
	gws.AddGateway(
		steamGW.RegisterWebHandlerFromEndpoint,
		steamGW.ServiceAddr,
	)
}
