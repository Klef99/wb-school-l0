package di

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"

	httphandler "github.com/klef99/wb-school-l0/internal/app/adapters/http"
	"github.com/klef99/wb-school-l0/internal/app/config"
)

type HTTPAdapter struct{}

var HHTPAdapterSet = wire.NewSet(
	ProvideEcho,
	ProvideHTTPAdapter,
)

func ProvideEcho(logger *slog.Logger) (*echo.Echo, func()) {
	e := echo.New()
	cleanup := func() {
		err := e.Shutdown(context.Background())
		if err != nil {
			logger.Error(fmt.Errorf("failed to shutdown echo: %w", err).Error())
		}
	}

	return e, cleanup
}

func ProvideHTTPAdapter(
	cfg *config.Config,
	logger *slog.Logger,
	e *echo.Echo,
	rootHandler *httphandler.RootHandlerV1,
) (HTTPAdapter, func(), error) {
	e.Use(middleware.RequestID())
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())
	e.Use(middleware.ContextTimeout(cfg.HTTP.Timeout))
	e.Use(middleware.Gzip())

	v1Group := e.Group("/v1")
	// Set up endpoints
	v1Group.GET("/health", rootHandler.Health.Handle)
	v1Group.GET("/orders/:id", rootHandler.GetOrderHandler.Handle)

	var errCh = make(chan error, 1)

	go func() {
		errCh <- e.Start(cfg.HTTP.Addr)
	}()

	select {
	case err := <-errCh:
		return HTTPAdapter{}, nil, fmt.Errorf("error start http server: %w", err)
	case <-time.After(500 * time.Millisecond):
	}

	logger.Info("HTTP started at ", cfg.HTTP.Addr)

	cleanup := func() {
		logger.Info("HTTP shutting down")

		defer func() {
			logger.Info("HTTP stopped")
		}()

		if err := e.Shutdown(context.Background()); err != nil {
			logger.Error(fmt.Errorf("failed to shutdown echo: %w", err).Error())
		}
	}

	return HTTPAdapter{}, cleanup, nil
}
