package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	// async listen server, run exit signal code
	go func() {
		log.Println("Start to listening the incoming requests on HTTP address" + addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Listen: ", err)
		}
	}()

	// receive exit signal
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server...")

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("Server Shutdown:", err)
	}

	log.Println("Server exited.")
}
