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

type DaysToShipLimits struct {
	// The maximum of dts，-1 means no dts.
	MaxLimit int `json:"max_limit"`
	// The minimum of dts, -1 means no dts.
	MinLimit int `json:"min_limit"`
}

type Category struct {
	// The Identify of each category.
	CategoryID int `json:"category_id"`
	// The Identify of the parent of the category.
	ParentID int `json:"parent_id"`
	// The name of each category.
	CategoryName string `json:"category_name"`
	// This is to indicate whether the category has children. Attributes can ONLY be fetched for the category_id WITHOUT children.
	HasChildren bool `json:"has_children"`
	// The limits of pre-order items' days_to_ship based on per category.
	DaysToShipLimits DaysToShipLimits `json:"days_to_ship_limits"`
}

type GetCategoriesResponse struct {
	// The category list.
	Categories []Category `json:"categories"`
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

type Value struct {
	// Value in original language. It's MANDATORY to use attributes in original_value to add items.
	OriginalValue string `json:"original_value"`
	// Value in translated language. As referrence only, CANNOT be used to add item. If the selected language is not supported in certain shop location, this field will be empty.
	TranslateValue string `json:"translate_value"`
}

type Attribute struct {
	// The Identify of each category.
	AttributeID int `json:"attribute_id"`
	// The name of each attribute.
	AttributeName int `json:"attribute_name"`
	// This is to indicate whether this attribute is mandantory.
	IsMandatory bool `json:"is_mandatory"`
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttribuiteType string `json:"attribuite_type"`
	// Enumerated type that defines the input type of the attribute. Applicable values: See Data Definition- AttributeInputType.
	InputType string `json:"input_type"`
	// All options that attribute contains.
	Options []string `json:"options"`
	// The option values in different language.
	Values Value `json:"values"`
}

type GetAttributesResponse struct {
	// The attributes list.
	Attributes []Attribute `json:"attributes"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id"`
}

type Variation struct {
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price. Max Length: 20 letters
	Name string `json:"name"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock"`
	// The current price of the variation in the listing currency.
	Price float64 `json:"price"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSku string `json:"variation_sku"`
}

type Image struct {
	// Url of items' image.The system would synchronous download the image one by one.if one of those image can not fetch, would get a warning in result.But can continue the AddItem proccessing.
	// if all image failed to fetch, would return an error.
	URL string `json:"url"`
}

type Attribute struct {
	// related to shopee.item.GetAttributes result.attributes.attribute_id
	AttributesID int `json:"attributes_id"`
	// related to shopee.item.GetAttributes one of result.attributes.options. Max length is 40 letters.
	Value string `json:"value"`
}

type Logistics struct {
	// related to shopee.logistics.GetLogistics result.logistics.logistic_id
	LogisticID int `json:"logistic_id"`
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool `json:"enabled"`
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64 `json:"shipping_fee"`
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID int `json:"size_id"`
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool `json:"is_free"`
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
	Variations []Variation `json:"variations"`
	// Image URLs of the item. Up to 9 images(12 images for TW mall seller), max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []Image `json:"images"`
	// This field is optional depending on the specific attribute under different categories. Should call shopee.item.GetAttributes to get attribute first. Must contain all all mandatory attribute.
	Attributes []Attribute `json:"attributes"`
	// Should call shopee.logistics.GetLogistics to get logistics first. Should contain all all logistics.
	Logistics []Logistics `json:"logistics"`
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
	WholeSales []WholeSale `json:"whole_sales"`
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

type WholeSale struct {
	// The min count of this tier wholesale. If the wholesale is not the first one, the min count must equal to max count of last tier plus one.
	Min int `json:"min"`
	// The max count of this tier wholesale.
	Max int `json:"max"`
	// The current price of the wholesale in the listing currency. The price must be cheaper than original price. And if the wholesale is not the first one, the price must be cheaper than previous tier.
	UnitPrice float64 `json:"unit_price"`
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name"`
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock"`
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status string `json:"status"`
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int `json:"create_time"`
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int `json:"update_time"`
	// The original price of the variation in the listing currency.
	OriginalPrice float `json:"original_price"`
	// The ID of discount activity the variation is currently in. One variation can only have one discount at a time. discount_id will be 0 if the variation has no discount applied.
	DiscountID int `json:"discount_id"`
}

type Attribute struct {
	// The Identify of each category.
	AttributeID int `json:"attribute_id"`
	// The name of each attribute.
	AttributeName int `json:"attribute_name"`
	// This is to indicate whether this attribute is mandantory.
	IsMandatory bool `json:"is_mandatory"`
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttributeType string `json:"attribute_type"`
	// The value of this item attribute.
	AtributeValue string `json:"atribute_value"`
}

type Logistics struct {
	// The identity of logistic channel.
	LogisticID int `json:"logistic_id"`
	// The name of logistic.
	LogisticName string `json:"logistic_name"`
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool `json:"enabled"`
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee floa64 `json:"shipping_fee"`
	// If specify logistic fee_type is SIZE_SELECTION size_id is required.
	SizeID int `json:"size_id"`
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool `json:"is_free"`
	// Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"`
}

type Wholesale struct {
	// The min count of this tier wholesale.
	Min int `json:"min"`
	// The max count of this tier wholesale.
	Max int `json:"max"`
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64 `json:"unit_price"`
}

type Item struct {
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku"`
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED, UNLIST and BANNED.
	Status string `json:"status"`
	// Name of the item in local language.
	Name string `json:"name"`
	// Description of the item in local language.
	Description string `json:"description"`
	// Image URLs of the item. It contains at most 9 URLs.
	Images []string `json:"images"`
	// The three-digit code representing the currency unit used for the item in Shopee Listings.
	Currency string `json:"currency"`
	// This is to indicate whether the item has variation(s).
	HasVariation bool `json:"has_variation"`
	// The current price of the item in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price"`
	// The current stock quantity of the item.
	Stock int `json:"stock"`
	// Timestamp that indicates the date and time that the item was created.
	CreateTime int `json:"create_time"`
	// Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
	UpdateTime int `json:"update_time"`
	// the net weight of this item, the unit is KG.
	Weight float64 `json:"weight"`
	// Could call shopee.item.GetCategories to get category detail.Related to result.categories.category_id.
	CategoryID int `json:"category_id"`
	// The original price of the item in the listing currency.
	OriginalPrice float64 `json:"original_price"`
	// The variation list of item.
	Variations []Variation `json:"variations"`
	//
	Attributes []Attributes `json:"attributes"`
	// The logistics list.
	Logistics []Logistic `json:"logistics"`
	// The wholesales tier list.
	Wholesales []Wholesale `json:"wholesales"`
	// The sales volume of item.
	Sales int `json:"sales"`
	// The page view of item.
	Views int `json:"views"`
	// The conllection number of item.
	Likes int `json:"likes"`
	// The length of package for this single item, the unit is CM
	PackageLength int `json:"package_length"`
	// The width of package for this single item, the unit is CM
	PackageWidth int `json:"package_width"`
	// The height of package for this single item, the unit is CM
	PackageHeight int `json:"package_height"`
	// The guaranteed days to ship orders. For pre-order, please input value from 7 to 30; for non pre-order, please exclude this field and it will default to the respective standard per your shop location.(e.g. 3 for CrossBorder)
	DaysToShip int `json:"days_to_ship"`
	// The rating star scores of this item.
	RatingStar float64 `json:"rating_star"`
	// Count of comments for the item.
	CMTCount int `json:"cmt_count"`
	// This indicates whether the item is secondhand.
	Condition string `json:"condition"`
	// The ID of discount activity the item is currently in. One item can only have one discount at a time. discount_id will be 0 if the item has no discount applied, or item has variation.
	DiscountID int `json:"discount_id"`
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool `json:"is_pre_order"`
}

type AddResponse struct {
	//
	ItemID int `json:"item_id"`
	// Item's info.
	Item Item `json:"item"`
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

type Item struct {
	// Item's unique identifier.
	ItemID int `json:"item_id"`
	// True: unlist this item; False: list this item.
	Unlist bool `json:"unlist"`
}

type UnlistItemRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// List of item_ids and expected status. Up to 50 items for one call.
	Items []Item `json:"items"`
}

type Failed struct {
	// Item's unique identifier.
	ItemID int `json:"item_id"`
	// Error message.
	ErrorDesciption string `json:"error_desciption"`
}

type Success struct {
	// Item's unique identifier.
	ItemID int `json:"item_id"`
	// True: item is unlisted; False: item is listed.
	Unlist bool `json:"unlist"`
}

type UnlistItemResponse struct {
	// List of item ids which failed to update status, including their reasons
	Failed []Failed `json:"failed"`
	// List of item ids which succeed to update status, including their current status.
	Success []Success `json:"success"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}

type Variation struct {
	// Name of the variation that belongs to the same item.A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock"`
	// The current price of the variation in the listing currency.
	Price float64 `json:"price"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
}

type AddVariationsRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// The variation of item is to list out all models of this product. For example, iPhone has model of White and Black, then its variations includes "White iPhone" and "Black iPhone".
	Variations []Variation `json:"variations"`
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
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name"`
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock"`
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status string `json:"status"`
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int `json:"create_time"`
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int `json:"update_time"`
	// The original price of the variation in the listing currency.
	OriginalPrice float64 `json:"original_price"`
}

type AddVariationsResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when stock of the variation is updated.
	ModifiedTime int `json:"modified_time"`
	// The variation list of item.
	Variations []Variation `json:"variations"`
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

type Variation struct {
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// The latest update time of the item.
	UpdateTime int `json:"update_time"`
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, BANNED and UNLIST.
	Status string `json:"status"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku"`
	// The variation list of item
	Variations []Variation `json:"variations"`
	// Whether 2-tier variation structure is activated for this item
	Is2TierItem bool `json:"is_2_tier_item"`
	// Only for TW seller. List of installments
	Tenures []int `json:"tenures"`
}

type GetItemsListResponse struct {
	//
	Items []Item `json:"items"`
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool `json:"more"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetItemDetailRequest struct {
	//
	ItemID
	//
	PartnerID
	//
	ShopID
	//
	Timestamp
}

type Variaion struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name"`
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price"`
	// The current stock quantity of the variation in the listing currency.
	Stock intt `json:"stock"`
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status string `json:"status"`
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int `json:"create_time"`
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int `json:"update_time"`
	// The original price of the variation in the listing currency.
	OriginalPrice floatt64 `json:"original_price"`
	// The ID of discount activity the variation is currently in. One variation can only have one discount at a time. discount_id will be 0 if the variation has no discount applied.
	DiscountID int `json:"discount_id"`
}

type Attribute struct {
	// The Identify of each category
	AttributeID int `json:"attribute_id"`
	// The name of each attribute
	AttributeName int `json:"attribute_name"`
	// This is to indicate whether this attribute is mandantory
	IsMandatory bool `json:"is_mandatory"`
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttributeType string `json:"attribute_type"`
	// The value of this item attribute.
	AttributeValue string `json:"attribute_value"`
}

type Logistic struct {
	// The identity of logistic channel
	LogisticID int `json:"logistic_id"`
	// The name of logistic
	LogisticName string `json:"logistic_name"`
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool `json:"enabled"`
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64 `json:"shipping_fee"`
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID int `json:"size_id"`
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool `json:"is_free"`
	// Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"`
}

type Wholesale struct {
	// The min count of this tier wholesale.
	Min int `json:"min"`
	// The max count of this tier wholesale.
	Max int `json:"max"`
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64 `json:"unit_price"`
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku"`
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED, BANNED and UNLIST.
	Status string `json:"status"`
	// Name of the item in local language.
	Name string `json:"name"`
	// Description of the item in local language.
	Description string `json:"description"`
	// Image URLs of the item. It contains at most 9 URLs.
	Images []string `json:"images"`
	// The three-digit code representing the currency unit used for the item in Shopee Listings.
	Currency string `json:"currency"`
	// This is to indicate whether the item has variation(s).
	HasVariaion bool `json:"has_variaion"`
	// The current price of the item in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price"`
	// The current stock quantity of the item.
	Stock int `json:"stock"`
	// Timestamp that indicates the date and time that the item was created.
	CreateTime int `json:"create_time"`
	// Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
	UpdateTime int `json:"update_time"`
	// the net weight of this item, the unit is KG.
	Weight float64 `json:"weight"`
	// Could call shopee.item.GetCategories to get category detail.Related to result.categories.category_id
	CategoryID int `json:"category_id"`
	// The original price of the item in the listing currency.
	OriginalPrice float64 `json:"original_price"`
	// The variation list of item
	Variations []Variation `json:"variations"`
	//
	Attributes []Attribute `json:"attributes"`
	// The logistics list.
	Logistics []Logistic `json:"logistics"`
	// The wholesales tier list.
	Wholesales []Wholesale `json:"wholesales"`
	// The rating star scores of this item.
	RatingStar float64 `json:"rating_star"`
	// Count of comments for the item.
	CMTCount int `json:"cmt_count"`
	// The sales volume of item.
	Sales int `json:"sales"`
	// The page view of item.
	Views int `json:"views"`
	// The conllection number of item.
	Likes int `json:"likes"`
	// The length of package for this single item, the unit is CM
	PackageLength float64 `json:"package_length"`
	// The width of package for this single item, the unit is CM
	PackageWidth float64 `json:"package_width"`
	// The height of package for this single item, the unit is CM
	PackageHeight float64 `json:"package_height"`
	// The days to ship.
	DaysToShip int `json:"days_to_ship"`
	// url of size chart image. Only particular categories support it.
	SizeChart string `json:"size_chart"`
	// This indicates whether the item is secondhand.
	Condition string `json:"condition"`
	// The ID of discount activity the item is currently in. One item can only have one discount at a time. discount_id will be 0 if the item has no discount applied, or item has variation.
	DiscountID int `json:"discount_id"`
	// Whether 2-tier variation structure is activated for this item
	Is2TierItem bool `json:"is_2_tier_item"`
	// Only for TW seller. List of installments
	Tenures []int `json:"tenures"`
	// Use this field to get the locked stock of item by promotions.
	ReservedStock int `json:"reserved_stock"`
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool `json:"is_pre_order"`
}

type GetItemDetailResponse struct {
	//
	Item Item `json:"item"`
	// Warning returned when the category or attributes are missing/invalid.
	Warning string `json:"warning"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU
}

type Attribute struct {
	// related to shopee.item.GetAttributes result.attributes.attribute_id
	AttributesID
	// related to shopee.item.GetAttributes one of result.attributes.options
	Value
}

type Wholesale struct {
	// The min count of this tier wholesale. If the wholesale is not the first one, the min count must equal to max count of last tier plus one.
	Min
	// The max count of this tier wholesale.
	Max
	// The current price of the wholesale in the listing currency. The price must be cheaper than original price. And if the wholesale is not the first one, the price must be cheaper than previous tier.'
	UnitPrice
}

type Logistic struct {
	// related to shopee.logistics.GetLogistics result.logistics.logistic_id
	LogisticID
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree
}

type UpdateItemRequest struct {
	// The identity of product item.
	ItemID
	// Should call shopee.item.GetItemDetail to get category first.Related to result.categories.category_id
	CategoryID
	// Name of the item in local language.
	Name
	// Description of the item in local language. HTML is not supported.
	Description
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU
	// The variation of item is to list out all models of this product, for example, iPhone has model of White and Black, then its variations includes "White iPhone" and "Black iPhone".
	Variations
	// Should call shopee.item.GetAttributes to get attribute first. Should contain all all mandatory attribute if change the category.
	Attributes
	// The guaranteed days to ship orders. Update value to less than 7 will default the value to the respective standard per your shop location and make this item non pre-order.(e.g. 3 for CrossBorder)
	DaysToShip
	// The wholesales tier list. If the item has already had wholesale info, the wholesale info will be replaced. Please put the wholesale tier info in order by min count.
	Wholesales
	// Should call shopee.logistics.GetLogistics to get logistics first. Should contain all all logistics.
	Logistics
	// the net weight of this item, the unit is KG.
	Weight
	// The length of package for this single item, the unit is CM
	PackageLength
	// The width of package for this single item, the unit is CM
	PackageWidth
	// The height of package for this single item, the unit is CM
	PackageHeight
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp
	// This indicates whether the item is secondhand.
	Condition
	// Url of size chart image. Only particular categories support it. max size: 500KB. 2000*2000 pixels
	SizeChart
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name"`
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock"`
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status sring `json:"status"`
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int `json:"create_time"`
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int `json:"update_time"`
	// The original price of the variation in the listing currency.
	OriginalPirce float64 `json:"original_pirce"`
	// The ID of discount activity the variation is currently in. One variation can only have one discount at a time. discount_id will be 0 if the variation has no discount applied.
	DiscountID int `json:"discount_id"`
}

type Attribute struct {
	// The Identify of each category
	AttributeID int `json:"attribute_id"`
	// The name of each attribute
	AttributeName int `json:"attribute_name"`
	// This is to indicate whether this attribute is mandantory
	IsMandatory bool `json:"is_mandatory"`
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttributeType string `json:"attribute_type"`
	// The value of this item attribute.
	AttribueValue string `json:"attribue_value"`
}

type Logistic struct {
	// The identity of logistic channel
	LogisticID int `json:"logistic_id"`
	// The name of logistic
	LogisticName string `json:"logistic_name"`
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool `json:"enabled"`
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64 `json:"shipping_fee"`
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID int `json:"size_id"`
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool `json:"is_free"`
	// Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"`
}

type Wholesale struct {
	// The min count of this tier wholesale.
	Min int `json:"min"`
	// The max count of this tier wholesale.
	Max int `json:"max"`
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64 `json:"unit_price"`
}

type Item struct {
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku"`
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED and BANNED.
	Status string `json:"status"`
	// Name of the item in local language.
	Name string `json:"name"`
	// Description of the item in local language.
	Description string `json:"description"`
	// Image URLs of the item. It contains at most 9 URLs.
	Images []string `json:"images"`
	// The three-digit code representing the currency unit used for the item in Shopee Listings.
	Currency string `json:"currency"`
	// This is to indicate whether the item has variation(s).
	HasVariation bool `json:"has_variation"`
	// The current price of the item in the listing currency. If item is in promotion, this value is discount price.
	Price float64 `json:"price"`
	// The current stock quantity of the item.
	Stock int `json:"stock"`
	// Timestamp that indicates the date and time that the item was created.
	CreateTime in `json:"create_time"`
	// Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
	UpdateTime in `json:"update_time"`
	// the net weight of this item, the unit is KG.
	Weight float64 `json:"weight"`
	// Could call shopee.item.GetCategories to get category detail.Related to result.categories.category_id
	CategoryID int `json:"category_id"`
	// The original price of the item in the listing currency.
	OriginalPrice float64 `json:"original_price"`
	// The variation list of item
	Variations []Variation `json:"variations"`
	//
	Attritube []Attribute `json:"attritube"`
	// The logistics list.
	Logistics []Logistic `json:"logistics"`
	// The wholesales tier list.
	Wholesales []Wholesale `json:"wholesales"`
	// The rating star scores of this item.
	RatingStar float64 `json:"rating_star"`
	// Count of comments for the item.
	CMTCount int `json:"cmt_count"`
	// The sales volume of item.
	Sales int `json:"sales"`
	// The page view of item.
	Views int `json:"views"`
	// The conllection number of item.
	Likes int `json:"likes"`
	// The length of package for this single item, the unit is CM
	PackageLength int `json:"package_length"`
	// The width of package for this single item, the unit is CM
	PackageWidth int `json:"package_width"`
	// The height of package for this single item, the unit is CM
	PackageHeight int `json:"package_height"`
	// The guaranteed days to ship orders. Update value to less than 7 will default the value to the respective standard per your shop location and make this item non pre-order.(e.g. 3 for CrossBorder)
	DaysToShip int `json:"days_to_ship"`
	// This indicates whether the item is secondhand.
	Condittion sttring `json:"condittion"`
	// The ID of discount activity the item is currently in. One item can only have one discount at a time. discount_id will be 0 if the item has no discount applied, or item has variation.
	DiscountID int `json:"discount_id"`
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool `json:"is_pre_order"`
}

type UpdateItemResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	//
	Item Item `json:"item"`
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
	PartnerID intt `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type AddItemImgResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Image URLs for fail download.
	FailImage string[] `json:"fail_image"`
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
	// Partner ID is assigned upon registration is successful.
	PartnerID intt `json:"partner_id"`
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request.
	Timestamp int `json:"timestamp"`
}

type UpdateItemImgResponse struct {
	// Image URLs of the item. Up to 9 images, max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []string `json:"images"`
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// This is to indicate the timestamp of the request.
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

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time"`
	// Specify the revised price of the item.
	Price float64 `json:"price"`
}

