package handler

import (
	"log"
	"net/http"
	"runtime"

	"dubinyang1993/go-restful/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func SendResponseData(c *gin.Context, data interface{}) {
	c.PureJSON(http.StatusOK, data)
}

func SendResponse(c *gin.Context, code int, msg string) {
	c.PureJSON(http.StatusOK, Response{
		ErrCode: code,
		ErrMsg:  msg,
	})
}

func TryCatchPanic(c *gin.Context, title string) {
	if err := recover(); err != nil {
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)
		log.Println(title, string(buf))
		SendResponse(c, errno.ApiErrorSystem.Code, errno.ApiErrorSystem.Message)
	}
}
