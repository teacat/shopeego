package shopeego

type AddShopCategoryRequest struct {
	// ShopCategory's name.
	Name string `json:"name,omitempty"`
	// ShopCategory's sort weight.
	SortWeight int `json:"sort_weight,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type AddShopCategoryResponse struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
}

type GetShopCategoriesRequest struct {
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset,omitempty"`
	// If many items are available to retrieve, you may need to call GetItemsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetShopCategoriesResponse struct {
	//
	ShopCategorys []GetShopCategoriesResponseCategory `json:"shop_categorys,omitempty"`
	// This is to indicate whether the shop categorys list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of shop categorys.
	More bool `json:"more,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type DeleteShopCategoryRequest struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type DeleteShopCategoryResponse struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	//
	Msg string `json:"msg,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateShopCategoryRequest struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	// ShopCategory's name.
	Name string `json:"name,omitempty"`
	// ShopCategory's sort weight.
	SortWeight int `json:"sort_weight,omitempty"`
	// ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
	Status string `json:"status,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdateShopCategoryResponse struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	// ShopCategory's name.
	Name string `json:"name,omitempty"`
	// ShopCategory's sort weight.
	SortWeight int `json:"sort_weight,omitempty"`
	// ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
	Status string `json:"status,omitempty"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
}

type AddItemsRequest struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	// Shopee's unique identifiers list for an item.
	Items []int `json:"items,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type AddItemsResponse struct {
	// List of invalid item ids.
	InvalidItemID []int `json:"invalid_item_id,omitempty"`
	// ShopCategory's unique identifier.
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	// Number of items in the shop category.
	Count int `json:"count,omitempty"`
	// The identifier for an API request for error tracking.
	Request string `json:"request,omitempty"`
}

type GetItemsRequest struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset,omitempty"`
	// If many items are available to retrieve, you may need to call GetItems multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 1000) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is used to specify the maximum number of entries to return in a single "page" of data. And the default will be 1000 as well.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page,omitempty"`
}

type GetItemsResponse struct {
	//
	Items []GetItemsResponseItem `json:"items,omitempty"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
}

type DeleteItemsRequest struct {
	// The id of the shop category
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	// The list of items need to be deleted
	Items []int `json:"items,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type DeleteItemsResponse struct {
	// The id of the shop category
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	// The list of item ids which are invalid
	InvalidItemID []int `json:"invalid_item_id,omitempty"`
	// count of items under this shop category after deletion
	Count int `json:"count,omitempty"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}
