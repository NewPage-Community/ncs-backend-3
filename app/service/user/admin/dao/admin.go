package dao

import (
	"backend/app/service/user/admin/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

func (d *dao) Info(uid int64) (res *model.Admin, err error) {
	res = &model.Admin{}

	// DB
	err = d.db.Where(uid).First(&res).Error
	if err == gorm.ErrRecordNotFound {
		ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	return
}
