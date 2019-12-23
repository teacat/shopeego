package shopeego

type GetCategoriesRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Indicate the translation language. Language values include: en(English), vi(Vietnamese), id(Indonesian), th(Thai), zh-Hant(Traditional Chinese), zh-Hans(Simplified Chinese), ms-my(Malaysian Malay). If the selected language is not supported in certain shop location, the response will be in default language.
	Language string
}

type DaysToShipLimits struct {
	// The maximum of dts，-1 means no dts.
	MaxLimit int
	// The minimum of dts, -1 means no dts.
	MinLimit int
}

type Category struct {
	// The Identify of each category.
	CategoryID int
	// The Identify of the parent of the category.
	ParentID int
	// The name of each category.
	CategoryName string
	// This is to indicate whether the category has children. Attributes can ONLY be fetched for the category_id WITHOUT children.
	HasChildren bool
	// The limits of pre-order items' days_to_ship based on per category.
	DaysToShipLimits DaysToShipLimits
}

type GetCategoriesResponse struct {
	// The category list.
	Categories []Category
	// The identifier for an API request for error tracking.
	RequestID string
}

type GetAttributesRequest struct {
	// The Identify of category. Should call shopee.item.GetCategories to get category_id first. Attributes can ONLY be fetched for the category_id WITHOUT children.
	CategoryID int
	// Indicate the translation language. Language values include: en(English), vi(Vietnamese), id(Indonesian), th(Thai), zh-Hant(Traditional Chinese), zh-Hans(Simplified Chinese), ms-my(Malaysian Malay). If the selected language is not supported in certain shop location, the response will be in default language.
	Language string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Two-characters country code(capital letter) for the attributes. Should be provided if no "shopid".
	Country string
	// Is cross-border or not. Should be provided if no "shopid".
	IsCB bool
	// Shopee's unique identifier for a shop. Should be provided if no "country" and "is_cb".
	ShopID int
}

type Value struct {
	// Value in original language. It's MANDATORY to use attributes in original_value to add items.
	OriginalValue string
	// Value in translated language. As referrence only, CANNOT be used to add item. If the selected language is not supported in certain shop location, this field will be empty.
	TranslateValue string
}

type Attribute struct {
	// The Identify of each category.
	AttributeID int
	// The name of each attribute.
	AttributeName int
	// This is to indicate whether this attribute is mandantory.
	IsMandatory bool
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttribuiteType string
	// Enumerated type that defines the input type of the attribute. Applicable values: See Data Definition- AttributeInputType.
	InputType string
	// All options that attribute contains.
	Options []string
	// The option values in different language.
	Values Value
}

type GetAttributesResponse struct {
	// The attributes list.
	Attributes []Attribute
	// The identifier for an API request for error tracking.
	RequestID string
}

type Variation struct {
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price. Max Length: 20 letters
	Name string
	// The current stock quantity of the variation in the listing currency.
	Stock int
	// The current price of the variation in the listing currency.
	Price float64
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSku string
}

type Image struct {
	// Url of items' image.The system would synchronous download the image one by one.if one of those image can not fetch, would get a warning in result.But can continue the AddItem proccessing.
	// if all image failed to fetch, would return an error.
	URL string
}

type Attribute struct {
	// related to shopee.item.GetAttributes result.attributes.attribute_id
	AttributesID int
	// related to shopee.item.GetAttributes one of result.attributes.options. Max length is 40 letters.
	Value string
}

type Logistics struct {
	// related to shopee.logistics.GetLogistics result.logistics.logistic_id
	LogisticID int
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID int
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool
}

