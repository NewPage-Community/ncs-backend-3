package dao

import (
	"backend/app/service/backpack/items/conf"
	"backend/app/service/backpack/items/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"

	"gorm.io/gorm"
)

type Dao interface {
	GetItems() ([]*model.Item, error)
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
	if err := d.db.AutoMigrate(&model.Item{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy()
}

func (d *dao) Close() {
}
