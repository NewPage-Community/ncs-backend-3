package model

type Donate struct {
	OutTradeNo string `gorm:"primary_key"`
	UID        int64
	Amount     int32
	Status     DonateStatus
	CreatedAt  int64 `gorm:"autoCreateTime"`
}

type DonateStatus int32

const (
	DonateUnPay = DonateStatus(iota)
	DonatePayed
	DonateCancel
)

// TableName return table name
func (*Donate) TableName() string {
	return "np_donate"
}
