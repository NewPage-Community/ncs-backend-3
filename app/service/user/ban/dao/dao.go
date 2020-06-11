package dao

import (
	"backend/app/service/user/ban/conf"
	"backend/app/service/user/ban/model"
	db "backend/pkg/database/mysql"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Info(uid int64) (*model.Ban, error)
	Add(info *model.Ban) error
	Remove(info *model.Ban) error
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
	d.db.AutoMigrate(&model.Ban{})
	return
}

func (d *dao) Close() {
	d.db.Close()
}
