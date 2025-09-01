package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/klef99/wb-school-l0/internal/dto"
	"github.com/klef99/wb-school-l0/internal/lib/logger/sl"
	"github.com/klef99/wb-school-l0/internal/service"
)

type OrderCreator interface {
	Store(ctx context.Context, order dto.Order) error
}

type Kafka struct {
	reader  *kafka.Reader
	writer  *kafka.Writer
	service OrderCreator
	logger  *slog.Logger
}

func NewKafka(
	service OrderCreator, logger *slog.Logger, address, consumerGroup, topic, topicDLQ string, opts ...Option,
) (*Kafka, error) {
	cfg := kafka.ReaderConfig{
		Brokers:        strings.Split(address, ","),
		GroupID:        consumerGroup,
		Topic:          topic,
		CommitInterval: 0,
	}

	writer := &kafka.Writer{
		Addr:                   kafka.TCP(strings.Split(address, ",")...),
		Topic:                  topicDLQ,
		AllowAutoTopicCreation: true,
	}

	// Setup options
	for _, opt := range opts {
		opt(&cfg)
	}

	reader := kafka.NewReader(cfg)

	return &Kafka{
		reader:  reader,
		service: service,
		logger:  logger,
		writer:  writer,
	}, nil
}

func (c *Kafka) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			c.logger.Info("consumer stopped")
			return
		default:
			msg, err := c.reader.FetchMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					c.logger.Info("consumer fetch canceled")
					return
				}

				c.logger.Error("failed to fetch message from kafka", sl.Err(err))

				continue
			}

			var order dto.Order

			err = json.Unmarshal(msg.Value, &order)
			if err != nil {
				c.logger.Error(
					"failed to unmarshal message",
					sl.Err(err),
					slog.String("raw", string(msg.Value)),
				)

				err := c.writer.WriteMessages(ctx, kafka.Message{Value: msg.Value})
				if err != nil {
					c.logger.Error("failed to write message to kafka", sl.Err(err))
				}

				continue
			}

			const maxRetries = 3

			var storeErr error
			for i := 1; i <= maxRetries; i++ {
				storeErr = c.service.Store(ctx, order)
				if storeErr == nil {
					break
				}

				if errors.Is(storeErr, service.ErrValidationFailed) {
					break
				}

				c.logger.Warn(
					"failed to store order, retrying...",
					slog.String("uid", order.OrderUID),
					slog.Int("attempt", i),
					sl.Err(storeErr),
				)
				time.Sleep(time.Second * time.Duration(i))
			}

			if storeErr != nil {
				c.logger.Error(
					"failed to store order after retries",
					slog.String("uid", order.OrderUID),
					sl.Err(storeErr),
				)

				err := c.writer.WriteMessages(ctx, kafka.Message{Value: msg.Value})
				if err != nil {
					c.logger.Error("failed to write message to kafka", sl.Err(err))
				}

				continue
			}

			c.logger.Info(
				"order stored successfully",
				slog.String("uid", order.OrderUID),
			)

			if err := c.reader.CommitMessages(ctx, msg); err != nil {
				c.logger.Error("failed to commit message in kafka", sl.Err(err))
			}
		}
	}
}

func (c *Kafka) Close() error {
	return errors.Join(c.reader.Close(), c.writer.Close())
}
