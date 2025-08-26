package di

import (
	"github.com/google/wire"

	v1 "github.com/klef99/wb-school-l0/internal/app/adapters/http/v1"
	"github.com/klef99/wb-school-l0/internal/service"
)

var ServiceSet = wire.NewSet(
	service.NewHealthService,
	wire.Bind(new(v1.HealthService), new(*service.HealthService)),
)
