package items

import (
	"github.com/klef99/wb-school-l0/internal/models"
)

type DBO struct {
	ChrtID      uint   `db:"chrt_id"`
	TrackNumber string `db:"track_number"`
	Price       uint   `db:"price"`
	RID         string `db:"rid"`
	Name        string `db:"name"`
	Sale        uint   `db:"sale"`
	Size        string `db:"size"`
	TotalPrice  uint   `db:"total_price"`
	NMID        uint   `db:"nm_id"`
	Brand       string `db:"brand"`
	Status      uint   `db:"status"`
}

func (i *DBO) Values() []any {
	return []any{
		i.ChrtID, i.TrackNumber, i.Price, i.RID, i.Name, i.Sale, i.Size, i.TotalPrice, i.NMID, i.Brand, i.Status,
	}
}

func NewDBO(i models.Item) DBO {
	return DBO{
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

func (i *DBO) ToMap() map[string]interface{} {
	return map[string]interface{}{
		columnChrtID:      i.ChrtID,
		columnTrackNumber: i.TrackNumber,
		columnPrice:       i.Price,
		columnRid:         i.RID,
		columnName:        i.Name,
		columnSale:        i.Sale,
		columnSize:        i.Size,
		columnTotalPrice:  i.TotalPrice,
		columnNMID:        i.NMID,
		columnBrand:       i.Brand,
		columnStatus:      i.Status,
	}
}

func NewDBOs(items []models.Item) []DBO {
	res := make([]DBO, len(items))

	for i, item := range items {
		res[i] = NewDBO(item)
	}

	return res
}
