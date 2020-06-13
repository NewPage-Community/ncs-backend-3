package service

import "backend/app/service/user/account/api/grpc"

type Service struct {
	account grpc.AccountClient
}

func Init() *Service {
	return &Service{
		account: grpc.InitClient(grpc.ServiceAddr),
	}
}

func (s *Service) Close() {
	grpc.Close()
}
