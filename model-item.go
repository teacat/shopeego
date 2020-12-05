package shopeego

type GetCategoriesRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Indicate the translation language. Language values include: en(English), vi(Vietnamese), id(Indonesian), th(Thai), zh-Hant(Traditional Chinese), zh-Hans(Simplified Chinese), ms-my(Malaysian Malay). If the selected language is not supported in certain shop location, the response will be in default language.
	Language string `json:"language,omitempty"`
}

type GetCategoriesResponse struct {
	// The category list.
	Categories []GetCategoriesResponseCategory `json:"categories,omitempty"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
}

type GetAttributesRequest struct {
	// The Identify of category. Should call shopee.item.GetCategories to get category_id first. Attributes can ONLY be fetched for the category_id WITHOUT children.
	CategoryID int64 `json:"category_id,omitempty"`
	// Indicate the translation language. Language values include: en(English), vi(Vietnamese), id(Indonesian), th(Thai), zh-Hant(Traditional Chinese), zh-Hans(Simplified Chinese), ms-my(Malaysian Malay). If the selected language is not supported in certain shop location, the response will be in default language.
	Language string `json:"language,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Two-characters country code(capital letter) for the attributes. Should be provided if no "shopid".
	Country string `json:"country,omitempty"`
	// Is cross-border or not. Should be provided if no "shopid".
	IsCB bool `json:"is_cb,omitempty"`
	// Shopee's unique identifier for a shop. Should be provided if no "country" and "is_cb".
	ShopID int64 `json:"shopid,omitempty"`
}

type GetAttributesResponse struct {
	// The attributes list.
	Attributes []GetAttributesResponseAttribute `json:"attributes,omitempty"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
}

type AddRequest struct {
	// Should call shopee.item.GetCategories to get category first.Related to result.categories.category_id
	CategoryID int64 `json:"category_id,omitempty"`
	// Name of the item in local language.
	Name string `json:"name,omitempty"`
	// Description of the item in local language. HTML is not supported.
	Description string `json:"description,omitempty"`
	// The current price of the item in the listing currency. This value will be ignored if there is variation level price input.
	Price float64 `json:"price,omitempty,string"`
	// The current stock quantity of the item. This value will be ignored if there is variation level stock input.
	Stock int `json:"stock,omitempty"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku,omitempty"`
	// Please skip this field and use the dedicated APIs to create 2-tier variation. More details: https://open.shopee.com/documents?module=63&type=2&id=54
	Variations []AddRequestVariation `json:"variations,omitempty"`
	// Image URLs of the item. Up to 9 images(12 images for TW mall seller), max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []AddRequestImage `json:"images,omitempty"`
	// This field is optional depending on the specific attribute under different categories. Should call shopee.item.GetAttributes to get attribute first. Must contain all all mandatory attribute.
	Attributes []AddRequestAttribute `json:"attributes,omitempty"`
	// Should call shopee.logistics.GetLogistics to get logistics first. Should contain all all logistics.
	Logistics []AddRequestLogistic `json:"logistics,omitempty"`
	// the net weight of this item, the unit is KG.
	Weight float64 `json:"weight,omitempty,string"`
	// The length of package for this single item, the unit is CM
	PackageLength int `json:"package_length,omitempty"`
	// The width of package for this single item, the unit is CM
	PackageWidth int `json:"package_width,omitempty"`
	// The height of package for this single item, the unit is CM
	PackageHeight int `json:"package_height,omitempty"`
	// The guaranteed days to ship orders. For pre-order, please input value from 7 to 30; for non pre-order, please exclude this field and it will default to the respective standard value per your shop location.(e.g. 3 for CrossBorder)
	DaysToShip int `json:"days_to_ship,omitempty"`
	// The wholesales tier list. Please put the wholesale tier info in order by min count.
	Wholesales []AddRequestWholesale `json:"wholesales,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Url of size chart image. Only particular categories support it. max size: 500KB. 2000*2000 pixels
	SizeChart string `json:"size_chart,omitempty"`
	// This indicates whether the item is secondhand.
	Condition string `json:"condition,omitempty"`
	// Use this field to indicate the initial status of the new item. Value can be one of 'NORMAL' or 'UNLIST'.
	Status string `json:"status,omitempty"`
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool `json:"is_pre_order,omitempty"`
}

type AddResponse struct {
	//
	ItemID int64 `json:"item_id,omitempty"`
	// Item's info.
	Item AddResponseItem `json:"item,omitempty"`
	//
	Warning string `json:"warning,omitempty"`
	// Image URLs for fail download.
	FailImage []string `json:"fail_image,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
	// Url of size chart image. Only particular categories support it. max size: 500KB. 2000*2000 pixels
	SizeChart string `json:"size_chart,omitempty"`
}

