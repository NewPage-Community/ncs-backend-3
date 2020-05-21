package dao

import (
	"backend/app/user/group/conf"
	"backend/app/user/group/model"
	cache "backend/pkg/cache/redis"
	db "backend/pkg/database/mysql"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

//TODO:Dao

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
	d.db.AutoMigrate(&model.Group{})
	d.db.AutoMigrate(&model.User{})
	return
}

func (d *dao) Close() {
	d.db.Close()
	d.cache.Close()
}
