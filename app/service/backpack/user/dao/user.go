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
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Where(uid).First(res).
		Clauses(clause.Locking{Strength: "UPDATE"}).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	if err != nil {
		return
	}

	err = res.CheckItem()
	if err != nil {
		return
	}

	err = d.db.Model(res).Updates(*res).Error
	return
}

func (d *dao) AddItem(uid int64, item model.Item, repeat bool) (err error) {
	user := &model.User{}
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Where(uid).First(user).
		Clauses(clause.Locking{Strength: "UPDATE"}).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	if err != nil {
		return
	}

	err = user.AddItem(item, repeat)
	if err != nil {
		return
	}

	err = d.db.Model(user).Updates(*user).Error
	return
}

func (d *dao) RemoveItem(uid int64, item model.Item, all bool) (err error) {
	user := &model.User{}
	defer func() {
		// To release lock
		d.db.Commit()
	}()

	// DB
	err = d.db.Where(uid).First(user).
		Clauses(clause.Locking{Strength: "UPDATE"}).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found UID(%d)", uid)
	}
	if err != nil {
		return
	}

	err = user.RemoveItem(item, all)
	if err != nil {
		return
	}

	err = d.db.Model(user).Updates(*user).Error
	return
}

func (d *dao) Create(uid int64) (res *model.User, err error) {
	res = &model.User{
		UID:   uid,
		Items: []byte("[]"),
	}
	err = d.db.Create(res).Error
	return
}
