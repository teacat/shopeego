package shopeego

type ConfirmReturnRequest struct {
	// The serial number of return.
	ReturnSN string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type ConfirmReturnResponse struct {
	// The serial number of return.
	ReturnSN string
	//
	Msg string
	// The identifier for an API request for error tracking
	RequestID string
}

type DisputeReturnRequest struct {
	// The serial number of return.
	ReturnSN string
	// Seller's email.
	Email string
	// The reason for seller dispute the return. Available value: NON_RECEIPT/OTHER/NOT_RECEIVED. See Data Definition- ReturnDisputeReason
	DisputeReason string
	// The reason description for seller dispute the return.
	DisputeTextReason string
	// Image URLs that seller provide. Seller can upload up 3 images at most.
	Images []string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type DisputeReturnResponse struct {
	// The serial number of return.
	ReturnSN string
	//
	Msg string
	// The identifier for an API request for error tracking
	RequestID string
}

type GetReturnListRequest struct {
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int
	// If many items are available to retrieve, you may need to call GetReturnList multiple times to retrieve all the data. Each result set is returned as a page of entries. Default is 40. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	PaginationEntriesPerPage int
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	CreateTimeFrom int
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	CreateTimeTo int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type User struct {
	// Buyer's nickname.
	Username int
	// Buyer's email.
	Email string
	// Buyer's portrait.
	Protrait string
}

type Item struct {
	// Item id.
	ItemID int
	// Shopee's unique identifier for a variation of an item.
	VariationID int
	// Name of item in local language.
	Name string
	// Image URLs of item.
	Images []string
	// Amount of this item.
	Amount int
	// The price of Item.
	ItemPrice float64
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool
	// The unique identity of an addon deal.
	AddOnDealID int
}

type Return struct {
	// Image URLs of return.
	Images []string
	// Reason for return product. Applicable values: See Data Definition- ReturnReason.
	Reason string
	// Reason that buyer provide.
	TextReason string
	// The serial number of return.
	ReturnSN int
	// Amount of the refund.
	RefundAmount float64
	// Currency of the return.
	Currency string
	// The time of return create.
	CreateTime int
	// The time of modify return.
	UpdateTime int
	// Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
	Status string
	// The last time seller deal with this return.
	DueDate int
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNumber string
	// The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
	DisputeReason string
	// The reason that seller provide. While the return has been disputed, this field is useful.
	DisputeTextReason string
	// Items to be sent back to seller. Can be either integrated/non-integrated.
	NeedsLogistics bool
	// Order price before discount.
	AmountBeforeDiscount floatt64
	//
	User User
	//
	Item []Item
	// Shopee's unique identifier for an order.
	OrderSN string
}

type GetReturnListResponse struct {
	Returns []Return
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool
	// The identifier for an API request for error tracking
	RequestID string
}

type GetReturnDetailRequest struct {
	// The serial number of return.
	ReturnSN int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type User struct {
	// Buyer's nickname.
	Username int
	// Buyer's email.
	Email string
	// Buyer's portrait.
	Portrait string
}

type Item struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int
	// Name of item in local language.
	Name string
	// Image URLs of item.
	Images []string
	// Amount of this item.
	Amount int
	// The price of item.
	ItemPrice float64
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool
	// The unique identity of an addon deal.
	AddOnDealID int
}

type GetReturnDetailResponse struct {
	// Image URLs of return.
	Images []string
	// Reason for return product. Applicable values: See Data Definition- ReturnReason.
	Reason string
	// Reason that buyer provide.
	TextReason string
	// The serial number of return.
	ReturnSN int
	// Amount of the refund.
	RefundAmount float64
	// Currency of the return.
	Currency string
	// The time of return create.
	CreateTime int
	// The time of modify return.
	UpdateTime int
	// Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
	Status string
	// The last time seller deal with this return.
	DueDate int
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNumber string
	// The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
	DisputeReason string
	// The reason that seller provide. While the return has been disputed, this field is useful.
	DisputeTextReason string
	// Items to be sent back to seller. Can be either integrated/non-integrated.
	NeedsLogistics bool
	// Order price before discount.
	AmountBeforeDiscount float64
	//
	User User
	//
	Item []Item
	// Shopee's unique identifier for an order.
	OrderSN string
	// The identifier for an API request for error tracking
	RequestID string
}
