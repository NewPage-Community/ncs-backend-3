package dao

import (
	"backend/app/user/vip/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
)

func (d *dao) Register(info *model.VIP) (err error) {
	err = d.db.Create(info).Error
	return
}

func (d *dao) Info(info *model.VIP) (err error) {
	dbRes := d.db.Where(info.UID).First(info)
	err = dbRes.Error
	if dbRes.RecordNotFound() {
		err = ecode.Errorf(codes.NotFound, "Can not found: UID(%v)", info.UID)
	}
	return
}

func (d *dao) ExpireTime(info *model.VIP) (err error) {
	err = d.db.Model(info).Update("expire_date", info.ExpireDate).Error
	return
}

func (d *dao) Point(info *model.VIP) (err error) {
	err = d.db.Model(info).Update("point", info.Point).Error
	return
}