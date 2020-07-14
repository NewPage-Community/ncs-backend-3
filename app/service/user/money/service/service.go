package service

import (
	"backend/app/service/user/money/conf"
	"backend/app/service/user/money/dao"
)

type Service struct {
	dao dao.Dao
}

func Init(config *conf.Config) *Service {
	return &Service{
		dao: dao.Init(config),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
}
