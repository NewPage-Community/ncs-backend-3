package dao

import (
	"backend/app/user/account/model"
)

// TODO 数据库逻辑

func (d *dao) UID(steamID int64) (uid int64, err error) {
	return
}

func (d *dao) Info(uid int64) (info *model.Info, err error) {
	return
}

func (d *dao) Register(steamID int64) (uid int64, err error) {
	return
}
