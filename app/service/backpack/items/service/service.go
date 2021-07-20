package service

import (
	pb "backend/app/service/backpack/items/api/grpc/v1"
	"backend/app/service/backpack/items/conf"
	"backend/app/service/backpack/items/dao"
)

type Service struct {
	dao dao.Dao
	pb.UnimplementedItemsServer
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
