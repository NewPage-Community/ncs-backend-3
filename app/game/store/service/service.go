package service

import (
	pb "backend/app/game/store/api/grpc/v1"
	"backend/app/game/store/conf"
	"backend/app/game/store/model"
	itemsService "backend/app/service/backpack/items/api/grpc/v1"
	userService "backend/app/service/backpack/user/api/grpc/v1"
	passService "backend/app/service/pass/user/api/grpc/v1"
	moneyService "backend/app/service/user/money/api/grpc/v1"
	vipService "backend/app/service/user/vip/api/grpc/v1"
)

type Service struct {
	items   itemsService.ItemsClient
	user    userService.UserClient
	money   moneyService.MoneyClient
	pass    passService.UserClient
	vip     vipService.VIPClient
	hotSale model.HotSale
	pb.UnimplementedStoreServer
}

func Init(config *conf.Config) *Service {
	return &Service{
		items:   itemsService.InitClient(itemsService.ServiceAddr),
		user:    userService.InitClient(userService.ServiceAddr),
		money:   moneyService.InitClient(moneyService.ServiceAddr),
		pass:    passService.InitClient(passService.ServiceAddr),
		vip:     vipService.InitClient(vipService.ServiceAddr),
		hotSale: *config.HotSale,
	}
}

func (s *Service) Healthy() bool {
	return true
}

func (s *Service) Close() {
	itemsService.Close()
	userService.Close()
	moneyService.Close()
	passService.Close()
	vipService.Close()
}
