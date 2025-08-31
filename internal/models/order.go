package models

import (
	"errors"
	"fmt"
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
	CustomerID        string
	DeliveryService   string
	ShardKey          int
	SMID              int
	DateCreated       time.Time
	OofShard          string
}

func (o *Order) Validate() error {
	var errs []error

	if o.OrderUID == "" {
		errs = append(errs, ValidationError{columnOrderUID, o.OrderUID, "empty OrderUID"})
	}

	if o.TrackNumber == "" {
		errs = append(errs, ValidationError{columnOrderTrackNumber, o.TrackNumber, "empty TrackNumber"})
	}

	if o.Entry == "" {
		errs = append(errs, ValidationError{columnOrderEntry, o.Entry, "empty Entry"})
	}

	if err := o.Delivery.Validate(); err != nil {
		errs = append(errs, ValidationError{columnOrderDelivery, "", "invalid Delivery: " + err.Error()})
	}

	if err := o.Payment.Validate(); err != nil {
		errs = append(errs, ValidationError{columnOrderPayment, "", "invalid Payment: " + err.Error()})
	}

	if len(o.Items) == 0 {
		errs = append(errs, ValidationError{columnOrderItems, "", "no Items"})
	} else {
		for idx, item := range o.Items {
			if err := item.Validate(); err != nil {
				errs = append(
					errs,
					ValidationError{columnOrderItems, fmt.Sprintf("index %d", idx), "invalid Item: " + err.Error()},
				)
			}
		}
	}

	if o.Locale == "" {
		errs = append(errs, ValidationError{columnOrderLocale, o.Locale, "empty Locale"})
	}

	if o.DeliveryService == "" {
		errs = append(errs, ValidationError{columnOrderDeliveryServ, o.DeliveryService, "empty DeliveryService"})
	}

	if o.SMID == 0 {
		errs = append(errs, ValidationError{columnOrderSMID, fmt.Sprint(o.SMID), "zero SMID"})
	}

	if o.DateCreated.IsZero() {
		errs = append(errs, ValidationError{columnOrderDateCreated, o.DateCreated.String(), "zero DateCreated"})
	}

	if o.OofShard == "" {
		errs = append(errs, ValidationError{columnOrderOofShard, o.OofShard, "empty OofShard"})
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failure: %w", errors.Join(errs...))
	}

	return nil
}

func NewOrder(
	orderUID, trackNumber, entry string,
	delivery Delivery, payment Payment, items []Item, locale, internalSignature, customerID, deliveryService string,
	shardKey int, smid int, dateCreated time.Time, oofShard string,
) (Order, error) {
	o := Order{
		OrderUID:          orderUID,
		TrackNumber:       trackNumber,
		Entry:             entry,
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            locale,
		InternalSignature: internalSignature,
		CustomerID:        customerID,
		DeliveryService:   deliveryService,
		ShardKey:          shardKey,
		DateCreated:       dateCreated,
		SMID:              smid,
		OofShard:          oofShard,
	}

	if err := o.Validate(); err != nil {
		return o, err
	}

	return o, nil
}

func NewOrderFromDBO(
	orderUID, trackNumber, entry string,
	delivery Delivery, payment Payment, items []Item, locale, internalSignature, customerID, deliveryService string,
	shardKey int, smid int, dateCreated time.Time, oofShard string,
) Order {
	o := Order{
		OrderUID:          orderUID,
		TrackNumber:       trackNumber,
		Entry:             entry,
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            locale,
		InternalSignature: internalSignature,
		CustomerID:        customerID,
		DeliveryService:   deliveryService,
		ShardKey:          shardKey,
		DateCreated:       dateCreated,
		SMID:              smid,
		OofShard:          oofShard,
	}

	return o
}
