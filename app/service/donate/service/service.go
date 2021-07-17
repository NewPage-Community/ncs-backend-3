package service

import (
	"backend/app/service/donate/conf"
	"backend/app/service/donate/dao"
	accountService "backend/app/service/user/account/api/grpc"
	moneyService "backend/app/service/user/money/api/grpc"
	"github.com/robfig/cron/v3"
)

type Service struct {
	account accountService.AccountClient
	money   moneyService.MoneyClient
	alipay  Payment
	wepay   Payment
	dao     dao.Dao
	cron    *cron.Cron
}

func Init(config *conf.Config) *Service {
	s := &Service{
		account: accountService.InitClient(accountService.ServiceAddr),
		money:   moneyService.InitClient(moneyService.ServiceAddr),
		alipay:  InitAlipay(config.Alipay),
		wepay:   InitWepay(config.Wepay),
		dao:     dao.Init(config),
		cron:    cron.New(),
	}
	s.regCron()
	s.cron.Start()
	return s
}

func (s *Service) Healthy() bool {
	return s.dao.Healthy()
}

func (s *Service) Close() {
	accountService.Close()
	moneyService.Close()
	s.dao.Close()
}
