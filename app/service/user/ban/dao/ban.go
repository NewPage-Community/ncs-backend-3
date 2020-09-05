package dao

import (
	"backend/app/service/user/ban/model"
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"time"
)

const (
	DayTime = 24 * time.Hour
)

func (d *dao) Info(uid uint64) (res *model.Ban, err error) {
	res = &model.Ban{}

	// DB
	err = d.db.Where(&model.Ban{UID: uid}).Order("expire_time desc").First(res).Error
	if err == gorm.ErrRecordNotFound {
		// Not found is not error
		// We should return empty info with uid
		err = nil
		res = &model.Ban{
			ID:         0,
			UID:        uid,
			CreateTime: 0,
			ExpireTime: 0,
			Type:       0,
			ServerID:   0,
			ModID:      0,
			GameID:     0,
			Reason:     "",
		}
	}
	return
}

func (d *dao) Add(info *model.Ban) (err error) {
	// DB
	if info.IsValid() {
		err = d.db.Create(info).Error
	}
	return
}

func (d *dao) Remove(info *model.Ban) (err error) {
	// IMPORTANT!!! id empty will delete all recode!!!
	if !info.IsValid() {
		err = ecode.Errorf(codes.InvalidArgument, "ID is empty")
		return
	}

	// DB
	DBres := d.db.Delete(info)
	err = DBres.Error
	if DBres.RowsAffected == 0 {
		err = ecode.Errorf(codes.NotFound, "Can not found recode ID(%d)", info.ID)
	}
	return
}

func (d *dao) IsBlockIP(ip string) (bool, error) {
	info := &model.Ban{}

	err := d.db.Where(
		"ip = ? AND create_time > ?",
		ip,
		time.Now().Add(-DayTime).Unix(),
	).First(info).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}
