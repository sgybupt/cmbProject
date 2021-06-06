package redisClient

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var VCAnswerRedisClient *redis.Client // 2  // questionId: answer

func SetVCAnswer(qId, answer string, expTime time.Duration) (err error) {
	_, err = VCAnswerRedisClient.Set(context.Background(), qId, answer, expTime).Result()
	return err
}

func CheckVCAnswer(qId, answer string) (pass bool, err error) {
	ans, err := VCAnswerRedisClient.Get(context.Background(), qId).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	if ans == answer {
		return true, nil
	}
	return false, err
}

func DeleteVCAnswer(qId string) (err error) {
	_, err = VCAnswerRedisClient.Del(context.Background(), qId).Result()
	return err
}
