package models

import (
	"time"
)

type Payment struct {
	Transaction  string
	RequestID    string
	Currency     string
	Provider     string
	Amount       uint
	PaymentDT    time.Time
	Bank         string
	DeliveryCost uint
	GoodsTotal   uint
	CustomFee    uint
}

func NewPayment(
	transaction, requestID, currency, provider string, amount uint, paymentDT time.Time, bank string,
	deliveryCost, goodsTotal, customFee uint,
) (Payment, error) {
	p := Payment{
		Transaction:  transaction,
		RequestID:    requestID,
		Currency:     currency,
		Provider:     provider,
		Amount:       amount,
		PaymentDT:    paymentDT,
		Bank:         bank,
		DeliveryCost: deliveryCost,
		GoodsTotal:   goodsTotal,
		CustomFee:    customFee,
	}
	return p, nil
}
