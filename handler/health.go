package handler

import (
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	SendResponseData(c, "go-restful is ok")
}
