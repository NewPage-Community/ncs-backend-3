package service

import (
	serverService "backend/app/game/server/api/grpc"
	"backend/app/service/user/ban/conf"
	"backend/app/service/user/ban/dao"
)

type Service struct {
	dao    dao.Dao
	server serverService.ServerClient
}

func Init(c *conf.Config) *Service {
	return &Service{
		dao:    dao.New(c),
		server: serverService.InitClient(serverService.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
	serverService.Close()
}
