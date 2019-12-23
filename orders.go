package shopeego

type GetOrdersListRequest struct {
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days. Must include only create_time or update_time in the request.
	CreateTimeFrom int
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_to field is the ending date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days. Must include only create_time or update_time in the request.
	CreateTimeTo int
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the order update time). The update_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days. Must include only create_time or update_time in the request.
	UpdateTimeFrom int
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the order update time). The update_time_to field is the ending date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days. Must include only create_time or update_time in the request.
	UpdateTimeTo int
	// If many orders are available to retrieve, you may need to call GetOrdersList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call.
	// This integer value is used to specify the maximum number of entries to return in a single "page" of data. Max entries per page is 100.
	PaginationEntriesPerPage int
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Order struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// Enumerated type that defines the current status of the order. Applicable values: See Data Definition- OrderStatus.
	OrderStatus string
	// Timestamp that indicates the last time that there was a change in value of order, such as order status changed from 'Paid' to 'Completed'.
	UpdateTime int
}

type GetOrdersListResponse struct {
	// The set of orders that match the ordersn or filter criteria specified.
	Orders []Order
	// This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	More bool
	// The identifier for an API request for error tracking
	RequestID string
}

type GetOrdersByStatusRequest struct {
	// The order_status filter for retriveing orders. Available value: ALL/UNPAID/READY_TO_SHIP/COMPLETED/IN_CANCEL/CANCELLED/TO_RETURN
	OrderStatus string
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days. Must include only create_time or update_time in the request.
	CreateTimeFrom int
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_to field is the ending date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days. Must include only create_time or update_time in the request.
	CreateTimeTo int
	// If many orders are available to retrieve, you may need to call GetOrdersList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call.
	// This integer value is used to specify the maximum number of entries to return in a single "page" of data. Max entries per page is 100.
	PaginationEntriesPerPage int
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Order struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// Enumerated type that defines the current status of the order. Applicable values: See Data Definition- OrderStatus.
	OrderStatus string
	// Timestamp that indicates the last time that there was a change in value of order, such as order status changed from 'Paid' to 'Completed'.
	UpdateTime int
}

type GetOrdersByStatusResponse struct {
	// The set of orders that match the ordersn or filter criteria specified.
	Orders []Order
	// This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	More bool
	// The identifier for an API request for error tracking
	RequestID string
}

type GetOrderDetailsRequest struct {
	// The set of ordersn. You can specify up to 50 ordersns in this call.
	OrderSNList []string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Address struct {
	// Recipient's name for the address.
	Name string
	// Recipient's phone number input when order was placed.
	Phone string
	// The town of the recipient's address. Whether there is a town will depend on the region and/or country.
	Town string
	// The district of the recipient's address. Whether there is a town will depend on the region and/or country.
	District string
	// The city of the recipient's address. Whether there is a town will depend on the region and/or country.
	City string
	// The state/province of the recipient's address. Whether there is a town will depend on the region and/or country.
	State string
	// The two-digit code representing the country of the Recipient.
	Country string
	// Recipient's postal code.
	Zipcode string
	// The full address of the recipient, including country, state, even street, and etc.
	FullAddress string
}

type Item struct {
	// ID of item
	ItemID int
	// Name of item
	ItemName string
	// A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string
	// ID of the variation that belongs to the same item.
	VariationID int
	// Name of the variation that belongs to the same item.
	// A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	VariationName string
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string
	// The number of identical items purchased at the same time by the same buyer from one listing/item.
	VariationQuantityPurchased int
	// The original price of the item in the listing currency.
	VariationOriginalPrice float64
	// The after-discount price of the item in the listing currency. If there is no discount, this value will be same as that of variation_original_price.
	// In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/variation level. Due to technical restriction, the value will return the price before bundle deal if we don't configure it to 0. Please call GetEscrowDetails if you want to calculate item-level discounted price for bundle deal item.
	VariationDiscountedPrice float64
	// This value indicates whether buyer buy the order item in wholesale price.
	IsWholesale bool
	// The weight of the item
	Weight float64
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool
	// The unique identity of an addon deal.
	AddOnDealID int
	// The type of the promotion,
	PromotionType string
	// The ID of the promotion.
	PromotionID int
}

type Order struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The two-digit code representing the country where the order was made.
	Country string
	// The three-digit code representing the currency unit for which the order was paid.
	Currency string
	// This value indicates whether the order was a COD (cash on delivery) order.
	COD bool
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string
	// Shipping preparation time set by the seller when listing item on Shopee.
	DaysToShip int
	// This object contains detailed breakdown for the recipient address.
	RecipientAddress Address
	// This object contains the detailed breakdown for the result of this API call.
	Items []Item
	// The time when the order status is updated from UNPAID to PAID. This value is NULL when order is not paid yet.
	PayTime int
	// For Indonesia orders only. The name of the dropshipper.
	Dropshipper string
	// Last 4 digits of the credit card
	CreditCardNumber string
	// The name of buyer
	BuyerUsername string
	// The phone number of dropshipper
	DropshipperPhone string
	// The deadline to ship out the parcel.
	ShipByDate int
	// To indicate whether this order is split to fullfil order(forder) level. Call GetForderInfo if it's "true".
	IsSplitUp bool
	// Cancel reason from buyer.
	BuyerCancelReason string
	// Could be one of buyer, seller or system
	CancelBy string
	// The first-mile tracking number.
	FMTN string
	// Use this field to get reason for buyer, seller, and system cancellation.
	CancelReason string
}

