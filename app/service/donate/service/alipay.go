package service

import (
	"backend/app/service/donate/conf"
	"backend/pkg/log"
	"context"
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
	res, err := s.client.TradeQuery(context.Background(), alipay.TradeQuery{
		OutTradeNo: outTradeNo,
	})
	if err == nil {
		if res.TradeStatus == alipay.TradeStatusSuccess {
			return true, nil
		}
	}
	return false, err
}

func (s *Alipay) CreateTrade(outTradeNo string, steamID int64, totalAmount int32) (qrCode string, err error) {
	res, err := s.client.TradePreCreate(context.Background(), alipay.TradePreCreate{
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
		log.Error(res.Msg, res.SubMsg)
		err = fmt.Errorf("%s - %s", res.Code, res.SubCode)
	} else {
		qrCode = res.QRCode
	}
	return
}

func (s *Alipay) CancelTrade(outTradeNo string) (err error) {
	_, err = s.client.TradeCancel(context.Background(), alipay.TradeCancel{
		OutTradeNo: outTradeNo,
	})
	return
}
