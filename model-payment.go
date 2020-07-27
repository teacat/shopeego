package shopeego

type GetTransactionListRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int64 `json:"pagination_offset,omitempty"`
	// If many transactions are available to retrieve, you may need to call GetTransactionList multiple times to retrieve all the data. Each result set is returned as a page of entries. Default is 40. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	PaginationEntriesPerPage int64 `json:"pagination_entries_per_page,omitempty"`
	// The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	CreateTimeFrom int64 `json:"create_time_from,omitempty"`
	// The create_time_to field is the ending date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	CreateTimeTo int64 `json:"create_time_to,omitempty"`
	// This field indicates the wallet type.
	WalletType string `json:"wallet_type,omitempty"`
	// This field indicates the transaction type.
	TransactionType string `json:"transaction_type,omitempty"`
	//
}

type GetTransactionListResponse struct {
	// This is to indicate whether the transaction list is more than one page. If this value is true, you may want to continue to check next page to retrieve orders.
	HasMore bool `json:"has_more,omitempty"`
	//
	TransactionList []GetTransactionListResponseTransactionList `json:"transaction_list,omitempty"`
}
