package redisClient

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var BusinessRedisClient *redis.Client // 0  // businessId: secretKey

func SetBusiness(businessId, businessKey string) (err error) {
	_, err = BusinessRedisClient.Set(context.Background(), businessId, businessKey, 0).Result()
	return
}

func GetBusinessSK(businessId string) (key string, err error) {
	key, err = BusinessRedisClient.Get(context.Background(), businessId).Result()
	if err != nil {
		return key, err
	}
	return
}

func DelBusiness(businessId string) (err error) {
	_, err = BusinessRedisClient.Del(context.Background(), businessId).Result()
	return err
}
