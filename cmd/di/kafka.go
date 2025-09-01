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
	*kafka.Kafka, func(), error,
) {
	consumer, err := kafka.NewKafka(
		service, logger, cfg.Kafka.Addrs, cfg.Kafka.GroupID, cfg.Kafka.Topic, cfg.Kafka.TopicDLQ,
		kafka.SessionTimeout(cfg.Kafka.Timeout),
		kafka.MaxAttempts(cfg.Kafka.MaxAttempts),
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

func ProvideKafkaAdapter(logger *slog.Logger, kafka *kafka.Kafka) (
	KafkaAdapter, func(), error,
) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		kafka.Start(ctx)
	}()

	cleanup := func() {
		logger.Info("shutting down kafka adapter")

		cancel()

		if err := kafka.Close(); err != nil {
			logger.Error("error shutting down kafka adapter", sl.Err(err))
		}

		logger.Info("kafka adapter shut down")
	}

	return KafkaAdapter{}, cleanup, nil
}
