package http

import (
	v1 "github.com/klef99/wb-school-l0/internal/app/adapters/http/v1"
)

type RootHandler struct {
	Health *v1.HealthHandler
}

func NewRootHandler(healthHandler *v1.HealthHandler) *RootHandler {
	return &RootHandler{
		Health: healthHandler,
	}
}
