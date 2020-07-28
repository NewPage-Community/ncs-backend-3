package dao

import (
	"backend/app/service/user/title/model"
	"gorm.io/gorm"
)

func (d *dao) Title(uid int64) (res *model.Title, err error) {
	res = &model.Title{}

	// DB
	DBRes := d.db.Where(uid).First(&res)
	err = DBRes.Error
	if err == gorm.ErrRecordNotFound {
		res.UID = uid
		err = d.db.Create(res).Error
	}
	return
}

func (d *dao) Update(title *model.Title) (err error) {
	// DB
	DBRes := d.db.Model(title).Updates(*title)
	err = DBRes.Error
	if DBRes.RowsAffected == 0 && err == nil {
		err = d.db.Create(title).Error
	}
	return
}
