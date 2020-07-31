package dao

import (
	"backend/app/service/user/money/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	defer func() {
		d.db.Commit()
	}()

	// DB
	err = d.db.Clauses(clause.Locking{
		Strength: "UPDATE",
	}).First(info, uid).Error
	if err != nil {
		// not found and create
		if err == gorm.ErrRecordNotFound {
			err = d.db.Create(info).Error
			if err != nil {
				return
			}
			err = ecode.Errorf(codes.Unknown, "No enough money!")
		}
		return
	}
	if info.RMB -= price; info.RMB < 0 {
		err = ecode.Errorf(codes.Unknown, "No enough money!")
		return
	}
	err = d.db.Model(info).Select("rmb").Updates(model.Money{RMB: info.RMB}).Error
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
