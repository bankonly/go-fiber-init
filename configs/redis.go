package configs

import (
	"time"

	"github.com/go-redis/redis/v9"
)

var (
	Cli *redis.Client
)

type RedisCli struct {
	Cli             *redis.Client
	DefaultMaxAlive time.Duration
}

func ConnectRedisDB() {
	Cli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetRedisCli() *RedisCli {
	return &RedisCli{Cli: Cli, DefaultMaxAlive: 10 * time.Second}
}
