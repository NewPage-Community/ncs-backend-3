package dao

import (
	"backend/app/user/admin/conf"
	"backend/app/user/admin/model"
	db "backend/pkg/database/mysql"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Info(uid int64) (*model.Admin, error)
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

func (d *dao) Close() {
	d.db.Close()
}
