package service

import (
	"backend/app/game/chat/conf"
	server "backend/app/game/server/api/grpc"
)

type Service struct {
	server server.ServerClient
}

func Init(config *conf.Config) *Service {
	return &Service{
		server: server.InitClient(server.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return true
}

func (s *Service) Close() {
	server.Close()
}
