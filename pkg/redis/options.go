package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type Options func(*redis.Options)

func MaxRetries(retries int) Options {
	return func(o *redis.Options) {
		o.MaxRetries = retries
	}
}

func DialTimeout(d time.Duration) Options {
	return func(o *redis.Options) {
		o.DialTimeout = d
	}
}
