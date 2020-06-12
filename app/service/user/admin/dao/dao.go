package dao

import (
	"backend/app/service/user/admin/conf"
	"backend/app/service/user/admin/model"
	db "backend/pkg/database/mysql"
	"backend/pkg/log"
	"github.com/jinzhu/gorm"
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
	d.db.AutoMigrate(&model.Admin{})
	return
}

func (d *dao) Healthy() bool {
	err := d.db.DB().Ping()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func (d *dao) Close() {
	d.db.Close()
}
