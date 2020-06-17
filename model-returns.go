package shopeego

type ConfirmReturnRequest struct {
	// The serial number of return.
	ReturnSN string `json:"return_sn,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type ConfirmReturnResponse struct {
	// The serial number of return.
	ReturnSN string `json:"return_sn,omitempty"`
	//
	Msg string `json:"msg,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type DisputeReturnRequest struct {
	// The serial number of return.
	ReturnSN string `json:"return_sn,omitempty"`
	// Seller's email.
	Email string `json:"email,omitempty"`
	// The reason for seller dispute the return. Available value: NON_RECEIPT/OTHER/NOT_RECEIVED. See Data Definition- ReturnDisputeReason
	DisputeReason string `json:"dispute_reason,omitempty"`
	// The reason description for seller dispute the return.
	DisputeTextReason string `json:"dispute_text_reason,omitempty"`
	// Image URLs that seller provide. Seller can upload up 3 images at most.
	Images []string `json:"images,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type DisputeReturnResponse struct {
	// The serial number of return.
	ReturnSN string `json:"return_sn,omitempty"`
	//
	Msg string `json:"msg,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetReturnListRequest struct {
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset,omitempty"`
	// If many items are available to retrieve, you may need to call GetReturnList multiple times to retrieve all the data. Each result set is returned as a page of entries. Default is 40. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page,omitempty"`
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	CreateTimeFrom int `json:"create_time_from,omitempty"`
	// The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	CreateTimeTo int `json:"create_time_to,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetReturnListResponse struct {
	Returns []GetReturnListResponseReturn `json:"returns,omitempty"`
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool `json:"more,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetReturnDetailRequest struct {
	// The serial number of return.
	ReturnSN int `json:"return_sn,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetReturnDetailResponse struct {
	// Image URLs of return.
	Images []string `json:"images,omitempty"`
	// Reason for return product. Applicable values: See Data Definition- ReturnReason.
	Reason string `json:"reason,omitempty"`
	// Reason that buyer provide.
	TextReason string `json:"text_reason,omitempty"`
	// The serial number of return.
	ReturnSN int `json:"return_sn,omitempty"`
	// Amount of the refund.
	RefundAmount float64 `json:"refund_amount,omitempty,string"`
	// Currency of the return.
	Currency string `json:"currency,omitempty"`
	// The time of return create.
	CreateTime int `json:"create_time,omitempty"`
	// The time of modify return.
	UpdateTime int `json:"update_time,omitempty"`
	// Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
	Status string `json:"status,omitempty"`
	// The last time seller deal with this return.
	DueDate int `json:"due_date,omitempty"`
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNumber string `json:"tracking_number,omitempty"`
	// The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
	DisputeReason string `json:"dispute_reason,omitempty"`
	// The reason that seller provide. While the return has been disputed, this field is useful.
	DisputeTextReason string `json:"dispute_text_reason,omitempty"`
	// Items to be sent back to seller. Can be either integrated/non-integrated.
	NeedsLogistics bool `json:"needs_logistics,omitempty"`
	// Order price before discount.
	AmountBeforeDiscount float64 `json:"amount_before_discount,omitempty,string"`
	//
	User GetReturnDetailResponseUser `json:"user,omitempty"`
	//
	Item []GetReturnDetailResponseItem `json:"item,omitempty"`
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}
