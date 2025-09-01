package di

import (
	"context"
	"log/slog"
)

type Command struct{}

func ProvideCommand(logger *slog.Logger, _ HTTPAdapter, _ KafkaAdapter, warmer *OrderCacheWarmer) (
	*Command, func(), error,
) {
	logger.Info("Application started")

	logger.Debug("Logger level: DEBUG")

	logger.Info("warming cache")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		if err := warmer.WarmUp(ctx); err != nil {
			logger.Error("warmup failed", "err", err)
		}
	}()

	return &Command{}, func() {
		logger.Info("Application stopped")
		cancel()
	}, nil
}
