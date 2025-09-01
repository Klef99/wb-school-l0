package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/klef99/wb-school-l0/internal/dto"
	"github.com/klef99/wb-school-l0/internal/service"
)

type GetOrderHandler struct {
	service OrderProvider
}

type OrderProvider interface {
	Get(ctx context.Context, uid string) (dto.Order, error)
}

func NewGetOrderHandler(service OrderProvider) *GetOrderHandler {
	return &GetOrderHandler{service: service}
}

func (h *GetOrderHandler) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	order, err := h.service.Get(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrOrderNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "Order not found")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
