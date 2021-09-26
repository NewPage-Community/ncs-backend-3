package dao

import (
	"backend/app/game/skin/conf"
	"backend/app/game/skin/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"

	"gorm.io/gorm"
)

type Dao interface {
	GetInfo(uid int64) (*model.SkinUser, error)
	SetUsedSkin(uid int64, usedSkin int32) error
	Create(uid int64) (*model.SkinUser, error)
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
	if err := d.db.AutoMigrate(&model.SkinUser{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy()
}

func (d *dao) Close() {
}
