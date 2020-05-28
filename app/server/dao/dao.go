package dao

import (
	"backend/app/server/conf"
	"backend/app/server/model"
	db "backend/pkg/database/mysql"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Info(address string) (*model.Info, error)
	UpdateRcon(server *model.Info) error
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
	d.db.AutoMigrate(&model.Info{})
	return
}

func (d *dao) Close() {
	d.db.Close()
}
