package service

type Payment interface {
	CreateTrade(outTradeNo string, steamID int64, totalAmount int32) (qrCode string, err error)
	CheckTrade(outTradeNo string) (success bool, err error)
	CancelTrade(outTradeNo string) (err error)
}
