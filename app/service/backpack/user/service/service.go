package service

import (
	itemsSrv "backend/app/service/backpack/items/api/grpc"
	"backend/app/service/backpack/user/conf"
	"backend/app/service/backpack/user/dao"
)

type Service struct {
	dao   dao.Dao
	items itemsSrv.ItemsClient
}

func Init(config *conf.Config) *Service {
	return &Service{
		dao:   dao.Init(config),
		items: itemsSrv.InitClient(itemsSrv.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
	itemsSrv.Close()
}
