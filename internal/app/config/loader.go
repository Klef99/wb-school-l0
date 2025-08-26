package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func LoadConfig() (*Config, error) {
	c := new(Config)
	if err := loadFromEnv(c); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return c, nil
}

func loadFromEnv(c *Config) error {
	if err := godotenv.Load(); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("error load env from .env file. %w", err)
		}
	}

	if err := env.ParseWithOptions(c, env.Options{}); err != nil {
		return fmt.Errorf("error parse env. %w", err)
	}

	return nil
}
