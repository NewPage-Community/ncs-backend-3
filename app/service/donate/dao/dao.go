package dao

import (
	"backend/app/service/donate/conf"
	"backend/app/service/donate/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
	"gorm.io/gorm"
)

type Dao interface {
	CreateDonate(uid int64, steamID int64, amount int32, payment model.DonatePayment) (outTradeNo string, err error)
	GetTradeInfo(outTradeNo string) (res *model.Donate, err error)
	FinishTrade(outTradeNo string) (err error)
	CancelTrade(outTradeNo string) (err error)
	GetDonateList(startTime int64, endTime int64, uid int64) (res []*model.Donate, err error)
	GetCheckTradeList() (res []*model.Donate, err error)
	Healthy() bool
	Close()
}

type dao struct {
	db *gorm.DB
}

func Init(config *conf.Config) (d *dao) {
	d = &dao{
		db: mysql.Init(config.Mysql),
	}
	if err := d.db.AutoMigrate(&model.Donate{}); err != nil {
		log.Error(err)
	}
	return
}

func (d *dao) Healthy() bool {
	return mysql.Healthy(d.db)
}

func (d *dao) Close() {
}