type AddRequest struct {
	// Should call shopee.item.GetCategories to get category first.Related to result.categories.category_id
	CategoryID int
	// Name of the item in local language.
	Name string
	// Description of the item in local language. HTML is not supported.
	Description string
	// The current price of the item in the listing currency. This value will be ignored if there is variation level price input.
	Price float64
	// The current stock quantity of the item. This value will be ignored if there is variation level stock input.
	Stock int
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string
	// Please skip this field and use the dedicated APIs to create 2-tier variation. More details: https://open.shopee.com/documents?module=63&type=2&id=54
	Variations []Variation
	// Image URLs of the item. Up to 9 images(12 images for TW mall seller), max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []Image
	// This field is optional depending on the specific attribute under different categories. Should call shopee.item.GetAttributes to get attribute first. Must contain all all mandatory attribute.
	Attributes []Attribute
	// Should call shopee.logistics.GetLogistics to get logistics first. Should contain all all logistics.
	Logistics []Logistics
	// the net weight of this item, the unit is KG.
	Weight float64
	// The length of package for this single item, the unit is CM
	PackageLength int
	// The width of package for this single item, the unit is CM
	PackageWidth int
	// The height of package for this single item, the unit is CM
	PackageHeight int
	// The guaranteed days to ship orders. For pre-order, please input value from 7 to 30; for non pre-order, please exclude this field and it will default to the respective standard value per your shop location.(e.g. 3 for CrossBorder)
	DaysToShip int
	// The wholesales tier list. Please put the wholesale tier info in order by min count.
	WholeSales []WholeSale
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Url of size chart image. Only particular categories support it. max size: 500KB. 2000*2000 pixels
	SizeChart string
	// This indicates whether the item is secondhand.
	Condition string
	// Use this field to indicate the initial status of the new item. Value can be one of 'NORMAL' or 'UNLIST'.
	Status string
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool
}

type WholeSale struct {
	// The min count of this tier wholesale. If the wholesale is not the first one, the min count must equal to max count of last tier plus one.
	Min int
	// The max count of this tier wholesale.
	Max int
	// The current price of the wholesale in the listing currency. The price must be cheaper than original price. And if the wholesale is not the first one, the price must be cheaper than previous tier.
	UnitPrice float64
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64
	// The current stock quantity of the variation in the listing currency.
	Stock int
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status string
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int
	// The original price of the variation in the listing currency.
	OriginalPrice float
	// The ID of discount activity the variation is currently in. One variation can only have one discount at a time. discount_id will be 0 if the variation has no discount applied.
	DiscountID int
}

type Attribute struct {
	// The Identify of each category.
	AttributeID int
	// The name of each attribute.
	AttributeName int
	// This is to indicate whether this attribute is mandantory.
	IsMandatory bool
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttributeType string
	// The value of this item attribute.
	AtributeValue string
}

type Logistics struct {
	// The identity of logistic channel.
	LogisticID int
	// The name of logistic.
	LogisticName string
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee floa64
	// If specify logistic fee_type is SIZE_SELECTION size_id is required.
	SizeID int
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool
	// Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
	EstimatedShippingFee float64
}

type Wholesale struct {
	// The min count of this tier wholesale.
	Min int
	// The max count of this tier wholesale.
	Max int
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64
}

type Item struct {
	// Shopee's unique identifier for a shop.
	ShopID int
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED, UNLIST and BANNED.
	Status string
	// Name of the item in local language.
	Name string
	// Description of the item in local language.
	Description string
	// Image URLs of the item. It contains at most 9 URLs.
	Images []string
	// The three-digit code representing the currency unit used for the item in Shopee Listings.
	Currency string
	// This is to indicate whether the item has variation(s).
	HasVariation bool
	// The current price of the item in the listing currency.If item is in promotion, this value is discount price.
	Price float64
	// The current stock quantity of the item.
	Stock int
	// Timestamp that indicates the date and time that the item was created.
	CreateTime int
	// Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
	UpdateTime int
	// the net weight of this item, the unit is KG.
	Weight float64
	// Could call shopee.item.GetCategories to get category detail.Related to result.categories.category_id.
	CategoryID int
	// The original price of the item in the listing currency.
	OriginalPrice float64
	// The variation list of item.
	Variations []Variation
	//
	Attributes []Attributes
	// The logistics list.
	Logistics []Logistic
	// The wholesales tier list.
	Wholesales []Wholesale
	// The sales volume of item.
	Sales int
	// The page view of item.
	Views int
	// The conllection number of item.
	Likes int
	// The length of package for this single item, the unit is CM
	PackageLength int
	// The width of package for this single item, the unit is CM
	PackageWidth int
	// The height of package for this single item, the unit is CM
	PackageHeight int
	// The guaranteed days to ship orders. For pre-order, please input value from 7 to 30; for non pre-order, please exclude this field and it will default to the respective standard per your shop location.(e.g. 3 for CrossBorder)
	DaysToShip int
	// The rating star scores of this item.
	RatingStar float64
	// Count of comments for the item.
	CMTCount int
	// This indicates whether the item is secondhand.
	Condition string
	// The ID of discount activity the item is currently in. One item can only have one discount at a time. discount_id will be 0 if the item has no discount applied, or item has variation.
	DiscountID int
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool
}