type UpdatePriceResponse struct {
	//
	Item Item `json:"item"`
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

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time"`
	// Specify the updated stock quantity.
	Stock int `json:"stock"`
}

type UpdateStockResponse struct {
	//
	Item Item `json:"item"`
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

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// Specify the revised price of one variation of the item.
	Price float64 `json:"price"`
}

type UpdateVariationPriceResponse struct {
	//
	Item Item `json:"item"`
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

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time"`
	// Specify the updated stock quantity.
	Stock int `json:"stock"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type UpdateVariationStockResponse struct {
	//
	Item Item `json:"item"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type Item struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// New price value for this item.
	Price int `json:"price"`
}

type UpdatePriceBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// List of items to update price. Up to 50 items in one call.
	Items []Item `json:"items"`
}

type Failure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int `json:"item_id"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description"`
}

type BatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []string `json:"modifications"`
	// Informations for failed stock updating.
	Failures []Failure `json:"failures"`
}

type UpdatePriceBatchResponse struct {
	// Result of batch updating.
	BatchResult []BatchResult `json:"batch_result"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type Item struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// New stock value for this item.
	Stock int `json:"stock"`
}

type UpdateStockBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// List of items to update stock. Up to 50 items in one call.
	Items []Item `json:"items"`
}

type Failure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int `json:"item_id"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description"`
}

type BatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []string `json:"modifications"`
	// Informations for failed stock updating.
	Failures []Failure `json:"failures"`
}

type UpdateStockBatchResponse struct {
	// Result of batch updating.
	BatchResult []BatchResult `json:"batch_result"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int `json:"variation_id"`
	// New price value of this variation.
	Price int `json:"price"`
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
}

type UpdateVariationPriceBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// List of variations to update price. Up to 50 variations in one call.
	Variations []Variation `json:"variations"`
}

type Failure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int `json:"item_id"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type Modification struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type BatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []Modification `json:"modifications"`
	// Informations for failed stock updating.
	Failures []Failure `json:"failures"`
}

type UpdateVariationPriceBatchResponse struct {
	// Result of batch updating.
	BatchResult []BatchResult `json:"batch_result"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int `json:"variation_id"`
	// New stock value of this variation.
	Stock int `json:"stock"`
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
}

type UpdateVariationStockBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// List of variations to update price. Up to 50 variations in one call.
	Variations []Variation `json:"variations"`
}

type Failure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int `json:"item_id"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type Modification struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type BatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []Modification `json:"modifications"`
	// Informations for failed stock updating.
	Failures []Failure `json:"failures"`
}

type UpdateVariationStockBatchResponse struct {
	// Result of batch updating.
	BatchResult []BatchResult `json:"batch_result"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type TierVariation struct {
	// Tier variation name.
	Name string `json:"name"`
	// Tier variation value options list. Option length should be under 20. Quantity of combinations of all 2 tier options is up to 50.
	Options []string `json:"options"`
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string `json:"images_url"`
}

type Variation struct {
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex []int `json:"tier_index"`
	// Stock value of this variation item. The original variation stock will be override when calling this API to initialize 2-tier structure for an existed item. 0 stock will make this variation a greyout option for buyer.
	Stock int `json:"stock"`
	// Price value of this variation item. The original variation price will be override when calling this API to initialize 2-tier structure for an existed item.
	Price float64 `json:"price"`
	// SKU string of this variation.SKU length should be under 100.
	VariationSKU string `json:"variation_sku"`
}

type InitTierVariationRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Tier_variation list. Up to 2 tiers.
	TierVariation []TierVariation `json:"tier_variation"`
	// 2-Tier variation list.
	Variation []Variation `json:"variation"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
}

