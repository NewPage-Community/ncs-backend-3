package dao

import "backend/app/service/auth/qq/model"

func (d *dao) GetUID(openID string) (res *model.QQConnect, err error) {
	res = &model.QQConnect{}
	err = d.db.Where("open_id = ?", openID).Find(res).Error
	return
}

func (d *dao) BindQQ(info model.QQConnect) error {
	return d.db.Create(&info).Error
}

func (d *dao) UnbindQQ(info model.QQConnect) error {
	return d.db.Delete(&info).Error
}

func (d *dao) GetStatus(uid int64) (res *model.QQConnect, err error) {
	res = &model.QQConnect{}
	err = d.db.First(res, uid).Error
	return
}
