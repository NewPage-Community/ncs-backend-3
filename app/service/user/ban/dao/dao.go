package dao

import (
	"backend/app/service/user/ban/conf"
	"backend/app/service/user/ban/model"
	db "backend/pkg/database/mysql"
	"backend/pkg/log"
	"gorm.io/gorm"
)

type Dao interface {
	Info(uid uint64) (*model.Ban, error)
	Add(info *model.Ban) error
	Remove(info *model.Ban) error
	IsBlockIP(ip string) (bool, error)
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
	if err := d.db.AutoMigrate(&model.Ban{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return db.Healthy(d.db)
}

func (d *dao) Close() {
}
