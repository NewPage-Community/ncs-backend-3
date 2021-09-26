package service

import (
	kaiheilaBot "backend/app/bot/kaiheila/api/grpc/v1"
	qqBot "backend/app/bot/qq/api/grpc/v1"
	pb "backend/app/game/chat/api/grpc/v1"
	"backend/app/game/chat/conf"
	"backend/app/game/chat/dao"
	server "backend/app/game/server/api/grpc/v1"
)

type Service struct {
	server   server.ServerClient
	kaiheila kaiheilaBot.KaiheilaClient
	qq       qqBot.QQClient
	pb.UnimplementedChatServer
	dao dao.Dao
}

func Init(config *conf.Config, service string) (s *Service) {
	s = &Service{
		server:   server.InitClient(server.ServiceAddr),
		kaiheila: kaiheilaBot.InitClient(kaiheilaBot.ServiceAddr),
		qq:       qqBot.InitClient(qqBot.ServiceAddr),
		dao:      dao.Init(config, service),
	}
	return
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	server.Close()
	kaiheilaBot.Close()
	qqBot.Close()
	s.dao.Close()
}
