package models

// Delivery model
const (
	columnDeliveryOrderUID = "orderUID"
	columnDeliveryName     = "name"
	columnDeliveryPhone    = "phone"
	columnDeliveryZIP      = "zip"
	columnDeliveryCity     = "city"
	columnDeliveryAddress  = "address"
	columnDeliveryRegion   = "region"
	columnDeliveryEmail    = "email"
)

// PaymentModel

const (
	columnPaymentTransaction  = "Transaction"
	columnPaymentRequestID    = "RequestID"
	columnPaymentCurrency     = "Currency"
	columnPaymentProvider     = "Provider"
	columnPaymentAmount       = "Amount"
	columnPaymentPaymentDT    = "PaymentDT"
	columnPaymentBank         = "Bank"
	columnPaymentDeliveryCost = "DeliveryCost"
	columnPaymentGoodsTotal   = "GoodsTotal"
	columnPaymentCustomFee    = "CustomFee"
)

const (
	columnItemChrtID      = "ChrtID"
	columnItemTrackNumber = "TrackNumber"
	columnItemPrice       = "Price"
	columnItemRid         = "RID"
	columnItemName        = "Name"
	columnItemSale        = "Sale"
	columnItemSize        = "Size"
	columnItemTotalPrice  = "TotalPrice"
	columnItemNMID        = "NMID"
	columnItemBrand       = "Brand"
	columnItemStatus      = "Status"
)

const (
	columnOrderUID          = "OrderUID"
	columnOrderTrackNumber  = "TrackNumber"
	columnOrderEntry        = "Entry"
	columnOrderDelivery     = "Delivery"
	columnOrderPayment      = "Payment"
	columnOrderItems        = "Items"
	columnOrderLocale       = "Locale"
	columnOrderInternalSign = "InternalSignature"
	columnOrderCustomerID   = "CustomerID"
	columnOrderDeliveryServ = "DeliveryService"
	columnOrderShardKey     = "ShardKey"
	columnOrderSMID         = "SMID"
	columnOrderDateCreated  = "DateCreated"
	columnOrderOofShard     = "OofShard"
)
