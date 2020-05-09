package dao

import (
	"backend/app/user/account/model"
	"backend/pkg/ecode"
	"backend/pkg/json"
	"fmt"
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
	var notFound interface{}
	res = &model.Info{}
	defer errorProcess(err, notFound)

	// Cache
	cacheRes, err := d.cache.HGet(UIDCache, strconv.FormatInt(steamID, 10)).Int64()
	if err == nil && cacheRes != 0 {
		res.UID = cacheRes
		return
	}

	// DB
	dbRes := d.db.Where(&model.Info{SteamID: steamID}).First(res)
	err = dbRes.Error
	// Not found
	if dbRes.RecordNotFound() || !res.IsValid() {
		notFound = fmt.Sprintf("SteamID(%v)", steamID)
		return
	}

	// Set cache
	err = d.cache.HSet(UIDCache, strconv.FormatInt(steamID, 10), res.UID).Err()
	return
}

func (d *dao) Info(uid int64) (res *model.Info, err error) {
	var notFound interface{}
	res = &model.Info{}
	defer errorProcess(err, notFound)

	// Cache (json)
	cacheRes, err := d.cache.HGet(InfoCache, strconv.FormatInt(uid, 10)).Result()
	if err == nil && cacheRes != "" {
		err = json.Unmarshal([]byte(cacheRes), res)
		if err == nil && res.IsValid() {
			return
		}
	}

	// DB
	dbRes := d.db.Where(&model.Info{UID: uid}).First(res)
	err = dbRes.Error
	// Not found
	if dbRes.RecordNotFound() || !res.IsValid() {
		notFound = fmt.Sprintf("UID(%v)", uid)
		return
	}

	// Set cache
	CacheData, err := json.Marshal(res)
	if err != nil {
		return
	}
	err = d.cache.HSet(InfoCache, strconv.FormatInt(uid, 10), string(CacheData)).Err()
	return
}

func (d *dao) Register(steamID int64) (res *model.Info, err error) {
	res = &model.Info{SteamID: steamID, Username: "", FirstJoin: time.Now().Unix()}
	defer errorProcess(err, nil)

	// DB
	if d.db.NewRecord(res) {
		dbRes := d.db.Create(res)
		err = dbRes.Error
	}

	// Set cache
	CacheData, err := json.Marshal(res)
	if err != nil {
		return
	}
	err = d.cache.HSet(InfoCache, strconv.FormatInt(res.UID, 10), string(CacheData)).Err()
	return
}

func (d *dao) ChangeName(info *model.Info) (err error) {
	defer errorProcess(err, nil)
	// DB
	err = d.db.Model(info).Update("username", info.Username).Error
	if err != nil {
		return
	}

	// Remove cache
	err = d.cache.HDel(InfoCache, strconv.FormatInt(info.UID, 10)).Err()
	return
}

func errorProcess(err error, notFound interface{}) {
	if notFound != nil {
		err = ecode.Errorf(codes.NotFound, "Can not found: %v", notFound)
	} else if err != nil && err != redis.Nil {
		err = ecode.Errorf(codes.Unknown, "dao: %v", err)
	}
}
