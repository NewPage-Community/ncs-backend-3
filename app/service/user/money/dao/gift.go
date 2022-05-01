package dao

import (
	"backend/app/service/user/money/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	GiftLimit = 200
)

func (d *dao) Gift(uid, target uint64, money uint32) (remaining uint32, err error) {
	err = d.db.Transaction(func(tx *gorm.DB) error {
		gifts := []model.Gift{}
		if err = tx.Where("created_at >= ? AND uid = ?", getZeroTime(time.Now()), uid).Find(&gifts).Error; err != nil {
			return err
		}

		var amount int32
		for _, v := range gifts {
			amount += v.Money
		}
		if uint32(amount)+money >= GiftLimit {
			err = fmt.Errorf("reach gift limit %d", amount)
			return err
		}
		remaining = uint32(GiftLimit - (amount + int32(money)))

		if err = d.pay(tx, int64(uid), money, fmt.Sprintf("赠送 UID:%d 软妹币", target)); err != nil {
			return err
		}
		if err = d.give(tx, int64(target), money, fmt.Sprintf("UID:%d 赠送软妹币", uid)); err != nil {
			return err
		}
		return tx.Create(&model.Gift{
			UID:    uid,
			Target: target,
			Money:  int32(money),
		}).Error
	})
	return
}

func getZeroTime(d time.Time) time.Time {
	local, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, local)
}
