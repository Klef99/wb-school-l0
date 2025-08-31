package commands

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v3"

	"github.com/klef99/wb-school-l0/internal/app/config"
)

var migrationsUpCommand = &cli.Command{
	Name:  "migrations-up",
	Flags: []cli.Flag{},
	Action: func(ctx context.Context, _ *cli.Command) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}
		db, err := sql.Open("pgx", cfg.Postgres.DSN())
		if err != nil {
			return err
		}

		if err := goose.SetDialect("postgres"); err != nil {
			return err
		}

		if err := goose.Up(db, "./migrations"); err != nil {
			return err
		}

		return nil
	},
}

var migrationsDownCommand = &cli.Command{
	Name:  "migrations-down",
	Flags: []cli.Flag{},
	Action: func(ctx context.Context, _ *cli.Command) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		db, err := sql.Open("pgx", cfg.Postgres.DSN())
		if err != nil {
			return err
		}

		if err := goose.SetDialect("postgres"); err != nil {
			return err
		}

		if err := goose.Down(db, "./migrations"); err != nil {
			return err
		}

		return nil
	},
}
