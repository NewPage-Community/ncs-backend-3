package dao

import (
	"backend/app/service/pass/user/conf"
	"backend/app/service/pass/user/model"
	"backend/pkg/database/mysql"
	"gorm.io/gorm"
)

type Dao interface {
	Info(uid int64) (*model.User, error)
	Update(info *model.User) error
	AddPoint(uid int64, addPoint int32) error
	Create(info *model.User) error
	Healthy() bool
	Close()
}

type dao struct {
	db *gorm.DB
}

func Init(config *conf.Config) (d *dao) {
	d = &dao{
		db: mysql.Init(config.Mysql),
	}
	d.db.AutoMigrate(&model.User{})
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy(d.db)
}

func (d *dao) Close() {
}
