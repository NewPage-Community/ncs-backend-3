package service

import (
	pb "backend/app/game/user/api/grpc/v1"
	"backend/app/game/user/conf"
	backpack "backend/app/service/backpack/user/api/grpc/v1"
	account "backend/app/service/user/account/api/grpc/v1"
	"backend/pkg/steam"
)

type Service struct {
	account  account.AccountClient
	backpack backpack.UserClient
	steam    steam.APIClient
	pb.UnimplementedGameServer
}

func Init(c *conf.Config) *Service {
	return &Service{
		account:  account.InitClient(account.ServiceAddr),
		backpack: backpack.InitClient(backpack.ServiceAddr),
		steam:    c.Steam,
	}
}

func (s *Service) Close() {
	account.Close()
	backpack.Close()
}
