package dao

import (
	"backend/app/service/user/money/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (d *dao) addRecord(tx *gorm.DB, uid int64, amount int32, reason string) (err error) {
	info := &model.Record{
		UID:    uint64(uid),
		Amount: amount,
		Reason: reason,
	}

	// DB
	return tx.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(info).Error
}

func (d *dao) GetRecords(uid int64, days uint32) (res *model.Records, err error) {
	res = &model.Records{}
	lastTime := time.Now().Add(time.Duration(days) * -24 * time.Hour).Unix()

	// DB
	err = d.db.Find(res, "created_at >= ?", lastTime).Error
	return
}
