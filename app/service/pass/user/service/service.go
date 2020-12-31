package service

import (
	server "backend/app/game/server/api/grpc"
	backpack "backend/app/service/backpack/user/api/grpc"
	reward "backend/app/service/pass/reward/api/grpc"
	"backend/app/service/pass/user/conf"
	"backend/app/service/pass/user/dao"
	money "backend/app/service/user/money/api/grpc"
)

type Service struct {
	dao             dao.Dao
	rewardService   reward.RewardClient
	backpackService backpack.UserClient
	moneyService    money.MoneyClient
	serverService   server.ServerClient
}

func Init(config *conf.Config) *Service {
	return &Service{
		dao:             dao.Init(config),
		rewardService:   reward.InitClient(reward.ServiceAddr),
		backpackService: backpack.InitClient(backpack.ServiceAddr),
		moneyService:    money.InitClient(money.ServiceAddr),
		serverService:   server.InitClient(server.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
	reward.Close()
	backpack.Close()
	money.Close()
	server.Close()
}
