package main

import (
	"log"
	"net/http"

	"dubinyang1993/go-restful/config"
	"dubinyang1993/go-restful/mysql"
	"dubinyang1993/go-restful/redis"
	"dubinyang1993/go-restful/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// init config
	config.Init()

	// init mysql
	mysql.Init()

	// init redis
	redis.Init()

	// gin setMode before gin.New()
	gin.SetMode(viper.GetString("runmode"))

	// create gin engine
	g := gin.New()

	// middlewares
	middlewares := []gin.HandlerFunc{
		gin.Recovery(),
	}

	// routers
	handler := router.Load(
		g,
		middlewares...,
	)

	// server
	addr := viper.GetString("port")
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Println("Start to listening the incoming requests on HTTP address" + addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("Listen: ", err)
	}
}