type Variation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int `json:"tier_index"`
	// The identity of the variation.
	VariationID int `json:"variation_id"`
}

type InitTierVariationResponse struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
	// Current variation ids under this item
	VariationIDList []Variation `json:"variation_id_list"`
}

type Variation struct {
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex []int `json:"tier_index"`
	// Stock value of this variation item. 0 stock will make this variation a greyout option for buyer.
	Stock int `json:"stock"`
	// Price value of this variation item.
	Price float64 `json:"price"`
	// SKU string of this variation item.
	VariationSKU string `json:"variation_sku"`
}

type AddTierVariationRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// 2-Tier variation list.
	Variation []Variation `json:"variation"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
}

type Variation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int `json:"tier_index"`
	// The identity of the variation.
	VariationID int `json:"variation_id"`
}

type AddTierVariationResponse struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
	// Current variations information under this item.
	VariationIDList []Variation `json:"variation_id_list"`
}

type GetVariationRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type TierVariation struct {
	// Tier variation name.
	Name string `json:"name"`
	// Tier variation value options list.
	Options []string `json:"options"`
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string `json:"images_url"`
}

type Variation struct {
	// Unique identifier of the variation.
	VariationID
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex
}

type GetVariationResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Tier_variation list.
	TierVariation []TierVariation `json:"tier_variation"`
	// Item's variation list.
	Variations []Variation `json:"variations"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}

