package service

import (
	pb "backend/app/service/auth/steam/api/grpc/v1"
	"backend/app/service/auth/steam/conf"
	accountSrv "backend/app/service/user/account/api/grpc/v1"
)

// Service 服务结构定义
type Service struct {
	config  *conf.Config
	account accountSrv.AccountClient
	pb.UnimplementedWebServer
}

// Init 服务初始化
func Init(config *conf.Config) *Service {
	return &Service{
		config:  config,
		account: accountSrv.InitClient(accountSrv.ServiceAddr),
	}
}

// Healthy 服务健康检查
func (s *Service) Healthy() bool {
	return true
}

// Close 服务关闭
func (s *Service) Close() {
	accountSrv.Close()
}
