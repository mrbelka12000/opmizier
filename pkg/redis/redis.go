package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

// New
func New() (*redis.Client, error) {
	redisAddr, ok := os.LookupEnv("REDIS_ADDR")
	if !ok {
		return nil, fmt.Errorf("environment variable REDIS_ADDR not set")
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return client, nil
}
