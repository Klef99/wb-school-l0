package items

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

func (r *Repository) Store(ctx context.Context, query postgres.Query, items []models.Item) error {
	dbos := NewDBOs(items)

	bld := r.builder.Insert(tableName).Columns(tableColumns...)

	for _, i := range dbos {
		bld = bld.Values(i.Values()...)
	}

	res, err := query.ExecWithResult(ctx, "items.Store", bld)
	if err != nil {
		return fmt.Errorf("error inserting items: %w", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected != int64(len(items)) {
		return fmt.Errorf("error inserting items: expected %d rows affected, got %d", len(items), rowsAffected)
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, query postgres.Query, trackNumber string) ([]models.Item, error) {
	var dbo []DBO

	bld := r.builder.Select(tableColumns...).From(tableName).Where(sq.Eq{columnTrackNumber: trackNumber})

	err := query.QueryAll(ctx, &dbo, "items.Get", bld)
	if err != nil {
		return []models.Item{}, fmt.Errorf("error getting items: %w", err)
	}

	itemsModels := make([]models.Item, len(dbo))
	for i, item := range dbo {
		itemsModels[i] = models.NewItemFromDB(
			item.ChrtID, item.TrackNumber, item.Price, item.RID, item.Name, item.Sale, item.Size, item.TotalPrice,
			item.NMID, item.Brand, item.Status,
		)
	}

	return itemsModels, nil
}

func (r *Repository) GetAll(ctx context.Context, query postgres.Query) ([]models.Item, error) {
	var dbos []DBO

	bld := r.builder.Select(tableColumns...).From(tableName)

	err := query.QueryAll(ctx, &dbos, "items.GetAll", bld)
	if err != nil {
		return nil, fmt.Errorf("error getting all deliveries: %w", err)
	}

	itemsModels := make([]models.Item, len(dbos))
	for i, item := range dbos {
		itemsModels[i] = models.NewItemFromDB(
			item.ChrtID, item.TrackNumber, item.Price, item.RID, item.Name, item.Sale, item.Size, item.TotalPrice,
			item.NMID, item.Brand, item.Status,
		)
	}

	return itemsModels, nil
}
