package handler

import (
	"context"
	"net/http"

	"goproj/redis"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var ctx = context.Background()

type Request struct {
	URL string `json:"url"`
}

func ShortenURL(c *gin.Context) {
	var req Request

	if err := c.BindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат"})
		return
	}

	shortID := uuid.New().String()[:6]

	err := redis.Client.Set(ctx, shortID, req.URL, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении в Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short": "http://localhost:8080/" + shortID,
	})
}
