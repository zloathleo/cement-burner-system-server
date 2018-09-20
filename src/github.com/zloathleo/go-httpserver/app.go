package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zloathleo/go-httpserver/context"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("server is starting...")
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Static("/public", "./public")
	context.Init(router)

	log.Println("server start at port 8080.")
	s := &http.Server{
		Addr:           ":8081",
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
