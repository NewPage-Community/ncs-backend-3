package dao

import (
	"backend/app/service/donate/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func (d *dao) CreateDonate(uid int64, steamID int64, amount int32) (outTradeNo string, err error) {
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

func (d *dao) GetDonateList(startTime int64, endTime int64) (res []*model.Donate, err error) {
	res = make([]*model.Donate, 0)
	err = d.db.Where(gorm.Expr("status = ? AND created_at BETWEEN ? AND ?", model.DonatePayed, startTime, endTime)).
		Order("created_at desc").Find(&res).Error
	return
}
