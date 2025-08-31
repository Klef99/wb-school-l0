package kafka

import (
	"context"
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/segmentio/kafka-go"

	"github.com/klef99/wb-school-l0/internal/dto"
	"github.com/klef99/wb-school-l0/internal/lib/logger/sl"
)

type OrderCreator interface {
	Store(ctx context.Context, order dto.Order) error
}

type Consumer struct {
	reader  *kafka.Reader
	service OrderCreator
	logger  *slog.Logger
	stop    bool
}

func NewConsumer(
	service OrderCreator, logger *slog.Logger, address, consumerGroup, topic string, opts ...Option,
) (*Consumer, error) {
	cfg := kafka.ReaderConfig{
		Brokers: strings.Split(address, ","),
		GroupID: consumerGroup,
		Topic:   topic,
	}

	// Setup options
	for _, opt := range opts {
		opt(&cfg)
	}

	reader := kafka.NewReader(cfg)

	return &Consumer{
		reader:  reader,
		service: service,
		logger:  logger,
		stop:    false,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) {
	for {
		if c.stop {
			break
		}

		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			c.logger.Error("failed to fetch message from kafka", sl.Err(err))
			continue
		}

		var order dto.Order

		err = json.Unmarshal(msg.Value, &order)
		if err != nil {
			c.logger.Error("failed to unmarshal message from kafka", sl.Err(err))
			continue
		}

		c.logger.Debug("get order from kafka", order)

		err = c.service.Store(ctx, order)
		if err != nil {
			c.logger.Error("failed to store message in db", sl.Err(err))
		}

		c.logger.Info("order stored successfully", slog.String("uid", order.OrderUID))

		err = c.reader.CommitMessages(ctx, msg)
		if err != nil {
			c.logger.Error("failed to commit message in kafka", sl.Err(err))
		}
	}
}

func (c *Consumer) Close() error {
	c.stop = true
	return c.reader.Close()
}
