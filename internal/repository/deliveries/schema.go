package deliveries

const tableName = "deliveries"

const (
	columnOrderUID = "order_uid"
	columnName     = "name"
	columnPhone    = "phone"
	columnZIP      = "zip"
	columnCity     = "city"
	columnAddress  = "address"
	columnRegion   = "region"
	columnEmail    = "email"
)

var tableColumns = []string{
	columnOrderUID,
	columnName,
	columnPhone,
	columnZIP,
	columnCity,
	columnAddress,
	columnRegion,
	columnEmail,
}