type AddResponse struct {
	//
	ItemID int
	// Item's info.
	Item Item
	//
	Warning string
	// Image URLs for fail download.
	FailImage []string
	// The identifier for an API request for error tracking
	RequestID string
	// Url of size chart image. Only particular categories support it. max size: 500KB. 2000*2000 pixels
	SizeChart string
}

type DeleteRequest struct {
	// The identity of product item.
	ItemID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type DeleteResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int
	//
	Msg string
	// The identifier for an API request for error tracking
	RequestID string
}

type Item struct {
	// Item's unique identifier.
	ItemID int
	// True: unlist this item; False: list this item.
	Unlist bool
}

type UnlistItemRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// List of item_ids and expected status. Up to 50 items for one call.
	Items []Item
}

type Failed struct {
	// Item's unique identifier.
	ItemID int
	// Error message.
	ErrorDesciption string
}

type Success struct {
	// Item's unique identifier.
	ItemID int
	// True: item is unlisted; False: item is listed.
	Unlist bool
}

type UnlistItemResponse struct {
	// List of item ids which failed to update status, including their reasons
	Failed []Failed
	// List of item ids which succeed to update status, including their current status.
	Success []Success
	// The identifier of the API request for error tracking
	RequestID string
}

type Variation struct {
	// Name of the variation that belongs to the same item.A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string
	// The current stock quantity of the variation in the listing currency.
	Stock int
	// The current price of the variation in the listing currency.
	Price float64
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string
}

type AddVariationsRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
	// The variation of item is to list out all models of this product. For example, iPhone has model of White and Black, then its variations includes "White iPhone" and "Black iPhone".
	Variations []Variation
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
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64
	// The current stock quantity of the variation in the listing currency.
	Stock int
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status string
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int
	// The original price of the variation in the listing currency.
	OriginalPrice float64
}

type AddVariationsResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// The time when stock of the variation is updated.
	ModifiedTime int
	// The variation list of item.
	Variations []Variation
	// The identifier for an API request for error tracking.
	RequestID string
}

type DeleteVariationRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type DeleteVariationResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// Shopee's unique identifier for a variation of an item.
	VariationID int
	// The time when stock of the variation is updated.
	ModifiedTime int
	// The identifier for an API request for error tracking
	RequestID string
}

