package service

import (
	pb "backend/app/game/pass/api/grpc/v1"
	"backend/app/game/pass/conf"
	itemService "backend/app/service/backpack/items/api/grpc/v1"
	userItemService "backend/app/service/backpack/user/api/grpc/v1"
	rewardService "backend/app/service/pass/reward/api/grpc/v1"
	userPassService "backend/app/service/pass/user/api/grpc/v1"
)

type Service struct {
	item     itemService.ItemsClient
	reward   rewardService.RewardClient
	userItem userItemService.UserClient
	userPass userPassService.UserClient
	pb.UnimplementedPassServer
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
