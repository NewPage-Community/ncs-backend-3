package dao

import (
	"backend/app/game/server/conf"
	"backend/app/game/server/event"
	"backend/app/game/server/model"
	"backend/pkg/database/mysql"
	"backend/pkg/database/redis"
	"backend/pkg/log"
	"context"

	goredis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Dao interface {
	Info(address string) (*model.Info, error)
	InfoWithID(id int32) (*model.Info, error)
	AllInfo() ([]*model.Info, error)
	UpdateRcon(server *model.Info) error
	CreateChangeMapEvent(ctx context.Context, data *event.ChangeMapEventData) error
	Healthy() bool
	Close()
}

type dao struct {
	db     *gorm.DB
	redis  *goredis.Client
	stream *redis.Stream
}

func New(config *conf.Config, service string) (d *dao) {
	d = &dao{
		db:    mysql.Init(config.Mysql),
		redis: redis.Init(config.Redis),
	}
	// Auto migrate db
	if err := d.db.AutoMigrate(&model.Info{}); err != nil {
		log.Error(err)
	}
	d.stream = redis.NewStream(d.redis, service)
	return
}

func (d *dao) Close() {
	redis.Close()
	mysql.Close()
}

func (d *dao) Healthy() bool {
	return mysql.Healthy() && redis.Healthy()
}
