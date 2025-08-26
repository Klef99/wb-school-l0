package storage

import (
	"context"

	"github.com/klef99/wb-school-l0/pkg/postgres"
)

type HealthStorage struct {
	conn *postgres.Postgres
}

func NewHealthStorage(conn *postgres.Postgres) *HealthStorage {
	return &HealthStorage{conn: conn}
}

func (s *HealthStorage) Health(ctx context.Context) error {
	return s.conn.Health(ctx)
}
