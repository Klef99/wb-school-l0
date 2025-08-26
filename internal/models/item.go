package models

type Item struct {
	ChrtID      uint
	TrackNumber string
	Price       uint
	RID         string
	Name        string
	Sale        uint
	Size        uint
	TotalPrice  uint
	NmID        uint
	Brand       string
	Status      int
}

func NewItem(
	chrtID uint, trackNumber string, price uint, RID, name string,
	sale, size, totalPrice, nmID uint, brand string, status int,
) (Item, error) {
	i := Item{
		ChrtID:      chrtID,
		TrackNumber: trackNumber,
		Price:       price,
		RID:         RID,
		Name:        name,
		Sale:        sale,
		Size:        size,
		TotalPrice:  totalPrice,
		NmID:        nmID,
		Brand:       brand,
		Status:      status,
	}

	return i, nil
}
