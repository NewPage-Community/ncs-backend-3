package dao

import (
	"backend/app/service/user/vip/conf"
	"backend/app/service/user/vip/model"
	db "backend/pkg/database/mysql"
	"gorm.io/gorm"
)

type Dao interface {
	Register(info *model.VIP) error
	Info(info *model.VIP) error
	ExpireTime(info *model.VIP) error
	Point(info *model.VIP) error
	Healthy() bool
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
}

func (d *dao) Healthy() bool {
	return db.Healthy(d.db)
}
