package dao

import (
	"backend/app/service/user/money/model"
	"backend/pkg/json"
	"github.com/go-redis/redis/v7"
	"strconv"
	"time"
)

const (
	EmptyRecordsJSON = "[]"
	KeyPrefix        = "ncs:user:money:"
	TenDays          = 10 * 24 * time.Hour
)

func (d *dao) AddRecord(uid int64, amount int32, reason string) (err error) {
	info := &model.Record{
		Timestamp: time.Now().Unix(),
		Amount:    amount,
		Reason:    reason,
	}

	// DB
	data, err := info.JSON()
	if err != nil {
		return
	}
	key := Key(uid)
	err = d.cache.LPush(key, string(data)).Err()
	if err != nil {
		return
	}
	err = d.cache.Expire(key, TenDays).Err()
	return
}

func (d *dao) GetRecords(uid int64) (res *model.Records, err error) {
	res = &model.Records{}

	// DB
	jsons, err := d.cache.LRange(Key(uid), 0, -1).Result()
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
	return KeyPrefix + strconv.FormatInt(uid, 10) + ":" + strconv.Itoa(time.Now().YearDay())
}
