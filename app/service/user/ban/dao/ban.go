package dao

import (
	"backend/app/service/user/ban/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

func (d *dao) Info(uid int64) (res *model.Ban, err error) {
	res = &model.Ban{}

	// DB
	err = d.db.Where(&model.Ban{UID: uid}).Order("expire_time desc").First(res).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found ban record UID(%d)", uid)
	}
	return
}

func (d *dao) Add(info *model.Ban) (err error) {
	// DB
	if info.IsValid() {
		err = d.db.Create(info).Error
	}
	return
}

func (d *dao) Remove(info *model.Ban) (err error) {
	// IMPORTANT!!! id empty will delete all recode!!!
	if !info.IsValid() {
		err = ecode.Errorf(codes.InvalidArgument, "ID is empty")
		return
	}

	// DB
	DBres := d.db.Delete(info)
	err = DBres.Error
	if DBres.RowsAffected == 0 {
		err = ecode.Errorf(codes.NotFound, "Can not found recode ID(%d)", info.ID)
	}
	return
}
