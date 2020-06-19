package dao

import (
	"backend/app/service/user/title/conf"
	"backend/app/service/user/title/model"
	"backend/pkg/database/mysql"
	"gorm.io/gorm"
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
}

func (d *dao) Healthy() bool {
	return mysql.Healthy(d.db)
}
