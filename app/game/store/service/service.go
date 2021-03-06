package service

import (
	"backend/app/game/store/conf"
	"backend/app/game/store/model"
	itemsService "backend/app/service/backpack/items/api/grpc"
	userService "backend/app/service/backpack/user/api/grpc"
	passService "backend/app/service/pass/user/api/grpc"
	moneyService "backend/app/service/user/money/api/grpc"
	vipService "backend/app/service/user/vip/api/grpc"
)

type Service struct {
	items   itemsService.ItemsClient
	user    userService.UserClient
	money   moneyService.MoneyClient
	pass    passService.UserClient
	vip     vipService.VIPClient
	hotSale model.HotSale
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
