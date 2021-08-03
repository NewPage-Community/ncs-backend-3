package service

import (
	pb "backend/app/game/cookie/api/grpc/v1"
	"backend/app/game/cookie/conf"
	"backend/app/game/cookie/dao"
	serverSrv "backend/app/game/server/api/grpc/v1"
	"backend/pkg/jwt"
)

type Service struct {
	dao dao.Dao
	jwt *jwt.JWT
	pb.UnimplementedCookieServer
	server serverSrv.ServerClient
}

func Init(config *conf.Config) *Service {
	return &Service{
		dao:    dao.Init(config),
		jwt:    config.JWT,
		server: serverSrv.InitClient(serverSrv.ServiceAddr),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	s.dao.Close()
	serverSrv.Close()
}
