package dao

import (
	"backend/app/game/cookie/conf"
	"backend/app/game/cookie/model"
	"backend/pkg/database/redis"
	"backend/pkg/log"

	goredis "github.com/go-redis/redis/v8"
)

type Dao interface {
	Get(uid int64) (*model.Cookie, error)
	Set(uid int64, key string, value string) error
	Healthy() bool
	Close()
}

type dao struct {
	db *goredis.Client
}

func Init(config *conf.Config) (d *dao) {
	d = &dao{
		db: redis.Init(config.Redis),
	}
	return
}

func (d *dao) Healthy() bool {
	return redis.Healthy(d.db)
}

func (d *dao) Close() {
	if err := d.db.Close(); err != nil {
		log.Error(err)
	}
}
