package redis

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v7"
)

func Init(opts *redis.Options) *redis.Client {
	client := redis.NewClient(opts)
	_, err := client.Ping().Result()
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
	return Init(&redis.Options{
		Network: "tcp",
		Addr:    test.Addr(),
	})
}
