package service

import (
	pb "backend/app/service/user/vip/api/grpc/v1"
	"backend/app/service/user/vip/conf"
	"backend/app/service/user/vip/dao"
)

type Service struct {
	dao dao.Dao
	pb.UnimplementedVIPServer
}

func Init(c *conf.Config) *Service {
	return &Service{
		dao: dao.New(c),
	}
}

func (s *Service) Close() {
	s.dao.Close()
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}
