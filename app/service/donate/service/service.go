package service

import (
	"backend/app/service/donate/conf"
	"backend/app/service/donate/dao"
	accountService "backend/app/service/user/account/api/grpc"
	moneyService "backend/app/service/user/money/api/grpc"
	"github.com/smartwalle/alipay/v3"
)

type Service struct {
	account accountService.AccountClient
	money   moneyService.MoneyClient
	alipay  *alipay.Client
	dao     dao.Dao
}

func Init(config *conf.Config) *Service {
	return &Service{
		account: accountService.InitClient(accountService.ServiceAddr),
		money:   moneyService.InitClient(moneyService.ServiceAddr),
		alipay:  InitAlipay(config.Alipay),
		dao:     dao.Init(config),
	}
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	accountService.Close()
	moneyService.Close()
	s.dao.Close()
}
