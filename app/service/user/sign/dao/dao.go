package dao

import (
	"backend/app/service/user/sign/conf"
	"backend/app/service/user/sign/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"

	"gorm.io/gorm"
)

type Dao interface {
	Info(uid int64) (*model.Sign, error)
	Sign(uid int64) (alreadySigned bool, err error)
	Healthy() bool
	Close()
}

type dao struct {
	db *gorm.DB
}

func New(config *conf.Config) (d *dao) {
	d = &dao{
		db: mysql.Init(config.Mysql),
	}
	if err := d.db.AutoMigrate(&model.Sign{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy()
}

func (d *dao) Close() {
}
