package dao

import (
	"backend/app/server/conf"
	"backend/app/user/ban/model"
	db "backend/pkg/database/mysql"
	"github.com/jinzhu/gorm"
)

// TODO: DAO interface
type Dao interface {
}

type dao struct {
	db *gorm.DB
}

func New(config *conf.Config) (d *dao) {
	d = &dao{
		db: db.Init(config.Mysql),
	}
	// Auto migrate db
	d.db.AutoMigrate(&model.Ban{})
	return
}

func (d *dao) Close() {
	d.db.Close()
}
