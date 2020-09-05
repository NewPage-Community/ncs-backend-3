package service

import (
	serverService "backend/app/game/server/api/grpc"
	accountService "backend/app/service/user/account/api/grpc"
	"backend/app/service/user/ban/conf"
	"backend/app/service/user/ban/dao"
	"backend/pkg/steam"
)

type Service struct {
	dao     dao.Dao
	server  serverService.ServerClient
	steam   steam.APIClient
	account accountService.AccountClient
}

func Init(c *conf.Config) *Service {
	return &Service{
		dao:     dao.New(c),
		server:  serverService.InitClient(serverService.ServiceAddr),
		steam:   steam.NewAPIClient(c.Steam.APIUrl, c.Steam.APIKey, c.Steam.Timeout),
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
