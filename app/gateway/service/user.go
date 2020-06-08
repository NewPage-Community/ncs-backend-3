package service

import (
	accountGW "backend/app/user/account/api/grpc"
	adminGW "backend/app/user/admin/api/grpc"
	banGW "backend/app/user/ban/api/grpc"
	signGW "backend/app/user/sign/api/grpc"
	vipGW "backend/app/user/vip/api/grpc"
)

func (s *Service) regUserService() (err error) {
	err = accountGW.RegisterAccountHandlerFromEndpoint(s.ctx, s.mux, accountGW.ServiceName, s.opts)
	err = adminGW.RegisterAdminHandlerFromEndpoint(s.ctx, s.mux, adminGW.ServiceName, s.opts)
	err = banGW.RegisterBanHandlerFromEndpoint(s.ctx, s.mux, banGW.ServiceName, s.opts)
	err = signGW.RegisterSignHandlerFromEndpoint(s.ctx, s.mux, signGW.ServiceName, s.opts)
	err = vipGW.RegisterVIPHandlerFromEndpoint(s.ctx, s.mux, vipGW.ServiceName, s.opts)
	return
}
