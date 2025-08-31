package orders

import (
	"time"

	"github.com/klef99/wb-school-l0/internal/models"
)

type DBO struct {
	UID               string    `db:"uid"`
	TrackNumber       string    `db:"track_number"`
	Entry             string    `db:"entry"`
	Locale            string    `db:"locale"`
	InternalSignature string    `db:"internal_signature"`
	CustomerID        string    `db:"customer_id"`
	DeliveryService   string    `db:"delivery_service"`
	ShardKey          int       `db:"shardkey"`
	SmID              int       `db:"sm_id"`
	DateCreated       time.Time `db:"date_created"`
	OofShard          string    `db:"oof_shard"`
}

func (o *DBO) ToMap() map[string]interface{} {
	return map[string]interface{}{
		columnUid:               o.UID,
		columnTrackNumber:       o.TrackNumber,
		columnEntry:             o.Entry,
		columnLocale:            o.Locale,
		columnInternalSignature: o.InternalSignature,
		columnCustomerID:        o.CustomerID,
		columnDeliveryService:   o.DeliveryService,
		columnShardKey:          o.ShardKey,
		columnSmID:              o.SmID,
		columnDateCreated:       o.DateCreated,
		columnOofShard:          o.OofShard,
	}
}

func NewDBO(o models.Order) DBO {
	return DBO{
		UID:               o.OrderUID,
		TrackNumber:       o.TrackNumber,
		Entry:             o.Entry,
		Locale:            o.Locale,
		InternalSignature: o.InternalSignature,
		CustomerID:        o.CustomerID,
		DeliveryService:   o.DeliveryService,
		ShardKey:          o.ShardKey,
		SmID:              o.SMID,
		DateCreated:       o.DateCreated,
		OofShard:          o.OofShard,
	}
}
