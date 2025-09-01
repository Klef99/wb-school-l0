package service

import (
	"context"
	"log/slog"

	"github.com/klef99/wb-school-l0/pkg/postgres"
)

type HealthService struct {
	logger *slog.Logger
	pg     postgres.StorageManager
}

func NewHealthService(logger *slog.Logger, storage postgres.StorageManager) *HealthService {
	return &HealthService{logger: logger, pg: storage}
}

func (s *HealthService) Health(ctx context.Context) error {
	return s.pg.GetStorage().Health(ctx)
}
