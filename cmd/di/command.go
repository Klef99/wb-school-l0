package di

import (
	"log/slog"
)

type Command struct{}

func ProvideCommand(logger *slog.Logger, _ HTTPAdapter, _ KafkaAdapter) (*Command, func(), error) {
	logger.Info("Application started")

	logger.Debug("Logger level: DEBUG")

	return &Command{}, func() {
		logger.Info("Application stopped")
	}, nil
}
