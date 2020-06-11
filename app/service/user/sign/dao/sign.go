package dao

import (
	"backend/app/service/user/sign/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
)

// Register register user and sign today
func (d *dao) Register(uid int64) (err error) {
	info := &model.Sign{
		UID:      uid,
		SignDays: 1,
	}
	info.SignTime = info.GetNowTime()

	// DB
	err = d.db.Create(info).Error
	return
}

// Info get info
func (d *dao) Info(uid int64) (res *model.Sign, err error) {
	res = &model.Sign{}

	// DB
	DBres := d.db.Where(uid).First(res)
	err = DBres.Error
	if DBres.RecordNotFound() {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	return
}

func (d *dao) Update(info *model.Sign) (err error) {
	// DB
	if !d.db.NewRecord(info) {
		DBres := d.db.Save(info)
		err = DBres.Error
		if DBres.RowsAffected == 0 {
			err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", info.UID)
		}
	}
	return
}
