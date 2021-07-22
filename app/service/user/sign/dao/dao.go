package dao

import (
	"backend/app/service/user/sign/conf"
	"backend/app/service/user/sign/model"
	db "backend/pkg/database/mysql"
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
		db: db.Init(),
	}
	if err := d.db.AutoMigrate(&model.Sign{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return db.Healthy(d.db)
}

func (d *dao) Close() {
}
