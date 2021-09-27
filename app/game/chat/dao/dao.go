package dao

import (
	"backend/app/game/chat/conf"
	chatEvent "backend/app/game/chat/event"
	"backend/pkg/database/redis"
	"context"

	goredis "github.com/go-redis/redis/v8"
)

type Dao interface {
	CreateAllChatEvent(ctx context.Context, data *chatEvent.AllChatEventData) (err error)
	ListenAllChatEvent(cb chatEvent.AllChatCallback) error
	Healthy() bool
	Close()
}

type dao struct {
	redis  *goredis.Client
	stream *redis.Stream
}

func Init(config *conf.Config, service string) (d *dao) {
	d = &dao{
		redis: redis.Init(config.Redis),
	}
	d.stream = redis.NewStream(d.redis, service)
	return
}

func (d *dao) Healthy() bool {
	return redis.Healthy()
}

func (d *dao) Close() {
	redis.Close()
}
