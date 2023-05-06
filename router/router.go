package router

import (
	"dubinyang1993/go-restful/handler"
	"dubinyang1993/go-restful/handler/user"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// health
	g.Any("/health", handler.HealthCheck)

	// user
	u := g.Group("/user")
	u.POST("/add", user.Add)

	return g
}
