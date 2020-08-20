package model

type Donate struct {
	OutTradeNo string `gorm:"primary_key"`
	UID        int64
	Amount     int32
	Payed      bool
	CreatedAt  int64 `gorm:"autoCreateTime"`
}

// TableName return table name
func (*Donate) TableName() string {
	return "np_donate"
}
