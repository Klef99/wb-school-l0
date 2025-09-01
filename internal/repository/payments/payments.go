package payments

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

func (r *Repository) Store(ctx context.Context, query postgres.Query, payment models.Payment) error {
	dbo := NewDBO(payment)

	bld := r.builder.Insert(tableName).SetMap(dbo.ToMap())

	res, err := query.ExecWithResult(ctx, "payments.Store", bld)
	if err != nil {
		return fmt.Errorf("error inserting payment %v: %w", payment, err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected != 1 {
		return fmt.Errorf("error inserting payments: expected 1 rows affected, got %d", rowsAffected)
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, query postgres.Query, transaction string) (models.Payment, error) {
	var dbo DBO

	bld := r.builder.Select(tableColumns...).From(tableName).Where(sq.Eq{columnTransaction: transaction})

	err := query.QueryOne(ctx, &dbo, "payments.Get", bld)
	if err != nil {
		return models.Payment{}, fmt.Errorf("error getting payment %v: %w", transaction, err)
	}

	return models.NewPaymentFromDB(
		dbo.Transaction, dbo.RequestID, dbo.Currency, dbo.Provider, dbo.Amount, dbo.PaymentDT, dbo.Bank,
		dbo.DeliveryCost, dbo.GoodsTotal, dbo.CustomFee,
	), nil
}

func (r *Repository) GetAll(ctx context.Context, query postgres.Query) ([]models.Payment, error) {
	var dbos []DBO

	bld := r.builder.Select(tableColumns...).From(tableName)

	err := query.QueryAll(ctx, &dbos, "payments.GetAll", bld)
	if err != nil {
		return nil, fmt.Errorf("error getting all payments: %w", err)
	}

	payments := make([]models.Payment, len(dbos))
	for i, dbo := range dbos {
		payments[i] = models.NewPaymentFromDB(
			dbo.Transaction, dbo.RequestID, dbo.Currency, dbo.Provider, dbo.Amount, dbo.PaymentDT, dbo.Bank,
			dbo.DeliveryCost, dbo.GoodsTotal, dbo.CustomFee,
		)
	}

	return payments, nil
}