type GetItemsListRequest struct {
	// Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PaginationOffset int
	// If many items are available to retrieve, you may need to call GetItemsList multiple times to retrieve all the data. Each result set is returned as a page of entries. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	PaginationEntriesPerPage int
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days.
	UpdateTimeFrom int
	// The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the item update time). The update_time_to field is the ending date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days.
	UpdateTimeTo int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Variation struct {
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string
	// Shopee's unique identifier for a variation of an item.
	VariationID int
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// Shopee's unique identifier for a shop.
	ShopID int
	// The latest update time of the item.
	UpdateTime int
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, BANNED and UNLIST.
	Status string
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string
	// The variation list of item
	Variations []Variation
	// Whether 2-tier variation structure is activated for this item
	Is2TierItem bool
	// Only for TW seller. List of installments
	Tenures []int
}

type GetItemsListResponse struct {
	//
	Items []Item
	// This is to indicate whether the item list is more than one page. If this value is true, you may want to continue to check next page to retrieve the rest of items.
	More bool
	// The identifier for an API request for error tracking
	RequestID string
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
	VariationID int
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64
	// The current stock quantity of the variation in the listing currency.
	Stock intt
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status string
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int
	// The original price of the variation in the listing currency.
	OriginalPrice floatt64
	// The ID of discount activity the variation is currently in. One variation can only have one discount at a time. discount_id will be 0 if the variation has no discount applied.
	DiscountID int
}

type Attribute struct {
	// The Identify of each category
	AttributeID int
	// The name of each attribute
	AttributeName int
	// This is to indicate whether this attribute is mandantory
	IsMandatory bool
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttributeType string
	// The value of this item attribute.
	AttributeValue string
}

type Logistic struct {
	// The identity of logistic channel
	LogisticID int
	// The name of logistic
	LogisticName string
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID int
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool
	// Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
	EstimatedShippingFee float64
}

type Wholesale struct {
	// The min count of this tier wholesale.
	Min int
	// The max count of this tier wholesale.
	Max int
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// Shopee's unique identifier for a shop.
	ShopID int
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED, BANNED and UNLIST.
	Status string
	// Name of the item in local language.
	Name string
	// Description of the item in local language.
	Description string
	// Image URLs of the item. It contains at most 9 URLs.
	Images []string
	// The three-digit code representing the currency unit used for the item in Shopee Listings.
	Currency string
	// This is to indicate whether the item has variation(s).
	HasVariaion bool
	// The current price of the item in the listing currency.If item is in promotion, this value is discount price.
	Price float64
	// The current stock quantity of the item.
	Stock int
	// Timestamp that indicates the date and time that the item was created.
	CreateTime int
	// Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
	UpdateTime int
	// the net weight of this item, the unit is KG.
	Weight float64
	// Could call shopee.item.GetCategories to get category detail.Related to result.categories.category_id
	CategoryID int
	// The original price of the item in the listing currency.
	OriginalPrice float64
	// The variation list of item
	Variations []Variation
	//
	Attributes []Attribute
	// The logistics list.
	Logistics []Logistic
	// The wholesales tier list.
	Wholesales []Wholesale
	// The rating star scores of this item.
	RatingStar float64
	// Count of comments for the item.
	CMTCount int
	// The sales volume of item.
	Sales int
	// The page view of item.
	Views int
	// The conllection number of item.
	Likes int
	// The length of package for this single item, the unit is CM
	PackageLength float64
	// The width of package for this single item, the unit is CM
	PackageWidth float64
	// The height of package for this single item, the unit is CM
	PackageHeight float64
	// The days to ship.
	DaysToShip int
	// url of size chart image. Only particular categories support it.
	SizeChart string
	// This indicates whether the item is secondhand.
	Condition string
	// The ID of discount activity the item is currently in. One item can only have one discount at a time. discount_id will be 0 if the item has no discount applied, or item has variation.
	DiscountID int
	// Whether 2-tier variation structure is activated for this item
	Is2TierItem bool
	// Only for TW seller. List of installments
	Tenures []int
	// Use this field to get the locked stock of item by promotions.
	ReservedStock int
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool
}

type GetItemDetailResponse struct {
	//
	Item Item
	// Warning returned when the category or attributes are missing/invalid.
	Warning string
	// The identifier for an API request for error tracking
	RequestID string
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
	VariationID int
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64
	// The current stock quantity of the variation in the listing currency.
	Stock int
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status sring
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int
	// The original price of the variation in the listing currency.
	OriginalPirce float64
	// The ID of discount activity the variation is currently in. One variation can only have one discount at a time. discount_id will be 0 if the variation has no discount applied.
	DiscountID int
}

type Attribute struct {
	// The Identify of each category
	AttributeID int
	// The name of each attribute
	AttributeName int
	// This is to indicate whether this attribute is mandantory
	IsMandatory bool
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttributeType string
	// The value of this item attribute.
	AttribueValue string
}

type Logistic struct {
	// The identity of logistic channel
	LogisticID int
	// The name of logistic
	LogisticName string
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID int
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool
	// Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
	EstimatedShippingFee float64
}

type Wholesale struct {
	// The min count of this tier wholesale.
	Min int
	// The max count of this tier wholesale.
	Max int
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64
}

type Item struct {
	// Shopee's unique identifier for a shop.
	ShopID int
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED and BANNED.
	Status string
	// Name of the item in local language.
	Name string
	// Description of the item in local language.
	Description string
	// Image URLs of the item. It contains at most 9 URLs.
	Images []string
	// The three-digit code representing the currency unit used for the item in Shopee Listings.
	Currency string
	// This is to indicate whether the item has variation(s).
	HasVariation bool
	// The current price of the item in the listing currency. If item is in promotion, this value is discount price.
	Price float64
	// The current stock quantity of the item.
	Stock int
	// Timestamp that indicates the date and time that the item was created.
	CreateTime in
	// Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
	UpdateTime in
	// the net weight of this item, the unit is KG.
	Weight float64
	// Could call shopee.item.GetCategories to get category detail.Related to result.categories.category_id
	CategoryID int
	// The original price of the item in the listing currency.
	OriginalPrice float64
	// The variation list of item
	Variations []Variation
	//
	Attritube []Attribute
	// The logistics list.
	Logistics []Logistic
	// The wholesales tier list.
	Wholesales []Wholesale
	// The rating star scores of this item.
	RatingStar float64
	// Count of comments for the item.
	CMTCount int
	// The sales volume of item.
	Sales int
	// The page view of item.
	Views int
	// The conllection number of item.
	Likes int
	// The length of package for this single item, the unit is CM
	PackageLength int
	// The width of package for this single item, the unit is CM
	PackageWidth int
	// The height of package for this single item, the unit is CM
	PackageHeight int
	// The guaranteed days to ship orders. Update value to less than 7 will default the value to the respective standard per your shop location and make this item non pre-order.(e.g. 3 for CrossBorder)
	DaysToShip int
	// This indicates whether the item is secondhand.
	Condittion sttring
	// The ID of discount activity the item is currently in. One item can only have one discount at a time. discount_id will be 0 if the item has no discount applied, or item has variation.
	DiscountID int
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool
}

type UpdateItemResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int
	//
	Item Item
	//
	Warning string
	// The identifier for an API request for error tracking
	RequestID string
}

