package service

import (
	"backend/app/game/cvar/conf"
	"backend/app/game/cvar/dao"
	serverService "backend/app/game/server/api/grpc"
	"github.com/robfig/cron/v3"
)

type Service struct {
	dao    dao.Dao
	cron   *cron.Cron
	server serverService.ServerClient
}

func Init(config *conf.Config) (s *Service) {
	s = &Service{
		dao:    dao.Init(config),
		cron:   cron.New(),
		server: serverService.InitClient(serverService.ServiceAddr),
	}
	s.regCron()
	return
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.cron.Stop()
	s.dao.Close()
	serverService.Close()
}
