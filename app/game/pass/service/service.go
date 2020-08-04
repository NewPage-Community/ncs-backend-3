package service

import (
	"backend/app/game/pass/conf"
	itemService "backend/app/service/backpack/items/api/grpc"
	userItemService "backend/app/service/backpack/user/api/grpc"
	rewardService "backend/app/service/pass/reward/api/grpc"
	userPassService "backend/app/service/pass/user/api/grpc"
)

type Service struct {
	item     itemService.ItemsClient
	reward   rewardService.RewardClient
	userItem userItemService.UserClient
	userPass userPassService.UserClient
}

func Init(config *conf.Config) *Service {
	return &Service{
		item:     itemService.InitClient(itemService.ServiceAddr),
		reward:   rewardService.InitClient(rewardService.ServiceAddr),
		userItem: userItemService.InitClient(userItemService.ServiceAddr),
		userPass: userPassService.InitClient(userPassService.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return true
}

func (s *Service) Close() {
	itemService.Close()
	rewardService.Close()
	userItemService.Close()
	userPassService.Close()
}
