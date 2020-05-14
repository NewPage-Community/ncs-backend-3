package service

import (
	"backend/app/user/vip/conf"
	"backend/app/user/vip/dao"
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
