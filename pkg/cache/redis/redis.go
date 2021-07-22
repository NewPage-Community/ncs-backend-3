package redis

import (
	"backend/pkg/conf"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v7"
	"time"
)

var _defaultConfig = redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "password",
}

func loadConf() *redis.Options {
	c := &struct {
		Redis *redis.Options
	}{}
	*(c.Redis) = _defaultConfig
	conf.Load(c)
	return c.Redis
}

func Init() *redis.Client {
	var err error
	opts := loadConf()
	client := redis.NewClient(opts)

	// Retry
	for i := 0; i < 3; i++ {
		_, err = client.Ping().Result()
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic(err)
	}
	return client
}

func InitMock() *redis.Client {
	test, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    test.Addr(),
	})
}

func Healthy(r *redis.Client) bool {
	return r.Ping().Err() == nil
}
