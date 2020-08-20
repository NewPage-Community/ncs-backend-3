package service

import (
	"backend/app/service/donate/conf"
	"backend/pkg/log"
	"github.com/smartwalle/alipay/v3"
	"time"
)

const (
	FilePath             = "/ncs/"
	AppPublicCertFile    = "appCertPublicKey.crt"
	AliPayRootCertFile   = "alipayRootCert.crt"
	AliPayPublicCertFile = "alipayCertPublicKey_RSA2.crt"
)

func InitAlipay(config *conf.Alipay) *alipay.Client {
	client, err := alipay.New(config.AppID, config.PrivateKey, config.IsProduction)
	if err != nil {
		log.Error(err)
	}
	// 加载应用公钥证书
	err = client.LoadAppPublicCertFromFile(FilePath + AppPublicCertFile)
	if err != nil {
		log.Error(err)
	}
	// 加载支付宝根证书
	err = client.LoadAliPayRootCertFromFile(FilePath + AliPayRootCertFile)
	if err != nil {
		log.Error(err)
	}
	// 加载支付宝公钥证书
	err = client.LoadAliPayPublicCertFromFile(FilePath + AliPayPublicCertFile)
	if err != nil {
		log.Error(err)
	}
	return client
}

func (s *Service) CheckTrade(outTradeNo string) {
	timeout := time.Now().Add(5 * time.Minute)
	go func() {
		// Trade timeout check
		for time.Now().Before(timeout) {
			// Wait...
			time.Sleep(time.Second * 5)

			// Query trade
			res, err := s.alipay.TradeQuery(alipay.TradeQuery{
				OutTradeNo: outTradeNo,
			})
			// Finish trade when it is success
			if err == nil {
				if res.IsSuccess() {
					// Finish it
					if err := s.FinishDonate(outTradeNo); err != nil {
						log.Error(err)
					}
					return
				}
			}
		}

		// Timeout to cancel trade
		_, err := s.alipay.TradeCancel(alipay.TradeCancel{
			OutTradeNo: outTradeNo,
		})
		if err != nil {
			log.Error(err)
		}
	}()
}
