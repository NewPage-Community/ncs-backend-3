package dao

import (
	"backend/app/service/pass/user/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *dao) Info(uid int64) (res *model.User, err error) {
	res = &model.User{}

	// DB
	err = d.db.Where(uid).First(res).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	return
}

func (d *dao) AddPoint(uid int64, addPoint int32) (res *model.User, upgrade bool, err error) {
	res = &model.User{}
	upgrade = false
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Where(uid).First(res).
		Clauses(clause.Locking{Strength: "UPDATE"}).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	if err != nil {
		return
	}

	_level := res.Level()
	res.Point += addPoint
	if res.Level() != _level {
		upgrade = true
	}

	err = d.db.Model(res).Updates(&model.User{Point: res.Point}).Error
	return
}

func (d *dao) Create(info *model.User) (err error) {
	// DB
	err = d.db.Create(info).Error
	return
}

func (d *dao) UpgradePass(uid int64) (err error) {
	// DB
	err = d.db.Model(&model.User{UID: uid}).Updates(&model.User{PassType: 1}).Error
	return
}
