package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_defaultConnAttempts = 10
	_defaultConnTimeout  = 1 * time.Second
)

type Postgres struct {
	*pgxpool.Pool
	connAttempts int
	connTimeout  time.Duration
}

func New(logger *slog.Logger, url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		connAttempts: _defaultConnAttempts,
	}

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	// Custom options
	for _, opt := range opts {
		opt(cfg)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	for pg.connAttempts > 0 {
		err = pool.Ping(context.Background())
		if err != nil {
			logger.Warn("failed to connect to postgres", slog.Int("attempts", pg.connAttempts))
			pg.connAttempts--
			time.Sleep(pg.connTimeout)

			continue
		}

		break
	}

	pg.Pool = pool

	return pg, nil
}