type AddItemImgRequest struct {
	// The identity of product item.
	ItemID int
	// Image URLs of the item. It contains at most 9 URLs. Could get the url by api GetItemDetail
	Images []string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID intt
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type AddItemImgResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// Image URLs for fail download.
	FailImage string[]
	// Image URLs of item.
	Images []string
	// The identifier for an API request for error tracking
	RequestID string
}

type UpdateItemImgRequest struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// Image URLs of the item. Up to 9 images(12 images for TW mall seller), max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []string
	// Partner ID is assigned upon registration is successful.
	PartnerID intt
	// Shopee's unique identifier for a shop.
	ShopID int
	// This is to indicate the timestamp of the request.
	Timestamp int
}

type UpdateItemImgResponse struct {
	// Image URLs of the item. Up to 9 images, max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max size: 2MB
	Images []string
	// Shopee's unique identifier for a shop.
	ShopID int
	// Partner ID is assigned upon registration is successful.
	PartnerID int
	// Shopee's unique identifier for an item.
	ItemID int
	// This is to indicate the timestamp of the request.
	Timestamp int
}

type InsertItemImgRequest struct {
	// The identity of product item.
	ItemID int
	// Image URL of the item.
	ImageURL string
	// The position that insert the image. It starts with 1 and the max number is 9. If the position is bigger than existing position, the image would be placed on the last position.
	ImagePosition int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type InsertItemImgResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// The time when stock of the variation is updated.
	ModifiedTime int
	// Image URLs of item.
	Images []string
	// The identifier for an API request for error tracking
	RequestID string
}

type DeleteItemImgRequest struct {
	// The identity of product item.
	ItemID int
	// Image URLs of the item. It contains at most 9 URLs.Could get the url by api GetItemDetail
	Images []string
	// Image positions of the item. Positions start with 1 and the max number is 9. If the position can not match the old image position, this position would be ignored It contains at most 9 positions. Param position and param images can not been empty at the same time.If there are images and positions at the same time, positions is ignored.
	Positions []int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type DeleteItemImgResponse struct {
	// The identifier for an API request for error tracking
	RequestID string
}

type UpdatePriceRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
	// Specify the revised price of the item. This value will be ignored if there is variation level price input.
	Price float64
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// The time when price of the item is updated.
	ModifiedTime int
	// Specify the revised price of the item.
	Price float64
}

type UpdatePriceResponse struct {
	//
	Item Item
	// The identifier for an API request for error tracking
	RequestID string
}

type UpdateStockRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
	// Specify the updated stock quantity. This value will be ignored if there is variation level stock input.
	Stock int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// The time when price of the item is updated.
	ModifiedTime int
	// Specify the updated stock quantity.
	Stock int
}

type UpdateStockResponse struct {
	//
	Item Item
	// The identifier for an API request for error tracking
	RequestID string
}

type UpdateVariationPriceRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
	// Specify the revised price of one variation of the item.
	Price float64
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// The time when price of the item is updated.
	ModifiedTime int
	// Shopee's unique identifier for a variation of an item.
	VariationID int
	// Specify the revised price of one variation of the item.
	Price float64
}

type UpdateVariationPriceResponse struct {
	//
	Item Item
	// The identifier for an API request for error tracking
	RequestID string
}

