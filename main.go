package main

import (
	"goproj/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Endpoint POST /shorten
	r.POST("/shorten", handler.ShortenURL)

	r.Run(":9090")
}
