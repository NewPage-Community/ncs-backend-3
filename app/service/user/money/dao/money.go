package dao

import (
	"backend/app/service/user/money/model"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	GiveMoneySQL = "INSERT INTO np_money(uid,rmb) VALUES(?,?) ON DUPLICATE KEY UPDATE rmb=rmb+?"
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
		err = d.create(uid)
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
			err = d.create(uid)
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
	// DB
	err = d.db.Exec(GiveMoneySQL, uid, money, money).Error
	return
}

func (d *dao) create(uid int64) (err error) {
	info := &model.Money{
		UID: uid,
	}
	return d.db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(info).Error
}