type UpdateVariationStockRequest struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
	// Specify the updated stock quantity.
	Stock int
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Item struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// The time when price of the item is updated.
	ModifiedTime int
	// Specify the updated stock quantity.
	Stock int
	// Shopee's unique identifier for a variation of an item.
	VariationID int
}

type UpdateVariationStockResponse struct {
	//
	Item Item
	// The identifier for an API request for error tracking
	RequestID string
}

type Item struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
	// New price value for this item.
	Price int
}

type UpdatePriceBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// List of items to update price. Up to 50 items in one call.
	Items []Item
}

type Failure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int
	// Detailed information for the failed updating.
	ErrorDescription string
}

type BatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []string
	// Informations for failed stock updating.
	Failures []Failure
}

type UpdatePriceBatchResponse struct {
	// Result of batch updating.
	BatchResult []BatchResult
	// The identifier for an API request for error tracking
	RequestID string
}

type Item struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
	// New stock value for this item.
	Stock int
}

type UpdateStockBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// List of items to update stock. Up to 50 items in one call.
	Items []Item
}

type Failure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int
	// Detailed information for the failed updating.
	ErrorDescription string
}

type BatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []string
	// Informations for failed stock updating.
	Failures []Failure
}

type UpdateStockBatchResponse struct {
	// Result of batch updating.
	BatchResult []BatchResult
	// The identifier for an API request for error tracking
	RequestID string
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int
	// New price value of this variation.
	Price int
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
}

type UpdateVariationPriceBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// List of variations to update price. Up to 50 variations in one call.
	Variations []Variation
}

type Failure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int
	// Detailed information for the failed updating.
	ErrorDescription string
	// Shopee's unique identifier for a variation of an item.
	VariationID int
}

type Modification struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// Shopee's unique identifier for a variation of an item.
	VariationID int
}

type BatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []Modification
	// Informations for failed stock updating.
	Failures []Failure
}

type UpdateVariationPriceBatchResponse struct {
	// Result of batch updating.
	BatchResult []BatchResult
	// The identifier for an API request for error tracking
	RequestID string
}

type Variation struct {
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int
	// New stock value of this variation.
	Stock int
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
}

type UpdateVariationStockBatchRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// List of variations to update price. Up to 50 variations in one call.
	Variations []Variation
}

type Failure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int
	// Detailed information for the failed updating.
	ErrorDescription string
	// Shopee's unique identifier for a variation of an item.
	VariationID int
}

type Modification struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// Shopee's unique identifier for a variation of an item.
	VariationID int
}

type BatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []Modification
	// Informations for failed stock updating.
	Failures []Failure
}

type UpdateVariationStockBatchResponse struct {
	// Result of batch updating.
	BatchResult []BatchResult
	// The identifier for an API request for error tracking
	RequestID string
}

type TierVariation struct {
	// Tier variation name.
	Name string
	// Tier variation value options list. Option length should be under 20. Quantity of combinations of all 2 tier options is up to 50.
	Options []string
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string
}

type Variation struct {
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex []int
	// Stock value of this variation item. The original variation stock will be override when calling this API to initialize 2-tier structure for an existed item. 0 stock will make this variation a greyout option for buyer.
	Stock int
	// Price value of this variation item. The original variation price will be override when calling this API to initialize 2-tier structure for an existed item.
	Price float64
	// SKU string of this variation.SKU length should be under 100.
	VariationSKU string
}

type InitTierVariationRequest struct {
	// The identity of product item.
	ItemID int
	// Tier_variation list. Up to 2 tiers.
	TierVariation []TierVariation
	// 2-Tier variation list.
	Variation []Variation
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
}

type Variation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int
	// The identity of the variation.
	VariationID int
}

type InitTierVariationResponse struct {
	// The identity of product item.
	ItemID int
	// The identifier of the API request for error tracking
	RequestID string
	// Current variation ids under this item
	VariationIDList []Variation
}

type Variation struct {
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex []int
	// Stock value of this variation item. 0 stock will make this variation a greyout option for buyer.
	Stock int
	// Price value of this variation item.
	Price float64
	// SKU string of this variation item.
	VariationSKU string
}

type AddTierVariationRequest struct {
	// The identity of product item.
	ItemID int
	// 2-Tier variation list.
	Variation []Variation
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
}

