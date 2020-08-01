package dao

import (
	"backend/app/service/pass/user/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *dao) Info(uid int64) (res *model.User, err error) {
	res = &model.User{}

	// DB
	err = d.db.Where(uid).First(res).Error
	if err == gorm.ErrRecordNotFound {
		err = d.db.Create(&model.User{
			UID: uid,
		}).Error
	}
	return
}

func (d *dao) AddPoint(uid int64, addPoint int32) (res *model.User, lastLevel int32, err error) {
	res = &model.User{
		UID:   uid,
		Point: addPoint,
	}

	// DB
	err = d.db.First(res).Error
	if err == gorm.ErrRecordNotFound {
		err = d.db.Create(res).Error
		lastLevel = 0
		return
	}

	err = d.db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(res).Error
		if err != nil {
			return
		}

		lastLevel = res.Level()
		res.Point += addPoint

		err = tx.Model(res).Updates(&model.User{Point: res.Point}).Error
		return
	})
	return
}

func (d *dao) UpgradePass(uid int64, passType int32) (err error) {
	// DB
	err = d.db.Model(&model.User{UID: uid}).Updates(&model.User{PassType: passType}).Error
	return
}
