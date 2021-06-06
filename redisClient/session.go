package redisClient

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var SessionRedisClient *redis.Client // 3  // sessionId: success

func SetSession(sId, res string, expTime time.Duration) (err error) {
	_, err = SessionRedisClient.Set(context.Background(), sId, res, expTime).Result()
	return err
}

func CheckSession(sId string) (pass bool, err error) {
	val, err := SessionRedisClient.Get(context.Background(), sId).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	if val == "success" {
		return true, nil
	}
	return false, nil
}

func DeleteSession(sId string) (err error) {
	_, err = SessionRedisClient.Del(context.Background(), sId).Result()
	if err != nil {
		return err
	}
	return nil
}
