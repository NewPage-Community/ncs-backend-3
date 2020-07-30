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
	res = &model.User{}
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where(uid).First(res).Error
	if err == gorm.ErrRecordNotFound {
		// Create and add point
		err = d.db.Create(&model.User{
			UID:   uid,
			Point: addPoint,
		}).Error
		return
	}
	if err != nil {
		return
	}

	lastLevel = res.Level()
	res.Point += addPoint

	err = d.db.Model(res).Updates(&model.User{Point: res.Point}).Error
	return
}

func (d *dao) UpgradePass(uid int64, passType int32) (err error) {
	// DB
	err = d.db.Model(&model.User{UID: uid}).Updates(&model.User{PassType: passType}).Error
	return
}