type GetOrderDetailsResponse struct {
	// The set of orders that match the ordersn or filter criteria specified.
	Orders []Order
	// Orders that encountered error
	Errors []string
	// The identifier for an API request for error tracking
	RequestID string
}

type GetEscrowDetailsRequest struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Income struct {
	// The three-digit code representing the currency unit used for all transactional amount under
	LocalCurrency string
	// The total amount paid by the buyer for the order. This amount includes the total sale price of items, shipping cost beared by buyer; and offset by Shopee promotions if applicable.
	TotalAmount float64
	// Final value of coins used by seller for the order.
	Coin float64
	// Final value of voucher provided by Shopee for the order.
	Voucher float64
	// Final value of voucher provided by Seller for the order.
	VoucherSeller float64
	// Final sum of each item Shopee discount of a specific order. This amount will rebate to seller.
	SellerRebate float64
	// The final shipping cost of order . For Non-integrated logistics channel is 0.
	ActualShippingCost float64
	// The platform shipping subsidy to the seller
	ShippingFeeRebate float64
	// The commission fee charged by Shopee platform if applicable.
	CommissionFee float64
	// The voucher code or promotion code the buyer used.
	VoucherCode float64
	// The voucher name or promotion name the buyer used.
	VoucherName float64
	// The total amount that the seller is expected to receive for the order and will change before order completed. escrow_amount=total_amount+voucher+credit_card_promotion+seller_rebate+coin-commission_fee-credit_card_transaction_fee-cross_border_tax-service_fee-buyer_shopee_kredit-seller_coin_cash_back+final_shipping_fee-seller_return_refund_amount.
	EscrowAmount float64
	// Amount incurred by Buyer for purchasing items outside of home country. Amount may change after Return Refund.
	CroossBorderTax float64
	// The amount offset via payment promotion.
	CreditCardPromotion float64
	// Include buyer transaction fee and seller transaction fee.
	CreditCardTransactionFee float64
	// Amount charged by Shopee to seller for additional services.
	ServiceFee float64
	// Amount charged by Shopee to Buyer for using ShopeeKredit for the order. Currently only applicable in ID.
	BuyerShopeeKredit float64
	// Value of coins provided by Seller for purchasing with his or her store for the order.
	SellerCoinCashBack float64
	// Final adjusted amount that seller has to bear as part of escrow. This amount could be negative or positive.
	FinalShippingFee float64
	// Amount returned to Seller in the event of partial return.
	SellerReturnRefundAmount float64
	// The amount offset via payment promotion. May include bank payment promotion and Shopee payment promotion.
	CreditCardPromotion float64
}

type BankAccount struct {
	// Name of the seller's receiving bank
	BankName string
	// Account number of the seller's receiving bank
	BankAccountNumber string
	// The two-digit code representing country of the seller's receiving bank account
	BankAccountCountry string
}

