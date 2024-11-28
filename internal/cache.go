package internal

import (
	"context"

	"github.com/go-redis/redis"
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

func (c *cache) List(ctx context.Context, query string) error {
	// TODO try to cache request data

	return c.adapter.List(ctx, query)
}
