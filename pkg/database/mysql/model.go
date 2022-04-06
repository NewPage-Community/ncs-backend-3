package mysql

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `gorm:"default:NULL" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"default:NULL;index" json:"-"`
}
