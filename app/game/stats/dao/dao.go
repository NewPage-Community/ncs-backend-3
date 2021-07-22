package dao

import (
	"backend/app/game/stats/conf"
	"backend/app/game/stats/model"
	cache "backend/pkg/cache/redis"
	"backend/pkg/log"
	"github.com/go-redis/redis/v7"
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
	redis *redis.Client
}

func Init(config *conf.Config) (d *dao) {
	d = &dao{
		redis: cache.Init(),
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
