package redisClient

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
	"strings"
	"time"
)

var VCRedisClient *redis.Client // 1  // questionId: answer_content

func RandomGetQuestion() (qId string, answer string, content string, err error) {
	qId, err = VCRedisClient.RandomKey(context.Background()).Result()
	if err != nil {
		return
	}
	answerRaw, err := VCRedisClient.Get(context.Background(), qId).Result()
	if err != nil {
		return
	}
	i := strings.IndexByte(answerRaw, '_')
	if i == -1 {
		log.Fatal(errors.New("error input vc"))
	}
	answer = answerRaw[:i]
	content = answerRaw[i+1:]
	//fmt.Println(answer, content)
	return qId, answer, content, nil
}

func SetQuestion(qId, answer, content string, expTime time.Duration) (err error) {
	_, err = VCRedisClient.Set(context.Background(), qId, answer+"_"+content, expTime).Result()
	return err
}

func DeleteQuestion(qId string) (err error) {
	_, err = VCRedisClient.Del(context.Background(), qId).Result()
	if err != nil {
		return err
	}
	return nil
}
