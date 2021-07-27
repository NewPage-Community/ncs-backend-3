package service

import (
	serverService "backend/app/game/server/api/grpc/v1"
	accountService "backend/app/service/user/account/api/grpc/v1"
	pb "backend/app/service/user/ban/api/grpc/v1"
	"backend/app/service/user/ban/conf"
	"backend/app/service/user/ban/dao"
	"backend/pkg/steam"
)

type Service struct {
	dao     dao.Dao
	server  serverService.ServerClient
	steam   steam.APIClient
	account accountService.AccountClient
	pb.UnimplementedBanServer
	pb.UnimplementedWebServer
}

func Init(c *conf.Config) *Service {
	return &Service{
		dao:     dao.New(c),
		server:  serverService.InitClient(serverService.ServiceAddr),
		steam:   c.Steam,
		account: accountService.InitClient(accountService.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
	serverService.Close()
	accountService.Close()
}