type Item struct {
	// ID of item.
	ItemID int
	// ID of the variation that belongs to the same item.
	VariationID int
	// The number of identical items purchased at the same time by the same buyer from one listing/item.
	QuantityPurchased int
	// The price used to participate activity. E.g. itemA original price is $10, promo price is $9, and bundle deal is buy 2 get 20% off equals to $14.4. The original_price value will be $9 in this case.
	OriginalPrice float64
}

type Activity struct {
	// ID of activity.
	ActivityID int
	// Type of activity. Currently only one type: bundle_deal
	ActivityType string
	// The original TOTAL price of ALL items in one activity(e.g. bundle deal. Define by activity_id) in the listing currency.
	OriginalPrice float64
	// The after-discocunt TOTAL price of ALL items in one activity(e.g. bundle deal. Define by activity_id) in the listing currency.
	DiscountedPrice float64
	// This object contains the items in this activity.
	Items []Item
}

type Item struct {
	// ID of item
	ItemID int
	// Name of item
	ItemName string
	// A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string
	// ID of the variation that belongs to the same item.
	VariationID int
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	VariationName string
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string
	// This value indicates the number of identical items purchased at the same time by the same buyer from one listing/item.
	QuantityPurchased int
	// The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.
	OriginalPrice float64
	// The after-discount price of the item in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1. If there is no discount, this value will be the same as that of original_price.
	// In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/variation level. Due to technical restriction, the value will return the price before bundle deal if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the bundle deal discount breakdown on item level.
	DiscountedPrice float64
	// The offset of this item when the buyer consumed Shopee Coins upon checkout.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	DiscountFromCoin float64
	// The offset of this item when the buyer use Shopee voucher.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	DiscountFromVoucher float64
	// The offset of this item when the buyer use seller-specific voucher.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	DiscountFromVoucherSeller float64
	// Platform subsidy to the seller for this item.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	SellerRebate float64
	// This value indicates the actual price the buyer pay.
	// In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/variation level. Due to technical restriction, the value will return the price before bundle deal if we don't configure it to 0. Please use the value under "income_details" and "activity" to calculate the bundle deal discount breakdown on item level.
	DealPrice float64
	// This value indicate the offset via credit card promotion.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	CreditCardPromotion float64
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool
	// The unique identity of an addon deal.
	AddOnDealID int
}

type Order struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The two-digit code representing the country where the order was made.
	Country string
	// This object contains detailed income breakdown for the order.
	IncomeDetails IncomeDetail
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string
	// The three-digit code representing the currency unit of total order amount (escorw_amount) at the point of payment to the seller.
	EscrowCurrency string
	// The exchange rate used by Shopee to convert local_currency to escrow_currency.
	ExchangeRate string
	// The payment channel that the seller selected to receive escrow for the order.
	EscrowChannel string
	// The unique identifier for a payee by the 3rd party payment service provider selected in escrow_channel.
	PayeeID int
	// This object contains detailed breakdown for bank account of the seller if selected escorw_channel is Bank Transfer.
	BankAccount BankAccount
	// This object contains the detailed breakdown for all the items in this order, including regular items(non-activity) and activity items.
	Items []Item
	// This object contains the activity in this order.
	Activity []Activity
}

type GetEscrowDetailsResponse struct {
	// My Income infomation
	Order Order
	// The identifier for an API request for error tracking
	RequestID string
}

type AddOrderNoteRequest struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The note seller made for own reference.
	Note string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type AddOrderNoteResponse struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The success or error message.
	Msg string
	// The identifier for an API request for error tracking
	RequestID string
}

type CancelOrderRequest struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The reason seller want to cancel this order. Applicable values: OUT_OF_STOCK, CUSTOMER_REQUEST, UNDELIVERABLE_AREA, COD_NOT_SUPPORTED.
	CancelReason string
	// ID of item. Required when cancel_reason is OUT_OF_STOCK.
	ItemID int
	// ID of the variation that belongs to the same item.Required when cancel_reason is OUT_OF_STOCK.
	VariationID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type CancelOrderResponse struct {
	// The time when the order is updated.
	ModifiedTime
	// The identifier for an API request for error tracking
	RequestID string
}

