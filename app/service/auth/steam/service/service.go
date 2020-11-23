package service

import (
	"backend/app/service/auth/steam/conf"
	account "backend/app/service/user/account/api/grpc"
)

// Service 服务结构定义
type Service struct {
	config         *conf.Config
	accountService account.AccountClient
}

// Init 服务初始化
func Init(config *conf.Config) *Service {
	return &Service{
		config:         config,
		accountService: account.InitClient(account.ServiceAddr),
	}
}

// Healthy 服务健康检查
func (s *Service) Healthy() bool {
	return true
}

// Close 服务关闭
func (s *Service) Close() {
	account.Close()
}
