package dao

import (
	"backend/app/service/user/admin/conf"
	"backend/app/service/user/admin/model"
	db "backend/pkg/database/mysql"
	"backend/pkg/log"
	"gorm.io/gorm"
)

type Dao interface {
	Info(uid int64) (*model.Admin, error)
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
	if err := d.db.AutoMigrate(&model.Admin{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return db.Healthy(d.db)
}

func (d *dao) Close() {
}