type AcceptBuyerCancellationRequest struct {
	// The order to be accepted cancellation request.
	OrderSN string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type AcceptBuyerCancellationResponse struct {
	// The time when the order is updated.
	ModifiedTime int
	// The identifier for an API request for error tracking
	RequestID string
}

type RejectBuyerCancellationRequest struct {
	// The order to be rejected cancellation request.
	OrderSN string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type RejectBuyerCancellationResponse struct {
	// The time when the order is updated.
	ModifiedTime int
	// The identifier for an API request for error tracking
	RequestID string
}

type GetForderInfoRequest struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Log struct {
	// The time when logistics info has been updated.
	Ctime int
	// The order logistics tracking info.
	Description string
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// Shopee's unique identifier for a variation of an item.
	VariationID int
	// The number of identical items/variations purchased at the same time by the same buyer from one listing/item.
	Num int
	// The original price of the item in the listing currency.
	ItemPrice float64
	// The original price of the variation in the listing currency.
	VariationPrice float64
}

type LogisticsInfo struct {
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string
	// Only work for cross-border order. This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	GoodsToDeclare bool
	// Only work for cross-border order. This code is required at some sorting hub. Please ensure the service_code is INCLUDED on your shipping label, otherwise the parcel cannot be processed by the warehouse. If you didn't retrieve service_code after you first called this API, please try few more times within 30 minutes.
	ServiceCode string
	// The tracking number of fullfill order assigned by the shipping carrier for item shipment.
	TrackingNo string
}

type Forder struct {
	// The unique identifier for a fulfill order.
	ForderID string
	// The fulfill order logistics status. Applicable values: See Data Definition - LogisticsStatus.
	Status string
	// Logistics tracking info.
	TrackingLog []Log
	// The items included in this fulfill order.
	Items []Item
	//
	LogisticsInfo []LogisticsInfo
	// The first-mile tracking number.
	FMTN string
}

type GetForderInfoResponse struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The fulfill order list
	ForderList []Forder
	// The identifier for an API request for error tracking
	RequestID string
}

type GetEscrowReleasedOrdersRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// If many orders are available to retrieve, you may need to call GetEscrowReleaseTime multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data. Max entries per page is 100.
	PaginationEntriesPerPage int
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int
	// Release time range to filter order escrow information.
	ReleaseTimeFrom int
	// Release time range to filter order escrow information. The value should be higher than release_time_from.
	ReleaseTimeTo int
}

type Order struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// Order's escrow amount.
	PayoutAmount float64
	// Timestamp of escrow amount transaction finished.
	EscrowReleaseTime int
}

type GetEscrowReleasedOrdersResponse struct {
	// Filtered orders' escrow information.
	Orders []Order
}

type Parcel struct {
	// Itemids that will be put into a fullfillment order.
	ItemID int
}

type SplitOrderRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Shopee's unique identifier for an order.
	OrderSN string
	// Item information contained in fulfilment orders.
	Parcels []Parcel
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
}

type Forder struct {
	// Shopee's unique identifier for a fulfillment order.
	ForderID string
	// Item information contained in fulfillment orders.Number of items must be greater than or equal to 2. eg.[[{"item_id": 123}],[{"item_id": 456}]]
	Items []Item
}

type SplitOrderResponse struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// Information of fulfillment orders.
	Forders []Forder
	// The identifier for an API request for error tracking
	RequestID string
}

type UndoSplitOrderRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Shopee's unique identifier for an order.
	OrderSN string
}

type UndoSplitOrderResponse struct {
	// Whether or not the split order has been cancelled.
	Result string
	// The identifier for an API request for error tracking
	RequestID string
}

type GetUnbindOrderListRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int
	// If many unbind orders are available to retrieve, you may need to call GetUnbindOrderList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PaginationEntriesPerPage int
}

type Order struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The Shopee logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.
	LogisticStatus string
	// The unique identifier for a fulfillment order.
	ForderID string
}

type GetUnbindOrderListResponse struct {
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool
	//
	Orders []Order
	// The identifier for an API request for error tracking
	RequestID string
}
