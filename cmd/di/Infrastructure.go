package di

import (
	"github.com/google/wire"
	"github.com/klef99/wb-school-l0/internal/app/config"
	"log"
	"log/slog"
)

var InfrastructureSet = wire.NewSet(
	config.LoadConfig,
	ProvideLogger,
)

func ProvideLogger(cfg *config.Config) (*slog.Logger, error) {}
