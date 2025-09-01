package postgres

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type storageTx struct {
	pgx.Tx
}

func (t *storageTx) Exec(ctx context.Context, _ string, bld Sqlizer) error {
	query, params, err := bld.ToSql()
	if err != nil {
		return err
	}

	_, err = t.Tx.Exec(ctx, query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (t *storageTx) ExecWithResult(
	ctx context.Context,
	_ string,
	bld Sqlizer,
) (pgconn.CommandTag, error) {
	query, params, err := bld.ToSql()
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	res, err := t.Tx.Exec(ctx, query, params...)
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return res, nil
}

func (t *storageTx) QueryOne(ctx context.Context, dest interface{}, _ string, bld Sqlizer) error {
	query, params, err := bld.ToSql()
	if err != nil {
		return err
	}

	err = pgxscan.Get(ctx, t, dest, query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (t *storageTx) QueryAll(ctx context.Context, dest interface{}, _ string, bld Sqlizer) error {
	query, params, err := bld.ToSql()
	if err != nil {
		return err
	}

	err = pgxscan.Select(ctx, t, dest, query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (t *storageTx) Health(ctx context.Context) error {
	return t.Health(ctx)
}

func (t *storageTx) Commit(ctx context.Context) error {
	err := t.Tx.Commit(ctx)
	if errors.Is(err, pgx.ErrTxClosed) {
		return ErrTxClosed
	}

	return err
}

func (t *storageTx) Rollback(ctx context.Context) error {
	err := t.Tx.Rollback(ctx)
	if errors.Is(err, pgx.ErrTxClosed) {
		return ErrTxClosed
	}

	return err
}
