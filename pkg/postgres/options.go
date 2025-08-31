package postgres

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Option func(*pgxpool.Config)

func MaxConnections(count int32) Option {
	return func(c *pgxpool.Config) {
		c.MaxConns = count
	}
}

func MinConnections(count int32) Option {
	return func(c *pgxpool.Config) {
		c.MinConns = count
	}
}

func MaxConnectionLifetime(d time.Duration) Option {
	return func(c *pgxpool.Config) {
		c.MaxConnLifetime = d
	}
}

func MaxConnectionIdleTime(d time.Duration) Option {
	return func(c *pgxpool.Config) {
		c.MaxConnIdleTime = d
	}
}
