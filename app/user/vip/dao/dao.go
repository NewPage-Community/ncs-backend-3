package dao

import (
	"backend/app/user/vip/conf"
	"backend/app/user/vip/model"
	db "backend/pkg/database/mysql"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Register(info *model.VIP) error
	Info(info *model.VIP) error
	ExpireTime(info *model.VIP) error
	Point(info *model.VIP) error
	Close()
}

type dao struct {
	db *gorm.DB
}

func New(config *conf.Config) (d *dao) {
	d = &dao{
		db: db.Init(config.Mysql),
	}
	// Auto migrate db
	d.db.AutoMigrate(&model.VIP{})
	return
}

func (d *dao) Close() {
	d.db.Close()
}
