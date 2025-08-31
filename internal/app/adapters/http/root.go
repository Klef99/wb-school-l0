package http

import (
	v1 "github.com/klef99/wb-school-l0/internal/app/adapters/http/v1"
)

type RootHandlerV1 struct {
	Health          *v1.HealthHandler
	GetOrderHandler *v1.GetOrderHandler
}

func NewRootHandlerV1(healthHandler *v1.HealthHandler, getOrderHandler *v1.GetOrderHandler) *RootHandlerV1 {
	return &RootHandlerV1{
		Health:          healthHandler,
		GetOrderHandler: getOrderHandler,
	}
}
