package dao

import (
	"backend/app/service/backpack/user/conf"
	"backend/app/service/backpack/user/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"

	"gorm.io/gorm"
)

type Dao interface {
	Create(uid int64) (*model.User, error)
	Get(uid int64) (*model.User, error)
	AddItems(uid int64, item *model.Items) error
	RemoveItem(uid int64, item model.Item, all bool) error
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
	if err := d.db.AutoMigrate(&model.UserModel{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy()
}

func (d *dao) Close() {
}
