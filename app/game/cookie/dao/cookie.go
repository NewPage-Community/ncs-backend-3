package dao

import (
	"backend/app/game/cookie/model"
	"github.com/go-redis/redis/v7"
)

func (d *dao) Get(uid int64) (res *model.Cookie, err error) {
	res = &model.Cookie{
		UID: uid,
	}
	json, err := d.db.HGet(res.Key(), res.GetUID()).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil
		}
		return
	}
	err = res.GetCookieFromJSON(json)
	return
}

func (d *dao) Set(uid int64, key string, value string) (err error) {
	c := &model.Cookie{
		UID:    uid,
		Cookie: make(map[string]string),
	}

	// Get cookie
	json, err := d.db.HGet(c.Key(), c.GetUID()).Result()
	if err != redis.Nil {
		if err != nil {
			return
		}
		err = c.GetCookieFromJSON(json)
		if err != nil {
			return
		}
	}

	// Set cookie
	c.Cookie[key] = value

	// Save cookie
	json, err = c.CookieJSON()
	if err != nil {
		return
	}
	err = d.db.HSet(c.Key(), c.GetUID(), json).Err()
	return
}
