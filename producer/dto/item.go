package dto

import (
	"github.com/brianvoe/gofakeit/v7"
)

type Item struct {
	ChrtID      uint   `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       uint   `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        uint   `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  uint   `json:"total_price"`
	NMID        uint   `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      uint   `json:"status"`
}

func RandomItems(track string) []Item {
	cnt := gofakeit.Number(1, 15)
	res := make([]Item, 0, cnt)
	for i := 0; i < cnt; i++ {
		res = append(
			res, Item{
				ChrtID:      uint(gofakeit.Uint16()),
				TrackNumber: track,
				Price:       uint(gofakeit.Uint16()),
				RID:         gofakeit.UUID(),
				Name:        gofakeit.BeerName(),
				Sale:        0,
				Size:        "500ml",
				TotalPrice:  uint(gofakeit.Uint16()),
				NMID:        uint(gofakeit.Uint16()),
				Brand:       gofakeit.BeerStyle(),
				Status:      200,
			},
		)
	}

	return res
}
