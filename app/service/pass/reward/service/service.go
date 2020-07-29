package service

import (
	"backend/app/service/pass/reward/conf"
	"backend/app/service/pass/reward/model"
)

type Service struct {
	reward *model.Reward
}

func Init(config *conf.Config) *Service {
	return &Service{
		reward: config.Reward,
	}
}

func (s *Service) Healthy() bool {
	return true
}

func (s *Service) Close() {
}
