package dao

import (
	"backend/app/service/user/title/conf"
	"backend/app/service/user/title/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Title(uid int64) (*model.Title, error)
	Update(title *model.Title) error
	Create(title *model.Title) error
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
	// Auto migrate db
	d.db.AutoMigrate(&model.Title{})
	return
}

func (d *dao) Close() {
	d.db.Close()
}

func (d *dao) Healthy() bool {
	err := d.db.DB().Ping()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}
