package service

import (
	pb "backend/app/service/user/sign/api/grpc/v1"
	"backend/app/service/user/sign/conf"
	"backend/app/service/user/sign/dao"
)

type Service struct {
	dao dao.Dao
	pb.UnimplementedSignServer
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
