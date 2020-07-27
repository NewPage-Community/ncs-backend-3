package dao

import (
	"backend/app/game/stats/model"
	"backend/pkg/ecode"
	"github.com/go-redis/redis/v7"
	"google.golang.org/grpc/codes"
	"strconv"
)

func (d *dao) Get(stats *model.Stats) (err error) {
	stats.Score, err = d.redis.ZScore(stats.Key(), stats.Member()).Result()
	if err != nil {
		if err == redis.Nil {
			return ecode.Errorf(codes.NotFound, "Can not found UID(%d)", stats.UID)
		}
		return ecode.Errorf(codes.Unknown, "Redis err: %s", err)
	}
	stats.Rank, err = d.redis.ZRevRank(stats.Key(), stats.Member()).Result()
	if err != nil {
		if err == redis.Nil {
			return ecode.Errorf(codes.NotFound, "Can not found UID(%d)", stats.UID)
		}
		return ecode.Errorf(codes.Unknown, "Redis err: %s", err)
	}
	return
}

func (d *dao) Gets(stats *model.Stats) (res []*model.Stats, err error) {
	z, err := d.redis.ZRevRangeWithScores(stats.Key(), 0, -1).Result()
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "Redis err: %s", err)
		return
	}

	var count int64
	for _, v := range z {
		uid, err := strconv.ParseInt(v.Member.(string), 10, 64)
		if err != nil {
			continue
		}
		res = append(res, &model.Stats{
			UID:   uid,
			Score: v.Score,
			Rank:  count,
		})
		count++
	}
	return
}

func (d *dao) Set(stats *model.Stats) (err error) {
	err = d.redis.ZAdd(stats.Key(), &redis.Z{
		Score:  stats.Score,
		Member: stats.Member(),
	}).Err()
	return
}

// Incr
// stats.Score is increment
func (d *dao) Incr(stats *model.Stats) (err error) {
	err = d.redis.ZIncrBy(stats.Key(), stats.Score, stats.Member()).Err()
	return
}
