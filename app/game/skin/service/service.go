package service

import (
	pb "backend/app/game/skin/api/grpc/v1"
	"backend/app/game/skin/conf"
	"backend/app/game/skin/dao"
	itemsService "backend/app/service/backpack/items/api/grpc/v1"
	userService "backend/app/service/backpack/user/api/grpc/v1"
)

type Service struct {
	dao  dao.Dao
	item itemsService.ItemsClient
	user userService.UserClient
	pb.UnimplementedSkinServer
}

func Init(config *conf.Config) *Service {
	return &Service{
		dao:  dao.Init(config),
		item: itemsService.InitClient(itemsService.ServiceAddr),
		user: userService.InitClient(userService.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
	itemsService.Close()
	userService.Close()
}
