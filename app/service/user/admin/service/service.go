package service

import (
	pb "backend/app/service/user/admin/api/grpc/v1"
	"backend/app/service/user/admin/conf"
	"backend/app/service/user/admin/dao"
)

type Service struct {
	dao dao.Dao
	pb.UnimplementedAdminServer
}

func Init(c *conf.Config) *Service {
	return &Service{
		dao: dao.New(c),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
}
