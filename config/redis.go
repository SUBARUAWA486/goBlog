package config

import (
	"context"
	"log"
	
	"github.com/redis/go-redis/v9"
	
)

var (
	RedisClient *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := RedisClient.Ping(Ctx).Result(); err != nil {
		log.Fatalf("Redis 连接失败: %v", err)
	}
}