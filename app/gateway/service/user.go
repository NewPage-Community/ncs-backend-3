package service

import (
	accountGW "backend/app/user/account/api/grpc"
	adminGW "backend/app/user/admin/api/grpc"
	banGW "backend/app/user/ban/api/grpc"
	signGW "backend/app/user/sign/api/grpc"
	vipGW "backend/app/user/vip/api/grpc"
)

func regUserService() []gateway {
	return []gateway{
		{
			handlder: accountGW.RegisterAccountHandlerFromEndpoint,
			endpoint: accountGW.ServiceAddr,
		},
		{
			handlder: adminGW.RegisterAdminHandlerFromEndpoint,
			endpoint: adminGW.ServiceAddr,
		},
		{
			handlder: banGW.RegisterBanHandlerFromEndpoint,
			endpoint: banGW.ServiceAddr,
		},
		{
			handlder: signGW.RegisterSignHandlerFromEndpoint,
			endpoint: signGW.ServiceAddr,
		},
		{
			handlder: vipGW.RegisterVIPHandlerFromEndpoint,
			endpoint: vipGW.ServiceAddr,
		},
	}
}
