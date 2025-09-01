package deliveries

import (
	"context"
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

func (r *Repository) Store(ctx context.Context, query postgres.Query, delivery models.Delivery) error {
	dbo := NewDBO(delivery)

	bld := r.builder.Insert(tableName).SetMap(dbo.ToMap())

	res, err := query.ExecWithResult(ctx, "deliveries.Store", bld)
	if err != nil {
		return fmt.Errorf("error inserting delivery %v: %w", delivery, err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected != 1 {
		return fmt.Errorf("error inserting deliveries: expected 1 rows affected, got %d", rowsAffected)
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, query postgres.Query, uid string) (models.Delivery, error) {
	var dbo DBO

	bld := r.builder.Select(tableColumns...).From(tableName).Where(sq.Eq{columnOrderUID: uid})

	err := query.QueryOne(ctx, &dbo, "deliveries.Get", bld)
	if err != nil {
		return models.Delivery{}, fmt.Errorf("error getting delivery %v: %w", uid, err)
	}

	return models.NewDeliveryFromDB(
		dbo.OrderUID, dbo.Name, dbo.Phone, dbo.Zip, dbo.City,
		dbo.Address, dbo.Region, dbo.Email,
	), nil
}

func (r *Repository) GetAll(ctx context.Context, query postgres.Query) ([]models.Delivery, error) {
	var dbos []DBO

	bld := r.builder.Select(tableColumns...).From(tableName)

	err := query.QueryAll(ctx, &dbos, "deliveries.GetAll", bld)
	if err != nil {
		return nil, fmt.Errorf("error getting all deliveries: %w", err)
	}

	deliveries := make([]models.Delivery, len(dbos))
	for i, dbo := range dbos {
		deliveries[i] = models.NewDeliveryFromDB(
			dbo.OrderUID, dbo.Name, dbo.Phone, dbo.Zip, dbo.City,
			dbo.Address, dbo.Region, dbo.Email,
		)
	}

	return deliveries, nil
}
