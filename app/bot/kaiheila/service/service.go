package service

import (
	pb "backend/app/bot/kaiheila/api/grpc/v1"
	"backend/app/bot/kaiheila/conf"
	"backend/app/bot/kaiheila/dao"
	"backend/pkg/log"

	"github.com/gunslinger23/kaiheila"
)

// Service 服务结构定义
type Service struct {
	kaiheilaClient *kaiheila.Client
	kaiheilaWSS    *kaiheila.WebSocketSession
	dao            dao.Dao
	kaiheilaConfig *conf.Kaiheila
	pb.UnimplementedKaiheilaServer
}

// Init 服务初始化
func Init(config *conf.Config, service string) *Service {
	srv := &Service{
		kaiheilaClient: kaiheila.NewClient("", kaiheila.TokenBot, config.Kaiheila.Token, 0),
		dao:            dao.Init(config, service),
		kaiheilaConfig: config.Kaiheila,
	}
	srv.kaiheilaWSS = srv.kaiheilaClient.WebSocketSession(srv.EventHandler)
	log.CheckErr(srv.dao.ListenAllChatEvent(srv.AllChatEvent))
	return srv
}

// Healthy 服务健康检查
func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

// Close 服务关闭
func (s *Service) Close() {
	s.kaiheilaWSS.Close()
	s.dao.Close()
}
