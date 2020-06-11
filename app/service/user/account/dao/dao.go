package dao

import (
	"backend/app/service/user/account/conf"
	"backend/app/service/user/account/model"
	cache "backend/pkg/cache/redis"
	db "backend/pkg/database/mysql"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	UID(steamID int64) (*model.Info, error)
	Info(uid int64) (*model.Info, error)
	Register(steamID int64) (*model.Info, error)
	ChangeName(info *model.Info) error
	Close()
}

type dao struct {
	db    *gorm.DB
	cache *redis.Client
}

func New(config *conf.Config) (d *dao) {
	d = &dao{
		db:    db.Init(config.Mysql),
		cache: cache.Init(config.Redis),
	}
	// Auto migrate db
	d.db.AutoMigrate(&model.Info{})
	return
}

func (d *dao) Close() {
	d.db.Close()
	d.cache.Close()
}
