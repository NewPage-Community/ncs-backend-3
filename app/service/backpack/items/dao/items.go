package dao

import "backend/app/service/backpack/items/model"

func (d *dao) GetItems() (res []*model.Item, err error) {
	res = []*model.Item{}

	// DB
	err = d.db.Find(&res).Error
	return
}
