package dao

import (
	"backend/app/user/account/model"
	"backend/pkg/ecode"
	"backend/pkg/json"
	"github.com/go-redis/redis/v7"
	"google.golang.org/grpc/codes"
	"strconv"
	"time"
)

const (
	CacheKeyPrefix = "ncs_account_"
	InfoCache      = CacheKeyPrefix + "info"
	UIDCache       = CacheKeyPrefix + "uid"
)

func (d *dao) UID(steamID int64) (res *model.Info, err error) {
	res = &model.Info{}

	// Cache
	cacheRes, err := d.cache.HGet(UIDCache, strconv.FormatInt(steamID, 10)).Int64()
	if err == nil && cacheRes != 0 {
		res.UID = cacheRes
		return
	}
	if err != nil && err != redis.Nil {
		err = ecode.Errorf(codes.Unknown, "Redis: %v", err)
	}

	// DB
	dbRes := d.db.Where(&model.Info{SteamID: steamID}).First(res)
	if dbRes.Error != nil {
		err = ecode.Errorf(codes.Unknown, "DB: %v", dbRes.Error)
		return
	}
	// Not found
	if dbRes.RecordNotFound() || !res.IsValid() {
		err = ecode.Errorf(codes.NotFound, "Can not found steamID: %v", steamID)
		return
	}

	// Set cache
	err = d.cache.HSet(UIDCache, strconv.FormatInt(steamID, 10), res.UID).Err()
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "Redis: %v", err)
	}
	return
}

func (d *dao) Info(uid int64) (res *model.Info, err error) {
	res = &model.Info{}

	// Cache (json)
	cacheRes, err := d.cache.HGet(InfoCache, strconv.FormatInt(uid, 10)).Result()
	if err == nil && cacheRes != "" {
		err = json.Unmarshal([]byte(cacheRes), res)
		if err == nil && res.IsValid() {
			return
		}
	}
	if err != nil && err != redis.Nil {
		err = ecode.Errorf(codes.Unknown, "Redis: %v", err)
	}

	// DB
	dbRes := d.db.Where(&model.Info{UID: uid}).First(res)
	if dbRes.Error != nil {
		err = ecode.Errorf(codes.Unknown, "DB: %v", dbRes.Error)
		return
	}
	// Not found
	if dbRes.RecordNotFound() || !res.IsValid() {
		err = ecode.Errorf(codes.NotFound, "Can not found uid: %v", uid)
		return
	}

	// Set cache
	CacheData, err := json.Marshal(res)
	if err != nil {
		return
	}
	err = d.cache.HSet(InfoCache, strconv.FormatInt(uid, 10), string(CacheData)).Err()
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "Redis: %v", err)
	}
	return
}

func (d *dao) Register(steamID int64) (res *model.Info, err error) {
	res = &model.Info{SteamID: steamID, Username: "", FirstJoin: time.Now().Unix()}

	// DB
	if d.db.NewRecord(res) {
		dbRes := d.db.Create(res)
		if dbRes.Error != nil {
			err = ecode.Errorf(codes.Unknown, "DB: %v", dbRes.Error)
			return
		}
	}

	// Set cache
	CacheData, err := json.Marshal(res)
	if err != nil {
		return
	}
	err = d.cache.HSet(InfoCache, strconv.FormatInt(res.UID, 10), string(CacheData)).Err()
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "Redis: %v", err)
	}
	return
}

func (d *dao) ChangeName(info *model.Info) (err error) {
	// DB
	err = d.db.Model(info).Update("username", info.Username).Error
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "DB: %v", err)
		return
	}

	// Remove cache
	err = d.cache.HDel(InfoCache, strconv.FormatInt(info.UID, 10)).Err()
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "Redis: %v", err)
	}
	return
}
