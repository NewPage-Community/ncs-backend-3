package service

import (
	accountGW "backend/app/user/account/api/grpc"
	adminGW "backend/app/user/admin/api/grpc"
	banGW "backend/app/user/ban/api/grpc"
	signGW "backend/app/user/sign/api/grpc"
	vipGW "backend/app/user/vip/api/grpc"
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
}
