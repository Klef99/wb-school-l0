package postgres

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgconn"
)

type Sqlizer interface {
	ToSql() (sql string, args []interface{}, err error)
}

type storage struct {
	DB *Postgres
}

func NewStorage(db *Postgres) *storage {
	return &storage{DB: db}
}

type Query interface {
	Exec(ctx context.Context, queryName string, bld Sqlizer) error
	ExecWithResult(ctx context.Context, queryName string, bld Sqlizer) (pgconn.CommandTag, error)
	QueryOne(ctx context.Context, dest interface{}, queryName string, bld Sqlizer) error
	QueryAll(ctx context.Context, dest interface{}, queryName string, bld Sqlizer) error
}

type Storage interface {
	Query
	Health(ctx context.Context) error
	Close()
	BeginTx(ctx context.Context, txName string) (StorageTx, error)
}

type StorageTx interface {
	Query
	Health(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

func (s *storage) Exec(ctx context.Context, _ string, bld Sqlizer) error {
	query, params, err := bld.ToSql()
	if err != nil {
		return err
	}

	_, err = s.DB.Exec(ctx, query, params...)

	return err
}

func (s *storage) ExecWithResult(
	ctx context.Context,
	_ string,
	bld Sqlizer,
) (pgconn.CommandTag, error) {
	query, params, err := bld.ToSql()
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	result, err := s.DB.Exec(ctx, query, params...)
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return result, nil
}

func (s *storage) QueryOne(ctx context.Context, dest interface{}, _ string, bld Sqlizer) error {
	query, params, err := bld.ToSql()
	if err != nil {
		return err
	}

	err = pgxscan.Get(ctx, s.DB, dest, query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) QueryAll(ctx context.Context, dest interface{}, _ string, bld Sqlizer) error {
	query, params, err := bld.ToSql()
	if err != nil {
		return err
	}
	err = pgxscan.Select(ctx, s.DB, dest, query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) BeginTx(ctx context.Context, _ string) (StorageTx, error) {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("pgxStorage BeginTx: %w", err)
	}

	return &storageTx{tx}, nil
}

func (s *storage) Close() {
	s.DB.Close()
}

func (s *storage) Health(ctx context.Context) error {
	return s.DB.Ping(ctx)
}

type StorageManager interface {
	GetStorage() Storage
	Close()
}

type storageManager struct {
	storage *storage
}

func (p *storageManager) GetStorage() Storage {
	return p.storage
}

func (p *storageManager) Close() {
	p.storage.Close()
}

func NewStorageManager(db *Postgres) StorageManager {
	return &storageManager{
		storage: NewStorage(db),
	}
}
