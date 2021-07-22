package redis

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v7"
	"time"
)

func Init(opts *redis.Options) *redis.Client {
	var err error
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
	return Init(&redis.Options{
		Network: "tcp",
		Addr:    test.Addr(),
	})
}

func Healthy(r *redis.Client) bool {
	return r.Ping().Err() == nil
}
