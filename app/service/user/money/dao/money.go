package dao

import (
	"backend/app/service/user/money/model"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrNoEnoughMoney = errors.New("no enough money")
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

func (d *dao) Pay(uid int64, price uint32, reason string) (err error) {
	return d.db.Transaction(func(tx *gorm.DB) error {
		return d.pay(tx, uid, price, reason)
	})
}

func (d *dao) pay(tx *gorm.DB, uid int64, price uint32, reason string) (err error) {
	info := &model.Money{
		UID: uid,
	}

	// DB
	err = tx.First(info, uid).Error
	if err != nil {
		// not found and create
		if err == gorm.ErrRecordNotFound {
			err = d.db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(info).Error
			if err != nil {
				return err
			}
			err = ErrNoEnoughMoney
		}
		return err
	}

	res := d.db.Model(info).Where(gorm.Expr("rmb >= ?", price)).
		Update("rmb", gorm.Expr("rmb - ?", price))
	err = res.Error
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		err = ErrNoEnoughMoney
		return err
	}

	// Record
	if err = d.addRecord(tx, uid, int32(-price), reason); err != nil {
		return err
	}
	return nil
}

func (d *dao) Give(uid int64, money uint32, reason string) (err error) {
	return d.db.Transaction(func(tx *gorm.DB) error {
		return d.give(tx, uid, money, reason)
	})
}

func (d *dao) give(tx *gorm.DB, uid int64, money uint32, reason string) (err error) {
	// DB
	err = tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "uid"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"rmb":        gorm.Expr("rmb + ?", int32(money)),
			"updated_at": time.Now(),
		}),
	}).Create(&model.Money{
		UID: uid,
		RMB: int32(money),
	}).Error
	if err != nil {
		return err
	}

	// Record
	if err = d.addRecord(tx, uid, int32(money), reason); err != nil {
		return err
	}
	return nil
}

func (d *dao) create(uid int64) (err error) {
	info := &model.Money{
		UID: uid,
	}
	return d.db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(info).Error
}
