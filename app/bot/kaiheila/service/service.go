package service

import (
	pb "backend/app/bot/kaiheila/api/grpc/v1"
	"backend/app/bot/kaiheila/conf"
	chatService "backend/app/game/chat/api/grpc/v1"
	"github.com/gunslinger23/kaiheila"
)

// Service 服务结构定义
type Service struct {
	kaiheilaClient *kaiheila.Client
	kaiheilaWSS    *kaiheila.WebSocketSession
	chat           chatService.ChatClient
	pb.UnimplementedKaiheilaServer
}

// Init 服务初始化
func Init(config *conf.Config) *Service {
	srv := &Service{
		kaiheilaClient: kaiheila.NewClient("", kaiheila.TokenBot, config.Kaiheila.Token, 0),
		chat:           chatService.InitClient(chatService.ServiceAddr),
	}
	srv.kaiheilaWSS = srv.kaiheilaClient.WebSocketSession(srv.EventHandler)
	return srv
}

// Healthy 服务健康检查
func (s *Service) Healthy() bool {
	return true
}

// Close 服务关闭
func (s *Service) Close() {
	s.kaiheilaWSS.Close()
	chatService.Close()
}