type TierVariation struct {
	// Tier variation name.
	Name string `json:"name"`
	// Tier variation value options list. Lenght should be under 20. Combinations of 2 level options should be under 50.
	Options []string `json:"options"`
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string `json:"images_url"`
}

type UpdateTierVariationListRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// Tier_variation list. Length must be 1 or 2.
	TierVariation []TierVariation `json:"tier_variation"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	PartnerID int `json:"partner_id"`
}

type UpdateTierVariationListResponse struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}

type Variation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int `json:"tier_index"`
	// The identity of product item variation.
	VariationID []int `json:"variation_id"`
}

type UpdateTierVariationIndexRequest struct {
	// The identity of product item.
	ItemID int `json:"item_id"`
	// 2-Tier variation list.
	Variation []Variation `json:"variation"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	PartnerID int `json:"partner_id"`
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
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type Failure struct {
	// to indicate error type.
	ErrorCode string `json:"error_code"`
	// Failed item id.
	ID int `json:"id"`
	// error description
	Description string `json:"description"`
}

type BatchResult struct {
	// A list of item ids which have been boosted successfully.
	Successes []int `json:"successes"`
	// A list of failed-to-boost items, including error details.
	Failures []Failure `json:"failures"`
}

type BoostItemResponse struct {
	// batch result details
	BatchResult BatchResult `json:"batch_result"`
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
}

type GetBoostedItemRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type Item struct {
	// boosted items' id.
	ItemID int `json:"item_id"`
	// Cooldown_second time is four hours after boost. After four hours you can boost this item again.
	CooldownSecond int `json:"cooldown_second"`
}

type GetBoostedItemResponse struct {
	// The identifier of the API request for error tracking
	RequestID string `json:"request_id"`
	//
	Items []Item `json:"items"`
}

type SetItemInstallmentTenuresRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// List of installments, applicable values: 3, 6, 12, 24. If the list is empty, it means you wanna close the installment.
	Tenures  []int `json:"tenures"`
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
	// The list of item_id. Up to 50 item_ids in one call.
	ItemIDList []int `json:"item_id_list"`
}

type Promotion struct {
	//
	PromotionType string `json:"promotion_type"`
	// The ID of promotion.
	PromotionID int `json:"promotion_id"`
	// ID of the variation that belongs to the same item.
	VariationID int `json:"variation_id"`
	// Start timestamp of promotion.
	StartTime int `json:"start_time"`
	// End timestamp of promotion.
	EndTime int `json:"end_time"`
	// The promotion price of item.
	PromotionPrice float64 `json:"promotion_price"`
	// The Locked stock of item by promotion.
	ReservedStock int `json:"reserved_stock"`
	// The sold out timestamp of promotion stock.
	StockoutTime int `json:"stockout_time"`
	// The stage at which the promotion goes. Available values: ongoing/upcoming.
	Staging string `json:"staging"`
}

type Error struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// Error Message.
	ErrorMsg string `json:"error_msg"`
}

type Item struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// Promotion information list.
	Promotions []Promotion `json:"promotions"`
	// The list of error items.
	Errors []Error `json:"errors"`
}

type GetPromotionInfoResponse struct {
	// The set of item's promotion list.
	Items []Item `json:"items"`
	// The identifier of the API request for error tracking.
	RequestID string `json:"request_id"`
}

type GetRecommendCatsRequest struct {
	// The title of a particular item, used to get recommended category ids.
	Name string `json:"name"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetRecommendCatsResponse struct {
	// List of recommended category ids.
	CategoryIDs []string `json:"category_i_ds"`
	// The identifier of the API request for error tracking.
	RequestID string `json:"request_id"`
}
