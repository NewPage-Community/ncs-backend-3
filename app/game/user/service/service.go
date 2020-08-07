package service

import (
	backpack "backend/app/service/backpack/user/api/grpc"
	account "backend/app/service/user/account/api/grpc"
)

type Service struct {
	account  account.AccountClient
	backpack backpack.UserClient
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
