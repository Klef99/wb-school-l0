package service

import (
	"context"
)

type HealthStorage interface {
	Health(ctx context.Context) error
}

type HealthService struct {
	healthStorage HealthStorage
}

func NewHealthService(storage HealthStorage) *HealthService {
	return &HealthService{healthStorage: storage}
}

func (s *HealthService) Health(ctx context.Context) error {
	return s.healthStorage.Health(ctx)
}
