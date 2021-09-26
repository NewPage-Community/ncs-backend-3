package service

import (
	pb "backend/app/bot/qq/api/grpc/v1"
	"backend/app/bot/qq/conf"
	chatService "backend/app/game/chat/api/grpc/v1"
	serverService "backend/app/game/server/api/grpc/v1"

	"github.com/miRemid/amy"
	"github.com/miRemid/amy/websocket/model"
)

// Service 服务结构定义
type Service struct {
	qqWSClient  *MsgClient
	qqAPIClient *amy.API
	serverSrv   serverService.ServerClient
	chatSrv     chatService.ChatClient
	qqGroups    []int
	pb.UnimplementedQQServer
}

// Init 服务初始化
func Init(config *conf.Config) *Service {
	srv := &Service{
		qqAPIClient: amy.NewAmyAPI(config.QQConfig.Address, config.QQConfig.APIPort),
		qqWSClient:  NewWSClient(config.QQConfig.Address, config.QQConfig.WSPort),
		serverSrv:   serverService.InitClient(serverService.ServiceAddr),
		chatSrv:     chatService.InitClient(chatService.ServiceAddr),
		qqGroups:    config.QQConfig.QQGroups,
	}
	srv.qqAPIClient.SetToken(config.QQConfig.Token)
	srv.qqWSClient.SetToken(config.QQConfig.Token)
	srv.qqWSClient.OnMessage(func(event model.CQEvent) {
		srv.OnMessage(event)
	})
	go srv.qqWSClient.Run()
	return srv
}

// Healthy 服务健康检查
func (s *Service) Healthy() bool {
	return true
}

// Close 服务关闭
func (s *Service) Close() {
	serverService.Close()
	chatService.Close()
}
