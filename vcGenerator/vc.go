package vcGenerator

import "cmbProject/redisClient"

func GetQuestion() (qId string, answer string, content string, err error) {
	return redisClient.RandomGetQuestion()
}
