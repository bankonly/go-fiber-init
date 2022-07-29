package configs

import (
	"github.com/go-redis/redis/v9"
)

var (
	RedisCli *redis.Client
)

func ConnectRedisDB() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