type DeleteRequest struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type DeleteResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	//
	Msg string `json:"msg,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UnlistItemRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// List of item_ids and expected status. Up to 50 items for one call.
	Items []UnlistItemRequestItem `json:"items,omitempty"`
}

type UnlistItemResponse struct {
	// List of item ids which failed to update status, including their reasons
	Failed []UnlistItemResponseFailed `json:"failed,omitempty"`
	// List of item ids which succeed to update status, including their current status.
	Success []UnlistItemResponseSuccess `json:"success,omitempty"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type AddVariationsRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
	// The variation of item is to list out all models of this product. For example, iPhone has model of White and Black, then its variations includes "White iPhone" and "Black iPhone".
	Variations []AddVariationsRequestVariation `json:"variations,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type AddVariationsResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// The time when stock of the variation is updated.
	ModifiedTime int `json:"modified_time,omitempty"`
	// The variation list of item.
	Variations []AddVariationsResponseVariation `json:"variations,omitempty"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
}

type DeleteVariationRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int64 `json:"variation_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type DeleteVariationResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// The time when stock of the variation is updated.
	ModifiedTime int `json:"modified_time,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetItemsListRequest struct {
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset"`
	// If many items are available to retrieve, you may need to call GetItemsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page"`
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days.
	UpdateTimeFrom int `json:"update_time_from,omitempty"`
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_to field is the ending date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days.
	UpdateTimeTo int `json:"update_time_to,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// true:return item in all status; false:return items without item deleted by seller; default: false.
	NeedDeletedItem bool `json:"need_deleted_item,omitempty"`
}

type GetItemsListResponse struct {
	//
	Items []GetItemsListResponseItem `json:"items,omitempty"`
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool `json:"more,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
	// The total count of items.
	Total int64 `json:"total,omitempty"`
}

