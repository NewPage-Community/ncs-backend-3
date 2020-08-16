package dao

import (
	"backend/app/service/user/sign/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Info get info
func (d *dao) Info(uid int64) (res *model.Sign, err error) {
	res = &model.Sign{}

	// DB
	err = d.db.Where(uid).First(res).Error
	if err == gorm.ErrRecordNotFound {
		res.UID = uid
		err = d.create(uid)
	}
	return
}

func (d *dao) Sign(uid int64) (alreadySigned bool, err error) {
	info := &model.Sign{
		UID: uid,
	}
	alreadySigned = true

	// DB
	err = d.db.First(info).Error
	if err == gorm.ErrRecordNotFound {
		err = d.create(uid)
	}
	if err != nil {
		return
	}

	err = d.db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where(uid).First(info).Error
		if err != nil {
			return
		}

		if alreadySigned = info.IsSigned(); alreadySigned {
			return
		}
		info.Sign()

		err = tx.Model(info).Updates(*info).Error
		return
	})
	return
}

func (d *dao) create(uid int64) (err error) {
	info := &model.Sign{
		UID: uid,
	}
	return d.db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(info).Error
}
