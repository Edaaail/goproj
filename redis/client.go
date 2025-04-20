package redis

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

func Init() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func SaveURL(code, longURL string) error {
	return Rdb.Set(Ctx, code, longURL, 0).Err()
}
