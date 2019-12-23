package shopeego

type GetCategoriesRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// Indicate the translation language. Language values include: en(English), vi(Vietnamese), id(Indonesian), th(Thai), zh-Hant(Traditional Chinese), zh-Hans(Simplified Chinese), ms-my(Malaysian Malay). If the selected language is not supported in certain shop location, the response will be in default language.
	Language string `json:"language"`
}

type GetCategoriesResponse struct {
	// The category list.
	Categories []GetCategoriesResponseCategory `json:"categories"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id"`
}

type GetAttributesRequest struct {
	// The Identify of category. Should call shopee.item.GetCategories to get category_id first. Attributes can ONLY be fetched for the category_id WITHOUT children.
	CategoryID int `json:"category_id"`
	// Indicate the translation language. Language values include: en(English), vi(Vietnamese), id(Indonesian), th(Thai), zh-Hant(Traditional Chinese), zh-Hans(Simplified Chinese), ms-my(Malaysian Malay). If the selected language is not supported in certain shop location, the response will be in default language.
	Language string `json:"language"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// Two-characters country code(capital letter) for the attributes. Should be provided if no "shopid".
	Country string `json:"country"`
	// Is cross-border or not. Should be provided if no "shopid".
	IsCB bool `json:"is_cb"`
	// Shopee's unique identifier for a shop. Should be provided if no "country" and "is_cb".
	ShopID int `json:"shop_id"`
}

type GetAttributesResponse struct {
	// The attributes list.
	Attributes []GetAttributesResponseAttribute `json:"attributes"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id"`
}

type AddRequest struct {
	// Should call shopee.item.GetCategories to get category first.Related to result.categories.category_id
	CategoryID int `json:"category_id"`
	// Name of the item in local language.
	Name string `json:"name"`
	// Description of the item in local language. HTML is not supported.
	Description string `json:"description"`
	// The current price of the item in the listing currency. This value will be ignored if there is variation level price input.
	Price float64 `json:"price"`
	// The current stock quantity of the item. This value will be ignored if there is variation level stock input.
	Stock int `json:"stock"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku"`
	// Please skip this field and use the dedicated APIs to create 2-tier variation. More details: https://open.shopee.com/documents?module=63&type=2&id=54
	Variations []AddRequestVariation `json:"variations"`
	// Image URLs of the item. Up to 9 images(12 images for TW mall seller), max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []AddRequestImage `json:"images"`
	// This field is optional depending on the specific attribute under different categories. Should call shopee.item.GetAttributes to get attribute first. Must contain all all mandatory attribute.
	Attributes []AddRequestAttribute `json:"attributes"`
	// Should call shopee.logistics.GetLogistics to get logistics first. Should contain all all logistics.
	Logistics []AddRequestLogistic `json:"logistics"`
	// the net weight of this item, the unit is KG.
	Weight float64 `json:"weight"`
	// The length of package for this single item, the unit is CM
	PackageLength int `json:"package_length"`
	// The width of package for this single item, the unit is CM
	PackageWidth int `json:"package_width"`
	// The height of package for this single item, the unit is CM
	PackageHeight int `json:"package_height"`
	// The guaranteed days to ship orders. For pre-order, please input value from 7 to 30; for non pre-order, please exclude this field and it will default to the respective standard value per your shop location.(e.g. 3 for CrossBorder)
	DaysToShip int `json:"days_to_ship"`
	// The wholesales tier list. Please put the wholesale tier info in order by min count.
	Wholesales []AddRequestWholesale `json:"wholesales"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// Url of size chart image. Only particular categories support it. max size: 500KB. 2000*2000 pixels
	SizeChart string `json:"size_chart"`
	// This indicates whether the item is secondhand.
	Condition string `json:"condition"`
	// Use this field to indicate the initial status of the new item. Value can be one of 'NORMAL' or 'UNLIST'.
	Status string `json:"status"`
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool `json:"is_pre_order"`
}

type AddResponse struct {
	//
	ItemID int `json:"item_id"`
	// Item's info.
	Item AddResponseItem `json:"item"`
	//
	Warning string `json:"warning"`
	// Image URLs for fail download.
	FailImage []string `json:"fail_image"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
	// Url of size chart image. Only particular categories support it. max size: 500KB. 2000*2000 pixels
	SizeChart string `json:"size_chart"`
}

type DeleteRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type DeleteResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	//
	Msg string `json:"msg"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UnlistItemRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// List of item_ids and expected status. Up to 50 items for one call.
	Items []UnlistItemRequestItem `json:"items"`
}

type UnlistItemResponse struct {
	// List of item ids which failed to update status, including their reasons
	Failed []UnlistItemResponseFailed `json:"failed"`
	// List of item ids which succeed to update status, including their current status.
	Success []UnlistItemResponseSuccess `json:"success"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}

type AddVariationsRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// The variation of item is to list out all models of this product. For example, iPhone has model of White and Black, then its variations includes "White iPhone" and "Black iPhone".
	Variations []AddVariationsRequestVariation `json:"variations"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type AddVariationsResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when stock of the variation is updated.
	ModifiedTime int `json:"modified_time"`
	// The variation list of item.
	Variations []AddVariationsResponseVariation `json:"variations"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id"`
}

type DeleteVariationRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int `json:"variation_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type DeleteVariationResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// The time when stock of the variation is updated.
	ModifiedTime int `json:"modified_time"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetItemsListRequest struct {
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int `json:"pagination_offset"`
	// If many items are available to retrieve, you may need to call GetItemsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	PaginationEntriesPerPage int `json:"pagination_entries_per_page"`
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days.
	UpdateTimeFrom int `json:"update_time_from"`
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_to field is the ending date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days.
	UpdateTimeTo int `json:"update_time_to"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetItemsListResponse struct {
	//
	Items []GetItemsListResponseItem `json:"items"`
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool `json:"more"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetItemDetailRequest struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetItemDetailResponse struct {
	//
	Item GetItemDetailResponseItem `json:"item"`
	// Warning returned when the category or attributes are missing/invalid.
	Warning string `json:"warning"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateItemRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Should call shopee.item.GetItemDetail to get category first.Related to result.categories.category_id
	CategoryID int `json:"category_id"`
	// Name of the item in local language.
	Name string `json:"name"`
	// Description of the item in local language. HTML is not supported.
	Description string `json:"description"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku"`
	// The variation of item is to list out all models of this product, for example, iPhone has model of White and Black, then its variations includes "White iPhone" and "Black iPhone".
	Variations []UpdateItemRequestVariation `json:"variations"`
	// Should call shopee.item.GetAttributes to get attribute first. Should contain all all mandatory attribute if change the category.
	Attributes []UpdateItemRequestAttribute `json:"attributes"`
	// The guaranteed days to ship orders. Update value to less than 7 will default the value to the respective standard per your shop location and make this item non pre-order.(e.g. 3 for CrossBorder)
	DaysToShip int `json:"days_to_ship"`
	// The wholesales tier list. If the item has already had wholesale info, the wholesale info will be replaced. Please put the wholesale tier info in order by min count.
	Wholesales []UpdateItemRequestWholesale `json:"wholesales"`
	// Should call shopee.logistics.GetLogistics to get logistics first. Should contain all all logistics.
	Logistics []UpdateItemRequestLogistic `json:"logistics"`
	// the net weight of this item, the unit is KG.
	Weight float64 `json:"weight"`
	// The length of package for this single item, the unit is CM
	PackageLength int `json:"package_length"`
	// The width of package for this single item, the unit is CM
	PackageWidth int `json:"package_width"`
	// The height of package for this single item, the unit is CM
	PackageHeight int `json:"package_height"`
	// This indicates whether the item is secondhand.
	Condition string `json:"condition"`
	// Url of size chart image. Only particular categories support it. max size: 500KB. 2000*2000 pixels
	SizeChart string `json:"size_chart"`
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool `json:"is_pre_order"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type UpdateItemResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	//
	Item UpdateItemResponseItem `json:"item"`
	//
	Warning string `json:"warning"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type AddItemImgRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Image URLs of the item. It contains at most 9 URLs. Could get the url by api GetItemDetail
	Images []string `json:"images"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type AddItemImgResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Image URLs for fail download.
	FailImage []string `json:"fail_image"`
	// Image URLs of item.
	Images []string `json:"images"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateItemImgRequest struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Image URLs of the item. Up to 9 images(12 images for TW mall seller), max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []string `json:"images"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type UpdateItemImgResponse struct {
	// Image URLs of the item. Up to 9 images, max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []string `json:"images"`
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type InsertItemImgRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Image URL of the item.
	ImageURL string `json:"image_url"`
	// The position that insert the image. It starts with 1 and the max number is 9. If the position is bigger than existing position, the image would be placed on the last position.
	ImagePosition int `json:"image_position"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type InsertItemImgResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when stock of the variation is updated.
	ModifiedTime int `json:"modified_time"`
	// Image URLs of item.
	Images []string `json:"images"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type DeleteItemImgRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Image URLs of the item. It contains at most 9 URLs.Could get the url by api GetItemDetail
	Images []string `json:"images"`
	// Image positions of the item. Positions start with 1 and the max number is 9. If the position can not match the old image position, this position would be ignored It contains at most 9 positions. Param position and param images can not been empty at the same time.If there are images and positions at the same time, positions is ignored.
	Positions []int `json:"positions"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type DeleteItemImgResponse struct {
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdatePriceRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// Specify the revised price of the item. This value will be ignored if there is variation level price input.
	Price float64 `json:"price"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type UpdatePriceResponse struct {
	//
	Item UpdatePriceResponseItem `json:"item"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateStockRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// Specify the updated stock quantity. This value will be ignored if there is variation level stock input.
	Stock int `json:"stock"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type UpdateStockResponse struct {
	//
	Item UpdateStockResponseItem `json:"item"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateVariationPriceRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// Specify the revised price of one variation of the item.
	Price float64 `json:"price"`
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int `json:"variation_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type UpdateVariationPriceResponse struct {
	//
	Item UpdateVariationPriceResponseItem `json:"item"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateVariationStockRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// Specify the updated stock quantity.
	Stock int `json:"stock"`
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int `json:"variation_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type UpdateVariationStockResponse struct {
	//
	Item UpdateVariationStockResponseItem `json:"item"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdatePriceBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// List of items to update price. Up to 50 items in one call.
	Items []UpdatePriceBatchRequestItem `json:"items"`
}

type UpdatePriceBatchResponse struct {
	// Result of batch updating.
	BatchResult []UpdatePriceBatchResponseBatchResult `json:"batch_result"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateStockBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// List of items to update stock. Up to 50 items in one call.
	Items []UpdateStockBatchRequestItem `json:"items"`
}

type UpdateStockBatchResponse struct {
	// Result of batch updating.
	BatchResult []UpdateStockBatchResponseBatchResult `json:"batch_result"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateVariationPriceBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// List of variations to update price. Up to 50 variations in one call.
	Variations []UpdateVariationPriceBatchRequestVariation `json:"variations"`
}

type UpdateVariationPriceBatchResponse struct {
	// Result of batch updating.
	BatchResult []UpdateVariationPriceBatchResponseBatchResult `json:"batch_result"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateVariationStockBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// List of variations to update price. Up to 50 variations in one call.
	Variations []UpdateVariationStockBatchRequestVariation `json:"variations"`
}

type UpdateVariationStockBatchResponse struct {
	// Result of batch updating.
	BatchResult []UpdateVariationStockBatchResponseBatchResult `json:"batch_result"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type InitTierVariationRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Tier_variation list. Up to 2 tiers.
	TierVariation []InitTierVariationRequestTierVariation `json:"tier_variation"`
	// 2-Tier variation list.
	Variation []InitTierVariationRequestVariation `json:"variation"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type InitTierVariationResponse struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
	// Current variation ids under this item
	VariationIDList []InitTierVariationResponseVariation `json:"variation_id_list"`
}

type AddTierVariationRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// 2-Tier variation list.
	Variation []AddTierVariationRequestVariation `json:"variation"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type AddTierVariationResponse struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
	// Current variations information under this item.
	VariationIDList []AddTierVariationResponseVariation `json:"variation_id_list"`
}

type GetVariationRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetVariationResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Tier_variation list.
	TierVariation []GetVariationResponseTierVariation `json:"tier_variation"`
	// Item's variation list.
	Variations []GetVariationResponseVariation `json:"variations"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateTierVariationListRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Tier_variation list. Length must be 1 or 2.
	TierVariation []UpdateTierVariationListRequestTierVariation `json:"tier_variation"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type UpdateTierVariationListResponse struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateTierVariationIndexRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// 2-Tier variation list.
	Variation []UpdateTierVariationIndexRequestVariation `json:"variation"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type UpdateTierVariationIndexResponse struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}

type BoostItemRequest struct {
	// A list of item ids to be boosted. You can input a maximum of 5 items per request.
	ItemID []int `json:"item_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type BoostItemResponse struct {
	// batch result details
	BatchResult BoostItemResponseBatchResult `json:"batch_result"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}

type GetBoostedItemRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetBoostedItemResponse struct {
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
	//
	Items []GetBoostedItemResponseItem `json:"items"`
}

type SetItemInstallmentTenuresRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// List of installments, applicable values: 3, 6, 12, 24. If the list is empty, it means you wanna close the installment.
	Tenures []int `json:"tenures"`
}

type SetItemInstallmentTenuresResponse struct {
	// List of installments
	Tenures []int `json:"tenures"`
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}

type GetPromotionInfoRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetPromotionInfoResponse struct {
	// The set of item's promotion list.
	Items []GetPromotionInfoResponseItem `json:"items"`
	// The identifier of the API request for error tracking.
	RequestID string `json:"request_id"`
}

type GetRecommendCatsRequest struct {
	// The title of a particular item, used to get recommended category ids.
	Name string `json:"name"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetRecommendCatsResponse struct {
	// List of recommended category ids.
	CategoryIDs []string `json:"category_i_ds"`
	// The identifier of the API request for error tracking.
	RequestID string `json:"request_id"`
}
