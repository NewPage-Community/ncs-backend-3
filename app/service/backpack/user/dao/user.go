package dao

import (
	"backend/app/service/backpack/user/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *dao) Get(uid int64) (res *model.User, err error) {
	res = &model.User{}
	userModel := &model.UserModel{}

	// DB
	err = d.db.First(userModel, uid).Error
	if err == gorm.ErrRecordNotFound {
		// not found and create
		return d.Create(uid)
	}
	if err != nil {
		return
	}

	res, err = userModel.GetUser()
	return
}

func (d *dao) AddItems(uid int64, items *model.Items) (err error) {
	userModel := &model.UserModel{}

	// DB
	err = d.db.First(userModel, uid).Error
	if err == gorm.ErrRecordNotFound {
		// not found and create
		_, err = d.Create(uid)
	}
	if err != nil {
		return
	}

	err = d.db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(userModel, uid).Error
		if err != nil {
			return
		}

		user, err := userModel.GetUser()
		if err != nil {
			return
		}

		// Delete expired item first
		user.CheckItem()
		user.AddItems(items)

		userModel, err = user.GetModel()
		if err != nil {
			return
		}
		err = tx.Model(userModel).Updates(*userModel).Error
		return
	})
	return
}

func (d *dao) RemoveItem(uid int64, item model.Item, all bool) (err error) {
	userModel := &model.UserModel{}

	// DB
	err = d.db.First(userModel, uid).Error
	if err == gorm.ErrRecordNotFound {
		// not found and create
		_, err = d.Create(uid)
	}
	if err != nil {
		return
	}

	err = d.db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(userModel, uid).Error
		if err != nil {
			return
		}

		user, err := userModel.GetUser()
		if err != nil {
			return
		}

		// Delete expired item first
		user.CheckItem()
		user.RemoveItem(item, all)

		userModel, err = user.GetModel()
		if err != nil {
			return
		}
		err = tx.Model(userModel).Updates(*userModel).Error
		return
	})
	return
}

func (d *dao) Create(uid int64) (res *model.User, err error) {
	res = &model.User{
		UID:   uid,
		Items: &model.Items{},
	}
	userModel, err := res.GetModel()
	if err != nil {
		return
	}
	err = d.db.Create(userModel).Error
	return
}
