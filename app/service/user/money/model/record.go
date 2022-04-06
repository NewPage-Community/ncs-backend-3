package model

import (
	"backend/pkg/database/mysql"
	"backend/pkg/json"
)

type Record struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	UID    uint64 `gorm:"not null" json:"uid"`
	Amount int32  `gorm:"not null" json:"amount"`
	Reason string `gorm:"not null" json:"reason"`
	mysql.Model
}

// TableName return table name
func (*Record) TableName() string {
	return "np_records"
}

// IsValid .
func (i *Record) IsValid() bool {
	return i.ID > 0
}

func (r *Record) JSON() ([]byte, error) {
	return json.Marshal(r)
}

type Records []*Record
