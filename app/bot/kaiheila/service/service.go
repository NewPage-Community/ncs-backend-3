package service

import (
	pb "backend/app/bot/kaiheila/api/grpc/v1"
	"backend/app/bot/kaiheila/conf"
	"backend/app/bot/kaiheila/dao"
	"backend/pkg/log"
	khlLogger "backend/pkg/log/kaiheila"

	"github.com/lonelyevil/khl"
)

// Service 服务结构定义
type Service struct {
	kaiheilaClient *khl.Session
	dao            dao.Dao
	kaiheilaConfig *conf.Kaiheila
	pb.UnimplementedKaiheilaServer
}

// Init 服务初始化
func Init(config *conf.Config, service string) *Service {
	srv := &Service{
		kaiheilaClient: khl.New(config.Kaiheila.Token, khlLogger.NewLogger(log.GetLogger())),
		dao:            dao.Init(config, service),
		kaiheilaConfig: config.Kaiheila,
	}
	srv.kaiheilaClient.AddHandler(srv.EventHandler)
	log.CheckErr(srv.kaiheilaClient.Open())
	log.CheckErr(srv.dao.ListenAllChatEvent(srv.AllChatEvent))
	log.CheckErr(srv.dao.ListenChangeMapEvent(srv.ChangeMapEvent))
	log.CheckErr(srv.dao.ListenDonateEvent(srv.DonateEvent))
	return srv
}

// Healthy 服务健康检查
func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

// Close 服务关闭
func (s *Service) Close() {
	_ = s.kaiheilaClient.Close()
	s.dao.Close()
}
