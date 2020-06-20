package service

import (
	"backend/app/service/backpack/user/conf"
	"backend/app/service/backpack/user/dao"
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
	return s.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
}
