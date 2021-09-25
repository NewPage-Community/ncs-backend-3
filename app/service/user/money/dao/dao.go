package dao

import (
	"backend/app/service/user/money/conf"
	"backend/app/service/user/money/model"
	"backend/pkg/database/mysql"
	"backend/pkg/database/redis"
	"backend/pkg/log"

	goredis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Dao interface {
	Get(uid int64) (*model.Money, error)
	Pay(uid int64, price int32) error
	Give(uid int64, money int32) error
	AddRecord(uid int64, amount int32, reason string) error
	GetRecords(uid int64) (*model.Records, error)
	Healthy() bool
	Close()
}

type dao struct {
	db    *gorm.DB
	cache *goredis.Client
}

func Init(config *conf.Config) (d *dao) {
	d = &dao{
		db:    mysql.Init(config.Mysql),
		cache: redis.Init(config.Redis),
	}
	if err := d.db.AutoMigrate(&model.Money{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy(d.db)
}

func (d *dao) Close() {
	if err := d.cache.Close(); err != nil {
		log.Error(err)
	}
}
