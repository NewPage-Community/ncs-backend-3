package dao

import (
	"backend/app/game/server/conf"
	"backend/app/game/server/model"
	db "backend/pkg/database/mysql"
	"gorm.io/gorm"
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
}

func (d *dao) Healthy() bool {
	return db.Healthy(d.db)
}
