package service

import (
	"backend/app/service/user/admin/conf"
	"backend/app/service/user/admin/dao"
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
