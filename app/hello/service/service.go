package service

import (
	"backend/app/hello/conf"
	"backend/app/hello/dao"
)

// Service 服务结构定义
type Service struct {
	dao dao.Dao
}

// Init 服务初始化
func Init(config *conf.Config) *Service {
	return &Service{
		dao: dao.Init(config),
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
