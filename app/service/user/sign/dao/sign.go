package dao

import (
	"backend/app/service/user/sign/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// Info get info
func (d *dao) Info(uid int64) (res *model.Sign, err error) {
	res = &model.Sign{}

	// DB
	err = d.db.Where(uid).First(res).Error
	if err == gorm.ErrRecordNotFound {
		res.UID = uid
		res.SignTime = time.Unix(0, 0)
		err = d.db.Create(res).Error
	}
	return
}

func (d *dao) Sign(uid int64) (err error) {
	info := &model.Sign{}
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where(uid).First(info).Error
	if err == gorm.ErrRecordNotFound {
		// Create and sign
		info.UID = uid
		info.Sign()
		return d.db.Create(info).Error
	}
	if err != nil {
		return
	}

	if info.IsSigned() {
		err = ecode.Errorf(codes.Unknown, "User already signed")
		return
	}
	info.Sign()

	err = d.db.Model(info).Updates(*info).Error
	return
}
