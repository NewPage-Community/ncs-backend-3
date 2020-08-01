package dao

import (
	"backend/app/service/user/money/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	EmptyRecordsJSON = "[]"
)

func (d *dao) AddRecord(uid int64, amount int32, reason string) (err error) {
	info := &model.RecordsDB{
		UID:     uid,
		Records: []byte(EmptyRecordsJSON),
	}

	// DB
	err = d.db.First(info).Error
	if err == gorm.ErrRecordNotFound {
		err = d.db.Create(info).Error
	}
	if err != nil {
		return
	}

	err = d.db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Clauses(clause.Locking{
			Strength: "UPDATE",
		}).First(info, uid).Error
		if err != nil {
			return
		}

		err = info.Add(amount, reason)
		if err != nil {
			return
		}

		err = tx.Model(info).Updates(*info).Error
		return
	})
	return
}

func (d *dao) GetRecords(uid int64) (*model.Records, error) {
	var err error
	info := &model.RecordsDB{}

	// DB
	res := d.db.Find(info, uid)
	err = res.Error
	if res.RowsAffected == 0 {
		info.Records = []byte(EmptyRecordsJSON)
	}
	if err != nil {
		return nil, err
	}
	return info.Get()
}
