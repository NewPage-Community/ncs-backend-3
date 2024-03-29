package dao

import (
	"backend/app/service/donate/event"
	"backend/app/service/donate/model"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *dao) CreateDonate(uid int64, steamID int64, amount int32, payment model.DonatePayment) (outTradeNo string, err error) {
	steam := strconv.FormatInt(steamID, 10)
	if len(steam) < 4 {
		err = errors.New("invalid steamid")
		return
	}

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	now := time.Now().In(beijing).Format("20060102150405")
	outTradeNo = fmt.Sprintf("%s%s", now, steam[len(steam)-4:])

	// DB
	info := &model.Donate{
		OutTradeNo: outTradeNo,
		UID:        uid,
		Amount:     amount,
		Payment:    payment,
	}
	err = d.db.Create(info).Error
	return
}

func (d *dao) GetTradeInfo(outTradeNo string) (res *model.Donate, err error) {
	res = &model.Donate{}

	// DB
	err = d.db.Find(res, outTradeNo).Error
	return
}

func (d *dao) FinishTrade(outTradeNo string) (err error) {
	info := &model.Donate{
		OutTradeNo: outTradeNo,
		Status:     model.DonatePayed,
	}
	err = d.db.Model(info).Updates(info).Error
	return
}

func (d *dao) CancelTrade(outTradeNo string) (err error) {
	info := &model.Donate{
		OutTradeNo: outTradeNo,
		Status:     model.DonateCancel,
	}
	err = d.db.Model(info).Updates(info).Error
	return
}

func (d *dao) GetDonateList(startTime int64, endTime int64, uid int64) (res []*model.Donate, err error) {
	res = make([]*model.Donate, 0)
	var expr clause.Expr
	if uid > 0 {
		expr = gorm.Expr("status = ? AND created_at BETWEEN ? AND ? AND uid = ?", model.DonatePayed, startTime, endTime, uid)
	} else {
		expr = gorm.Expr("status = ? AND created_at BETWEEN ? AND ?", model.DonatePayed, startTime, endTime)
	}
	err = d.db.Where(expr).Order("created_at desc").Find(&res).Error
	return
}

func (d *dao) GetCheckTradeList() (res []*model.Donate, err error) {
	res = make([]*model.Donate, 0)
	err = d.db.Where(gorm.Expr("status = ?", model.DonateUnPay)).Find(&res).Error
	return
}

func (d *dao) CreateDonateEvent(ctx context.Context, data *event.DonateEventData) (err error) {
	return event.NewDonateEvent(d.stream).Create(ctx, data)
}

func (d *dao) ListenDonateEvent(cb event.DonateCallback) error {
	return event.NewDonateEvent(d.stream).Listen(cb)
}
