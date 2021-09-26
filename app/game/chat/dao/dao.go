package dao

import (
	"backend/app/game/chat/conf"
	chatEvent "backend/app/game/chat/event"
	"backend/pkg/database/redis"
	"backend/pkg/log"
	"context"

	goredis "github.com/go-redis/redis/v8"
)

type Dao interface {
	CreateChatEvent(ctx context.Context, data *chatEvent.AllChatEventData) (err error)
	ListenChatEvent(cb chatEvent.AllChatCallback) error
	Healthy() bool
	Close()
}

type dao struct {
	db     *goredis.Client
	stream *redis.Stream
}

func Init(config *conf.Config, service string) (d *dao) {
	d = &dao{
		db: redis.Init(config.Redis),
	}
	d.stream = redis.NewStream(d.db, service)
	return
}

func (d *dao) Healthy() bool {
	return redis.Healthy(d.db)
}

func (d *dao) Close() {
	d.stream.Close()
	if err := d.db.Close(); err != nil {
		log.Error(err)
	}
}
