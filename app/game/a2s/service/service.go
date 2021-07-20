package service

import (
	pb "backend/app/game/a2s/api/grpc/v1"
	"backend/app/game/a2s/conf"
)

// Service 服务结构定义
type Service struct {
	pb.UnimplementedA2SServer
}

// Init 服务初始化
func Init(config *conf.Config) *Service {
	return &Service{}
}

// Healthy 服务健康检查
func (s *Service) Healthy() bool {
	return true
}

// Close 服务关闭
func (s *Service) Close() {
}
