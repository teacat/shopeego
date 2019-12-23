package shopeego

type Variation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int
	// The discount price of the item.
	VariationPromotionPrice float64
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	//
	Variations []Variation
	// The discount price of the item. If the item has no variation, this param is necessary.
	ItemPromotionPrice float64
	// The max number of this product in the promotion price.
	PurchaseLimit int
}

type AddDiscountRequest struct {
	// Title of the discount.
	DiscountName string
	// The time when discount activity start.
	StartTime int
	// The time when discount activity end. The end time must be 1 hour later than start time.
	EndTime int
	// Max item to upload is 50 in one API call.
	Items []Item
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type AddDiscountResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// The number of items that add successfully.
	Count int
	//
	Warning string
	// The identifier for an API request for error tracking
	RequestID string
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int
	// The discount price of the item.
	VariationPromotionPrice float64
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	//
	Variations []Variation
	// The discount price of the item. If the item has no variation, this param is necessary.
	ItemPromotionPrice float64
	// The max number of this product in the promotion price.
	PurchaseLimit int
}

type AddDiscountItemRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// Max item to upload is 50 in one API call.
	Items []Item
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type AddDiscountItemResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// The number of items that add successfully.
	Count int
	// The identifier for an API request for error tracking
	RequestID string
}

type DeleteDiscountRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type DeleteDiscountResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// The time when discount has been deleted.
	ModifyTime int
	// The identifier for an API request for error tracking
	RequestID string
}

type DeleteDiscountItemRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// Shopee's unique identifier for an item.
	ItemID int
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type DeleteDiscountItemResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// Shopee's unique identifier for an item.
	ItemID int
	// Shopee's unique identifier for a variation of an item.
	VariationID int
	// The time when item is deleted.
	ModifyTime int
	// The identifier for an API request for error tracking
	RequestID string
}

type GetDiscountDetailRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// Specifies the starting entry of data to return in the current call. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int
	// If many items of this discount are available to retrieve, you may need to call GetDiscountDetail multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PaginationEntriesPerPage int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int
	// Name of the variation that belongs to the same item.
	VariationName string
	// The original price before discount of the variation.
	VariationOriginalPrice float64
	// The discount price of the variation.
	VariationPromotionPrice float64
	// The current stock quantity of the variation.
	VariationStock int
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// Name of the item in local language.
	ItemName string
	// The max number of this product in the promotion price.
	PurchaseLimit int
	// The original price before discount of the item. If there is variation, this value is 0.
	ItemOriginalPrice float64
	// The discount price of the item. If there is variation, this value is 0.
	ItemPromotionPrice float64
	// The current stock quantity of the item.
	Stock int
	//
	Variations []Variation
}

type GetDiscountDetailResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// Title of the discount.
	DiscountName int
	// The time when discount activity start.
	StartTime int
	// The time when discount activity end.
	EndTime int
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool
	//
	Items []Item
	// The status of discount, applicable values: expired, ongoing, upcoming.
	Status string
	// The identifier for an API request for error tracking
	RequestID string
}

type GetDiscountsListRequest struct {
	// The status filter for retriveing discount list. Available value: UPCOMING/ONGOING/EXPIRED/ALL.
	DiscountStatus string
	// Specifies the starting entry of data to return in the current call. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int
	// If many items are available to retrieve, you may need to call GetDiscountsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data.
	PaginationEntriesPerPage int
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the discount update time). The create_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days.
	UpdateTimeFrom int
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the discount update time). The create_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days.
	UpdateTimeTo int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Discount struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// Title of the discount.
	DiscountName string
	// The time when discount activity start.
	StartTime int
	// The time when discount activity end.
	EndTime int
	// The status of discount, applicable values: expired, ongoing, upcoming.
	Status string
}

type GetDiscountsListResponse struct {
	//
	Discount []Discount
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool
	// The identifier for an API request for error tracking
	RequestID string
}

type UpdateDiscountRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// Title of the discount.
	DiscountName string
	// The time when discount activity end. The end time must be 1 hour later than start time. End time can only be shortened.
	EndTime int
	// The time when discount activity start. The new start time must later than original start time. Start time cannot be changed after discount starts.
	StartTime int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type UpdateDiscountResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// The time when items are updated.
	ModifyTime int
	// The identifier for an API request for error tracking
	RequestID string
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int
	// The discount price of the item.
	VariationPromotionPrice float64
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// The max number of this product in the promotion price.
	PurchaseLimit int
	// The discount price of the item.
	ItemOriginalPrice float64
	//
	Variations []Variation
}

type UpdateDiscountItemsRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	//
	Items []Item
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type UpdateDiscountItemsResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int
	// The number of items that modify successfully.
	Count int
	// The time when items are updated.
	ModifyTime int
	// The identifier for an API request for error tracking
	RequestID string
}
