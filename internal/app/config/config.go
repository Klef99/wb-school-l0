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
}

type HTTPConfig struct {
	Addr    string        `env:"ADDR" envDefault:":8080"`
	Timeout time.Duration `env:"TIMEOUT" envDefault:"10s"`
}

type KafkaConfig struct {
	Addrs   string        `env:"ADDR" envDefault:"localhost:9092"`
	Timeout time.Duration `env:"TIMEOUT" envDefault:"10s"`
	GroupID string        `env:"GROUP_ID" envDefault:"orderServices"`
	Topic   string        `env:"TOPIC" envDefault:"orders"`
}

type PostgresConfig struct {
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	Database string `env:"DB" envDefault:"postgres"`
}

func (c PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}
