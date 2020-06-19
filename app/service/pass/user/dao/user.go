package dao

import (
	"backend/app/service/pass/user/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
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

func (d *dao) Update(info *model.User) (err error) {
	// DB
	err = d.db.Model(info).Updates(*info).Error
	return
}

func (d *dao) AddPoint(uid int64, addPoint int32) (err error) {
	res := &model.User{}

	// DB
	err = d.db.Where(uid).First(res).Set("gorm:query_option", "FOR UPDATE").Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	if err != nil {
		return
	}
	err = d.db.Model(res).Update("point", res.Point+addPoint).Error
	return
}

func (d *dao) Create(info *model.User) (err error) {
	// DB
	err = d.db.Create(info).Error
	return
}
