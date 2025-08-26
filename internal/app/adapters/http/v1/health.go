package v1

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	service HealthService
}

type HealthService interface {
	Health(ctx context.Context) error
}

func NewHealthHandler(service HealthService) *HealthHandler {
	return &HealthHandler{service: service}
}

func (h *HealthHandler) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	err := h.service.Health(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
