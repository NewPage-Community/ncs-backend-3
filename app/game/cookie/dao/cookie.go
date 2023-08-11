package dao

import (
	"backend/app/game/cookie/model"
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func (d *dao) Get(uid int64) (res *model.Cookie, err error) {
	res = &model.Cookie{
		UID: uid,
	}
	json, err := d.db.HGet(ctx, res.Key(), res.GetUID()).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil
		}
		return
	}
	err = res.MergeCookieData(json)
	return
}

func (d *dao) Set(uid int64, key string, value string, timestamp int64) (err error) {
	c := &model.Cookie{
		UID:    uid,
		Cookie: make(map[string]model.CookieModel),
	}

	// Get cookie
	json, err := d.db.HGet(ctx, c.Key(), c.GetUID()).Result()
	if err != redis.Nil {
		if err != nil {
			return
		}
		err = c.MergeCookieData(json)
		if err != nil {
			return
		}
	}

	// Set cookie
	if c.Cookie[key].LastUpdated < timestamp {
		c.Cookie[key] = model.CookieModel{
			Value:       value,
			LastUpdated: timestamp,
		}
	}

	// Save cookie
	json, err = c.CookieJSON()
	if err != nil {
		return
	}
	err = d.db.HSet(ctx, c.Key(), c.GetUID(), json).Err()
	return
}
