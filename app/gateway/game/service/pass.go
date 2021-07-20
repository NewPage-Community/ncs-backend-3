package service

import (
	rewardGW "backend/app/service/pass/reward/api/grpc/v1"
	userGW "backend/app/service/pass/user/api/grpc/v1"
	"backend/pkg/rpc"
)

func regPassService(gws *rpc.Gateways) {
	gws.AddGateway(
		userGW.RegisterUserHandlerFromEndpoint,
		userGW.ServiceAddr,
	)
	gws.AddGateway(
		rewardGW.RegisterRewardHandlerFromEndpoint,
		rewardGW.ServiceAddr,
	)
}
