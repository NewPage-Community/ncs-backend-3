package model

import "backend/pkg/database/mysql"

type Money struct {
	UID int64 `gorm:"primary_key" json:"uid"`
	RMB int32 `gorm:"not null;INDEX" json:"rmb"`
	mysql.Model
}

// TableName return table name
func (*Money) TableName() string {
	return "np_money"
}

// IsValid .
func (i *Money) IsValid() bool {
	return i.UID > 0
}
