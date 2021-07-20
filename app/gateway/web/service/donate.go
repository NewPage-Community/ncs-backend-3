package service

import (
	donateGW "backend/app/service/donate/api/grpc/v1"
	"backend/pkg/rpc"
)

func regDonateService(gws *rpc.Gateways) {
	gws.AddGateway(
		donateGW.RegisterDonateHandlerFromEndpoint,
		donateGW.ServiceAddr,
	)
}
