package model

import (
	"time"
)

type Sign struct {
	UID      int64     `gorm:"primary_key;not null" json:"uid"`
	SignTime time.Time `gorm:"not null;INDEX" json:"sign_time"`
	SignDays int       `gorm:"not null;INDEX" json:"sign_days"`
}

// TableName return table name
func (*Sign) TableName() string {
	return "np_sign"
}

func (s *Sign) IsValid() bool {
	return s.UID > 0
}

// GetNowTime .
func (*Sign) GetNowTime() time.Time {
	return time.Now().UTC()
}

// Sign .
func (s *Sign) Sign() {
	if !s.IsSigned() {
		s.SignTime = s.GetNowTime()
		s.SignDays += 1
	}
}

// IsSigned check if user signed
func (s *Sign) IsSigned() bool {
	now := s.GetNowTime()
	sign := s.SignTime
	if now.Year() == sign.Year() && now.YearDay() == sign.YearDay() {
		return true
	}
	return false
}
