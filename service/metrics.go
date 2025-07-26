package service

import (
	"time"
	
	"goBlog/dao"
)

func SyncViewCountsToDB() {

	lockKey := "lock:sync:viewcount"
	lockVal, ok := dao.RedisAcquireViewSyncLock(lockKey, 3*time.Minute)
	if !ok {
		return
	}
	defer dao.RedisReleaseViewSyncLock(lockKey, lockVal)

	keys, err := dao.RedisScanPostViewKeys()
	if err != nil {
		return
	}

	for _, key := range keys {
		postID := dao.RedisExtractPostIDFromKey(key)
		count, err := dao.RedisGetPostViewCount(postID)
		if err != nil || count <= 0 {
			continue
		}

		err = dao.IncrementPostViewCountInDB(postID, count)
		if err != nil {
			continue
		}

		_ = dao.RedisDeletePostViewKey(postID)
	}
}
