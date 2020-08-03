package service

import (
	backpack "backend/app/service/backpack/user/api/grpc"
	"backend/app/service/user/account/api/grpc"
)

type Service struct {
	account  grpc.AccountClient
	backpack backpack.UserClient
}

func Init() *Service {
	return &Service{
		account:  grpc.InitClient(grpc.ServiceAddr),
		backpack: backpack.InitClient(grpc.ServiceAddr),
	}
}

func (s *Service) Close() {
	grpc.Close()
	backpack.Close()
}
