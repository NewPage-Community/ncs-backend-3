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
	info := &model.VIP{
		UID:   uid,
		Point: addPoint,
	}

	// DB
	err = d.db.First(info).Error
	if err == gorm.ErrRecordNotFound {
		// Create and add point
		return info.Point,
			d.db.Create(info).Error
	}

	err = d.db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(info).Error
		if err != nil {
			return
		}

		info.Point += addPoint

		err = tx.Model(info).
			Update("point", gorm.Expr("point + ?", addPoint)).Error
		if err == nil {
			point = info.Point
		}
		return
	})
	return
}

func (d *dao) Renewal(uid int64, length int64) (exprTime int64, err error) {
	info := &model.VIP{
		UID: uid,
	}

	// DB
	err = d.db.First(info).Error
	if err == gorm.ErrRecordNotFound {
		info.Renewal(length)
		return info.ExpireDate,
			d.db.Create(info).Error
	}

	err = d.db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(info).Error
		if err != nil {
			return
		}

		info.Renewal(length)

		err = tx.Model(info).
			Update("expire_date", info.ExpireDate).Error
		if err == nil {
			exprTime = info.ExpireDate
		}
		return
	})
	return
}
