package main

import (
	"goproj/handler"
	"goproj/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	redis.Init()

	r.POST("/shorten", handler.ShortenURL)

	r.Run(":8080")
}
