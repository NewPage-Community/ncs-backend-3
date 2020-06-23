package service

import "backend/app/service/pass/reward/conf"

type Service struct {
	config *conf.Config
}

func Init(config *conf.Config) *Service {
	return &Service{
		config: config,
	}
}

func (s *Service) Healthy() bool {
	return true
}

func (s *Service) Close() {
}
