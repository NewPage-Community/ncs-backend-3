package dao

import (
	"backend/app/service/user/money/conf"
	"backend/app/service/user/money/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"

	"gorm.io/gorm"
)

type Dao interface {
	Get(uid int64) (*model.Money, error)
	Pay(uid int64, price uint32, reason string) error
	Give(uid int64, money uint32, reason string) error
	GetRecords(uid int64, days uint32) (*model.Records, error)
	Gift(uid, target uint64, money uint32) (remaining uint32, err error)
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
	if err := d.db.AutoMigrate(&model.Money{}); err != nil {
		log.Error(err)
	}
	if err := d.db.AutoMigrate(&model.Record{}); err != nil {
		log.Error(err)
	}
	if err := d.db.AutoMigrate(&model.Gift{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy()
}

func (d *dao) Close() {
}
