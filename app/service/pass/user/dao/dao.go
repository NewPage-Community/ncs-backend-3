package dao

import (
	"backend/app/service/pass/user/conf"
	"backend/app/service/pass/user/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
	"gorm.io/gorm"
)

type Dao interface {
	Info(uid int64) (*model.User, error)
	UpgradePass(uid int64, passType int32) error
	AddPoint(uid int64, addPoint int32) (res *model.User, lastLevel int32, err error)
	Healthy() bool
	Close()
}

type dao struct {
	db *gorm.DB
}

func Init(config *conf.Config) (d *dao) {
	d = &dao{
		db: mysql.Init(),
	}
	if err := d.db.AutoMigrate(&model.User{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy(d.db)
}

func (d *dao) Close() {
}
