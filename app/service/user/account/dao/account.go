package dao

import (
	"backend/app/service/user/account/model"
	"backend/pkg/ecode"
	"backend/pkg/json"
	"gorm.io/gorm"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
	"google.golang.org/grpc/codes"
)

func (d *dao) UID(steamID int64) (res *model.Info, err error) {
	res = &model.Info{}

	// Cache
	cacheRes, err := d.cache.HGet(res.UIDKey(), strconv.FormatInt(steamID, 10)).Int64()
	if err == nil && cacheRes != 0 {
		res.UID = cacheRes
		return
	}
	if err == redis.Nil {
		err = nil
	}

	// DB
	err = d.db.Where(&model.Info{SteamID: steamID}).First(res).Error
	// Not found
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found: SteamID(%v)", steamID)
		return
	}

	// Set cache
	err = d.cache.HSet(res.UIDKey(), strconv.FormatInt(steamID, 10), res.UID).Err()
	return
}

func (d *dao) Info(uid int64) (res *model.Info, err error) {
	res = &model.Info{}

	// Cache (json)
	cacheRes, err := d.cache.HGet(res.InfoKey(), strconv.FormatInt(uid, 10)).Result()
	if err == nil && cacheRes != "" {
		err = json.Unmarshal([]byte(cacheRes), res)
		if err == nil && res.IsValid() {
			return
		}
	}
	if err == redis.Nil {
		err = nil
	}

	// DB
	err = d.db.Where(&model.Info{UID: uid}).First(res).Error
	// Not found
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found: UID(%v)", uid)
		return
	}

	// Set cache
	CacheData, err := json.Marshal(res)
	if err != nil {
		return
	}
	err = d.cache.HSet(res.InfoKey(), strconv.FormatInt(uid, 10), string(CacheData)).Err()
	return
}

func (d *dao) Register(steamID int64) (res *model.Info, err error) {
	res = &model.Info{SteamID: steamID, FirstJoin: time.Now().UTC().Unix()}

	// DB
	dbRes := d.db.Create(res)
	err = dbRes.Error
	if err != nil {
		return
	}

	// Set cache
	CacheData, err := json.Marshal(res)
	if err != nil {
		return
	}
	err = d.cache.HSet(res.InfoKey(), strconv.FormatInt(res.UID, 10), string(CacheData)).Err()
	return
}

func (d *dao) ChangeName(info *model.Info) (err error) {
	// DB
	err = d.db.Model(info).Update("username", info.Username).Error
	if err != nil {
		return
	}

	// Remove cache
	err = d.cache.HDel(info.InfoKey(), strconv.FormatInt(info.UID, 10)).Err()
	return
}

func (d *dao) GetAllUID() (res *[]int64, err error) {
	res = &[]int64{}
	var users []model.Info

	// DB
	err = d.db.Find(&users).Error

	for _, user := range users {
		*res = append(*res, user.UID)
	}
	return
}
