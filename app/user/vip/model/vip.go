package model

import "time"

type VIP struct {
	UID        int64 `gorm:"primary_key;unique;not null" json:"UID"`
	Point      int   `gorm:"not null;INDEX" json:"Point"`
	ExpireDate int64 `gorm:"not null;INDEX" json:"ExpireDate"`
}

// TableName return table name
func (*VIP) TableName() string {
	return "np_vip"
}

// IsValid .
func (v *VIP) IsValid() bool {
	return v.UID > 0
}

// IsExpired .
func (v *VIP) IsExpired() bool {
	return time.Now().Unix() > v.ExpireDate
}
