package di

import (
	"context"
	"log/slog"

	"github.com/google/wire"

	"github.com/klef99/wb-school-l0/internal/app/adapters/kafka"
	"github.com/klef99/wb-school-l0/internal/app/config"
	"github.com/klef99/wb-school-l0/internal/lib/logger/sl"
)

type KafkaAdapter struct{}

var KafkaAdapterSet = wire.NewSet(
	ProvideKafkaConsumer,
	ProvideKafkaAdapter,
)

func ProvideKafkaConsumer(service kafka.OrderCreator, cfg *config.Config, logger *slog.Logger) (
	*kafka.Consumer, func(), error,
) {
	consumer, err := kafka.NewConsumer(
		service, logger, cfg.Kafka.Addrs, cfg.Kafka.GroupID, cfg.Kafka.Topic, kafka.SessionTimeout(cfg.Kafka.Timeout),
	)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		logger.Info("shutting down kafka consumer")

		if err := consumer.Close(); err != nil {
			logger.Error("error shutting down kafka consumer", sl.Err(err))
		}

		logger.Info("kafka consumer shut down")
	}

	return consumer, cleanup, nil
}

func ProvideKafkaAdapter(cfg *config.Config, logger *slog.Logger, consumer *kafka.Consumer) (
	KafkaAdapter, func(), error,
) {
	go func() {
		consumer.Start(context.Background())
	}()

	cleanup := func() {
		logger.Info("shutting down kafka adapter")

		if err := consumer.Close(); err != nil {
			logger.Error("error shutting down kafka adapter", sl.Err(err))
		}

		logger.Info("kafka adapter shut down")
	}

	return KafkaAdapter{}, cleanup, nil
}
