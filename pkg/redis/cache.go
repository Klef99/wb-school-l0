package redis

import (
	"context"
	"encoding/json"
)

type cache struct {
	conn *Redis
}

func NewCache(conn *Redis) *cache {
	return &cache{conn: conn}
}

type Cache interface {
	Set(ctx context.Context, key string, val interface{}) error
	Get(ctx context.Context, key string, dest interface{}) error
	Health(ctx context.Context) error
	Close() error
}

func (c *cache) Health(ctx context.Context) error {
	return c.conn.client.Ping(ctx).Err()
}

func (c *cache) Close() error {
	return c.conn.client.Close()
}

func (c *cache) Set(ctx context.Context, key string, val interface{}) error {
	err := c.Health(ctx)
	if err != nil {
		return err
	}

	msg, err := json.Marshal(val)
	if err != nil {
		return err
	}

	res := c.conn.client.Set(ctx, key, msg, c.conn.TTL)

	return res.Err()
}

func (c *cache) Get(ctx context.Context, key string, dest interface{}) error {
	err := c.Health(ctx)
	if err != nil {
		return err
	}

	res, err := c.conn.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, dest)
	if err != nil {
		return err
	}

	return nil
}

type CacheManager interface {
	GetCache() Cache
	Close() error
}

type cacheManager struct {
	cache *cache
}

func (c *cacheManager) GetCache() Cache {
	return c.cache
}

func (c *cacheManager) Close() error {
	return c.cache.Close()
}

func NewCacheManager(ch *Redis) CacheManager {
	return &cacheManager{
		cache: NewCache(ch),
	}
}
