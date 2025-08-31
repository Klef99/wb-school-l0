package models

import (
	"errors"
	"fmt"
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

func (p *Payment) Validate() error {
	var errs []error

	if p.Transaction == "" {
		errs = append(errs, ValidationError{columnPaymentTransaction, p.Transaction, "empty Transaction"})
	}

	if p.Currency == "" {
		errs = append(errs, ValidationError{columnPaymentCurrency, p.Currency, "empty Currency"})
	}

	if p.Provider == "" {
		errs = append(errs, ValidationError{columnPaymentProvider, p.Provider, "empty Provider"})
	}

	if p.Amount == 0 {
		errs = append(errs, ValidationError{columnPaymentAmount, fmt.Sprint(p.Amount), "zero Amount"})
	}

	if p.PaymentDT.IsZero() {
		errs = append(errs, ValidationError{columnPaymentPaymentDT, p.PaymentDT.String(), "zero PaymentDT"})
	}

	if p.Bank == "" {
		errs = append(errs, ValidationError{columnPaymentBank, p.Bank, "empty Bank"})
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failure: %w", errors.Join(errs...))
	}

	return nil
}

func NewPayment(
	transaction, requestID, currency, provider string,
	amount uint, paymentDT time.Time, bank string,
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

	if err := p.Validate(); err != nil {
		return Payment{}, err
	}

	return p, nil
}

func NewPaymentFromDB(
	transaction, requestID, currency, provider string,
	amount uint, paymentDT time.Time, bank string,
	deliveryCost, goodsTotal, customFee uint,
) Payment {
	return Payment{
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
}
