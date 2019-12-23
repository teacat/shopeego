package shopeego

type Info struct {
	// The full address of the seller.
	Address string
	// Seller's name for the address.
	Name string
	// Seller's postal code.
	Zipcode string
	// Seller's location.
	Area string
	// Seller's phone number.
	Phone string
}

type GenerateFMTrackingNoRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// This field is used for seller to specify the declare time.
	DeclareDate string
	// The number of first-mile tracking numbers generated.
	Quantity int
	// This object contains detailed breakdown for the seller address.
	SellerInfo Info
}

type GenerateFMTrackingNoResponse struct {
	// The first-mile tracking number.
	FMTNList []string
	// The identifier for an API request for error tracking
	RequestID string
}

type GetShopFMTrackingNoRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// The start time of declare_date.
	FromDate string
	// The end time of declare_date
	ToDate string
	// If a lot of first-mile tracking numbers are available to retrieve, you may need to call GetUnbindOrderList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PaginationEntriesPerPage int
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int
}

type List struct {
	// The specified delivery date.
	DeclareDate string
	// The logistics status for bound orders.
	Status string
	// The first-mile tracking number.
	FMTN string
}

type GetShopFMTrackingNoResponse struct {
	// This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	HasMore bool
	//
	FMTNList []List
	// The identifier for an API request for error tracking
	RequestID string
}

type Order struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The unique identifier for a fulfillment order.
	ForderID string
}

type FirstMileCodeBindOrderRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// The set of ordersn. You can specify up to 50 ordersns in this call.
	OrderList []Order
	// The first-mile tracking number.
	FMTN string
	// The shipment method for bound orders, should be pickup or dropoff.
	ShipmentMethod string
	// The identity of logistic channel
	LogisticID int
	// The volume of the parcel.
	Volume float64
	// The weight of the parcel.
	Weight float64
	// The width of the parcel.
	Width float64
	// The length of the parcel.
	Length float64
	// The height of the parcel.
	Height float64
	// Use this field to specify the region you want to ship parcel.
	Area string
}

type List struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The unique identifier for a fulfillment order.
	ForderID string
	// The reason why the order/fulfillment order cannot be bound.
	Reason string
}

type FirstMileCodeBindOrderResponse struct {
	// This is to indicate whether orders are bound successfully.
	Success bool
	// The first-mile tracking number.
	FMTN string
	//
	FailList []List
	// The identifier for an API request for error tracking
	RequestID string
}

type GetFmTnDetailRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// The first-mile tracking number.
	FMTN sttring
}

type Order struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The unique identifier for a fulfillment order.
	ForderID string
	// The tracking number of SLS for orders/forders.
	SLSTN string
}

type GetFmTnDetailResponse struct {
	// The identity of logistic channel
	LogisticID int
	// The first-mile tracking number.
	FMTN string
	// The shipment method for bound orders, should be pickup or dropoff.
	ShipmentMethod string
	// The logistics status for first-mile tracking number.
	Status string
	//
	Orders []Order
	// The specified delivery date.
	DeclareDate string
	// The identifier for an API request for error tracking
	RequestID string
}

type GetFMTrackingNoWaybillRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// The set of first-mile tracking numbers. Up to 50 tracking numbers in one call.
	FMTNList []string
	// Option to get batch airway bills in single file. Default value is false.
	IsBatch bool
}

type Error struct {
	//
	ErrorCode string
	// The detail information of this error.
	ErrorDescription string
	// The first-mile tracking number.
	FMTN string
}

type Waybill struct {
	// The first-mile tracking number.
	FMTN string
	// The url of retrieving waybill.
	Waybill string
}

type Result struct {
	// This Object contains the waybill to each tracking number.
	Waybills []Waybill
	// The number of Tracking Number to get waybills in this call.
	TotalCount int
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []Error
}

type Batch struct {
	// The list contains urls of retrieving waybill in PDF format. Each url contains the airway bills which is in the same logistics channel.
	Waybills []string
	// The number of Tracking Number to get waybills in this call.
	TotalCount int
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []Error
}

type GetFMTrackingNoWaybillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult Batch
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result Result
	// The number of Tracking Number to get waybills in this call.
	TotalCount int
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []Error
	// The identifier for an API request for error tracking
	RequestID string
}

type GetShopFirstMileChannelRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Use this field to specify the region you want to ship parcel.
	Area string
}

type Logistic struct {
	// The identity of logistic channel.
	LogisticID int
	// The name of logistic.
	LogisticName string
	// The shipment method for bound orders, should be pickup or dropoff.
	ShipmentMethod string
}

type GetShopFirstMileChannelResponse struct {
	//
	Logistics []Logistic
	// The identifier for an API request for error tracking
	RequestID string
}
