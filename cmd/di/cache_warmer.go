package di

import (
	"context"

	"github.com/google/wire"

	"github.com/klef99/wb-school-l0/internal/service"
	"github.com/klef99/wb-school-l0/pkg/redis"
)

var CacheWarmerSet = wire.NewSet(
	NewOrderCacheWarmer,
)

type OrderCacheWarmer struct {
	orderService *service.OrderService
	cache        redis.CacheManager
}

func NewOrderCacheWarmer(orderService *service.OrderService, cache redis.CacheManager) *OrderCacheWarmer {
	return &OrderCacheWarmer{
		orderService: orderService,
		cache:        cache,
	}
}

func (c *OrderCacheWarmer) WarmUp(ctx context.Context) error {
	orders, err := c.orderService.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, order := range orders {
		err := c.cache.GetCache().Set(ctx, order.OrderUID, order)
		if err != nil {
			return err
		}
	}

	return nil
}
