package internal

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"

	"github.com/mrbelka12000/optimizer/internal/models"
)

type Cache struct {
	redis   *redis.Client
	adapter Adapter
}

func newCache(redis *redis.Client, adapter Adapter) *Cache {
	return &Cache{
		redis:   redis,
		adapter: adapter,
	}
}

func (c *Cache) List(ctx context.Context, pars models.Data) error {
	fmt.Println("popal")
	return c.adapter.List(ctx, pars)
}
