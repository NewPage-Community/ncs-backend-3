package service

import (
	accountGW "backend/app/user/account/api/grpc"
	adminGW "backend/app/user/admin/api/grpc"
	banGW "backend/app/user/ban/api/grpc"
	signGW "backend/app/user/sign/api/grpc"
	vipGW "backend/app/user/vip/api/grpc"
)

func (s *Service) regUserService() (err error) {
	err = accountGW.RegisterAccountHandlerFromEndpoint(s.ctx, s.mux, accountGW.ServiceAddr, s.opts)
	err = adminGW.RegisterAdminHandlerFromEndpoint(s.ctx, s.mux, adminGW.ServiceAddr, s.opts)
	err = banGW.RegisterBanHandlerFromEndpoint(s.ctx, s.mux, banGW.ServiceAddr, s.opts)
	err = signGW.RegisterSignHandlerFromEndpoint(s.ctx, s.mux, signGW.ServiceAddr, s.opts)
	err = vipGW.RegisterVIPHandlerFromEndpoint(s.ctx, s.mux, vipGW.ServiceAddr, s.opts)
	return
}
