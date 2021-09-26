package dao

import (
	"backend/app/service/user/account/conf"
	"backend/app/service/user/account/model"
	"backend/pkg/database/mysql"
	"backend/pkg/database/redis"
	"backend/pkg/log"

	goredis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Dao interface {
	UID(steamID int64) (*model.Info, error)
	Info(uid int64) (*model.Info, error)
	Register(steamID int64) (*model.Info, error)
	ChangeName(info *model.Info) error
	GetAllUID() (*[]int64, error)
	Healthy() bool
	Close()
}

type dao struct {
	db    *gorm.DB
	cache *goredis.Client
}

func New(config *conf.Config) (d *dao) {
	d = &dao{
		db:    mysql.Init(config.Mysql),
		cache: redis.Init(config.Redis),
	}
	// Auto migrate db
	if err := d.db.AutoMigrate(&model.Info{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Close() {
	if err := d.cache.Close(); err != nil {
		log.Error(err)
	}
}

func (d *dao) Healthy() bool {
	return mysql.Healthy()
}
