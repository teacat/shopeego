package shopeego

type Info struct {
	// The full address of the seller.
	Address string `json:"address"`
	// Seller's name for the address.
	Name string `json:"name"`
	// Seller's postal code.
	Zipcode string `json:"zipcode"`
	// Seller's location.
	Area string `json:"area"`
	// Seller's phone number.
	Phone string `json:"phone"`
}

type GenerateFMTrackingNoRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// This field is used for seller to specify the declare time.
	DeclareDate string `json:"declare_date"`
	// The number of first-mile tracking numbers generated.
	Quantity int `json:"quantity"`
	// This object contains detailed breakdown for the seller address.
	SellerInfo Info `json:"seller_info"`
}

type GenerateFMTrackingNoResponse struct {
	// The first-mile tracking number.
	FMTNList []string `json:"fmtn_list"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetShopFMTrackingNoRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// The start time of declare_date.
	FromDate string `json:"from_date"`
	// The end time of declare_date
	ToDate string `json:"to_date"`
	// If a lot of first-mile tracking numbers are available to retrieve, you may need to call GetUnbindOrderList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page"`
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset"`
}

type List struct {
	// The specified delivery date.
	DeclareDate string `json:"declare_date"`
	// The logistics status for bound orders.
	Status string `json:"status"`
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
}

type GetShopFMTrackingNoResponse struct {
	// This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	HasMore bool `json:"has_more"`
	//
	FMTNList []List `json:"fmtn_list"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type Order struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
}

type FirstMileCodeBindOrderRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// The set of ordersn. You can specify up to 50 ordersns in this call.
	OrderList []Order `json:"order_list"`
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
	// The shipment method for bound orders, should be pickup or dropoff.
	ShipmentMethod string `json:"shipment_method"`
	// The identity of logistic channel
	LogisticID int `json:"logistic_id"`
	// The volume of the parcel.
	Volume float64 `json:"volume"`
	// The weight of the parcel.
	Weight float64 `json:"weight"`
	// The width of the parcel.
	Width float64 `json:"width"`
	// The length of the parcel.
	Length float64 `json:"length"`
	// The height of the parcel.
	Height float64 `json:"height"`
	// Use this field to specify the region you want to ship parcel.
	Area string `json:"area"`
}

type List struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
	// The reason why the order/fulfillment order cannot be bound.
	Reason string `json:"reason"`
}

type FirstMileCodeBindOrderResponse struct {
	// This is to indicate whether orders are bound successfully.
	Success bool `json:"success"`
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
	//
	FailList []List `json:"fail_list"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetFmTnDetailRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// The first-mile tracking number.
	FMTN sttring `json:"fmtn"`
}

type Order struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
	// The tracking number of SLS for orders/forders.
	SLSTN string `json:"slstn"`
}

type GetFmTnDetailResponse struct {
	// The identity of logistic channel
	LogisticID int `json:"logistic_id"`
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
	// The shipment method for bound orders, should be pickup or dropoff.
	ShipmentMethod string `json:"shipment_method"`
	// The logistics status for first-mile tracking number.
	Status string `json:"status"`
	//
	Orders []Order `json:"orders"`
	// The specified delivery date.
	DeclareDate string `json:"declare_date"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetFMTrackingNoWaybillRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// The set of first-mile tracking numbers. Up to 50 tracking numbers in one call.
	FMTNList []string `json:"fmtn_list"`
	// Option to get batch airway bills in single file. Default value is false.
	IsBatch bool `json:"is_batch"`
}

type Error struct {
	//
	ErrorCode string `json:"error_code"`
	// The detail information of this error.
	ErrorDescription string `json:"error_description"`
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
}

type Waybill struct {
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
	// The url of retrieving waybill.
	Waybill string `json:"waybill"`
}

type Result struct {
	// This Object contains the waybill to each tracking number.
	Waybills []Waybill `json:"waybills"`
	// The number of Tracking Number to get waybills in this call.
	TotalCount int `json:"total_count"`
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
}

type Batch struct {
	// The list contains urls of retrieving waybill in PDF format. Each url contains the airway bills which is in the same logistics channel.
	Waybills []string `json:"waybills"`
	// The number of Tracking Number to get waybills in this call.
	TotalCount int `json:"total_count"`
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
}

type GetFMTrackingNoWaybillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult Batch `json:"batch_result"`
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result Result `json:"result"`
	// The number of Tracking Number to get waybills in this call.
	TotalCount int `json:"total_count"`
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetShopFirstMileChannelRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// Use this field to specify the region you want to ship parcel.
	Area string `json:"area"`
}

type Logistic struct {
	// The identity of logistic channel.
	LogisticID int `json:"logistic_id"`
	// The name of logistic.
	LogisticName string `json:"logistic_name"`
	// The shipment method for bound orders, should be pickup or dropoff.
	ShipmentMethod string `json:"shipment_method"`
}

type GetShopFirstMileChannelResponse struct {
	//
	Logistics []Logistic `json:"logistics"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}
