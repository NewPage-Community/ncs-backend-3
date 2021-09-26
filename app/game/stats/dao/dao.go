package dao

import (
	"backend/app/game/stats/conf"
	"backend/app/game/stats/model"
	"backend/pkg/database/redis"

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
		redis: redis.Init(config.Redis),
	}
	return
}

func (d *dao) Healthy() bool {
	return redis.Healthy()
}

func (d *dao) Close() {
	redis.Close()
}
