package service

import (
	backpack "backend/app/service/backpack/user/api/grpc"
	reward "backend/app/service/pass/reward/api/grpc"
	"backend/app/service/pass/user/conf"
	"backend/app/service/pass/user/dao"
)

type Service struct {
	dao             dao.Dao
	rewardService   reward.RewardClient
	backpackService backpack.UserClient
}

func Init(config *conf.Config) *Service {
	return &Service{
		dao:             dao.Init(config),
		rewardService:   reward.InitClient(reward.ServiceAddr),
		backpackService: backpack.InitClient(backpack.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
	reward.Close()
	backpack.Close()
}
