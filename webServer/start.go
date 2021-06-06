package webServer

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

var (
	corsAllowHeaders     = "*"
	corsExposeHeaders    = "*"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func ErrorHandler(err error, ignoreErr bool, c *gin.Context, statusCode int) bool {
	if err != nil {
		log.Println("[Error]: ", err)
		if ignoreErr {
			return false
		}

		type stopStruct struct {
			ErrorStr string `json:"error"`
		}
		//_, _ = c.WriteString(err.Error())
		c.AbortWithStatusJSON(statusCode, stopStruct{ErrorStr: err.Error()})
		return true
	}
	return false
}

func requestDataNilCheck(r interface{}) error {
	t := reflect.TypeOf(r)
	v := reflect.ValueOf(r)
	log.Println(t, v)
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		fi := t.Field(i)
		ti := t.Field(i).Type
		vi := v.Field(i)
		switch ti.Kind() {
		case reflect.Ptr:
			if vi.IsNil() {
				return errors.New(fmt.Sprintf("param %s is needed", fi.Name))
			} else {
				vi = vi.Elem()
				switch vi.Kind() {
				case reflect.Struct:
					if vi.CanInterface() {
						err := requestDataNilCheck(vi.Interface())
						if err != nil {
							return err
						}
					} else {
						return errors.New("reflect cannot interface")
					}
				}
			}
		}
	}
	return nil
}

func ServerStart() {
	app := gin.New()
	app.POST("/validation/getQuestion", GetQuestion)
	app.POST("/validation/checkAnswer", CheckAnswer)
	err := app.Run(fmt.Sprintf("0.0.0.0:%d", 8180))
	if err != nil {
		panic(err)
	}
}
