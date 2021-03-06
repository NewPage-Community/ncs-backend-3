package dao

import (
	"backend/app/game/server/conf"
	"backend/app/game/server/model"
	db "backend/pkg/database/mysql"
	"backend/pkg/log"
	"gorm.io/gorm"
)

type Dao interface {
	Info(address string) (*model.Info, error)
	InfoWithID(id int32) (*model.Info, error)
	AllInfo() ([]*model.Info, error)
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
	if err := d.db.AutoMigrate(&model.Info{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Close() {
}

func (d *dao) Healthy() bool {
	return db.Healthy(d.db)
}
