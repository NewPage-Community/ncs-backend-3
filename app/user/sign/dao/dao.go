package dao

import (
	"backend/app/user/sign/conf"
	"backend/app/user/sign/model"
	db "backend/pkg/database/mysql"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Register(uid int64) error
	Info(uid int64) (*model.Sign, error)
	Update(info *model.Sign) error
	Close()
}

type dao struct {
	db *gorm.DB
}

func New(config *conf.Config) (d *dao) {
	d = &dao{
		db: db.Init(config.Mysql),
	}
	d.db.AutoMigrate(&model.Sign{})
	return
}

func (d *dao) Close() {
	d.db.Close()
}
