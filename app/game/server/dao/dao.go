package dao

import (
	"backend/app/game/server/conf"
	"backend/app/game/server/model"
	db "backend/pkg/database/mysql"
	"backend/pkg/log"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Info(address string) (*model.Info, error)
	UpdateRcon(server *model.Info) error
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
	d.db.AutoMigrate(&model.Info{})
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
