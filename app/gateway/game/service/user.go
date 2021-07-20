package service

import (
	accountGW "backend/app/service/user/account/api/grpc/v1"
	adminGW "backend/app/service/user/admin/api/grpc/v1"
	banGW "backend/app/service/user/ban/api/grpc/v1"
	moneyGW "backend/app/service/user/money/api/grpc/v1"
	signGW "backend/app/service/user/sign/api/grpc/v1"
	titleGW "backend/app/service/user/title/api/grpc/v1"
	vipGW "backend/app/service/user/vip/api/grpc/v1"
	"backend/pkg/rpc"
)

func regUserService(gws *rpc.Gateways) {
	gws.AddGateway(
		accountGW.RegisterAccountHandlerFromEndpoint,
		accountGW.ServiceAddr,
	)
	gws.AddGateway(
		adminGW.RegisterAdminHandlerFromEndpoint,
		adminGW.ServiceAddr,
	)
	gws.AddGateway(
		banGW.RegisterBanHandlerFromEndpoint,
		banGW.ServiceAddr,
	)
	gws.AddGateway(
		signGW.RegisterSignHandlerFromEndpoint,
		signGW.ServiceAddr,
	)
	gws.AddGateway(
		vipGW.RegisterVIPHandlerFromEndpoint,
		vipGW.ServiceAddr,
	)
	gws.AddGateway(
		titleGW.RegisterTitleHandlerFromEndpoint,
		titleGW.ServiceAddr,
	)
	gws.AddGateway(
		moneyGW.RegisterMoneyHandlerFromEndpoint,
		moneyGW.ServiceAddr,
	)
}
