package model

import (
	"backend/pkg/database/mysql"
)

type Gift struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	UID    uint64 `gorm:"not null" json:"uid"`
	Target uint64 `gorm:"not null" json:"target"`
	Money  int32  `gorm:"not null" json:"money"`
	mysql.Model
}

// TableName return table name
func (*Gift) TableName() string {
	return "np_gift"
}

// IsValid .
func (i *Gift) IsValid() bool {
	return i.ID > 0
}
