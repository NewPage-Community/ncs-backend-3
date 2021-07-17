package service

import (
	"backend/app/service/donate/conf"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	wepay "github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"time"
)

type Wepay struct {
	client *wepay.Client
	appID  string
	mchID  string
}

func InitWepay(config *conf.Wepay) *Wepay {
	ctx := context.Background()
	block, _ := pem.Decode([]byte(config.MchPrivateKey))
	if block == nil || block.Type != "PRIVATE KEY" {
		panic("failed to decode PEM block containing public key")
	}

	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	mchPrivateKey := pri.(*rsa.PrivateKey)

	opts := []wepay.ClientOption{
		// 一次性设置 签名/验签/敏感字段加解密，并注册 平台证书下载器，自动定时获取最新的平台证书
		option.WithWechatPayAutoAuthCipher(config.MchID, config.MchCertificateSerialNumber, mchPrivateKey, config.MchAPIv3Key),
	}
	client, err := wepay.NewClient(ctx, opts...)
	if err != nil {
		panic(err)
	}
	return &Wepay{
		client: client,
		appID:  config.AppID,
		mchID:  config.MchID,
	}
}

func (s *Wepay) CheckTrade(outTradeNo string) (bool, error) {
	svc := native.NativeApiService{Client: s.client}
	resp, _, err := svc.QueryOrderByOutTradeNo(context.Background(),
		native.QueryOrderByOutTradeNoRequest{
			OutTradeNo: wepay.String(outTradeNo),
			Mchid:      wepay.String(s.mchID),
		},
	)
	if err == nil {
		if resp.TradeState == wepay.String("SUCCESS") {
			return true, nil
		}
	}
	return false, err
}

func (s *Wepay) CreateTrade(outTradeNo string, steamID int64, totalAmount int32) (qrCode string, err error) {
	svc := native.NativeApiService{Client: s.client}
	resp, _, err := svc.Prepay(context.Background(),
		native.PrepayRequest{
			Appid:       wepay.String(s.appID),
			Mchid:       wepay.String(s.mchID),
			Description: wepay.String(fmt.Sprintf(TradeBody, steamID)),
			OutTradeNo:  wepay.String(outTradeNo),
			TimeExpire:  wepay.Time(time.Now().Add(5 * time.Minute)),
			Amount: &native.Amount{
				Total: wepay.Int64(int64(totalAmount * 100)),
			},
		},
	)
	if err != nil {
		return
	}

	qrCode = *resp.CodeUrl
	return
}

func (s *Wepay) CancelTrade(outTradeNo string) (err error) {
	svc := native.NativeApiService{Client: s.client}
	_, err = svc.CloseOrder(context.Background(),
		native.CloseOrderRequest{
			OutTradeNo: wepay.String(outTradeNo),
			Mchid:      wepay.String(s.mchID),
		},
	)
	return
}
