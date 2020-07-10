package dao

import (
	"backend/app/service/user/vip/conf"
	"backend/app/service/user/vip/model"
	db "backend/pkg/database/mysql"
	"backend/pkg/log"
	"gorm.io/gorm"
)

type Dao interface {
	Register(info *model.VIP) error
	Info(info *model.VIP) error
	Renewal(uid int64, length int64) (exprTime int64, err error)
	AddPoint(uid int64, addPoint int) (point int, err error)
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
	if err := d.db.AutoMigrate(&model.VIP{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Close() {
}

func (d *dao) Healthy() bool {
	return db.Healthy(d.db)
}
