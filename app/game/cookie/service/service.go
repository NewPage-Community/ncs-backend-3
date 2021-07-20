package service

import (
	pb "backend/app/game/cookie/api/grpc/v1"
	"backend/app/game/cookie/conf"
	"backend/app/game/cookie/dao"
)

type Service struct {
	dao dao.Dao
	pb.UnimplementedCookieServer
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
