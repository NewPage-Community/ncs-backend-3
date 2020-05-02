package redis

import "github.com/go-redis/redis/v7"

func Init(opts *redis.Options) *redis.Client {
	client := redis.NewClient(opts)
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
