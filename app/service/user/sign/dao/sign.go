package dao

import (
	"backend/app/service/user/sign/model"
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

func (d *dao) Sign(uid int64) (alreadySigned bool, err error) {
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
		return false, d.db.Create(info).Error
	}
	if err != nil {
		return
	}

	if info.IsSigned() {
		return true, nil
	}
	info.Sign()

	err = d.db.Model(info).Updates(*info).Error
	return
}
