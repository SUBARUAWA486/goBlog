package dao

import (
	"strings"
	"time"
	"fmt"

	"goBlog/config"
	"goBlog/utils"
)

func RedisDeleteToken(userID string) error {
	key := fmt.Sprintf("login:token:%s", userID)
	return config.RedisClient.Del(config.Ctx, key).Err()
}

func RedisIncrPostView(postID string) error {
	key := "post:view:" + postID
	return config.RedisClient.Incr(config.Ctx, key).Err()
}

func RedisScanPostViewKeys() ([]string, error) {
	var keys []string
	var cursor uint64
	for {
		scanKeys, nextCursor, err := config.RedisClient.Scan(config.Ctx, cursor, "post:view:*", 100).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, scanKeys...)
		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}
	return keys, nil
}

// 获取分布式锁
func RedisAcquireViewSyncLock(key string, ttl time.Duration) (string, bool) {
	return utils.AcquireLock(key, ttl)
}

// 释放锁
func RedisReleaseViewSyncLock(key, val string) {
	utils.ReleaseLock(key, val)
}

// 获取某个 key 的浏览数
func RedisGetPostViewCount(postID string) (int, error) {
	key := "post:view:" + postID
	return config.RedisClient.Get(config.Ctx, key).Int()
}

// 删除 key
func RedisDeletePostViewKey(postID string) error {
	key := "post:view:" + postID
	return config.RedisClient.Del(config.Ctx, key).Err()
}

// 获取 postID 从 key
func RedisExtractPostIDFromKey(key string) string {
	return strings.TrimPrefix(key, "post:view:")
}