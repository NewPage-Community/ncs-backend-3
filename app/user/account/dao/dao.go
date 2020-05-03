package dao

import (
	"backend/app/user/account/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	UID(steamID int64) (uid int64, err error)
	Info(uid int64) (info *model.Info, err error)
	Register(steamID int64) (uid int64, err error)
	Close()
}

type dao struct {
	db *gorm.DB
}

func New(config *mysql.Config) Dao {
	return &dao{
		db: mysql.Init(config),
	}
}

func (d *dao) Close() {
	err := d.db.Close()
	if err != nil {
		log.Error(err)
	}
}
