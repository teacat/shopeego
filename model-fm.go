package shopeego

type GenerateFMTrackingNoRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// This field is used for seller to specify the declare time.
	DeclareDate string `json:"declare_date,omitempty"`
	// The number of first-mile tracking numbers generated.
	Quantity int `json:"quantity,omitempty"`
	// This object contains detailed breakdown for the seller address.
	SellerInfo GenerateFMTrackingNoRequestSellerInfo `json:"seller_info,omitempty"`
}

type GenerateFMTrackingNoResponse struct {
	// The first-mile tracking number.
	FMTNList []string `json:"fmtn_list,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetShopFMTrackingNoRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// The start time of declare_date.
	FromDate string `json:"from_date,omitempty"`
	// The end time of declare_date
	ToDate string `json:"to_date,omitempty"`
	// If a lot of first-mile tracking numbers are available to retrieve, you may need to call GetUnbindOrderList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page,omitempty"`
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset,omitempty"`
}

type GetShopFMTrackingNoResponse struct {
	// This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	HasMore bool `json:"has_more,omitempty"`
	//
	FMTNList []GetShopFMTrackingNoResponseFMTNList `json:"fmtn_list,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type FirstMileCodeBindOrderRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// The set of ordersn. You can specify up to 50 ordersns in this call.
	OrderList []FirstMileCodeBindOrderRequestOrder `json:"order_list,omitempty"`
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
	// The shipment method for bound orders, should be pickup or dropoff.
	ShipmentMethod string `json:"shipment_method,omitempty"`
	// The identity of logistic channel
	LogisticID int64 `json:"logistic_id,omitempty"`
	// The volume of the parcel.
	Volume float64 `json:"volume,omitempty,string"`
	// The weight of the parcel.
	Weight float64 `json:"weight,omitempty,string"`
	// The width of the parcel.
	Width float64 `json:"width,omitempty,string"`
	// The length of the parcel.
	Length float64 `json:"length,omitempty,string"`
	// The height of the parcel.
	Height float64 `json:"height,omitempty,string"`
	// Use this field to specify the region you want to ship parcel.
	Area string `json:"area,omitempty"`
}

type FirstMileCodeBindOrderResponse struct {
	// This is to indicate whether orders are bound successfully.
	Success bool `json:"success,omitempty"`
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
	//
	FailList []FirstMileCodeBindOrderResponseFail `json:"fail_list,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetFmTnDetailRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
	// This field indicates the current page number. By default, each page returns 100 orders, if the current page is less than 100, it returns all
	Page int `json:"page,omitempty"`
}

type GetFmTnDetailResponse struct {
	// The identity of logistic channel
	LogisticID int64 `json:"logistic_id,omitempty"`
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
	// The shipment method for bound orders, should be pickup or dropoff.
	ShipmentMethod string `json:"shipment_method,omitempty"`
	// The logistics status for first-mile tracking number.
	Status string `json:"status,omitempty"`
	//
	Orders []GetFmTnDetailResponseOrder `json:"orders,omitempty"`
	// The specified delivery date.
	DeclareDate string `json:"declare_date,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
	// This field represents the total number of pages calculated by 100 orders per page.
	TotalPage int `json:"total_page,omitempty"`
}

type GetFMTrackingNoWaybillRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// The set of first-mile tracking numbers. Up to 50 tracking numbers in one call.
	FMTNList []string `json:"fmtn_list,omitempty"`
	// Option to get batch airway bills in single file. Default value is false.
	IsBatch bool `json:"is_batch,omitempty"`
}

type GetFMTrackingNoWaybillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult GetFMTrackingNoWaybillResponseBatchResult `json:"batch_result,omitempty"`
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result GetFMTrackingNoWaybillResponseResult `json:"result,omitempty"`
	// The number of Tracking Number to get waybills in this call.
	TotalCount int `json:"total_count,omitempty"`
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []GetFMTrackingNoWaybillResponseError `json:"errors,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetShopFirstMileChannelRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Use this field to specify the region you want to ship parcel.
	Area string `json:"area,omitempty"`
}

type GetShopFirstMileChannelResponse struct {
	//
	Logistics []GetShopFirstMileChannelResponseLogistic `json:"logistics,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type FirstMileUnbindRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
	//
	OrderList FirstMileUnbindRequestOrderList `json:"order_list,omitempty"`
}

type FirstMileUnbindResponse struct {
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
	// This is to indicate whether orders are unbound successfully.
	Success bool `json:"success,omitempty"`
	//
	FailList FirstMileUnbindResponseFailList `json:"fail_list,omitempty"`
}
