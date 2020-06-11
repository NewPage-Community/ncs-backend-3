package dao

import (
	"backend/app/service/user/admin/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
)

func (d *dao) Info(uid int64) (res *model.Admin, err error) {
	res = &model.Admin{}

	// DB
	DBRes := d.db.Where(uid).First(&res)
	err = DBRes.Error
	if DBRes.RecordNotFound() {
		ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	return
}
