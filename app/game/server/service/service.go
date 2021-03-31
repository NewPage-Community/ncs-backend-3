package service

import (
	qqBot "backend/app/bot/qq/api/grpc/v1"
	"backend/app/game/server/conf"
	"backend/app/game/server/dao"
)

type Service struct {
	dao dao.Dao
	qq  qqBot.QQClient
}

func Init(c *conf.Config) *Service {
	return &Service{
		dao: dao.New(c),
		qq:  qqBot.InitClient(qqBot.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
	qqBot.Close()
}
