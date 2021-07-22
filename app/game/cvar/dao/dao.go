package dao

import (
	"backend/app/game/cvar/conf"
	"backend/app/game/cvar/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
	"gorm.io/gorm"
)

type Dao interface {
	GetCVars() (res []*model.CVar, err error)
	UpdatedCVar(id int) (err error)
	Healthy() bool
	Close()
}

type dao struct {
	db *gorm.DB
}

func Init(config *conf.Config) (d *dao) {
	d = &dao{
		db: mysql.Init(),
	}
	if err := d.db.AutoMigrate(&model.CVar{}); err != nil {
		log.Error(err)
	}
	return d
}

func (d *dao) Healthy() bool {
	return mysql.Healthy(d.db)
}

func (d *dao) Close() {
}
