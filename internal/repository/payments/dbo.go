package payments

import (
	"time"

	"github.com/klef99/wb-school-l0/internal/models"
)

type DBO struct {
	Transaction  string    `db:"transaction"`
	RequestID    string    `db:"request_id"`
	Currency     string    `db:"currency"`
	Provider     string    `db:"provider"`
	Amount       uint      `db:"amount"`
	PaymentDT    time.Time `db:"payment_dt"`
	Bank         string    `db:"bank"`
	DeliveryCost uint      `db:"delivery_cost"`
	GoodsTotal   uint      `db:"goods_total"`
	CustomFee    uint      `db:"custom_fee"`
}

func (p *DBO) Values() []any {
	return []any{
		p.Transaction, p.RequestID, p.Currency, p.Provider, p.Amount, p.PaymentDT, p.Bank, p.DeliveryCost, p.GoodsTotal,
		p.CustomFee,
	}
}

func (p *DBO) ToMap() map[string]interface{} {
	return map[string]interface{}{
		columnTransaction:  p.Transaction,
		columnRequestID:    p.RequestID,
		columnCurrency:     p.Currency,
		columnProvider:     p.Provider,
		columnAmount:       p.Amount,
		columnPaymentDT:    p.PaymentDT,
		columnBank:         p.Bank,
		columnDeliveryCost: p.DeliveryCost,
		columnGoodsTotal:   p.GoodsTotal,
		columnCustomFee:    p.CustomFee,
	}
}

func NewDBO(p models.Payment) DBO {
	return DBO{
		Transaction:  p.Transaction,
		RequestID:    p.RequestID,
		Currency:     p.Currency,
		Provider:     p.Provider,
		Amount:       p.Amount,
		PaymentDT:    p.PaymentDT,
		Bank:         p.Bank,
		DeliveryCost: p.DeliveryCost,
		GoodsTotal:   p.GoodsTotal,
		CustomFee:    p.CustomFee,
	}
}
