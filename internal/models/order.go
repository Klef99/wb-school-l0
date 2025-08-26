package models

import (
	"time"
)

type Order struct {
	OrderUID          string
	TrackNumber       string
	Entry             string
	Delivery          Delivery
	Payment           Payment
	Items             []Item
	Locale            string
	InternalSignature string
	CustomesID        string
	DeliveryService   string
	ShardKey          int
	SMID              int
	DateCreated       time.Time
	OofShard          int
}

func NewOrder(
	orderUID, trackNumer, entry string,
	delivery Delivery, payment Payment, items []Item, locale, internalSignature, customerID, deliveryService string,
	shardKey int, dateCreated time.Time, oofShard int,
) (Order, error) {
	o := Order{
		OrderUID:          orderUID,
		TrackNumber:       trackNumer,
		Entry:             entry,
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            locale,
		InternalSignature: internalSignature,
		CustomesID:        customerID,
		DeliveryService:   deliveryService,
		ShardKey:          shardKey,
		DateCreated:       dateCreated,
		OofShard:          oofShard,
	}

	return o, nil
}
