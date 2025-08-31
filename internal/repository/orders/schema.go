package orders

const tableName = "orders"

const (
	columnUid               = "uid"
	columnTrackNumber       = "track_number"
	columnEntry             = "entry"
	columnLocale            = "locale"
	columnInternalSignature = "internal_signature"
	columnCustomerID        = "customer_id"
	columnDeliveryService   = "delivery_service"
	columnShardKey          = "shardkey"
	columnSmID              = "sm_id"
	columnDateCreated       = "date_created"
	columnOofShard          = "oof_shard"
)

var tableColumns = []string{
	columnUid,
	columnTrackNumber,
	columnEntry,
	columnLocale,
	columnInternalSignature,
	columnCustomerID,
	columnDeliveryService,
	columnShardKey,
	columnSmID,
	columnDateCreated,
	columnOofShard,
}
