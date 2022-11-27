package queue

import "github.com/go-redis/redis/v8"

func ConnectQueue(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}