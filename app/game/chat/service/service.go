package service

import (
	pb "backend/app/game/chat/api/grpc/v1"
	"backend/app/game/chat/conf"
	"backend/app/game/chat/dao"
	server "backend/app/game/server/api/grpc/v1"
	"backend/pkg/log"
)

type Service struct {
	server server.ServerClient
	dao    dao.Dao
	pb.UnimplementedChatServer
}

func Init(config *conf.Config, service string) (s *Service) {
	s = &Service{
		server: server.InitClient(server.ServiceAddr),
		dao:    dao.Init(config, service),
	}
	log.CheckErr(s.dao.ListenAllChatEvent(s.AllChatEvent))
	return
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	server.Close()
	s.dao.Close()
}
