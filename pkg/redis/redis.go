package redis

import "github.com/go-redis/redis"

func NewRedis(address string) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: address, Password: "", DB: 0})
}
