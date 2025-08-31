package di

import (
	"github.com/google/wire"

	httphandler "github.com/klef99/wb-school-l0/internal/app/adapters/http"
	v1 "github.com/klef99/wb-school-l0/internal/app/adapters/http/v1"
)

var HTTPHandlerSet = wire.NewSet(
	v1.NewHealthHandler,
	v1.NewGetOrderHandler,
	httphandler.NewRootHandlerV1,
)
