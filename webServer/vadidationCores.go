package webServer

import (
	"cmbProject/redisClient"
	"cmbProject/vcGenerator"
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
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

func GetQuestion(ctx *gin.Context) {
	var err error
	var req GetQuestionReq
	err = ctx.BindJSON(&req)
	if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
		return
	}

	err = requestDataNilCheck(req)
	if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
		return
	}

	if time.Since(time.Unix(0, *req.TimeStamp)) > time.Second*60 {
		err = errors.New("expiration")
		if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
			return
		}
	}

	_, answer, content, err := vcGenerator.GetQuestion()
	if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
		return
	}

	uu, _ := uuid.NewV4()
	qId := uu.String()

	var rsp GetQuestionRsp
	rsp.ReqStr = *req.ReqStr
	rsp.Mode = *req.Mode
	rsp.Content = content
	rsp.QuestionId = qId
	ctx.JSON(200, rsp)
	if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
		return
	}

	err = redisClient.SetVCAnswer(qId, answer, time.Second*60)
	if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
		return
	}
	return
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

func CheckAnswer(ctx *gin.Context) {
	var err error
	var req CheckAnswerReq
	err = ctx.BindJSON(&req)
	if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
		return
	}

	err = requestDataNilCheck(req)
	if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
		return
	}

	if time.Since(time.Unix(0, *req.TimeStamp)) > time.Second*60 {
		err = errors.New("expiration")
		if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
			return
		}
	}

	// pre check
	pass, err := redisClient.CheckVCAnswer(*req.QuestionId, *req.Answer)
	if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
		return
	}
	var rsp CheckAnswerRsp
	if !pass {
		err = errors.New("check failed")
		if ErrorHandler(err, false, ctx, iris.StatusBadRequest) {
			return
		}
		ctx.JSON(200, rsp)
	}
	rsp.Result = 1
	rsp.ReqStr = *req.ReqStr

	ctx.JSON(200, rsp)

	_ = redisClient.DeleteVCAnswer(*req.QuestionId)
	return
}
