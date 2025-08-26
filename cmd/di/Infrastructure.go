package di

import (
	"log/slog"
	"os"

	"github.com/google/wire"

	"github.com/klef99/wb-school-l0/internal/app/config"
)

var InfrastructureSet = wire.NewSet(
	config.LoadConfig,
	ProvideLogger,
)

func ProvideLogger(cfg *config.Config) *slog.Logger {
	var logger *slog.Logger

	switch cfg.Environment {
	case config.EnvironmentDevelopment:
		logger = slog.New(
			slog.NewTextHandler(
				os.Stdout, &slog.HandlerOptions{
					AddSource: true,
					Level:     slog.LevelDebug,
				},
			),
		)
	case config.EnvironmentStage:
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdout, &slog.HandlerOptions{
					AddSource: true,
					Level:     slog.LevelDebug,
				},
			),
		)
	case config.EnvironmentProduction:
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdout, &slog.HandlerOptions{
					AddSource: true,
					Level:     slog.LevelInfo,
				},
			),
		)
	}

	return logger
}
