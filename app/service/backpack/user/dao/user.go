package dao

import (
	"backend/app/service/backpack/user/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *dao) Get(uid int64) (res *model.User, err error) {
	res = &model.User{}
	userModel := &model.UserModel{}
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Where(uid).First(userModel).
		Clauses(clause.Locking{Strength: "UPDATE"}).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	if err != nil {
		return
	}

	res, err = userModel.GetUser()
	if err != nil {
		return
	}
	res.CheckItem()

	userModel, err = res.GetModel()
	if err != nil {
		return
	}
	err = d.db.Model(userModel).Updates(*userModel).Error
	return
}

func (d *dao) AddItems(uid int64, items *model.Items) (err error) {
	userModel := &model.UserModel{}
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Where(uid).First(userModel).
		Clauses(clause.Locking{Strength: "UPDATE"}).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	if err != nil {
		return
	}

	user, err := userModel.GetUser()
	if err != nil {
		return
	}
	user.AddItems(items)

	userModel, err = user.GetModel()
	if err != nil {
		return
	}
	err = d.db.Model(userModel).Updates(*userModel).Error
	return
}

func (d *dao) RemoveItem(uid int64, item model.Item, all bool) (err error) {
	userModel := &model.UserModel{}
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Where(uid).First(userModel).
		Clauses(clause.Locking{Strength: "UPDATE"}).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	if err != nil {
		return
	}

	user, err := userModel.GetUser()
	if err != nil {
		return
	}
	user.RemoveItem(item, all)

	userModel, err = user.GetModel()
	if err != nil {
		return
	}
	err = d.db.Model(userModel).Updates(*userModel).Error
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