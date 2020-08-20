package service

import (
	donateGW "backend/app/service/donate/api/grpc"
	"backend/pkg/rpc"
)

func regDonateService(gws *rpc.Gateways) {
	gws.AddGateway(
		donateGW.RegisterDonateHandlerFromEndpoint,
		donateGW.ServiceAddr,
	)
}
