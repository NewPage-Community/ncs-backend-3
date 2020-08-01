package dao

import (
	"backend/app/service/user/money/model"
	"errors"
	"gorm.io/gorm"
)

var (
	NoEnoughMoney = errors.New("no enough money")
)

func (d *dao) Get(uid int64) (res *model.Money, err error) {
	res = &model.Money{}

	// DB
	err = d.db.First(res, uid).Error
	if err == gorm.ErrRecordNotFound {
		res.UID = uid
		err = d.db.Create(res).Error
	}
	return
}

func (d *dao) Pay(uid int64, price int32) (err error) {
	info := &model.Money{
		UID: uid,
	}

	// DB
	err = d.db.First(info, uid).Error
	if err != nil {
		// not found and create
		if err == gorm.ErrRecordNotFound {
			err = d.db.Create(info).Error
			if err != nil {
				return
			}
			err = NoEnoughMoney
		}
		return
	}

	res := d.db.Model(info).Where(gorm.Expr("rmb >= ?", price)).
		Update("rmb", gorm.Expr("rmb - ?", price))
	err = res.Error
	if res.RowsAffected == 0 {
		err = NoEnoughMoney
	}
	return
}

func (d *dao) Give(uid int64, money int32) (err error) {
	info := &model.Money{
		UID: uid,
		RMB: money,
	}

	// DB
	res := d.db.Model(info).Update("rmb", gorm.Expr("rmb + ?", money))
	if err = res.Error; err != nil {
		return
	}
	if res.RowsAffected == 0 {
		// not found and create
		err = d.db.Create(info).Error
	}
	return
}
