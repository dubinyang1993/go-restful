package handler

import (
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	defer TryCatchPanic(c, "HealthCheck")

	SendResponseData(c, "go-restful is ok")
}
