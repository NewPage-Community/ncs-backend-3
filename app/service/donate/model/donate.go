package model

type Donate struct {
	OutTradeNo string `gorm:"primary_key"`
	UID        int64
	Amount     int32
	Status     DonateStatus
	Payment    DonatePayment
	CreatedAt  int64 `gorm:"autoCreateTime"`
}

type DonateStatus int32
type DonatePayment int32

const (
	DonateUnPay = DonateStatus(iota)
	DonatePayed
	DonateCancel
)

const (
	Alipay = DonatePayment(iota + 1)
	Wepay
)

// TableName return table name
func (*Donate) TableName() string {
	return "np_donate"
}
