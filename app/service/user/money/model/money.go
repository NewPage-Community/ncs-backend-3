package model

type Money struct {
	UID int64 `gorm:"primary_key" json:"uid"`
	RMB int32 `gorm:"not null;INDEX" json:"rmb"`
}

// TableName return table name
func (*Money) TableName() string {
	return "np_money"
}

// IsValid .
func (i *Money) IsValid() bool {
	return i.UID > 0
}
