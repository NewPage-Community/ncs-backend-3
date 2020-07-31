package dao

import (
	"backend/app/service/user/vip/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *dao) Register(info *model.VIP) (err error) {
	err = d.db.Create(info).Error
	return
}

func (d *dao) Info(info *model.VIP) (err error) {
	err = d.db.Where(info.UID).First(info).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found: UID(%v)", info.UID)
	}
	return
}

func (d *dao) AddPoint(uid int64, addPoint int) (point int, err error) {
	info := &model.VIP{}
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where(uid).First(info).Error
	if err == gorm.ErrRecordNotFound {
		// Create and add point
		info.UID = uid
		info.Point = addPoint
		return info.Point,
			d.db.Create(info).Error
	}
	if err != nil {
		return
	}

	info.Point += addPoint

	err = d.db.Model(info).Updates(&model.VIP{Point: info.Point}).Error
	if err == nil {
		point = info.Point
	}
	return
}

func (d *dao) Renewal(uid int64, length int64) (exprTime int64, err error) {
	info := &model.VIP{}
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where(uid).First(info).Error
	if err == gorm.ErrRecordNotFound {
		// Create and renewal
		info.UID = uid
		info.Renewal(length)
		return info.ExpireDate,
			d.db.Create(info).Error
	}
	if err != nil {
		return
	}

	info.Renewal(length)

	err = d.db.Model(info).Updates(&model.VIP{ExpireDate: info.ExpireDate}).Error
	if err == nil {
		exprTime = info.ExpireDate
	}
	return
}
