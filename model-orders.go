package shopeego

type GetOrdersListRequest struct {
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days. Must include only create_time or update_time in the request.
	CreateTimeFrom int `json:"create_time_from,omitempty"`
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_to field is the ending date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days. Must include only create_time or update_time in the request.
	CreateTimeTo int `json:"create_time_to,omitempty"`
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the order update time). The update_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days. Must include only create_time or update_time in the request.
	UpdateTimeFrom int `json:"update_time_from,omitempty"`
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the order update time). The update_time_to field is the ending date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days. Must include only create_time or update_time in the request.
	UpdateTimeTo int `json:"update_time_to,omitempty"`
	// If many orders are available to retrieve, you may need to call GetOrdersList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call.
	// This integer value is used to specify the maximum number of entries to return in a single "page" of data. Max entries per page is 100.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page,omitempty"`
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetOrdersListResponse struct {
	// The set of orders that match the ordersn or filter criteria specified.
	Orders []GetOrdersListResponseOrder `json:"orders,omitempty"`
	// This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	More bool `json:"more,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetOrdersByStatusRequest struct {
	// The order_status filter for retriveing orders. Available value: ALL/UNPAID/READY_TO_SHIP/COMPLETED/IN_CANCEL/CANCELLED/TO_RETURN
	OrderStatus string `json:"order_status,omitempty"`
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days. Must include only create_time or update_time in the request.
	CreateTimeFrom int `json:"create_time_from,omitempty"`
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_to field is the ending date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days. Must include only create_time or update_time in the request.
	CreateTimeTo int `json:"create_time_to,omitempty"`
	// If many orders are available to retrieve, you may need to call GetOrdersList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call.
	// This integer value is used to specify the maximum number of entries to return in a single "page" of data. Max entries per page is 100.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page,omitempty"`
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetOrdersByStatusResponse struct {
	// The set of orders that match the ordersn or filter criteria specified.
	Orders []GetOrdersByStatusResponseOrder `json:"orders,omitempty"`
	// This is to indicate whether the order list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	More bool `json:"more,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetOrderDetailsRequest struct {
	// The set of ordersn. You can specify up to 50 ordersns in this call.
	OrderSNList []string `json:"ordersn_list,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetOrderDetailsResponse struct {
	// The set of orders that match the ordersn or filter criteria specified.
	Orders []GetOrderDetailsResponseOrder `json:"orders,omitempty"`
	// Orders that encountered error
	Errors []string `json:"errors,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetEscrowDetailsRequest struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetEscrowDetailsResponse struct {
	// My Income infomation
	Order GetEscrowDetailsResponseOrder `json:"order,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type AddOrderNoteRequest struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The note seller made for own reference.
	Note string `json:"note,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type AddOrderNoteResponse struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The success or error message.
	Msg string `json:"msg,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type CancelOrderRequest struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The reason seller want to cancel this order. Applicable values: OUT_OF_STOCK, CUSTOMER_REQUEST, UNDELIVERABLE_AREA, COD_NOT_SUPPORTED.
	CancelReason string `json:"cancel_reason,omitempty"`
	// ID of item. Required when cancel_reason is OUT_OF_STOCK.
	ItemID int64 `json:"item_id,omitempty"`
	// ID of the variation that belongs to the same item.Required when cancel_reason is OUT_OF_STOCK.
	VariationID int64 `json:"variation_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type CancelOrderResponse struct {
	// The time when the order is updated.
	ModifiedTime int `json:"modified_time,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type AcceptBuyerCancellationRequest struct {
	// The order to be accepted cancellation request.
	OrderSN string `json:"ordersn,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type AcceptBuyerCancellationResponse struct {
	// The time when the order is updated.
	ModifiedTime int `json:"modified_time,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type RejectBuyerCancellationRequest struct {
	// The order to be rejected cancellation request.
	OrderSN string `json:"ordersn,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type RejectBuyerCancellationResponse struct {
	// The time when the order is updated.
	ModifiedTime int `json:"modified_time,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetForderInfoRequest struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetForderInfoResponse struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The fulfill order list
	ForderList []GetForderInfoResponseForder `json:"forder_list,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetEscrowReleasedOrdersRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// If many orders are available to retrieve, you may need to call GetEscrowReleaseTime multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data. Max entries per page is 100.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page,omitempty"`
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset,omitempty"`
	// Release time range to filter order escrow information.
	ReleaseTimeFrom int `json:"release_time_from,omitempty"`
	// Release time range to filter order escrow information. The value should be higher than release_time_from.
	ReleaseTimeTo int `json:"release_time_to,omitempty"`
}

type GetEscrowReleasedOrdersResponse struct {
	// Filtered orders' escrow information.
	Orders []GetEscrowReleasedOrdersResponseOrder `json:"orders,omitempty"`
}

type SplitOrderRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// Item information contained in fulfilment orders.
	Parcels []SplitOrderRequestParcel `json:"parcels,omitempty"`
}

type SplitOrderResponse struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// Information of fulfillment orders.
	Forders []SplitOrderResponseForder `json:"forders,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UndoSplitOrderRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
}

type UndoSplitOrderResponse struct {
	// Whether or not the split order has been cancelled.
	Result string `json:"result,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetUnbindOrderListRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset,omitempty"`
	// If many unbind orders are available to retrieve, you may need to call GetUnbindOrderList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page,omitempty"`
}

type GetUnbindOrderListResponse struct {
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool `json:"more,omitempty"`
	//
	Orders []GetUnbindOrderListResponseOrder `json:"orders,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type MyIncomeRequest struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type MyIncomeResponse struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The username of buyer.
	BuyerUserName string `json:"buyer_user_name,omitempty"`
	// The list of the serial number of return.
	ReturnSNList []string `json:"returnsn_list,omitempty"`
	//
	OrderIncome MyIncomeResponseOrderIncome `json:"order_income,omitempty"`
}
