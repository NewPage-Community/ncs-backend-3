package dao

import (
	"backend/app/service/user/title/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
)

func (d *dao) Title(uid int64) (res *model.Title, err error) {
	res = &model.Title{}

	// DB
	DBRes := d.db.Where(uid).First(&res)
	err = DBRes.Error
	if DBRes.RecordNotFound() {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	return
}

func (d *dao) Update(title *model.Title) (err error) {
	// DB
	if !d.db.NewRecord(title) {
		DBRes := d.db.Model(title).Update(*title)
		err = DBRes.Error
	}
	return
}

func (d *dao) Create(title *model.Title) (err error) {
	// DB
	DBRes := d.db.Create(title)
	err = DBRes.Error
	return
}
