package dao

import (
	"backend/app/service/user/sign/conf"
	"backend/app/service/user/sign/model"
	db "backend/pkg/database/mysql"
	"gorm.io/gorm"
)

type Dao interface {
	Register(uid int64) error
	Info(uid int64) (*model.Sign, error)
	Update(info *model.Sign) error
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
	d.db.AutoMigrate(&model.Sign{})
	return
}

func (d *dao) Healthy() bool {
	return db.Healthy(d.db)
}

func (d *dao) Close() {
}
