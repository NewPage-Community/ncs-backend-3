package service

import (
	"backend/app/service/user/ban/conf"
	"backend/app/service/user/ban/dao"
)

type Service struct {
	dao dao.Dao
}

func Init(c *conf.Config) *Service {
	return &Service{
		dao: dao.New(c),
	}
}

func (s *Service) Close() {
	s.dao.Close()
}
