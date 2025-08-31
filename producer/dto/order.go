package dto

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type Order struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomesID        string    `json:"customes_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          int       `json:"shard_key"`
	SMID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

func RandomOrder() Order {
	uid := gofakeit.UUID()
	track := gofakeit.UUID()
	amount := uint(gofakeit.Price(2, 10000))
	deliveryCost := uint(gofakeit.Price(1, float64(amount)-1))
	order := Order{
		OrderUID:    uid,
		TrackNumber: track,
		Entry:       "WBIL",
		Delivery: Delivery{
			Name:    gofakeit.Name(),
			Phone:   "+7" + gofakeit.Phone(),
			Zip:     gofakeit.Zip(),
			City:    gofakeit.City(),
			Address: gofakeit.Street(),
			Region:  gofakeit.Country(),
			Email:   gofakeit.Email(),
		},
		Payment: Payment{
			Transaction:  uid,
			RequestID:    "",
			Currency:     gofakeit.CurrencyShort(),
			Provider:     "wbpay",
			Amount:       amount,
			PaymentDT:    gofakeit.Date().Unix(),
			Bank:         gofakeit.BankName(),
			DeliveryCost: deliveryCost,
			GoodsTotal:   amount - deliveryCost,
			CustomFee:    0,
		},
		Items:             RandomItems(track),
		Locale:            "en",
		InternalSignature: "",
		CustomesID:        "test",
		DeliveryService:   gofakeit.BS(),
		ShardKey:          9,
		SMID:              99,
		DateCreated:       gofakeit.Date(),
		OofShard:          "1",
	}

	return order
}
