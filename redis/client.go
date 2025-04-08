package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // Подключение к локальному Redis
})

var ctx = context.Background()
