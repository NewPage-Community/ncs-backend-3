package redis

import (
	"context"
	"time"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/go-redis/redis/v8"
)

func Init(opts *goredis.Options) *goredis.Client {
	var err error
	client := goredis.NewClient(opts)

	// Retry
	for i := 0; i < 3; i++ {
		err = client.Ping(context.Background()).Err()
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

func InitMock() *goredis.Client {
	test, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return Init(&goredis.Options{
		Network: "tcp",
		Addr:    test.Addr(),
	})
}

func Healthy(r *goredis.Client) bool {
	return r.Ping(context.Background()).Err() == nil
}
