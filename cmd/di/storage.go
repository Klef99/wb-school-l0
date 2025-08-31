package di

import (
	"fmt"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/wire"

	"github.com/klef99/wb-school-l0/internal/app/config"
	"github.com/klef99/wb-school-l0/internal/repository/deliveries"
	"github.com/klef99/wb-school-l0/internal/repository/items"
	"github.com/klef99/wb-school-l0/internal/repository/orders"
	"github.com/klef99/wb-school-l0/internal/repository/payments"
	"github.com/klef99/wb-school-l0/internal/service"
	"github.com/klef99/wb-school-l0/pkg/postgres"
)

var StorageSet = wire.NewSet(
	ProvidePostgres,
	ProvideStatementBuilder,
	postgres.NewStorageManager,
	payments.NewRepository,
	deliveries.NewRepository,
	items.NewRepository,
	orders.NewRepository,
	wire.Bind(new(service.PaymentsRepository), new(*payments.Repository)),
	wire.Bind(new(service.DeliveriesRepository), new(*deliveries.Repository)),
	wire.Bind(new(service.ItemsRepository), new(*items.Repository)),
	wire.Bind(new(service.OrdersRepository), new(*orders.Repository)),
)

func ProvidePostgres(cfg *config.Config, logger *slog.Logger) (*postgres.Postgres, func(), error) {
	pg, err := postgres.New(logger, cfg.Postgres.DSN())
	if err != nil {
		return nil, nil, fmt.Errorf("unable connect to postgres: %w", err)
	}

	cleanup := func() {
		logger.Info("Closing database connection")

		err := pg.Close()
		if err != nil {
			logger.Error("Unable to close postgres connection: %v", err)
		}

		logger.Info("Database connection closed")
	}

	return pg, cleanup, nil
}

func ProvideStatementBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
