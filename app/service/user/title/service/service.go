package service

import (
	pb "backend/app/service/user/title/api/grpc/v1"
	"backend/app/service/user/title/conf"
	"backend/app/service/user/title/dao"
)

type Service struct {
	dao dao.Dao
	pb.UnimplementedTitleServer
}

func Init(config *conf.Config) *Service {
	return &Service{
		dao: dao.Init(config),
	}
}

func (s *Service) Close() {
	s.dao.Close()
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}
