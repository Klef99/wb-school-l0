package payments

const tableName = "payments"

const (
	columnTransaction  = "transaction"
	columnRequestID    = "request_id"
	columnCurrency     = "currency"
	columnProvider     = "provider"
	columnAmount       = "amount"
	columnPaymentDT    = "payment_dt"
	columnBank         = "bank"
	columnDeliveryCost = "delivery_cost"
	columnGoodsTotal   = "goods_total"
	columnCustomFee    = "custom_fee"
)

var tableColumns = []string{
	columnTransaction,
	columnRequestID,
	columnCurrency,
	columnProvider,
	columnAmount,
	columnPaymentDT,
	columnBank,
	columnDeliveryCost,
	columnGoodsTotal,
	columnCustomFee,
}
