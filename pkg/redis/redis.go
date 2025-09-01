package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
	logger *slog.Logger
	TTL    time.Duration
}

func NewRedisClient(logger *slog.Logger, url string, ttl time.Duration, opts ...Options) (*Redis, error) {
	cfg, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %w", err)
	}

	for _, opt := range opts {
		opt(cfg)
	}
	cfg.DialTimeout = 5 * time.Millisecond
	client := redis.NewClient(cfg)

	if err := client.Ping(context.Background()).Err(); err != nil {
		logger.Warn("failed to connect to redis")
	}

	return &Redis{client: client, TTL: ttl, logger: logger}, nil
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *Redis) Close() error {
	return r.client.Close()
}
