package service

import (
	accountGW "backend/app/service/user/account/api/grpc"
	adminGW "backend/app/service/user/admin/api/grpc"
	banGW "backend/app/service/user/ban/api/grpc"
	signGW "backend/app/service/user/sign/api/grpc"
	titleGW "backend/app/service/user/title/api/grpc"
	vipGW "backend/app/service/user/vip/api/grpc"
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
}
