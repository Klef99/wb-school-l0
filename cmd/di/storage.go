package di

import (
	"fmt"
	"log"

	"github.com/google/wire"

	"github.com/klef99/wb-school-l0/internal/app/config"
	"github.com/klef99/wb-school-l0/internal/service"
	"github.com/klef99/wb-school-l0/internal/storage"
	"github.com/klef99/wb-school-l0/pkg/postgres"
)

var StorageSet = wire.NewSet(
	ProvidePostgres,
	storage.NewHealthStorage,
	wire.Bind(new(service.HealthStorage), new(*storage.HealthStorage)),
)

func ProvidePostgres(cfg *config.Config) (*postgres.Postgres, func(), error) {
	pg, err := postgres.New(cfg.Postgres.DSN())
	if err != nil {
		return nil, nil, fmt.Errorf("unable connect to postgres: %w", err)
	}

	cleanup := func() {
		log.Println("Closing database connection")

		err := pg.Close()
		if err != nil {
			log.Printf("Unable to close postgres connection: %v", err)
		}

		log.Println("Database connection closed")
	}

	return pg, cleanup, nil
}