type Variation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int
	// The identity of the variation.
	VariationID int
}

type AddTierVariationResponse struct {
	// The identity of product item.
	ItemID int
	// The identifier of the API request for error tracking
	RequestID string
	// Current variations information under this item.
	VariationIDList []Variation
}

type GetVariationRequest struct {
	// The identity of product item.
	ItemID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type TierVariation struct {
	// Tier variation name.
	Name string
	// Tier variation value options list.
	Options []string
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string
}

type Variation struct {
	// Unique identifier of the variation.
	VariationID
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex
}

type GetVariationResponse struct {
	// Shopee's unique identifier for an item.
	ItemID int
	// Tier_variation list.
	TierVariation []TierVariation
	// Item's variation list.
	Variations []Variation
	// The identifier of the API request for error tracking
	RequestID string
}

type TierVariation struct {
	// Tier variation name.
	Name string
	// Tier variation value options list. Lenght should be under 20. Combinations of 2 level options should be under 50.
	Options []string
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string
}

type UpdateTierVariationListRequest struct {
	// The identity of product item.
	ItemID int
	// Tier_variation list. Length must be 1 or 2.
	TierVariation []TierVariation
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	PartnerID int
}

type UpdateTierVariationListResponse struct {
	// The identity of product item.
	ItemID int
	// The identifier of the API request for error tracking
	RequestID string
}

type Variation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int
	// The identity of product item variation.
	VariationID []int
}

type UpdateTierVariationIndexRequest struct {
	// The identity of product item.
	ItemID int
	// 2-Tier variation list.
	Variation []Variation
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	PartnerID int
}

type UpdateTierVariationIndexResponse struct {
	// The identity of product item.
	ItemID int
	// The identifier of the API request for error tracking
	RequestID string
}

type BoostItemRequest struct {
	// A list of item ids to be boosted. You can input a maximum of 5 items per request.
	ItemID []int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Failure struct {
	// to indicate error type.
	ErrorCode string
	// Failed item id.
	ID int
	// error description
	Description string
}

type BatchResult struct {
	// A list of item ids which have been boosted successfully.
	Successes []int
	// A list of failed-to-boost items, including error details.
	Failures []Failure
}

type BoostItemResponse struct {
	// batch result details
	BatchResult BatchResult
	// The identifier of the API request for error tracking
	RequestID string
}

type GetBoostedItemRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Item struct {
	// boosted items' id.
	ItemID int
	// Cooldown_second time is four hours after boost. After four hours you can boost this item again.
	CooldownSecond int
}

type GetBoostedItemResponse struct {
	// The identifier of the API request for error tracking
	RequestID string
	//
	Items []Item
}

type SetItemInstallmentTenuresRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Shopee's unique identifier for an item.
	ItemID int
	// List of installments, applicable values: 3, 6, 12, 24. If the list is empty, it means you wanna close the installment.
	Tenures  []int
}

type SetItemInstallmentTenuresResponse struct {
	// List of installments
	Tenures []int
	// Shopee's unique identifier for an item.
	ItemID int
	// The identifier of the API request for error tracking
	RequestID string
}

type GetPromotionInfoRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// The list of item_id. Up to 50 item_ids in one call.
	ItemIDList []int
}

type Promotion struct {
	//
	PromotionType string
	// The ID of promotion.
	PromotionID int
	// ID of the variation that belongs to the same item.
	VariationID int
	// Start timestamp of promotion.
	StartTime int
	// End timestamp of promotion.
	EndTime int
	// The promotion price of item.
	PromotionPrice float64
	// The Locked stock of item by promotion.
	ReservedStock int
	// The sold out timestamp of promotion stock.
	StockoutTime int
	// The stage at which the promotion goes. Available values: ongoing/upcoming.
	Staging string
}

type Error struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
	// Error Message.
	ErrorMsg string
}

type Item struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int
	// Promotion information list.
	Promotions []Promotion
	// The list of error items.
	Errors []Error
}

type GetPromotionInfoResponse struct {
	// The set of item's promotion list.
	Items []Item
	// The identifier of the API request for error tracking.
	RequestID string
}

type GetRecommendCatsRequest struct {
	// The title of a particular item, used to get recommended category ids.
	Name string
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type GetRecommendCatsResponse struct {
	// List of recommended category ids.
	CategoryIDs []string
	// The identifier of the API request for error tracking.
	RequestID string
}
