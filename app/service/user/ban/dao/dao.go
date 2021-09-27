package dao

import (
	"backend/app/service/user/ban/conf"
	"backend/app/service/user/ban/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"

	"gorm.io/gorm"
)

type Dao interface {
	Info(uid uint64) (*model.Ban, error)
	Add(info *model.Ban) error
	Remove(info *model.Ban) error
	IsBlockIP(ip string) (bool, error)
	List(uid uint64) ([]*model.Ban, error)
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
	// Auto migrate db
	if err := d.db.AutoMigrate(&model.Ban{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy()
}

func (d *dao) Close() {
}
