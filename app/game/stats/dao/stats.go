package dao

import (
	"backend/app/game/stats/model"
	"backend/pkg/ecode"
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
)

var ctx = context.Background()

func (d *dao) Get(stats *model.Stats) (err error) {
	stats.Score, err = d.redis.ZScore(ctx, stats.Key(), stats.Member()).Result()
	if err != nil {
		if err == redis.Nil {
			return ecode.Errorf(codes.NotFound, "Can not found UID(%d)", stats.UID)
		}
		return ecode.Errorf(codes.Unknown, "Redis err: %s", err)
	}
	stats.Rank, err = d.redis.ZRevRank(ctx, stats.Key(), stats.Member()).Result()
	// ZSet rank start from 0
	stats.Rank++
	if err != nil {
		if err == redis.Nil {
			return ecode.Errorf(codes.NotFound, "Can not found UID(%d)", stats.UID)
		}
		return ecode.Errorf(codes.Unknown, "Redis err: %s", err)
	}
	return
}

func (d *dao) GetAll(stats *model.Stats) (res []*model.Stats, err error) {
	z, err := d.redis.ZRevRangeWithScores(ctx, stats.Key(), 0, -1).Result()
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "Redis err: %s", err)
		return
	}

	var count int64 = 1
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
	err = d.redis.ZAdd(ctx, stats.Key(), &redis.Z{
		Score:  stats.Score,
		Member: stats.Member(),
	}).Err()
	return
}

// Incr
// stats.Score is increment
func (d *dao) Incr(stats *model.Stats) (err error) {
	err = d.redis.ZIncrBy(ctx, stats.Key(), stats.Score, stats.Member()).Err()
	return
}

func (d *dao) GetPartly(stats *model.Stats, start int64, end int64) (res []*model.Stats, total int64, err error) {
	total, err = d.redis.ZCard(ctx, stats.Key()).Result()
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "Redis err: %s", err)
		return
	}

	z, err := d.redis.ZRevRangeWithScores(ctx, stats.Key(), start-1, end-1).Result()
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "Redis err: %s", err)
		return
	}

	count := start
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
