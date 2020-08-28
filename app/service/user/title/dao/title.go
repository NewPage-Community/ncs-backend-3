package dao

import (
	"backend/app/service/user/title/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *dao) Title(uid int64) (res *model.Title, err error) {
	res = &model.Title{}

	// DB
	DBRes := d.db.Where(uid).First(&res)
	err = DBRes.Error
	if err == gorm.ErrRecordNotFound {
		res.UID = uid
		err = d.db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(res).Error
	}
	return
}

func (d *dao) Update(title *model.Title) (err error) {
	// DB
	err = d.db.Model(title).Select("custom_title", "title_type").Updates(*title).Error
	return
}
