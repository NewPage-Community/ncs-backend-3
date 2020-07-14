package dao

import (
	"backend/app/service/user/money/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *dao) AddRecord(uid int64, amount int32, reason string) (err error) {
	info := &model.RecordsDB{
		UID:     uid,
		Records: []byte("[]"),
	}
	defer func() {
		d.db.Commit()
	}()

	// DB
	err = d.db.Clauses(clause.Locking{
		Strength: "UPDATE",
	}).First(info, uid).Error
	if err1 := info.Add(amount, reason); err1 != nil {
		return err1
	}
	if err != nil {
		// not found and create
		if err == gorm.ErrRecordNotFound {
			err = d.db.Create(info).Error
		}
		return
	}
	err = d.db.Model(info).Updates(*info).Error
	return
}

func (d *dao) GetRecords(uid int64) (*model.Records, error) {
	var err error
	info := &model.RecordsDB{}

	// DB
	res := d.db.Find(info, uid)
	err = res.Error
	if res.RowsAffected == 0 {
		info.Records = []byte("[]")
	}
	if err != nil {
		return nil, err
	}
	return info.Get()
}
