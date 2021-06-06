package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type GetQuestionReq struct {
	ReqStr    *string `json:"req_str"`    // business request id, 用于方便客户端表示会话
	ClientIP  *string `json:"client_ip"`  // 可能的 跳转功能
	Mode      *int    `json:"mode"`       // 模式
	TimeStamp *int64  `json:"time_stamp"` // 时间戳
}

type GetQuestionRsp struct {
	ReqStr     string `json:"req_str"`
	Mode       int    `json:"mode"` // 模式
	Content    string `json:"content"`
	QuestionId string `json:"question_id"`
}

type CheckAnswerReq struct {
	ReqStr     *string `json:"req_str"`    // business request id, 用于方便客户端表示会话
	TimeStamp  *int64  `json:"time_stamp"` // 时间戳
	QuestionId *string `json:"question_id"`
	Answer     *string `json:"answer"`
	SecretId   *string `json:"secret_id"`
}

type CheckAnswerRsp struct {
	ReqStr string `json:"req_str"`
	Result int    `json:"result"`
}

func postJson1(u string, a GetQuestionReq) GetQuestionRsp {
	jsonBytes, _ := json.Marshal(a)
	url := u
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var rsp GetQuestionRsp
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		panic(err)
	}
	return rsp
}

func postJson2(u string, a CheckAnswerReq) {
	jsonBytes, _ := json.Marshal(a)
	url := u
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		for j := 0; j < 5; j++ {
			wg.Add(1)
			go func() {

				reqS := ""
				clientIp := ""
				mode := 1
				timeStamp := int64(1123123)

				a := GetQuestionReq{
					ReqStr:    &reqS,
					ClientIP:  &clientIp,
					Mode:      &mode,
					TimeStamp: &timeStamp,
				}

				a1 := postJson1("http://test-8180nlb-c580d437de8d633c.elb.cn-northwest-1.amazonaws.com.cn:8180/validation/getQuestion", a)
				answer := "123456"
				postJson2("http://test-8180nlb-c580d437de8d633c.elb.cn-northwest-1.amazonaws.com.cn:8180/validation/checkAnswer", CheckAnswerReq{
					ReqStr:     &a1.ReqStr,
					TimeStamp:  a.TimeStamp,
					QuestionId: &a1.QuestionId,
					Answer:     &answer,
					SecretId:   &answer,
				})
				wg.Done()
			}()
		}
		wg.Wait()
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
