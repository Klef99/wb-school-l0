package orders

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/klef99/wb-school-l0/internal/models"
	"github.com/klef99/wb-school-l0/pkg/postgres"
)

type Repository struct {
	builder sq.StatementBuilderType
}

func NewRepository(builder sq.StatementBuilderType) *Repository {
	return &Repository{
		builder: builder,
	}
}

func (r *Repository) StoreBare(ctx context.Context, query postgres.Query, bareOrder models.Order) error {
	dbo := NewDBO(bareOrder)

	bld := r.builder.Insert(tableName).SetMap(dbo.ToMap())

	res, err := query.ExecWithResult(ctx, "orders.Store", bld)
	if err != nil {
		return fmt.Errorf("error inserting prder %v: %w", bareOrder.OrderUID, err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected != 1 {
		return fmt.Errorf("error inserting order: expected 1 rows affected, got %d", rowsAffected)
	}

	return nil
}

func (r *Repository) GetBare(ctx context.Context, query postgres.Query, uid string) (models.Order, error) {
	var dbo DBO

	bld := r.builder.Select(tableColumns...).From(tableName).Where(sq.Eq{columnUid: uid})

	err := query.QueryOne(ctx, &dbo, "orders.Get", bld)
	if err != nil {
		return models.Order{}, fmt.Errorf("error getting order %v: %w", uid, err)
	}

	return models.NewOrderFromDBO(
		dbo.UID, dbo.TrackNumber, dbo.Entry, models.Delivery{}, models.Payment{}, []models.Item{}, dbo.Locale,
		dbo.InternalSignature, dbo.CustomerID, dbo.DeliveryService,
		dbo.ShardKey, dbo.SmID, dbo.DateCreated, dbo.OofShard,
	), nil
}

func (r *Repository) Exist(ctx context.Context, query postgres.Query, uid string) (bool, error) {
	var orderUID string

	bld := r.builder.Select(columnUid).From(tableName).Where(sq.Eq{columnUid: uid})

	err := query.QueryOne(ctx, &orderUID, "orders.Exist", bld)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("error getting order %v: %w", uid, err)
	}

	return orderUID == uid, err
}

func (r *Repository) GetBareAll(ctx context.Context, query postgres.Query) ([]models.Order, error) {
	var dbos []DBO

	bld := r.builder.Select(tableColumns...).From(tableName)

	err := query.QueryAll(ctx, &dbos, "orders.GetAll", bld)
	if err != nil {
		return nil, fmt.Errorf("error getting all bare orders: %w", err)
	}

	bareOrders := make([]models.Order, len(dbos))
	for i, dbo := range dbos {
		bareOrders[i] = models.NewOrderFromDBO(
			dbo.UID, dbo.TrackNumber, dbo.Entry, models.Delivery{}, models.Payment{}, []models.Item{}, dbo.Locale,
			dbo.InternalSignature, dbo.CustomerID, dbo.DeliveryService,
			dbo.ShardKey, dbo.SmID, dbo.DateCreated, dbo.OofShard,
		)
	}

	return bareOrders, nil
}
