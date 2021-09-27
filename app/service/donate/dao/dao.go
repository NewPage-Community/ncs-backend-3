package dao

import (
	"backend/app/service/donate/conf"
	"backend/app/service/donate/event"
	"backend/app/service/donate/model"
	"backend/pkg/database/mysql"
	"backend/pkg/database/redis"
	"backend/pkg/log"
	"context"

	goredis "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Dao interface {
	CreateDonate(uid int64, steamID int64, amount int32, payment model.DonatePayment) (outTradeNo string, err error)
	GetTradeInfo(outTradeNo string) (res *model.Donate, err error)
	FinishTrade(outTradeNo string) (err error)
	CancelTrade(outTradeNo string) (err error)
	GetDonateList(startTime int64, endTime int64, uid int64) (res []*model.Donate, err error)
	GetCheckTradeList() (res []*model.Donate, err error)
	CreateDonateEvent(ctx context.Context, data *event.DonateEventData) (err error)
	ListenDonateEvent(cb event.DonateCallback) error
	Healthy() bool
	Close()
}

type dao struct {
	db     *gorm.DB
	redis  *goredis.Client
	stream *redis.Stream
}

func Init(config *conf.Config, service string) (d *dao) {
	d = &dao{
		db:    mysql.Init(config.Mysql),
		redis: redis.Init(config.Redis),
	}
	if err := d.db.AutoMigrate(&model.Donate{}); err != nil {
		log.Error(err)
	}
	d.stream = redis.NewStream(d.redis, service)
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy() && redis.Healthy()
}

func (d *dao) Close() {
	redis.Close()
}
