package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // connecting redis
})

func main() {
	r := gin.Default()

	r.POST("/shorten", func(c *gin.Context) {

		var json struct {
			URL string `json:"url"`
		}
		if err := c.BindJSON(&json); err != nil || json.URL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат"})
			return
		}

		shortID := uuid.New().String()[:6]

		err := rdb.Set(ctx, shortID, json.URL, 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка Redis"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"short": "http://localhost:8080/" + shortID,
		})
	})

	r.Run(":8080")
}
