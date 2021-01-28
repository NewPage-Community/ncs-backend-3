package service

import (
	kaiheilaBot "backend/app/bot/kaiheila/api/grpc/v1"
	"backend/app/game/chat/conf"
	server "backend/app/game/server/api/grpc"
)

type Service struct {
	server   server.ServerClient
	kaiheila kaiheilaBot.KaiheilaClient
}

func Init(config *conf.Config) *Service {
	return &Service{
		server:   server.InitClient(server.ServiceAddr),
		kaiheila: kaiheilaBot.InitClient(kaiheilaBot.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return true
}

func (s *Service) Close() {
	server.Close()
	kaiheilaBot.Close()
}
