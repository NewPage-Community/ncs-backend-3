package dao

import (
	"backend/app/service/user/admin/model"
	"gorm.io/gorm"
)

func (d *dao) Info(uid int64) (res *model.Admin, err error) {
	res = &model.Admin{}

	// DB
	err = d.db.Where(uid).First(&res).Error
	if err == gorm.ErrRecordNotFound {
		// Not found is not error
		// We should return empty flag and 0 imm
		err = nil
		res = &model.Admin{
			UID:      uid,
			Flag:     "",
			Immunity: 0,
		}
	}
	return
}
