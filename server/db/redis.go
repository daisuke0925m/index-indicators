package db

import "github.com/go-redis/redis/v7"

var Redis *redis.Client

// InitRedis redis起動
func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: "redis_container:6379",
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}
