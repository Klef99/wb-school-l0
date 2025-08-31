package deliveries

import (
	"github.com/klef99/wb-school-l0/internal/models"
)

type DBO struct {
	OrderUID string `db:"order_uid"`
	Name     string `db:"name"`
	Phone    string `db:"phone"`
	Zip      string `db:"zip"`
	City     string `db:"city"`
	Address  string `db:"address"`
	Region   string `db:"region"`
	Email    string `db:"email"`
}

func (d *DBO) Values() []any {
	return []any{d.OrderUID, d.Name, d.Phone, d.Zip, d.City, d.Address, d.Region, d.Email}
}

func (d *DBO) ToMap() map[string]interface{} {
	return map[string]interface{}{
		columnOrderUID: d.OrderUID,
		columnName:     d.Name,
		columnPhone:    d.Phone,
		columnZIP:      d.Zip,
		columnCity:     d.City,
		columnAddress:  d.Address,
		columnRegion:   d.Region,
		columnEmail:    d.Email,
	}
}

func NewDBO(d models.Delivery) DBO {
	return DBO{
		OrderUID: d.OrderUID,
		Name:     d.Name,
		Phone:    d.Phone,
		Zip:      d.Zip,
		City:     d.City,
		Address:  d.Address,
		Region:   d.Region,
		Email:    d.Email,
	}
}
