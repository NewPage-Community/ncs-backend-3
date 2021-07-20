package service

import (
	pb "backend/app/service/auth/qq/api/grpc/v1"
	"backend/app/service/auth/qq/conf"
	"backend/app/service/auth/qq/dao"
)

// Service 服务结构定义
type Service struct {
	config *conf.Config
	dao    dao.Dao
	pb.UnimplementedWebServer
}

// Init 服务初始化
func Init(config *conf.Config) *Service {
	return &Service{
		config: config,
		dao:    dao.Init(config),
	}
}

// Healthy 服务健康检查
func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

// Close 服务关闭
func (s *Service) Close() {
	s.dao.Close()
}
