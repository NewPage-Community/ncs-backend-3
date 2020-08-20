package service

import (
	rewardGW "backend/app/service/pass/reward/api/grpc"
	userGW "backend/app/service/pass/user/api/grpc"
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
