package router

import (
	"dubinyang1993/go-restful/handler"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// health
	g.Any("/health", handler.HealthCheck)

	// user
	// user := g.Group("/user")

	return g
}
