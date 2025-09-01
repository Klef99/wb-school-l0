package di

import (
	"fmt"
	"log/slog"

	"github.com/google/wire"

	"github.com/klef99/wb-school-l0/internal/app/config"
	"github.com/klef99/wb-school-l0/pkg/redis"
)

var CacheSet = wire.NewSet(
	ProvideRedis,
	redis.NewCacheManager,
)

func ProvideRedis(cfg *config.Config, logger *slog.Logger) (*redis.Redis, func(), error) {
	rds, err := redis.NewRedisClient(logger, cfg.Redis.DSN(), cfg.Redis.TTL, redis.MaxRetries(cfg.Redis.MaxRetries))
	if err != nil {
		return nil, nil, fmt.Errorf("unable to connect to redis: %w", err)
	}

	cleanup := func() {
		logger.Info("Closing redis connection")

		if err := rds.Close(); err != nil {
			logger.Error("Unable to close redis connection: %v", err)
		}

		logger.Info("Redis connection closed")
	}

	return rds, cleanup, nil
}