type GetItemDetailRequest struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetItemDetailResponse struct {
	//
	Item GetItemDetailResponseItem `json:"item,omitempty"`
	// Warning returned when the category or attributes are missing/invalid.
	Warning string `json:"warning,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateItemRequest struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// Should call shopee.item.GetItemDetail to get category first.Related to result.categories.category_id
	CategoryID int64 `json:"category_id,omitempty"`
	// Name of the item in local language.
	Name string `json:"name,omitempty"`
	// Description of the item in local language. HTML is not supported.
	Description string `json:"description,omitempty"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku,omitempty"`
	// The variation of item is to list out all models of this product, for example, iPhone has model of White and Black, then its variations includes "White iPhone" and "Black iPhone".
	Variations []UpdateItemRequestVariation `json:"variations,omitempty"`
	// Should call shopee.item.GetAttributes to get attribute first. Should contain all all mandatory attribute if change the category.
	Attributes []UpdateItemRequestAttribute `json:"attributes,omitempty"`
	// The guaranteed days to ship orders. Update value to less than 7 will default the value to the respective standard per your shop location and make this item non pre-order.(e.g. 3 for CrossBorder)
	DaysToShip int `json:"days_to_ship,omitempty"`
	// The wholesales tier list. If the item has already had wholesale info, the wholesale info will be replaced. Please put the wholesale tier info in order by min count.
	Wholesales []UpdateItemRequestWholesale `json:"wholesales,omitempty"`
	// Should call shopee.logistics.GetLogistics to get logistics first. Should contain all all logistics.
	Logistics []UpdateItemRequestLogistic `json:"logistics,omitempty"`
	// the net weight of this item, the unit is KG.
	Weight float64 `json:"weight,omitempty,string"`
	// The length of package for this single item, the unit is CM
	PackageLength int `json:"package_length,omitempty"`
	// The width of package for this single item, the unit is CM
	PackageWidth int `json:"package_width,omitempty"`
	// The height of package for this single item, the unit is CM
	PackageHeight int `json:"package_height,omitempty"`
	// This indicates whether the item is secondhand.
	Condition string `json:"condition,omitempty"`
	// Url of size chart image. Only particular categories support it. max size: 500KB. 2000*2000 pixels
	SizeChart string `json:"size_chart,omitempty"`
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool `json:"is_pre_order,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdateItemResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	//
	Item UpdateItemResponseItem `json:"item,omitempty"`
	//
	Warning string `json:"warning,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type AddItemImgRequest struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// Image URLs of the item. It contains at most 9 URLs. Could get the url by api GetItemDetail
	Images []string `json:"images,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type AddItemImgResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Image URLs for fail download.
	FailImage []string `json:"fail_image,omitempty"`
	// Image URLs of item.
	Images []string `json:"images,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateItemImgRequest struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Image URLs of the item. Up to 9 images(12 images for TW mall seller), max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []string `json:"images,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdateItemImgResponse struct {
	// Image URLs of the item. Up to 9 images, max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []string `json:"images,omitempty"`
	// Shopee's unique identifier for a shop.
	ShopID int64 `json:"shopid,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type InsertItemImgRequest struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// Image URL of the item.
	ImageURL string `json:"image_url,omitempty"`
	// The position that insert the image. It starts with 1 and the max number is 9. If the position is bigger than existing position, the image would be placed on the last position.
	ImagePosition int `json:"image_position,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type InsertItemImgResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// The time when stock of the variation is updated.
	ModifiedTime int `json:"modified_time,omitempty"`
	// Image URLs of item.
	Images []string `json:"images,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type DeleteItemImgRequest struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// Image URLs of the item. It contains at most 9 URLs.Could get the url by api GetItemDetail
	Images []string `json:"images,omitempty"`
	// Image positions of the item. Positions start with 1 and the max number is 9. If the position can not match the old image position, this position would be ignored It contains at most 9 positions. Param position and param images can not been empty at the same time.If there are images and positions at the same time, positions is ignored.
	Positions []int `json:"positions,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type DeleteItemImgResponse struct {
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdatePriceRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
	// Specify the revised price of the item. This value will be ignored if there is variation level price input.
	Price float64 `json:"price,omitempty,string"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdatePriceResponse struct {
	//
	Item UpdatePriceResponseItem `json:"item,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateStockRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
	// Specify the updated stock quantity. This value will be ignored if there is variation level stock input.
	Stock int `json:"stock,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdateStockResponse struct {
	//
	Item UpdateStockResponseItem `json:"item,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateVariationPriceRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
	// Specify the revised price of one variation of the item.
	Price float64 `json:"price,omitempty,string"`
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int64 `json:"variation_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdateVariationPriceResponse struct {
	//
	Item UpdateVariationPriceResponseItem `json:"item,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateVariationStockRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
	// Specify the updated stock quantity.
	Stock int `json:"stock,omitempty"`
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int64 `json:"variation_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdateVariationStockResponse struct {
	//
	Item UpdateVariationStockResponseItem `json:"item,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdatePriceBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// List of items to update price. Up to 50 items in one call.
	Items []UpdatePriceBatchRequestItem `json:"items,omitempty"`
}

type UpdatePriceBatchResponse struct {
	// Result of batch updating.
	BatchResult []UpdatePriceBatchResponseBatchResult `json:"batch_result,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateStockBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// List of items to update stock. Up to 50 items in one call.
	Items []UpdateStockBatchRequestItem `json:"items,omitempty"`
}

type UpdateStockBatchResponse struct {
	// Result of batch updating.
	BatchResult []UpdateStockBatchResponseBatchResult `json:"batch_result,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateVariationPriceBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// List of variations to update price. Up to 50 variations in one call.
	Variations []UpdateVariationPriceBatchRequestVariation `json:"variations,omitempty"`
}

type UpdateVariationPriceBatchResponse struct {
	// Result of batch updating.
	BatchResult []UpdateVariationPriceBatchResponseBatchResult `json:"batch_result,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateVariationStockBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// List of variations to update price. Up to 50 variations in one call.
	Variations []UpdateVariationStockBatchRequestVariation `json:"variations,omitempty"`
}

type UpdateVariationStockBatchResponse struct {
	// Result of batch updating.
	BatchResult []UpdateVariationStockBatchResponseBatchResult `json:"batch_result,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type InitTierVariationRequest struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// Tier_variation list. Up to 2 tiers.
	TierVariation []InitTierVariationRequestTierVariation `json:"tier_variation,omitempty"`
	// 2-Tier variation list.
	Variation []InitTierVariationRequestVariation `json:"variation,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type InitTierVariationResponse struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
	// Current variation ids under this item
	VariationIDList []InitTierVariationResponseVariation `json:"variation_id_list,omitempty"`
}

type AddTierVariationRequest struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// 2-Tier variation list.
	Variation []AddTierVariationRequestVariation `json:"variation,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type AddTierVariationResponse struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
	// Current variations information under this item.
	VariationIDList []AddTierVariationResponseVariation `json:"variation_id_list,omitempty"`
}

type GetVariationRequest struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetVariationResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Tier_variation list.
	TierVariation []GetVariationResponseTierVariation `json:"tier_variation,omitempty"`
	// Item's variation list.
	Variations []GetVariationResponseVariation `json:"variations,omitempty"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateTierVariationListRequest struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// Tier_variation list. Length must be 1 or 2.
	TierVariation []UpdateTierVariationListRequestTierVariation `json:"tier_variation,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdateTierVariationListResponse struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateTierVariationIndexRequest struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// 2-Tier variation list.
	Variation []UpdateTierVariationIndexRequestVariation `json:"variation,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdateTierVariationIndexResponse struct {
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type BoostItemRequest struct {
	// A list of item ids to be boosted. You can input a maximum of 5 items per request.
	ItemID []int `json:"item_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type BoostItemResponse struct {
	// batch result details
	BatchResult BoostItemResponseBatchResult `json:"batch_result,omitempty"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetBoostedItemRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetBoostedItemResponse struct {
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
	//
	Items []GetBoostedItemResponseItem `json:"items,omitempty"`
}

type SetItemInstallmentTenuresRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// List of installments, applicable values: 3, 6, 12, 24. If the list is empty, it means you wanna close the installment.
	Tenures []int `json:"tenures,omitempty"`
}

type SetItemInstallmentTenuresResponse struct {
	// List of installments
	Tenures []int `json:"tenures,omitempty"`
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetPromotionInfoRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// The list of item_id. Up to 50 item_ids in one call.
	ItemIDList []int64 `json:"item_id_list,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetPromotionInfoResponse struct {
	// The set of item's promotion list.
	Items []GetPromotionInfoResponseItem `json:"items,omitempty"`
	// The identifier of the API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
}

type GetRecommendCatsRequest struct {
	// The title of a particular item, used to get recommended category ids.
	Name string `json:"name,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetRecommendCatsResponse struct {
	// List of recommended category ids.
	CategoryIDs []string `json:"category_i_ds,omitempty"`
	// The identifier of the API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
}

type GetCommentRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// The identity of product item.
	ItemID int64 `json:"item_id,omitempty"`
	// The identity of comment.
	CMTID int64 `json:"cmt_id,omitempty"`
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset,omitempty"`
	// If many items are available to retrieve, you may need to call GetItemsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page,omitempty"`
}

type GetCommentResponse struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate whether the comment list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of comments.
	More bool `json:"more,omitempty"`
	//
	ItemCMTList []GetCommentResponseItemCMTList `json:"item_cmt_list,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type ReplyCommentsRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	//
	CMTList ReplyCommentsRequestCMTList `json:"cmt_list,omitempty"`
}

type ReplyCommentsResponse struct {
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
	//
	SuccList []ReplyCommentsResponseSuccList `json:"succ_list,omitempty"`
	//
	Errors []ReplyCommentsResponseError `json:"errors,omitempty"`
}
