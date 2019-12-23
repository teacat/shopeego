package shopeego

type AddShopCategoryRequest struct {
	// ShopCategory's name.
	Name string `json:"name"`
	// ShopCategory's sort weight.
	SortWeight int `json:"sort_weight"`
	//
	RequestBase
}

type AddShopCategoryResponse struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int `json:"shop_category_id"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id"`
}

type GetShopCategoriesRequest struct {
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset"`
	// If many items are available to retrieve, you may need to call GetItemsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page"`
	//
	RequestBase
}

type GetShopCategoriesResponse struct {
	//
	ShopCategorys []ShopCategory `json:"shop_categorys"`
	// This is to indicate whether the shop categorys list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of shop categorys.
	More bool `json:"more"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type DeleteShopCategoryRequest struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int `json:"shop_category_id"`
	//
	RequestBase
}

type DeleteShopCategoryResponse struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int `json:"shop_category_id"`
	//
	Msg string `json:"msg"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateShopCategoryRequest struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int `json:"shop_category_id"`
	// ShopCategory's name.
	Name string `json:"name"`
	// ShopCategory's sort weight.
	SortWeight int `json:"sort_weight"`
	// ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
	Status string `json:"status"`
	//
	RequestBase
}

type UpdateShopCategoryResponse struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int `json:"shop_category_id"`
	// ShopCategory's name.
	Name string `json:"name"`
	// ShopCategory's sort weight.
	SortWeight int `json:"sort_weight"`
	// ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
	Status string `json:"status"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id"`
}

type AddItemsRequest struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int `json:"shop_category_id"`
	// Shopee's unique identifiers list for an item.
	Items []int `json:"items"`
	//
	RequestBase
}

type AddItemsResponse struct {
	// List of invalid item ids.
	InvalidItemID []int `json:"invalid_item_id"`
	// ShopCategory's unique identifier.
	ShopCategoryID int `json:"shop_category_id"`
	// Number of items in the shop category.
	Count int `json:"count"`
	// The identifier for an API request for error tracking.
	Request string `json:"request"`
}

type GetItemsRequest struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int `json:"shop_category_id"`
	//
	RequestBase
}

type GetItemsResponse struct {
	//
	Items []Item `json:"items"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id"`
}

type DeleteItemsRequest struct {
	// The id of the shop category
	ShopCategoryID int `json:"shop_category_id"`
	// The list of items need to be deleted
	Items []int `json:"items"`
	//
	RequestBase
}

type DeleteItemsResponse struct {
	// The id of the shop category
	ShopCategoryID int `json:"shop_category_id"`
	// The list of item ids which are invalid
	InvalidItemID []int `json:"invalid_item_id"`
	// count of items under this shop category after deletion
	Count int `json:"count"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}
