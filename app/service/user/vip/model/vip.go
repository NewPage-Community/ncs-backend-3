package model

import "time"

type VIP struct {
	UID        int64 `gorm:"primary_key;unique;not null" json:"uid"`
	Point      int   `gorm:"not null;INDEX" json:"point"`
	ExpireDate int64 `gorm:"not null;INDEX" json:"expire_date"`
}

// TableName return table name
func (*VIP) TableName() string {
	return "np_vip"
}

// IsValid .
func (v *VIP) IsValid() bool {
	return v.UID > 0
}

// Level .
func (v *VIP) Level() int {
	if v.Point == 0 {
		return 0
	}
	return v.Point/200 + 1
}

// Renewal .
func (v *VIP) Renewal(length int64) {
	now := time.Now().UTC().Unix()
	if v.ExpireDate > now {
		v.ExpireDate += length
	} else {
		v.ExpireDate = now + length
	}
}
