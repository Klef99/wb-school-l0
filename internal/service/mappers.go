package service

import (
	"fmt"
	"time"

	"github.com/klef99/wb-school-l0/internal/dto"
	"github.com/klef99/wb-school-l0/internal/models"
)

// DTO -> Models

func toPaymentModel(p dto.Payment) (models.Payment, error) {
	payment, err := models.NewPayment(
		p.Transaction, p.RequestID, p.Currency, p.Provider, p.Amount,
		time.Unix(p.PaymentDT, 0), p.Bank, p.DeliveryCost, p.GoodsTotal, p.CustomFee,
	)
	if err != nil {
		return models.Payment{}, fmt.Errorf("failed to create payment model: %w", err)
	}

	return payment, nil
}

func toDeliveryModel(d dto.Delivery, orderUID string) (models.Delivery, error) {
	delivery, err := models.NewDelivery(orderUID, d.Name, d.Phone, d.Zip, d.City, d.Address, d.Region, d.Email)
	if err != nil {
		return models.Delivery{}, fmt.Errorf("failed to create deliveries model: %w", err)
	}

	return delivery, nil
}

func toItemModel(it dto.Item) (models.Item, error) {
	item, err := models.NewItem(
		it.ChrtID, it.TrackNumber, it.Price, it.RID, it.Name, it.Sale,
		it.Size, it.TotalPrice, it.NMID, it.Brand,
		it.Status,
	)
	if err != nil {
		return models.Item{}, fmt.Errorf("failed to create item model: %w", err)
	}

	return item, nil
}

func toItemModels(its []dto.Item) ([]models.Item, error) {
	items := make([]models.Item, len(its))

	for i, it := range its {
		item, err := toItemModel(it)
		if err != nil {
			return nil, err
		}

		items[i] = item
	}

	return items, nil
}

func toOrderModel(o dto.Order) (models.Order, error) {
	payment, err := toPaymentModel(o.Payment)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to create order model: %w", err)
	}

	items, err := toItemModels(o.Items)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to create order models: %w", err)
	}

	delivery, err := toDeliveryModel(o.Delivery, o.OrderUID)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to create order models: %w", err)
	}

	order, err := models.NewOrder(
		o.OrderUID, o.TrackNumber, o.Entry, delivery, payment, items, o.Locale, o.InternalSignature, o.CustomesID,
		o.DeliveryService, o.ShardKey, o.SMID, o.DateCreated, o.OofShard,
	)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to create order models: %w", err)
	}

	return order, nil
}

// Models -> DTO

func toPaymentDTO(p models.Payment) dto.Payment {
	return dto.Payment{
		Transaction:  p.Transaction,
		RequestID:    p.RequestID,
		Currency:     p.Currency,
		Provider:     p.Provider,
		Amount:       p.Amount,
		PaymentDT:    p.PaymentDT.Unix(),
		Bank:         p.Bank,
		DeliveryCost: p.DeliveryCost,
		GoodsTotal:   p.GoodsTotal,
		CustomFee:    p.CustomFee,
	}
}

func toDeliveryDTO(d models.Delivery) dto.Delivery {
	return dto.Delivery{
		Name:    d.Name,
		Phone:   d.Phone,
		Zip:     d.Zip,
		City:    d.City,
		Address: d.Address,
		Region:  d.Region,
		Email:   d.Email,
	}
}

func toItemDTO(i models.Item) dto.Item {
	return dto.Item{
		ChrtID:      i.ChrtID,
		TrackNumber: i.TrackNumber,
		Price:       i.Price,
		RID:         i.RID,
		Name:        i.Name,
		Sale:        i.Sale,
		Size:        i.Size,
		TotalPrice:  i.TotalPrice,
		NMID:        i.NMID,
		Brand:       i.Brand,
		Status:      i.Status,
	}
}

func toItemsDTOs(items []models.Item) []dto.Item {
	itemDTOs := make([]dto.Item, len(items))
	for i, item := range items {
		itemDTOs[i] = toItemDTO(item)
	}

	return itemDTOs
}

func toOrderDTO(o models.Order) dto.Order {
	return dto.Order{
		OrderUID:          o.OrderUID,
		TrackNumber:       o.TrackNumber,
		Entry:             o.Entry,
		Delivery:          toDeliveryDTO(o.Delivery),
		Payment:           toPaymentDTO(o.Payment),
		Items:             toItemsDTOs(o.Items),
		Locale:            o.Locale,
		InternalSignature: o.InternalSignature,
		CustomesID:        o.CustomerID,
		DeliveryService:   o.DeliveryService,
		ShardKey:          o.ShardKey,
		SMID:              o.SMID,
		DateCreated:       o.DateCreated,
		OofShard:          o.OofShard,
	}
}
