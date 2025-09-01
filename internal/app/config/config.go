package config

import (
	"fmt"
	"time"
)

const (
	EnvironmentDevelopment = "development"
	EnvironmentStage       = "stage"
	EnvironmentProduction  = "production"
)

type Config struct {
	Environment string         `env:"ENVIRONMENT" envDefault:"development"`
	HTTP        HTTPConfig     `envPrefix:"HTTP_"`
	Kafka       KafkaConfig    `envPrefix:"KAFKA_"`
	Postgres    PostgresConfig `envPrefix:"POSTGRES_"`
	Redis       RedisConfig    `envPrefix:"REDIS_"`
}

type HTTPConfig struct {
	Addr    string        `env:"ADDR" envDefault:":8080"`
	Timeout time.Duration `env:"TIMEOUT" envDefault:"10s"`
}

type KafkaConfig struct {
	Addrs       string        `env:"ADDR" envDefault:"localhost:9092"`
	Timeout     time.Duration `env:"TIMEOUT" envDefault:"10s"`
	MaxAttempts int           `env:"MAX_ATTEMPTS" envDefault:"10"`
	GroupID     string        `env:"GROUP_ID" envDefault:"orderServices"`
	Topic       string        `env:"TOPIC" envDefault:"orders"`
	TopicDLQ    string        `env:"TOPIC_DLQ" envDefault:"ordersDLQ"`
}

type PostgresConfig struct {
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	Database string `env:"DB" envDefault:"postgres"`
}

func (c *PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}

type RedisConfig struct {
	Host        string        `env:"HOST" envDefault:"localhost"`
	Port        string        `env:"PORT" envDefault:"6379"`
	User        string        `env:"USER" envDefault:""`
	Password    string        `env:"PASSWORD" envDefault:"redis"`
	DB          string        `env:"DB" envDefault:"0"`
	TTL         time.Duration `env:"TTL" envDefault:"5m"`
	MaxRetries  int           `env:"MAX_RETRIES" envDefault:"0"`
	DialTimeout time.Duration `env:"DIAL_TIMEOUT" envDefault:"5ms"`
	WarmCache   bool          `env:"WARM_CACHE" envDefault:"false"`
}

func (c *RedisConfig) DSN() string {
	return fmt.Sprintf("redis://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.DB)
}
