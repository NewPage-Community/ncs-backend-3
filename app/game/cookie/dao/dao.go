package dao

import (
	"backend/app/game/cookie/conf"
	"backend/app/game/cookie/model"
	cache "backend/pkg/cache/redis"
	"backend/pkg/log"
	"github.com/go-redis/redis/v7"
)

type Dao interface {
	Get(uid int64) (*model.Cookie, error)
	Set(uid int64, key string, value string) error
	Healthy() bool
	Close()
}

type dao struct {
	db *redis.Client
}

func Init(config *conf.Config) (d *dao) {
	d = &dao{
		db: cache.Init(),
	}
	return
}

func (d *dao) Healthy() bool {
	return cache.Healthy(d.db)
}

func (d *dao) Close() {
	if err := d.db.Close(); err != nil {
		log.Error(err)
	}
}
