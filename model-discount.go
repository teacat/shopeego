package shopeego

type AddDiscountRequest struct {
	// Title of the discount.
	DiscountName string `json:"discount_name"`
	// The time when discount activity start.
	StartTime int `json:"start_time"`
	// The time when discount activity end. The end time must be 1 hour later than start time.
	EndTime int `json:"end_time"`
	// Max item to upload is 50 in one API call.
	Items []AddDiscountRequestItem `json:"items"`
	//
	RequestBase
}

type AddDiscountResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// The number of items that add successfully.
	Count int `json:"count"`
	//
	Warning string `json:"warning"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type AddDiscountItemRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// Max item to upload is 50 in one API call.
	Items []AddDiscountItemRequestItem `json:"items"`
	//
	RequestBase
}

type AddDiscountItemResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// The number of items that add successfully.
	Count int `json:"count"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type DeleteDiscountRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	//
	RequestBase
}

type DeleteDiscountResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// The time when discount has been deleted.
	ModifyTime int `json:"modify_time"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type DeleteDiscountItemRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int `json:"variation_id"`
	//
	RequestBase
}

type DeleteDiscountItemResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// The time when item is deleted.
	ModifyTime int `json:"modify_time"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetDiscountDetailRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// Specifies the starting entry of data to return in the current call. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset"`
	// If many items of this discount are available to retrieve, you may need to call GetDiscountDetail multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page"`
	//
	RequestBase
}

type GetDiscountDetailResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// Title of the discount.
	DiscountName int `json:"discount_name"`
	// The time when discount activity start.
	StartTime int `json:"start_time"`
	// The time when discount activity end.
	EndTime int `json:"end_time"`
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool `json:"more"`
	//
	Items []GetDiscountDetailResponseItem `json:"items"`
	// The status of discount, applicable values: expired, ongoing, upcoming.
	Status string `json:"status"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetDiscountsListRequest struct {
	// The status filter for retriveing discount list. Available value: UPCOMING/ONGOING/EXPIRED/ALL.
	DiscountStatus string `json:"discount_status"`
	// Specifies the starting entry of data to return in the current call. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset"`
	// If many items are available to retrieve, you may need to call GetDiscountsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page"`
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the discount update time). The create_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days.
	UpdateTimeFrom int `json:"update_time_from"`
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the discount update time). The create_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days.
	UpdateTimeTo int `json:"update_time_to"`
	//
	RequestBase
}

type GetDiscountsListResponse struct {
	//
	Discount []GetDiscountsListResponseDiscount `json:"discount"`
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool `json:"more"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateDiscountRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// Title of the discount.
	DiscountName string `json:"discount_name"`
	// The time when discount activity end. The end time must be 1 hour later than start time. End time can only be shortened.
	EndTime int `json:"end_time"`
	// The time when discount activity start. The new start time must later than original start time. Start time cannot be changed after discount starts.
	StartTime int `json:"start_time"`
	//
	RequestBase
}

type UpdateDiscountResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// The time when items are updated.
	ModifyTime int `json:"modify_time"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateDiscountItemsRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	//
	Items []UpdateDiscountItemsRequestItem `json:"items"`
	//
	RequestBase
}

type UpdateDiscountItemsResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// The number of items that modify successfully.
	Count int `json:"count"`
	// The time when items are updated.
	ModifyTime int `json:"modify_time"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}
