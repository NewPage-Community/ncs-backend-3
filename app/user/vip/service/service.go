package service

import "backend/app/user/vip/conf"

type Service struct {
}

func Init(c *conf.Config) *Service {
	return &Service{}
}

func (s *Service) Close() {

}
