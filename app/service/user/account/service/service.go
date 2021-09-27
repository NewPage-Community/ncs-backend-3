package service

import (
	pb "backend/app/service/user/account/api/grpc/v1"
	"backend/app/service/user/account/conf"
	"backend/app/service/user/account/dao"
)

type Service struct {
	dao dao.Dao
	pb.UnimplementedAccountServer
	pb.UnimplementedAccountPublicServer
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
