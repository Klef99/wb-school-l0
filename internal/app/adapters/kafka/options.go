package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type Option func(c *kafka.ReaderConfig)

func SessionTimeout(timeout time.Duration) Option {
	return func(c *kafka.ReaderConfig) {
		c.SessionTimeout = timeout
	}
}

func MaxAttempts(attempts int) Option {
	return func(c *kafka.ReaderConfig) {
		c.MaxAttempts = attempts
	}
}
