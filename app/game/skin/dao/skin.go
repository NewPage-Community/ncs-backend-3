package dao

import (
	"backend/app/game/skin/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

func (d *dao) GetInfo(uid int64) (res *model.SkinUser, err error) {
	res = &model.SkinUser{}

	// DB
	err = d.db.Where(uid).First(res).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	return
}

func (d *dao) SetUsedSkin(uid int64, usedSkin int32) (err error) {
	info := &model.SkinUser{
		UID:      uid,
		UsedSkin: usedSkin,
	}

	// DB
	err = d.db.Model(info).Update("used_skin", usedSkin).Error
	return
}

func (d *dao) Create(uid int64) (res *model.SkinUser, err error) {
	res = &model.SkinUser{
		UID:      uid,
		UsedSkin: 0,
	}

	// DB
	err = d.db.Create(res).Error
	return
}
