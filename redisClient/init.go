package redisClient

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

//var BusinessRedisClient *redis.Client // 0  // businessId: secretKey
//var VCRedisClient *redis.Client       // 1  // questionId: answer_content
//var VCAnswerRedisClient *redis.Client // 2  // questionId: answer
//var SessionRedisClient *redis.Client // 3  // sessionId: success

func init() {
	var err error
	brc, err := createClient(0)
	if err != nil {
		log.Fatal(err)
	}
	BusinessRedisClient = brc

	vcrc, err := createClient(1)
	if err != nil {
		log.Fatal(err)
	}
	VCRedisClient = vcrc

	vcarc, err := createClient(2)
	if err != nil {
		log.Fatal(err)
	}
	VCAnswerRedisClient = vcarc

	src, err := createClient(3)
	if err != nil {
		log.Fatal(err)
	}
	SessionRedisClient = src

	_, err = BusinessRedisClient.FlushAll(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	// init
	_, err = BusinessRedisClient.Set(context.Background(),
		"9f2d6615-59be-4af0-96d9-b15f841d6ca3",
		"076ba636-b761-46ed-9b9e-8bfdfa07ff95",
		0).Result()
	if err != nil {
		log.Fatal(err)
	}
	_, err = VCRedisClient.Set(context.Background(),
		"1b7129a1-a1c5-4d69-81e4-76f3c48b24bf",
		"123456"+"_"+`<img class="lnXdpd" alt="Google" height="92" src="/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png" srcset="/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png 1x, /images/branding/googlelogo/2x/googlelogo_color_272x92dp.png 2x" width="272" data-atf="1">`, 0).Result()
	if err != nil {
		log.Fatal(err)
	}
}

func createClient(i int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "sc-cluster.zepyep.ng.0001.cnw1.cache.amazonaws.com.cn:6379",
		Password: "",
		DB:       i,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
