package service

import (
	"backend/app/service/donate/conf"
	"backend/pkg/log"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"strconv"
)

type Alipay struct {
	client *alipay.Client
}

func InitAlipay(config *conf.Alipay) *Alipay {
	client, err := alipay.New(config.AppID, config.PrivateKey, config.IsProduction)
	if err != nil {
		panic(err)
	}
	// 加载支付宝公钥
	err = client.LoadAliPayPublicKey(config.AlipayPublicKey)
	if err != nil {
		panic(err)
	}
	return &Alipay{client}
}

func (s *Alipay) CheckTrade(outTradeNo string) (bool, error) {
	// Query trade
	res, err := s.client.TradeQuery(alipay.TradeQuery{
		OutTradeNo: outTradeNo,
	})
	// Finish trade when it is success
	if err == nil {
		if res.Content.TradeStatus == alipay.TradeStatusSuccess {
			return true, nil
		}
	}

	return false, err
}

func (s *Alipay) CreateTrade(outTradeNo string, steamID int64, totalAmount int32) (qrCode string, err error) {
	res, err := s.client.TradePreCreate(alipay.TradePreCreate{
		Trade: alipay.Trade{
			Subject:        TradeSubject,
			OutTradeNo:     outTradeNo,
			TotalAmount:    strconv.FormatInt(int64(totalAmount), 10),
			Body:           fmt.Sprintf(TradeBody, steamID),
			TimeoutExpress: Timeout,
		},
	})
	if err != nil {
		return
	}
	if !res.IsSuccess() {
		log.Error(res.Content.Msg, res.Content.SubMsg)
		err = fmt.Errorf("%s - %s", res.Content.Code, res.Content.SubCode)
	} else {
		qrCode = res.Content.QRCode
	}
	return
}

func (s *Alipay) CancelTrade(outTradeNo string) (err error) {
	_, err = s.client.TradeCancel(alipay.TradeCancel{
		OutTradeNo: outTradeNo,
	})
	return
}
