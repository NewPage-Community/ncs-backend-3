package service

import (
	qqBot "backend/app/bot/qq/api/grpc/v1"
	a2sSrv "backend/app/game/a2s/api/grpc/v1"
	pb "backend/app/game/server/api/grpc/v1"
	"backend/app/game/server/conf"
	"backend/app/game/server/dao"
)

type Service struct {
	dao dao.Dao
	qq  qqBot.QQClient
	a2s a2sSrv.A2SClient
	pb.UnimplementedServerServer
}

func Init(c *conf.Config, service string) *Service {
	return &Service{
		dao: dao.New(c, service),
		qq:  qqBot.InitClient(qqBot.ServiceAddr),
		a2s: a2sSrv.InitClient(a2sSrv.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
	qqBot.Close()
	a2sSrv.Close()
}
