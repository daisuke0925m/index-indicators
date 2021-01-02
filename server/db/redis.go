package db

import (
	"os"

	"github.com/go-redis/redis/v7"
)

// Redis RedisClient
var Redis *redis.Client

// InitRedis redis起動
func InitRedis() {
	port := os.Getenv("REDIS_HOST")
	Redis = redis.NewClient(&redis.Options{
		Addr: port + ":6379",
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}
