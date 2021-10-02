package redis

import (
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/openmesh/booking"
)

func NewRedisCache(opts *redis.Options) booking.Cache {
	cl := redis.NewClient(opts)
	ca := cache.New(&cache.Options{
		Redis: cl,
	})

	return &redisCache{cache: ca, client: cl}
}

type redisCache struct {
	client *redis.Client
	cache  *cache.Cache
}

func (c *redisCache) Get(ctx context.Context, key string, val interface{}) error {
	return c.cache.Get(ctx, key, val)
}

func (c *redisCache) Refresh(ctx context.Context, key string) error {
	panic("not implemented") // TODO: Implement
}

func (c *redisCache) Remove(ctx context.Context, key string) error {
	return c.cache.Delete(ctx, key)
}

func (c *redisCache) RemoveMany(ctx context.Context, match string) error {
	iter := c.client.Scan(ctx, 0, match, 0).Iterator()
	for iter.Next(ctx) {
		err := c.cache.Delete(ctx, iter.Val())
		if err != nil {
			return err
		}
	}
	err := iter.Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *redisCache) Set(ctx context.Context, key string, val interface{}) error {
	return c.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: val,
	})
}

// func (c *redisCache) Scan(ctx context.Context, match string) ([]string, error) {
// 	var keys []string
// 	iter := c.client.Scan(ctx, 0, match, 0).Iterator()
// 	for iter.Next(ctx) {
// 		keys = append(keys, iter.Val())
// 	}
// 	err := iter.Err()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return keys, nil
// }
