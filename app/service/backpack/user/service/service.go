package service

import (
	itemsSrv "backend/app/service/backpack/items/api/grpc/v1"
	pb "backend/app/service/backpack/user/api/grpc/v1"
	"backend/app/service/backpack/user/conf"
	"backend/app/service/backpack/user/dao"
	"backend/pkg/jwt"
)

type Service struct {
	dao   dao.Dao
	items itemsSrv.ItemsClient
	jwt   *jwt.JWT
	pb.UnimplementedUserServer
	pb.UnimplementedUserPublicServer
}

func Init(config *conf.Config) *Service {
	return &Service{
		dao:   dao.Init(config),
		items: itemsSrv.InitClient(itemsSrv.ServiceAddr),
		jwt:   config.JWT,
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
	itemsSrv.Close()
}
