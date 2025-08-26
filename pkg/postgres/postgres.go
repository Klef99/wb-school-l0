package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	Conn    *sql.DB
	Builder squirrel.StatementBuilderType
	// Params
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
}

func New(url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	var err error

	for pg.connAttempts > 0 {
		pg.Conn, err = sql.Open("pgx", url)
		if err == nil {
			pg.Conn.SetMaxOpenConns(pg.maxPoolSize)
			pg.Conn.SetMaxIdleConns(pg.maxPoolSize)

			err = pg.Conn.Ping()
			if err == nil {
				break
			}
		}

		log.Printf("postgres is trying to connect, attempts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(pg.Conn)

	return pg, nil
}

func (p *Postgres) Close() error {
	return p.Conn.Close()
}

func (p *Postgres) Health(ctx context.Context) error {
	return p.Conn.PingContext(ctx)
}
