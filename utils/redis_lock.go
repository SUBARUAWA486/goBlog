package utils

import (
	"time"
	"goBlog/config"
	
	"github.com/google/uuid"
)

// 获取锁，返回锁标识（value）和是否成功
func AcquireLock(key string, ttl time.Duration) (string, bool) {
	ctx := config.Ctx
	val := uuid.New().String() // 用 UUID 作为唯一标识
	ok, err := config.RedisClient.SetNX(ctx, key, val, ttl).Result()
	if err != nil || !ok {
		return "", false
	}
	return val, true
}

// 释放锁（需要判断是不是自己加的锁）
func ReleaseLock(key, value string) {
	ctx := config.Ctx
	script := `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
	`
	_ = config.RedisClient.Eval(ctx, script, []string{key}, value).Err()
}
