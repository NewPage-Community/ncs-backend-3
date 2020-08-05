package model

import (
	"time"
)

type Sign struct {
	UID      int64 `gorm:"primary_key;not null" json:"uid"`
	SignDate int   `gorm:"not null;INDEX" json:"sign_date"`
	SignDays int   `gorm:"not null;INDEX" json:"sign_days"`
}

// TableName return table name
func (*Sign) TableName() string {
	return "np_sign"
}

func (s *Sign) IsValid() bool {
	return s.UID > 0
}

// GetNowDate return now CST date
func (*Sign) GetNowDate() int {
	cst := time.FixedZone("CST", 8*3600)
	now := time.Now().In(cst)
	return GetDate(now)
}

// Sign .
func (s *Sign) Sign() {
	if !s.IsSigned() {
		now := s.GetNowDate()
		// Continuous
		if now-s.SignDate == 1 {
			s.SignDays += 1
		} else {
			s.SignDays = 1
		}
		s.SignDate = now
	}
}

// IsSigned check if user signed
func (s *Sign) IsSigned() bool {
	return s.GetNowDate() == s.SignDate
}

// GetDate return Date 19700101
func GetDate(t time.Time) int {
	return t.Year()*10000 + int(t.Month())*100 + t.Day()
}
