package dao

import (
	"backend/app/service/user/money/model"
	"backend/pkg/json"
	"backend/pkg/log"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
)

const (
	KeyPrefix = "ncs:user:money:"
	TenDays   = 10 * 24 * time.Hour
)

func (d *dao) AddRecord(uid int64, amount int32, reason string) (err error) {
	now := time.Now().Unix()
	info := &model.Record{
		Timestamp: now,
		Amount:    amount,
		Reason:    reason,
	}

	// DB
	data, err := info.JSON()
	if err != nil {
		return
	}
	key := Key(uid)
	err = d.cache.ZAdd(key, &redis.Z{
		Score:  float64(now),
		Member: string(data),
	}).Err()
	return
}

func (d *dao) GetRecords(uid int64) (res *model.Records, err error) {
	res = &model.Records{}
	now := strconv.FormatInt(time.Now().Add(-TenDays).Unix(), 10)

	// DB
	key := Key(uid)
	err = d.cache.ZRemRangeByScore(key, "0", now).Err()
	if err != nil {
		log.Error(err)
	}
	jsons, err := d.cache.ZRevRange(key, 0, -1).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil
		}
		return
	}
	for _, v := range jsons {
		data := &model.Record{}
		err = json.Unmarshal([]byte(v), data)
		if err == nil {
			*res = append(*res, data)
		}
	}
	return
}

func Key(uid int64) string {
	return KeyPrefix + strconv.FormatInt(uid, 10)
}
