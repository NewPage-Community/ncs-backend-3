package service

import (
	pb "backend/app/bot/qq/api/grpc/v1"
	"backend/app/bot/qq/conf"
	"backend/app/bot/qq/dao"
	serverService "backend/app/game/server/api/grpc/v1"
	"backend/pkg/log"

	"github.com/miRemid/amy"
	"github.com/miRemid/amy/websocket/model"
)

// Service 服务结构定义
type Service struct {
	qqWSClient  *MsgClient
	qqAPIClient *amy.API
	serverSrv   serverService.ServerClient
	qqGroups    []int
	dao         dao.Dao
	pb.UnimplementedQQServer
}

// Init 服务初始化
func Init(config *conf.Config, service string) *Service {
	srv := &Service{
		qqAPIClient: amy.NewAmyAPI(config.QQConfig.Address, config.QQConfig.APIPort),
		qqWSClient:  NewWSClient(config.QQConfig.Address, config.QQConfig.WSPort),
		serverSrv:   serverService.InitClient(serverService.ServiceAddr),
		qqGroups:    config.QQConfig.QQGroups,
		dao:         dao.Init(config, service),
	}
	srv.qqAPIClient.SetToken(config.QQConfig.Token)
	srv.qqWSClient.SetToken(config.QQConfig.Token)
	srv.qqWSClient.OnMessage(func(event model.CQEvent) {
		srv.OnMessage(event)
	})
	go srv.qqWSClient.Run()
	//srv.dao.ListenAllChatEvent(srv.AllChatEvent)
	log.CheckErr(srv.dao.ListenChangeMapEvent(srv.ChangeMapEvent))
	return srv
}

// Healthy 服务健康检查
func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

// Close 服务关闭
func (s *Service) Close() {
	serverService.Close()
	s.dao.Close()
}
