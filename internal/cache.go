package internal

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-redis/redis"
)

// cache decorator for retrieve database data
type cache struct {
	redis   *redis.Client
	adapter adapter
	log     *slog.Logger
}

func newCache(redis *redis.Client, adapter adapter, log *slog.Logger) *cache {
	return &cache{
		redis:   redis,
		adapter: adapter,
		log:     log,
	}
}

func (c *cache) List(ctx context.Context, query string) error {

	value, err := c.redis.Get(query).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("failed to get value for key %q: %w", query, err)
	}

	if value != "" {
		c.log.Info("got from cache")
		return nil
	}

	if err := c.adapter.List(ctx, query); err != nil {
		return fmt.Errorf("list: %w", err)
	}

	go c.redis.Set(query, "got", 10*time.Second)

	return nil
}
