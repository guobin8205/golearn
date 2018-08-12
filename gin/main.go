package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guobin8205/api_demo/middleware"
	"github.com/guobin8205/api_demo/utils/config"
	"github.com/DeanThompson/ginpprof"
	"github.com/guobin8205/api_demo/routers"
	"net/http"
	"time"
)

func main(){
	router := gin.Default()
	router.Static("data", "./data")

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ParseRequest())
	if conf.RunMode == "pro" {
		gin.SetMode(gin.ReleaseMode)
	}
	ginpprof.Wrap(router)
	routers.SetRouters(router)

	s := &http.Server{
		Addr:           conf.HttpPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
