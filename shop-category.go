package shopeego

type AddShopCategoryRequest struct {
	// Shopee's unique identifier for a shop.
	ShopID int
	// ShopCategory's name.
	Name string
	// ShopCategory's sort weight.
	SortWeight int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type AddShopCategoryResponse struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int
	// The identifier for an API request for error tracking.
	RequestID string
}

type GetShopCategoriesRequest struct {
	// Shopee's unique identifier for a shop.
	ShopID int
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int
	// If many items are available to retrieve, you may need to call GetItemsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	PaginationEntriesPerPage int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type ShopCategory struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int
	// ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
	Status string
	// ShopCategory's name.
	Name string
	// ShopCategory's sort weight.
	SortWeight int
}

type GetShopCategoriesResponse struct {
	//
	ShopCategorys []ShopCategory
	// This is to indicate whether the shop categorys list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of shop categorys.
	More bool
	// The identifier for an API request for error tracking
	RequestID string
}

type DeleteShopCategoryRequest struct {
	// Shopee's unique identifier for a shop.
	ShopID int
	// ShopCategory's unique identifier.
	ShopCategoryID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type DeleteShopCategoryResponse struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int
	//
	Msg string
	// The identifier for an API request for error tracking
	RequestID string
}

type UpdateShopCategoryRequest struct {
	// Shopee's unique identifier for a shop.
	ShopID int
	// ShopCategory's unique identifier.
	ShopCategoryID int
	// ShopCategory's name.
	Name string
	// ShopCategory's sort weight.
	SortWeight int
	// ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
	Status string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type UpdateShopCategoryResponse struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int
	// ShopCategory's name.
	Name string
	// ShopCategory's sort weight.
	SortWeight int
	// ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
	Status string
	// The identifier for an API request for error tracking.
	RequestID string
}

type AddItemsRequest struct {
	// Shopee's unique identifier for a shop.
	ShopID int
	// ShopCategory's unique identifier.
	ShopCategoryID int
	// Shopee's unique identifiers list for an item.
	Items []int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type AddItemsResponse struct {
	// List of invalid item ids.
	InvalidItemID []int
	// ShopCategory's unique identifier.
	ShopCategoryID int
	// Number of items in the shop category.
	Count int
	// The identifier for an API request for error tracking.
	Request string
}

type GetItemsRequest struct {
	// Shopee's unique identifier for a shop.
	ShopID int
	// ShopCategory's unique identifier.
	ShopCategoryID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
}

type GetItemsResponse struct {
	//
	Items []Item
	// The identifier for an API request for error tracking.
	RequestID string
}

type DeleteItemsRequest struct {
	// The id of the shop category
	ShopCategoryID int
	// The list of items need to be deleted
	Items []int
	// Shopee's unique identifier for a shop.
	ShopID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type DeleteItemsResponse struct {
	// The id of the shop category
	ShopCategoryID int
	// The list of item ids which are invalid
	InvalidItemID []int
	// count of items under this shop category after deletion
	Count int
	// The identifier of the API request for error tracking
	RequestID string
}
