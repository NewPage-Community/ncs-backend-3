package dao

import (
	"backend/app/game/stats/conf"
	"backend/app/game/stats/model"
	cache "backend/pkg/database/redis"
	"backend/pkg/log"
	goredis "github.com/go-redis/redis/v8"
)

type Dao interface {
	Get(stats *model.Stats) error
	GetAll(stats *model.Stats) ([]*model.Stats, error)
	Set(stats *model.Stats) error
	Incr(stats *model.Stats) error
	Healthy() bool
	Close()
}

type dao struct {
	redis *goredis.Client
}

func Init(config *conf.Config) (d *dao) {
	d = &dao{
		redis: cache.Init(config.Redis),
	}
	return
}

func (d *dao) Healthy() bool {
	return cache.Healthy(d.redis)
}

func (d *dao) Close() {
	if err := d.redis.Close(); err != nil {
		log.Error(err)
	}
}
