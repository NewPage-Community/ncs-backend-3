package service

import (
	pb "backend/app/game/user/api/grpc/v1"
	backpack "backend/app/service/backpack/user/api/grpc/v1"
	account "backend/app/service/user/account/api/grpc/v1"
)

type Service struct {
	account  account.AccountClient
	backpack backpack.UserClient
	pb.UnimplementedGameServer
}

func Init() *Service {
	return &Service{
		account:  account.InitClient(account.ServiceAddr),
		backpack: backpack.InitClient(backpack.ServiceAddr),
	}
}

func (s *Service) Close() {
	account.Close()
	backpack.Close()
}
