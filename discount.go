package shopeego

type Variation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int `json:"variation_id"`
	// The discount price of the item.
	VariationPromotionPrice float64 `json:"variation_promotion_price"`
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	//
	Variations []Variation `json:"variations"`
	// The discount price of the item. If the item has no variation, this param is necessary.
	ItemPromotionPrice float64 `json:"item_promotion_price"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit"`
}

type AddDiscountRequest struct {
	// Title of the discount.
	DiscountName string `json:"discount_name"`
	// The time when discount activity start.
	StartTime int `json:"start_time"`
	// The time when discount activity end. The end time must be 1 hour later than start time.
	EndTime int `json:"end_time"`
	// Max item to upload is 50 in one API call.
	Items []Item `json:"items"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
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

type Variation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int `json:"variation_id"`
	// The discount price of the item.
	VariationPromotionPrice float64 `json:"variation_promotion_price"`
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	//
	Variations []Variation `json:"variations"`
	// The discount price of the item. If the item has no variation, this param is necessary.
	ItemPromotionPrice float64 `json:"item_promotion_price"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit"`
}

type AddDiscountItemRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// Max item to upload is 50 in one API call.
	Items []Item `json:"items"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
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
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
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
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
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
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// Name of the variation that belongs to the same item.
	VariationName string `json:"variation_name"`
	// The original price before discount of the variation.
	VariationOriginalPrice float64 `json:"variation_original_price"`
	// The discount price of the variation.
	VariationPromotionPrice float64 `json:"variation_promotion_price"`
	// The current stock quantity of the variation.
	VariationStock int `json:"variation_stock"`
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Name of the item in local language.
	ItemName string `json:"item_name"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit"`
	// The original price before discount of the item. If there is variation, this value is 0.
	ItemOriginalPrice float64 `json:"item_original_price"`
	// The discount price of the item. If there is variation, this value is 0.
	ItemPromotionPrice float64 `json:"item_promotion_price"`
	// The current stock quantity of the item.
	Stock int `json:"stock"`
	//
	Variations []Variation `json:"variations"`
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
	Items []Item `json:"items"`
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
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type Discount struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// Title of the discount.
	DiscountName string `json:"discount_name"`
	// The time when discount activity start.
	StartTime int `json:"start_time"`
	// The time when discount activity end.
	EndTime int `json:"end_time"`
	// The status of discount, applicable values: expired, ongoing, upcoming.
	Status string `json:"status"`
}

type GetDiscountsListResponse struct {
	//
	Discount []Discount `json:"discount"`
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
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type UpdateDiscountResponse struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	// The time when items are updated.
	ModifyTime int `json:"modify_time"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int `json:"variation_id"`
	// The discount price of the item.
	VariationPromotionPrice float64 `json:"variation_promotion_price"`
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit"`
	// The discount price of the item.
	ItemOriginalPrice float64 `json:"item_original_price"`
	//
	Variations []Variation `json:"variations"`
}

type UpdateDiscountItemsRequest struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int `json:"discount_id"`
	//
	Items []Item `json:"items"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
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
