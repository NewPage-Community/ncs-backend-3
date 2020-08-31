package dao

import "backend/app/game/cvar/model"

func (d *dao) GetCVars() (res []*model.CVar, err error) {
	res = make([]*model.CVar, 0)

	// DB
	err = d.db.Find(res).Error
	return
}

func (d *dao) UpdatedCVar(id int) (err error) {
	info := &model.CVar{ID: id}
	// DB
	err = d.db.Model(info).Updates(model.CVar{
		NeedUpdate: false,
	}).Error
	return
}
