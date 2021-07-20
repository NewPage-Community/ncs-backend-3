package service

import (
	kaiheilaBot "backend/app/bot/kaiheila/api/grpc/v1"
	qqBot "backend/app/bot/qq/api/grpc/v1"
	pb "backend/app/game/chat/api/grpc/v1"
	"backend/app/game/chat/conf"
	server "backend/app/game/server/api/grpc/v1"
)

type Service struct {
	server   server.ServerClient
	kaiheila kaiheilaBot.KaiheilaClient
	qq       qqBot.QQClient
	pb.UnimplementedChatServer
}

func Init(config *conf.Config) *Service {
	return &Service{
		server:   server.InitClient(server.ServiceAddr),
		kaiheila: kaiheilaBot.InitClient(kaiheilaBot.ServiceAddr),
		qq:       qqBot.InitClient(qqBot.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return true
}

func (s *Service) Close() {
	server.Close()
	kaiheilaBot.Close()
	qqBot.Close()
}
