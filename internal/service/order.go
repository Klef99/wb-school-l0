package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/klef99/wb-school-l0/internal/dto"
	"github.com/klef99/wb-school-l0/internal/lib/logger/sl"
	"github.com/klef99/wb-school-l0/internal/models"
	"github.com/klef99/wb-school-l0/pkg/postgres"
	"github.com/klef99/wb-school-l0/pkg/redis"
)

type PaymentsRepository interface {
	Store(ctx context.Context, query postgres.Query, payment models.Payment) error
	Get(ctx context.Context, query postgres.Query, transaction string) (models.Payment, error)
}

type DeliveriesRepository interface {
	Store(ctx context.Context, query postgres.Query, delivery models.Delivery) error
	Get(ctx context.Context, query postgres.Query, uid string) (models.Delivery, error)
}

type ItemsRepository interface {
	Store(ctx context.Context, query postgres.Query, items []models.Item) error
	Get(ctx context.Context, query postgres.Query, trackNumber string) ([]models.Item, error)
}

type OrdersRepository interface {
	StoreBare(ctx context.Context, query postgres.Query, bareOrder models.Order) error
	GetBare(ctx context.Context, query postgres.Query, uid string) (models.Order, error)
	Exist(ctx context.Context, query postgres.Query, uid string) (bool, error)
}

type OrderService struct {
	logger         *slog.Logger
	pg             postgres.StorageManager
	cache          redis.CacheManager
	paymentsRepo   PaymentsRepository
	deliveriesRepo DeliveriesRepository
	itemsRepo      ItemsRepository
	ordersRepo     OrdersRepository
}

func NewOrderService(
	logger *slog.Logger, pg postgres.StorageManager, cache redis.CacheManager, paymentsRepo PaymentsRepository,
	deliveriesRepo DeliveriesRepository, itemsRepo ItemsRepository, ordersRepo OrdersRepository,
) *OrderService {
	return &OrderService{
		logger:         logger,
		pg:             pg,
		cache:          cache,
		paymentsRepo:   paymentsRepo,
		deliveriesRepo: deliveriesRepo,
		itemsRepo:      itemsRepo,
		ordersRepo:     ordersRepo,
	}
}

func (s *OrderService) Store(ctx context.Context, order dto.Order) error {
	orderModel, err := toOrderModel(order)
	if err != nil {
		var perr models.ValidationError

		if errors.As(err, &perr) {
			return errors.Join(ErrValidationFailed, err)
		}

		return fmt.Errorf("failed convert dto to order model: %w", err)
	}

	tx, errTx := s.pg.GetStorage().BeginTx(ctx, "tx.order.Store")
	if errTx != nil {
		s.logger.Error("failed to begin tx", sl.Err(errTx))

		return fmt.Errorf("failed start tx: %w", err)
	}

	defer func() {
		if err != nil {
			if errTx = tx.Rollback(ctx); errTx != nil && !errors.Is(err, postgres.ErrTxClosed) {
				err = errors.Join(errTx, err)
				s.logger.Error("failed to roll back transaction", sl.Err(err))
			}
		}
	}()

	ok, err := s.ordersRepo.Exist(ctx, tx, order.OrderUID)
	if err != nil {
		return fmt.Errorf("failed check orders: %w", err)
	}

	if ok {
		s.logger.Info("order exist in db", slog.String("order_uid", order.OrderUID))
		return nil
	}

	err = s.ordersRepo.StoreBare(ctx, tx, orderModel)
	if err != nil {
		return fmt.Errorf("failed store order: %w", err)
	}

	err = s.paymentsRepo.Store(ctx, tx, orderModel.Payment)
	if err != nil {
		return fmt.Errorf("failed store payment: %w", err)
	}

	err = s.deliveriesRepo.Store(ctx, tx, orderModel.Delivery)
	if err != nil {
		return fmt.Errorf("failed store delivery: %w", err)
	}

	err = s.itemsRepo.Store(ctx, tx, orderModel.Items)
	if err != nil {
		return fmt.Errorf("failed store items: %w", err)
	}

	if err = tx.Commit(ctx); err != nil && !errors.Is(err, postgres.ErrTxClosed) {
		s.logger.Error("failed to commit tx", sl.Err(err))

		return fmt.Errorf("failed commit tx: %w", err)
	}

	err = s.cache.GetCache().Set(ctx, order.OrderUID, order)
	if err != nil {
		s.logger.Warn("failed to set cache", sl.Err(err))
	}

	return nil
}

func (s *OrderService) Get(ctx context.Context, uid string) (dto.Order, error) {
	var err error

	if uid == "" {
		return dto.Order{}, errors.New("invalid uid")
	}

	var order dto.Order

	if err = s.cache.GetCache().Get(ctx, uid, order); err != nil {
		s.logger.Info("failed to get cache", sl.Err(err))
	} else {
		return order, nil
	}

	tx, errTx := s.pg.GetStorage().BeginTx(ctx, "tx.order.Store")
	if errTx != nil {
		s.logger.Error("failed to begin tx", sl.Err(errTx))

		return dto.Order{}, fmt.Errorf("failed start tx: %w", err)
	}

	defer func() {
		if err != nil {
			if errTx = tx.Rollback(ctx); errTx != nil {
				err = errors.Join(errTx, err)
				s.logger.Error("failed to roll back transaction", sl.Err(err))
			}
		} else {
			if errTx = tx.Commit(ctx); errTx != nil {
				err = errTx
				s.logger.Error("failed to commit transaction", sl.Err(err))
			}
		}
	}()

	ok, err := s.ordersRepo.Exist(ctx, tx, uid)
	if err != nil {
		return dto.Order{}, fmt.Errorf("failed check orders: %w", err)
	}

	if !ok {
		return dto.Order{}, ErrOrderNotFound
	}

	bareOrder, err := s.ordersRepo.GetBare(ctx, tx, uid)
	if err != nil {
		return dto.Order{}, fmt.Errorf("failed get order: %w", err)
	}

	payment, err := s.paymentsRepo.Get(ctx, tx, uid)
	if err != nil {
		return dto.Order{}, fmt.Errorf("failed get payment: %w", err)
	}

	delivery, err := s.deliveriesRepo.Get(ctx, tx, uid)
	if err != nil {
		return dto.Order{}, fmt.Errorf("failed get delivery: %w", err)
	}

	items, err := s.itemsRepo.Get(ctx, tx, bareOrder.TrackNumber)
	if err != nil {
		return dto.Order{}, fmt.Errorf("failed get items: %w", err)
	}

	bareOrder.Payment = payment
	bareOrder.Delivery = delivery
	bareOrder.Items = items

	order = toOrderDTO(bareOrder)

	err = s.cache.GetCache().Set(ctx, order.OrderUID, order)
	if err != nil {
		s.logger.Info("failed to set cache", sl.Err(err))
	}

	return order, nil
}
