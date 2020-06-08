package service

import serverGW "backend/app/server/api/grpc"

func (s *Service) regServerService() (err error) {
	err = serverGW.RegisterServerHandlerFromEndpoint(s.ctx, s.mux, serverGW.ServiceAddr, s.opts)
	return
}
