package internal

import (
	"context"

	"github.com/go-redis/redis"

	"github.com/mrbelka12000/optimizer/internal/models"
)

// cache decorator for retrieve database data
type cache struct {
	redis   *redis.Client
	adapter adapter
}

func newCache(redis *redis.Client, adapter adapter) *cache {
	return &cache{
		redis:   redis,
		adapter: adapter,
	}
}

func (c *cache) List(ctx context.Context, req models.Request) error {
	// TODO try to cache request data

	return c.adapter.List(ctx, req)
}
