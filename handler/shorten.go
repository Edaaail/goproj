package handler

import (
	"net/http"

	"goproj/redis"
	"goproj/utils"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

func ShortenURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.LongURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	shortCode := utils.GenerateShortCode(req.LongURL)
	shortURL := "http://localhost:8080/" + shortCode

	err := redis.SaveURL(shortCode, req.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}
