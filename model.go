package shopeego

//
type RequestLiteBase struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

//
type RequestBase struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

//=======================================================
// GetShopInfoResponse
//=======================================================

type GetShopInfoResponseShop struct {
	// Affiliate shop's id.
	AShopID string `json:"a_shop_id"`
	// Affiliate Shop's area.
	Country string `json:"country"`
}

//=======================================================
// PerformanceResponse
//=======================================================

type PerformanceResponsePerformance struct {
	// The threshold used to compare shop's actual performance to the target performance. It has four types: lt(less than), gt(greater than), lte(less than or equal), gte(greater than or equal).
	ThresholdType string `json:"threshold_type"`
	// Null, not applicable.
	Unit string `json:"unit"`
	// Your target performance index.
	Target int `json:"target"`
	// Your actual performance index.
	My float64 `json:"my"`
}

//=======================================================
// GetShopCategoriesResponse
//=======================================================

type GetShopCategoriesResponseCategory struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int `json:"shop_category_id"`
	// ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
	Status string `json:"status"`
	// ShopCategory's name.
	Name string `json:"name"`
	// ShopCategory's sort weight.
	SortWeight int `json:"sort_weight"`
}

//=======================================================
// GetItemsResponse
//=======================================================

type GetItemsResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
}

//=======================================================
// GetCategoriesResponse
//=======================================================

type GetCategoriesResponseLimits struct {
	// The maximum of dts，-1 means no dts.
	MaxLimit int `json:"max_limit"`
	// The minimum of dts, -1 means no dts.
	MinLimit int `json:"min_limit"`
}

type GetCategoriesResponseCategory struct {
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

//=======================================================
// GetAttributesResponse
//=======================================================

type GetAttributesResponseValue struct {
	// Value in original language. It's MANDATORY to use attributes in original_value to add items.
	OriginalValue string `json:"original_value"`
	// Value in translated language. As referrence only, CANNOT be used to add item. If the selected language is not supported in certain shop location, this field will be empty.
	TranslateValue string `json:"translate_value"`
}

type GetAttributesResponseAttribute struct {
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
	Values []Value `json:"values"`
}

//=======================================================
// AddRequest
//=======================================================

type AddRequestVariation struct {
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price. Max Length: 20 letters
	Name string `json:"name"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock"`
	// The current price of the variation in the listing currency.
	Price float64 `json:"price"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSku string `json:"variation_sku"`
}

type AddRequestImage struct {
	// Url of items' image.The system would synchronous download the image one by one.if one of those image can not fetch, would get a warning in result.But can continue the AddItem proccessing.
	// if all image failed to fetch, would return an error.
	URL string `json:"url"`
}

type AddRequestAttribute struct {
	// related to shopee.item.GetAttributes result.attributes.attribute_id
	AttributesID int `json:"attributes_id"`
	// related to shopee.item.GetAttributes one of result.attributes.options. Max length is 40 letters.
	Value string `json:"value"`
}

type AddRequestLogistic struct {
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

//=======================================================
// AddResponse
//=======================================================

type AddResponseWholesale struct {
	// The min count of this tier wholesale. If the wholesale is not the first one, the min count must equal to max count of last tier plus one.
	Min int `json:"min"`
	// The max count of this tier wholesale.
	Max int `json:"max"`
	// The current price of the wholesale in the listing currency. The price must be cheaper than original price. And if the wholesale is not the first one, the price must be cheaper than previous tier.
	UnitPrice float64 `json:"unit_price"`
}

type AddResponseVariation struct {
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

type AddResponseAttribute struct {
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

type AddResponseLogistic struct {
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

type AddResponseWholesale struct {
	// The min count of this tier wholesale.
	Min int `json:"min"`
	// The max count of this tier wholesale.
	Max int `json:"max"`
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64 `json:"unit_price"`
}

type AddResponseItem struct {
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

//=======================================================
// UnlistItemRequest
//=======================================================

type UnlistItemRequestItem struct {
	// Item's unique identifier.
	ItemID int `json:"item_id"`
	// True: unlist this item; False: list this item.
	Unlist bool `json:"unlist"`
}

//=======================================================
// UnlistItemResponse
//=======================================================

type UnlistItemResponseFailed struct {
	// Item's unique identifier.
	ItemID int `json:"item_id"`
	// Error message.
	ErrorDesciption string `json:"error_desciption"`
}

type UnlistItemResponseSuccess struct {
	// Item's unique identifier.
	ItemID int `json:"item_id"`
	// True: item is unlisted; False: item is listed.
	Unlist bool `json:"unlist"`
}

//=======================================================
// AddVariationsRequest
//=======================================================

type AddVariationsRequestVariation struct {
	// Name of the variation that belongs to the same item.A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock"`
	// The current price of the variation in the listing currency.
	Price float64 `json:"price"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
}

//=======================================================
// AddVariationsResponse
//=======================================================

type AddVariationsResponseVariation struct {
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

//=======================================================
// GetItemsListResponse
//=======================================================

type GetItemsListResponseVariation struct {
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type GetItemsListResponseItem struct {
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

//=======================================================
// GetItemDetailResponse
//=======================================================

type GetItemDetailResponseVariaion struct {
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

type GetItemDetailResponseAttribute struct {
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

type GetItemDetailResponseLogistic struct {
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

type GetItemDetailResponseWholesale struct {
	// The min count of this tier wholesale.
	Min int `json:"min"`
	// The max count of this tier wholesale.
	Max int `json:"max"`
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64 `json:"unit_price"`
}

type GetItemDetailResponseItem struct {
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

//=======================================================
// UpdateItemRequest
//=======================================================

type UpdateItemRequestVariation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
}

type UpdateItemRequestAttribute struct {
	// related to shopee.item.GetAttributes result.attributes.attribute_id
	AttributesID int `json:"attributtes_id"`
	// related to shopee.item.GetAttributes one of result.attributes.options
	Value string `json:"value"`
}

type UpdateItemRequestWholesale struct {
	// The min count of this tier wholesale. If the wholesale is not the first one, the min count must equal to max count of last tier plus one.
	Min int `json:"min"`
	// The max count of this tier wholesale.
	Max int `json:"max"`
	// The current price of the wholesale in the listing currency. The price must be cheaper than original price. And if the wholesale is not the first one, the price must be cheaper than previous tier.'
	UnitPrice float64 `json:"unity_price"`
}

type UpdateItemRequestLogistic struct {
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

//=======================================================
// UpdateItemResponse
//=======================================================

type UpdateItemResponseVariation struct {
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

type UpdateItemResponseAttribute struct {
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

type UpdateItemResponseLogistic struct {
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

type UpdateItemResponseWholesale struct {
	// The min count of this tier wholesale.
	Min int `json:"min"`
	// The max count of this tier wholesale.
	Max int `json:"max"`
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64 `json:"unit_price"`
}

type UpdateItemResponseItem struct {
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

//=======================================================
// UpdatePriceResponse
//=======================================================

type UpdatePriceResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time"`
	// Specify the revised price of the item.
	Price float64 `json:"price"`
}

//=======================================================
// UpdateStockResponse
//=======================================================

type UpdateStockResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time"`
	// Specify the updated stock quantity.
	Stock int `json:"stock"`
}

//=======================================================
// UpdateVariationPriceResponse
//=======================================================

type UpdateVariationPriceResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// Specify the revised price of one variation of the item.
	Price float64 `json:"price"`
}

//=======================================================
// UpdateVariationStockResponse
//=======================================================

type UpdateVariationStockResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time"`
	// Specify the updated stock quantity.
	Stock int `json:"stock"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

//=======================================================
// UpdatePriceBatchRequest
//=======================================================

type UpdatePriceBatchRequestItem struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// New price value for this item.
	Price int `json:"price"`
}

//=======================================================
// UpdatePriceBatchResponse
//=======================================================

type UpdatePriceBatchResponseFailure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int `json:"item_id"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description"`
}

type UpdatePriceBatchResponseBatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []string `json:"modifications"`
	// Informations for failed stock updating.
	Failures []Failure `json:"failures"`
}

//=======================================================
// UpdateStockBatchRequest
//=======================================================

type UpdateStockBatchRequestItem struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// New stock value for this item.
	Stock int `json:"stock"`
}

//=======================================================
// UpdateStockBatchResponse
//=======================================================

type UpdateStockBatchResponseFailure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int `json:"item_id"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description"`
}

type UpdateStockBatchResponseBatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []string `json:"modifications"`
	// Informations for failed stock updating.
	Failures []Failure `json:"failures"`
}

//=======================================================
// UpdateVariationPriceBatchRequest
//=======================================================

type UpdateVariationPriceBatchRequestVariation struct {
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int `json:"variation_id"`
	// New price value of this variation.
	Price int `json:"price"`
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
}

//=======================================================
// UpdateVariationPriceBatchResponse
//=======================================================

type UpdateVariationPriceBatchResponseFailure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int `json:"item_id"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type UpdateVariationPriceBatchResponseModification struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type UpdateVariationPriceBatchResponseBatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []Modification `json:"modifications"`
	// Informations for failed stock updating.
	Failures []Failure `json:"failures"`
}

//=======================================================
// UpdateVariationStockBatchRequest
//=======================================================

type UpdateVariationStockBatchRequestVariation struct {
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int `json:"variation_id"`
	// New stock value of this variation.
	Stock int `json:"stock"`
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
}

//=======================================================
// UpdateVariationStockBatchResponse
//=======================================================

type UpdateVariationStockBatchResponseFailure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int `json:"item_id"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type UpdateVariationStockBatchResponseModification struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
}

type UpdateVariationStockBatchResponseBatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []Modification `json:"modifications"`
	// Informations for failed stock updating.
	Failures []Failure `json:"failures"`
}

//=======================================================
// InitTierVariationRequest
//=======================================================

type InitTierVariationRequestTierVariation struct {
	// Tier variation name.
	Name string `json:"name"`
	// Tier variation value options list. Option length should be under 20. Quantity of combinations of all 2 tier options is up to 50.
	Options []string `json:"options"`
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string `json:"images_url"`
}

type InitTierVariationRequestVariation struct {
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex []int `json:"tier_index"`
	// Stock value of this variation item. The original variation stock will be override when calling this API to initialize 2-tier structure for an existed item. 0 stock will make this variation a greyout option for buyer.
	Stock int `json:"stock"`
	// Price value of this variation item. The original variation price will be override when calling this API to initialize 2-tier structure for an existed item.
	Price float64 `json:"price"`
	// SKU string of this variation.SKU length should be under 100.
	VariationSKU string `json:"variation_sku"`
}

//=======================================================
// InitTierVariationResponse
//=======================================================

type InitTierVariationResponseVariation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int `json:"tier_index"`
	// The identity of the variation.
	VariationID int `json:"variation_id"`
}

//=======================================================
// AddTierVariationRequest
//=======================================================

type AddTierVariationRequestVariation struct {
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex []int `json:"tier_index"`
	// Stock value of this variation item. 0 stock will make this variation a greyout option for buyer.
	Stock int `json:"stock"`
	// Price value of this variation item.
	Price float64 `json:"price"`
	// SKU string of this variation item.
	VariationSKU string `json:"variation_sku"`
}

//=======================================================
// AddTierVariationResponse
//=======================================================

type AddTierVariationResponseVariation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int `json:"tier_index"`
	// The identity of the variation.
	VariationID int `json:"variation_id"`
}

//=======================================================
// GetVariationResponse
//=======================================================

type GetVariationResponseTierVariation struct {
	// Tier variation name.
	Name string `json:"name"`
	// Tier variation value options list.
	Options []string `json:"options"`
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string `json:"images_url"`
}

type GetVariationResponseVariation struct {
	// Unique identifier of the variation.
	VariationID
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex
}

//=======================================================
// UpdateTierVariationListRequest
//=======================================================

type UpdateTierVariationListRequestTierVariation struct {
	// Tier variation name.
	Name string `json:"name"`
	// Tier variation value options list. Lenght should be under 20. Combinations of 2 level options should be under 50.
	Options []string `json:"options"`
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string `json:"images_url"`
}

//=======================================================
// UpdateTierVariationIndexRequest
//=======================================================

type UpdateTierVariationIndexRequestVariation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int `json:"tier_index"`
	// The identity of product item variation.
	VariationID []int `json:"variation_id"`
}

//=======================================================
// BoostItemResponse
//=======================================================

type BoostItemResponseFailure struct {
	// to indicate error type.
	ErrorCode string `json:"error_code"`
	// Failed item id.
	ID int `json:"id"`
	// error description
	Description string `json:"description"`
}

type BoostItemResponseBatchResult struct {
	// A list of item ids which have been boosted successfully.
	Successes []int `json:"successes"`
	// A list of failed-to-boost items, including error details.
	Failures []Failure `json:"failures"`
}

//=======================================================
// GetBoostedItemResponse
//=======================================================

type GetBoostedItemResponseItem struct {
	// boosted items' id.
	ItemID int `json:"item_id"`
	// Cooldown_second time is four hours after boost. After four hours you can boost this item again.
	CooldownSecond int `json:"cooldown_second"`
}

//=======================================================
// GetPromotionInfoResponse
//=======================================================

type GetPromotionInfoResponsePromotion struct {
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

type GetPromotionInfoResponseError struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// Error Message.
	ErrorMsg string `json:"error_msg"`
}

type GetPromotionInfoResponseItem struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int `json:"item_id"`
	// Promotion information list.
	Promotions []Promotion `json:"promotions"`
	// The list of error items.
	Errors []Error `json:"errors"`
}

//=======================================================
// GetPromotionInfoResponse
//=======================================================

type GetPromotionInfoResponseUploadImgResponse struct {
	// origin image url
	ImageURL string `json:"image_url"`
	// Shopee image url
	ShopeeImageURL string `json:"shopee_image_url"`
}

//=======================================================
// AddDiscountRequest
//=======================================================

type AddDiscountRequestVariation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int `json:"variation_id"`
	// The discount price of the item.
	VariationPromotionPrice float64 `json:"variation_promotion_price"`
}

type AddDiscountRequestItem struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	//
	Variations []Variation `json:"variations"`
	// The discount price of the item. If the item has no variation, this param is necessary.
	ItemPromotionPrice float64 `json:"item_promotion_price"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit"`
}

//=======================================================
// AddDiscountItemRequest
//=======================================================

type AddDiscountItemRequestVariation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int `json:"variation_id"`
	// The discount price of the item.
	VariationPromotionPrice float64 `json:"variation_promotion_price"`
}

type AddDiscountItemRequestItem struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	//
	Variations []Variation `json:"variations"`
	// The discount price of the item. If the item has no variation, this param is necessary.
	ItemPromotionPrice float64 `json:"item_promotion_price"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit"`
}

//=======================================================
// GetDiscountDetailResponse
//=======================================================

type GetDiscountDetailResponseVariation struct {
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

type GetDiscountDetailResponseItem struct {
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

//=======================================================
// GetDiscountsListResponse
//=======================================================

type GetDiscountsListResponseDiscount struct {
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

//=======================================================
// UpdateDiscountResponse
//=======================================================

type UpdateDiscountResponseVariation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int `json:"variation_id"`
	// The discount price of the item.
	VariationPromotionPrice float64 `json:"variation_promotion_price"`
}

type UpdateDiscountResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit"`
	// The discount price of the item.
	ItemOriginalPrice float64 `json:"item_original_price"`
	//
	Variations []Variation `json:"variations"`
}

//=======================================================
// GetOrdersListResponse
//=======================================================

type GetOrdersListResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// Enumerated type that defines the current status of the order. Applicable values: See Data Definition- OrderStatus.
	OrderStatus string `json:"order_status"`
	// Timestamp that indicates the last time that there was a change in value of order, such as order status changed from 'Paid' to 'Completed'.
	UpdateTime int `json:"update_time"`
}

//=======================================================
// GetOrdersByStatusResponse
//=======================================================

type GetOrdersByStatusResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// Enumerated type that defines the current status of the order. Applicable values: See Data Definition- OrderStatus.
	OrderStatus string `json:"order_status"`
	// Timestamp that indicates the last time that there was a change in value of order, such as order status changed from 'Paid' to 'Completed'.
	UpdateTime int `json:"update_time"`
}

//=======================================================
// GetOrderDetailsRequest
//=======================================================

type GetOrderDetailsRequestAddress struct {
	// Recipient's name for the address.
	Name string `json:"name"`
	// Recipient's phone number input when order was placed.
	Phone string `json:"phone"`
	// The town of the recipient's address. Whether there is a town will depend on the region and/or country.
	Town string `json:"town"`
	// The district of the recipient's address. Whether there is a town will depend on the region and/or country.
	District string `json:"district"`
	// The city of the recipient's address. Whether there is a town will depend on the region and/or country.
	City string `json:"city"`
	// The state/province of the recipient's address. Whether there is a town will depend on the region and/or country.
	State string `json:"state"`
	// The two-digit code representing the country of the Recipient.
	Country string `json:"country"`
	// Recipient's postal code.
	Zipcode string `json:"zipcode"`
	// The full address of the recipient, including country, state, even street, and etc.
	FullAddress string `json:"full_address"`
}

type GetOrderDetailsRequestItem struct {
	// ID of item
	ItemID int `json:"item_id"`
	// Name of item
	ItemName string `json:"item_name"`
	// A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku"`
	// ID of the variation that belongs to the same item.
	VariationID int `json:"variation_id"`
	// Name of the variation that belongs to the same item.
	// A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	VariationName string `json:"variation_name"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
	// The number of identical items purchased at the same time by the same buyer from one listing/item.
	VariationQuantityPurchased int `json:"variation_quantity_purchased"`
	// The original price of the item in the listing currency.
	VariationOriginalPrice float64 `json:"variation_original_price"`
	// The after-discount price of the item in the listing currency. If there is no discount, this value will be same as that of variation_original_price.
	// In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/variation level. Due to technical restriction, the value will return the price before bundle deal if we don't configure it to 0. Please call GetEscrowDetails if you want to calculate item-level discounted price for bundle deal item.
	VariationDiscountedPrice float64 `json:"variation_discounted_price"`
	// This value indicates whether buyer buy the order item in wholesale price.
	IsWholesale bool `json:"is_wholesale"`
	// The weight of the item
	Weight float64 `json:"weight"`
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool `json:"is_add_on_deal"`
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool `json:"is_main_item"`
	// The unique identity of an addon deal.
	AddOnDealID int `json:"add_on_deal_id"`
	// The type of the promotion,
	PromotionType string `json:"promotion_type"`
	// The ID of the promotion.
	PromotionID int `json:"promotion_id"`
}

type GetOrderDetailsRequestOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The two-digit code representing the country where the order was made.
	Country string `json:"country"`
	// The three-digit code representing the currency unit for which the order was paid.
	Currency string `json:"currency"`
	// This value indicates whether the order was a COD (cash on delivery) order.
	COD bool `json:"cod"`
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no"`
	// Shipping preparation time set by the seller when listing item on Shopee.
	DaysToShip int `json:"days_to_ship"`
	// This object contains detailed breakdown for the recipient address.
	RecipientAddress Address `json:"recipient_address"`
	// This object contains the detailed breakdown for the result of this API call.
	Items []Item `json:"items"`
	// The time when the order status is updated from UNPAID to PAID. This value is NULL when order is not paid yet.
	PayTime int `json:"pay_time"`
	// For Indonesia orders only. The name of the dropshipper.
	Dropshipper string `json:"dropshipper"`
	// Last 4 digits of the credit card
	CreditCardNumber string `json:"credit_card_number"`
	// The name of buyer
	BuyerUsername string `json:"buyer_username"`
	// The phone number of dropshipper
	DropshipperPhone string `json:"dropshipper_phone"`
	// The deadline to ship out the parcel.
	ShipByDate int `json:"ship_by_date"`
	// To indicate whether this order is split to fullfil order(forder) level. Call GetForderInfo if it's "true".
	IsSplitUp bool `json:"is_split_up"`
	// Cancel reason from buyer.
	BuyerCancelReason string `json:"buyer_cancel_reason"`
	// Could be one of buyer, seller or system
	CancelBy string `json:"cancel_by"`
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
	// Use this field to get reason for buyer, seller, and system cancellation.
	CancelReason string `json:"cancel_reason"`
}

//=======================================================
// GetEscrowDetailsResponse
//=======================================================

type GetEscrowDetailsResponseIncome struct {
	// The three-digit code representing the currency unit used for all transactional amount under
	LocalCurrency string `json:"local_currency"`
	// The total amount paid by the buyer for the order. This amount includes the total sale price of items, shipping cost beared by buyer; and offset by Shopee promotions if applicable.
	TotalAmount float64 `json:"total_amount"`
	// Final value of coins used by seller for the order.
	Coin float64 `json:"coin"`
	// Final value of voucher provided by Shopee for the order.
	Voucher float64 `json:"voucher"`
	// Final value of voucher provided by Seller for the order.
	VoucherSeller float64 `json:"voucher_seller"`
	// Final sum of each item Shopee discount of a specific order. This amount will rebate to seller.
	SellerRebate float64 `json:"seller_rebate"`
	// The final shipping cost of order . For Non-integrated logistics channel is 0.
	ActualShippingCost float64 `json:"actual_shipping_cost"`
	// The platform shipping subsidy to the seller
	ShippingFeeRebate float64 `json:"shipping_fee_rebate"`
	// The commission fee charged by Shopee platform if applicable.
	CommissionFee float64 `json:"commission_fee"`
	// The voucher code or promotion code the buyer used.
	VoucherCode float64 `json:"voucher_code"`
	// The voucher name or promotion name the buyer used.
	VoucherName float64 `json:"voucher_name"`
	// The total amount that the seller is expected to receive for the order and will change before order completed. escrow_amount=total_amount+voucher+credit_card_promotion+seller_rebate+coin-commission_fee-credit_card_transaction_fee-cross_border_tax-service_fee-buyer_shopee_kredit-seller_coin_cash_back+final_shipping_fee-seller_return_refund_amount.
	EscrowAmount float64 `json:"escrow_amount"`
	// Amount incurred by Buyer for purchasing items outside of home country. Amount may change after Return Refund.
	CroossBorderTax float64 `json:"crooss_border_tax"`
	// The amount offset via payment promotion.
	CreditCardPromotion float64 `json:"credit_card_promotion"`
	// Include buyer transaction fee and seller transaction fee.
	CreditCardTransactionFee float64 `json:"credit_card_transaction_fee"`
	// Amount charged by Shopee to seller for additional services.
	ServiceFee float64 `json:"service_fee"`
	// Amount charged by Shopee to Buyer for using ShopeeKredit for the order. Currently only applicable in ID.
	BuyerShopeeKredit float64 `json:"buyer_shopee_kredit"`
	// Value of coins provided by Seller for purchasing with his or her store for the order.
	SellerCoinCashBack float64 `json:"seller_coin_cash_back"`
	// Final adjusted amount that seller has to bear as part of escrow. This amount could be negative or positive.
	FinalShippingFee float64 `json:"final_shipping_fee"`
	// Amount returned to Seller in the event of partial return.
	SellerReturnRefundAmount float64 `json:"seller_return_refund_amount"`
	// The amount offset via payment promotion. May include bank payment promotion and Shopee payment promotion.
	CreditCardPromotion float64 `json:"credit_card_promotion"`
}

type GetEscrowDetailsResponseBankAccount struct {
	// Name of the seller's receiving bank
	BankName string `json:"bank_name"`
	// Account number of the seller's receiving bank
	BankAccountNumber string `json:"bank_account_number"`
	// The two-digit code representing country of the seller's receiving bank account
	BankAccountCountry string `json:"bank_account_country"`
}

type GetEscrowDetailsResponseItem struct {
	// ID of item.
	ItemID int `json:"item_id"`
	// ID of the variation that belongs to the same item.
	VariationID int `json:"variation_id"`
	// The number of identical items purchased at the same time by the same buyer from one listing/item.
	QuantityPurchased int `json:"quantity_purchased"`
	// The price used to participate activity. E.g. itemA original price is $10, promo price is $9, and bundle deal is buy 2 get 20% off equals to $14.4. The original_price value will be $9 in this case.
	OriginalPrice float64 `json:"original_price"`
}

type GetEscrowDetailsResponseActivity struct {
	// ID of activity.
	ActivityID int `json:"activity_id"`
	// Type of activity. Currently only one type: bundle_deal
	ActivityType string `json:"activity_type"`
	// The original TOTAL price of ALL items in one activity(e.g. bundle deal. Define by activity_id) in the listing currency.
	OriginalPrice float64 `json:"original_price"`
	// The after-discocunt TOTAL price of ALL items in one activity(e.g. bundle deal. Define by activity_id) in the listing currency.
	DiscountedPrice float64 `json:"discounted_price"`
	// This object contains the items in this activity.
	Items []Item `json:"items"`
}

type GetEscrowDetailsResponseItem struct {
	// ID of item
	ItemID int `json:"item_id"`
	// Name of item
	ItemName string `json:"item_name"`
	// A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku"`
	// ID of the variation that belongs to the same item.
	VariationID int `json:"variation_id"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	VariationName string `json:"variation_name"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku"`
	// This value indicates the number of identical items purchased at the same time by the same buyer from one listing/item.
	QuantityPurchased int `json:"quantity_purchased"`
	// The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.
	OriginalPrice float64 `json:"original_price"`
	// The after-discount price of the item in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1. If there is no discount, this value will be the same as that of original_price.
	// In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/variation level. Due to technical restriction, the value will return the price before bundle deal if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the bundle deal discount breakdown on item level.
	DiscountedPrice float64 `json:"discounted_price"`
	// The offset of this item when the buyer consumed Shopee Coins upon checkout.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	DiscountFromCoin float64 `json:"discount_from_coin"`
	// The offset of this item when the buyer use Shopee voucher.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	DiscountFromVoucher float64 `json:"discount_from_voucher"`
	// The offset of this item when the buyer use seller-specific voucher.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	DiscountFromVoucherSeller float64 `json:"discount_from_voucher_seller"`
	// Platform subsidy to the seller for this item.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	SellerRebate float64 `json:"seller_rebate"`
	// This value indicates the actual price the buyer pay.
	// In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/variation level. Due to technical restriction, the value will return the price before bundle deal if we don't configure it to 0. Please use the value under "income_details" and "activity" to calculate the bundle deal discount breakdown on item level.
	DealPrice float64 `json:"deal_price"`
	// This value indicate the offset via credit card promotion.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	CreditCardPromotion float64 `json:"credit_card_promotion"`
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool `json:"is_add_on_deal"`
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool `json:"is_main_item"`
	// The unique identity of an addon deal.
	AddOnDealID int `json:"add_on_deal_id"`
}

type GetEscrowDetailsResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The two-digit code representing the country where the order was made.
	Country string `json:"country"`
	// This object contains detailed income breakdown for the order.
	IncomeDetails IncomeDetail `json:"income_details"`
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string `json:"shipping_carrier"`
	// The three-digit code representing the currency unit of total order amount (escorw_amount) at the point of payment to the seller.
	EscrowCurrency string `json:"escrow_currency"`
	// The exchange rate used by Shopee to convert local_currency to escrow_currency.
	ExchangeRate string `json:"exchange_rate"`
	// The payment channel that the seller selected to receive escrow for the order.
	EscrowChannel string `json:"escrow_channel"`
	// The unique identifier for a payee by the 3rd party payment service provider selected in escrow_channel.
	PayeeID int `json:"payee_id"`
	// This object contains detailed breakdown for bank account of the seller if selected escorw_channel is Bank Transfer.
	BankAccount BankAccount `json:"bank_account"`
	// This object contains the detailed breakdown for all the items in this order, including regular items(non-activity) and activity items.
	Items []Item `json:"items"`
	// This object contains the activity in this order.
	Activity []Activity `json:"activity"`
}

//=======================================================
// GetForderInfoResponse
//=======================================================

type GetForderInfoResponseLog struct {
	// The time when logistics info has been updated.
	Ctime int `json:"ctime"`
	// The order logistics tracking info.
	Description string `json:"description"`
}

type GetForderInfoResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// The number of identical items/variations purchased at the same time by the same buyer from one listing/item.
	Num int `json:"num"`
	// The original price of the item in the listing currency.
	ItemPrice float64 `json:"item_price"`
	// The original price of the variation in the listing currency.
	VariationPrice float64 `json:"variation_price"`
}

type GetForderInfoResponseLogisticsInfo struct {
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string `json:"shipping_carrier"`
	// Only work for cross-border order. This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	GoodsToDeclare bool `json:"goods_to_declare"`
	// Only work for cross-border order. This code is required at some sorting hub. Please ensure the service_code is INCLUDED on your shipping label, otherwise the parcel cannot be processed by the warehouse. If you didn't retrieve service_code after you first called this API, please try few more times within 30 minutes.
	ServiceCode string `json:"service_code"`
	// The tracking number of fullfill order assigned by the shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no"`
}

type GetForderInfoResponseForder struct {
	// The unique identifier for a fulfill order.
	ForderID string `json:"forder_id"`
	// The fulfill order logistics status. Applicable values: See Data Definition - LogisticsStatus.
	Status string `json:"status"`
	// Logistics tracking info.
	TrackingLog []Log `json:"tracking_log"`
	// The items included in this fulfill order.
	Items []Item `json:"items"`
	//
	LogisticsInfo []LogisticsInfo `json:"logistics_info"`
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
}

//=======================================================
// GetEscrowReleasedOrdersResponse
//=======================================================

type GetEscrowReleasedOrdersResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// Order's escrow amount.
	PayoutAmount float64 `json:"payout_amount"`
	// Timestamp of escrow amount transaction finished.
	EscrowReleaseTime int `json:"escrow_release_time"`
}

//=======================================================
// SplitOrderRequest
//=======================================================

type SplitOrderRequestParcel struct {
	// Itemids that will be put into a fullfillment order.
	ItemID int `json:"item_id"`
}

//=======================================================
// SplitOrderResponse
//=======================================================

type SplitOrderResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int `json:"item_id"`
}

type SplitOrderResponseForder struct {
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
	// Item information contained in fulfillment orders.Number of items must be greater than or equal to 2. eg.[[{"item_id": 123}],[{"item_id": 456}]]
	Items []Item `json:"items"`
}

//=======================================================
// GetUnbindOrderListResponse
//=======================================================

type GetUnbindOrderListResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The Shopee logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.
	LogisticStatus string `json:"logistic_status"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
}

//=======================================================
// GetLogisticsResponse
//=======================================================

type GetLogisticsResponseSize struct {
	// The identity of size.
	SizeID int `json:"size_id"`
	// The name of size.
	Name int `json:"name"`
	// The pre-defined shipping fee for the specific size.
	DefaultPrice int `json:"default_price"`
}

type GetLogisticsResponseLimit struct {
	// The max weight for an item on this logistic channel.If the value is 0 or null, that means there is no limit.
	ItemMaxWeight
	// The min weight for an item on this logistic channel. If the value is 0 or null, that means there is no limit.
	ItemMinWeight
}

type GetLogisticsResponseDimension struct {
	// The max height limit.
	Height float64 `json:"height"`
	// The max width limit.
	Width float64 `json:"width"`
	// The max length limit.
	Length float64 `json:"length"`
	// The unit for the limit.
	Unit string `json:"unit"`
}

type GetLogisticsResponseLogistic struct {
	// The identity of logistic channel
	LogisticID int `json:"logistic_id"`
	// The name of logistic channel
	LogisticName string `json:"logistic_name"`
	// This is to indicate whether this logistic channel supports COD
	HasCOD bool `json:"has_cod"`
	// Whether this logistic channel is enabled on shop level.
	Enabled bool `json:"enabled"`
	// See Define FeeType, related to FeeType Value
	FeeType string `json:"fee_type"`
	// Only for fee_type is SIZE_SELECTION
	Sizes Size `json:"sizes"`
	// The weight limit for this logistic channel.
	WeightLimits Limit `json:"weight_limits"`
	// The dimension limit for this logistic channel.
	ItemMaxDimension Dimension `json:"item_max_dimension"`
}

//=======================================================
// GetAddressResponse
//=======================================================

type GetAddressResponseAddress struct {
	// The identity of address
	AddressID int `json:"address_id"`
	// The country of specify address
	Country string `json:"country"`
	// The state of specify address
	State string `json:"state"`
	// The city of specify address
	City string `json:"city"`
	// The address description of specify address
	Address string `json:"address"`
	// The zipcode of specify address
	Zipcode string `json:"zipcode"`
	// The district of specify address
	District string `json:"district"`
	// The town of specify address
	Town string `json:"town"`
}

//=======================================================
// GetTimeSlotResponse
//=======================================================

type GetTimeSlotResponseTime struct {
	// The identity of pickuptime.
	PickupTimeID string `json:"pickup_time_id"`
	// The date of pickup time. In timestamp.
	Date int `json:"date"`
	// The text description of pickup time. Only applicable for certain channels.
	TimeText string `json:"time_text"`
}

//=======================================================
// GetBranchResponse
//=======================================================

type GetBranchResponseBranch struct {
	// The identity of branch.
	BranchID int `json:"branch_id"`
	// The country of specify branch.
	Country string `json:"country"`
	// The state of specify branch.
	State string `json:"state"`
	// The city of specify branch.
	City string `json:"city"`
	// The address description of specify branch.
	Address string `json:"address"`
	// The zipcode of specify branch.
	Zipcode string `json:"zipcode"`
	// The district of specify branch.
	District string `json:"district"`
	// The town of specify branch.
	Town string `json:"town"`
}

//=======================================================
// GetLogisticInfoResponse
//=======================================================

type GetLogisticInfoResponseAddress struct {
	// The identity of address.
	AddressID int `json:"address_id"`
	// The country of specify branch.
	Country string `json:"country"`
	// The state of specify branch.
	State string `json:"state"`
	// The city of specify branch.
	City string `json:"city"`
	// The address description of specify branch.
	Address string `json:"address"`
	// The zipcode of specify branch.
	Zipcode string `json:"zipcode"`
	// The district of specify branch.
	District string `json:"district"`
	// The town of specify branch.
	Town string `json:"town"`
	// List of pickup_time information corresponding to the address_id.
	TimeSlotList []string `json:"time_slot_list"`
	// The identity of pickuptime.
	PickupTimeID string `json:"pickup_time_id"`
	// The date of pickup time. In timestamp.
	Date int `json:"date"`
	// The text description of pickup time. Only applicable for certain channels.
	TimeText string `json:"time_text"`
}

type GetLogisticInfoResponseBranch struct {
	// The identity of branch.
	BranchID int `json:"branch_id"`
	// The country of specify branch.
	Country string `json:"country"`
	// The state of specify branch.
	State string `json:"state"`
	// The city of specify branch.
	City string `json:"city"`
	// The address description of specify branch.
	Address string `json:"address"`
	// The zipcode of specify branch.
	Zipcode string `json:"zipcode"`
	// The district of specify branch.
	District string `json:"district"`
	// The town of specify branch.
	Town string `json:"town"`
}

type GetLogisticInfoResponsePickup struct {
	// List of available pickup address info.
	AddressList []Address `json:"address_list"`
}

type GetLogisticInfoResponseDropoff struct {
	// List of available dropoff branches info.
	BranchList []Branch `json:"branch_list"`
}

type GetLogisticInfoResponseInfo struct {
	// Logistics information for pickup mode order.
	Pickup []string `json:"pickup"`
	// Logistics information for dropoff mode order.
	Dropoff []string `json:"dropoff"`
	// The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init.
	NonIntegrated []string `json:"non_integrated"`
}

//=======================================================
// InitRequest
//=======================================================

type InitRequestPickup struct {
	// The identity of address. Retrieved from shopee.logistics.GetAddress.
	AddressID int `json:"address_id"`
	// The pickup time id. Retrieved from shopee.logistics.GetTimeSlot.
	PickupItemID string `json:"pickup_item_id"`
}

type InitRequestDropoff struct {
	// The identity of branch. Retrieved from shopee.logistics.GetBranch branch.
	BranchID int `json:"branch_id"`
	// The real name of sender.
	SenderRealName string `json:"sender_real_name"`
	// Need input this field when "tracking_no" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no"`
}

type InitRequestNonIntegrated struct {
	// Optional parameter for non-integrated channel order. The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no"`
}

//=======================================================
// GetAirwayBillResponse
//=======================================================

type GetAirwayBillResponseAirwayBill struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The url of retrieving airway bill.
	AirwayBill string `json:"airway_bill"`
}

type GetAirwayBillResponseError struct {
	// The ordersn of orders which occurred error.
	OrderSN string `json:"order_sn"`
	//
	ErrorCode string `json:"error_code"`
	// The detail information of this error.
	ErrorDescription string `json:"error_description"`
}

type GetAirwayBillResponseResult struct {
	// The number of ordersn to get airway bills in this call.
	TotalCount int `json:"total_count"`
	// This Object contains the airway bill to each order.
	AirwayBills []AirwayBill `json:"airway_bills"`
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
}

type GetAirwayBillResponseBatchResult struct {
	// The number of orderSN to get airway bills in this call.
	TotalCount int `json:"total_count"`
	// The list contains urls of retrieving airway bill in PDF format. Each url contains the airway bills which is in the same logistics channel.
	AirwayBills []string `json:"airway_bills"`
}

//=======================================================
// GetOrderLogisticsResponse
//=======================================================

type GetOrderLogisticsResponseAddress struct {
	// Recipient's name for the address.
	Name string `json:"name"`
	// Recipient's phone number input when order was placed.
	Phone string `json:"phone"`
	// The town of the recipient's address. Whether there is a town will depend on the region and/or country.
	Town string `json:"town"`
	// The district of the recipient's address. Whether there is a town will depend on the region and/or country.
	District string `json:"district"`
	// The city of the recipient's address. Whether there is a town will depend on the region and/or country.
	City string `json:"city"`
	// The state/province of the recipient's address. Whether there is a town will depend on the region and/or country.
	State string `json:"state"`
	// The two-digit code representing the country of the Recipient.
	Country string `json:"country"`
	// Recipient's postal code.
	Zipcode string `json:"zipcode"`
	// The full address of the recipient, including country, state, even street, and etc.
	FullAddress string `json:"full_address"`
}

type GetOrderLogisticsResponseLogistic struct {
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string `json:"shipping_carrier"`
	// The identity of logistic channel.
	LogisticID int `json:"logistic_id"`
	// Only work for cross-border order. This code is required at some sorting hub. Please ensure the service_code is INCLUDED on your shipping label, otherwise the parcel cannot be processed by the warehouse. If you didn't retrieve service_code after you first called this API, please try few more times within 30 minutes.
	ServiceCode string `json:"service_code"`
	// Only work for cross-border order.The name of the carrier ships cross countries.
	FirstMileName string `json:"first_mile_name"`
	// Only work for cross-border order.The name of the carrier delivers the parcels in local country.
	LastMileName string `json:"last_mile_name"`
	// Only work for cross-border order.This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	GoodsToDeclare bool `json:"goods_to_declare"`
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no"`
	//
	Zone string `json:"zone"`
	// Only work for cross-border order. The string use for waybill printing. The format is "S - country_code and lane_number". For example, S-TH01, S-TH02
	LaneCode string `json:"lane_code"`
	// Only work for cross-border order in some special shop. The address info of the warehouse.
	WarehouseAddress string `json:"warehouse_address"`
	// Only work for cross-border order in some special shop. The ID of the warehouse.
	WarehouseID int `json:"warehouse_id"`
	// This object contains detailed breakdown for the recipient address.
	RecipientAddress Address `json:"recipient_address"`
	// This value indicates whether the order was a COD (cash on delivery) order.
	COD bool `json:"cod"`
}

//=======================================================
// GetLogisticsMessageResponse
//=======================================================

type GetLogisticsMessageResponseInfo struct {
	// The time when logistics info has been updated.
	CTime int `json:"c_time"`
	// The order logistics tracking info.
	Description string `json:"description"`
	// The 3PL logistics status for the order. Applicable values: See Data Definition - TrackingLogisticsStatus.
	Status string `json:"status"`
}

//=======================================================
// GetForderWaybillRequest
//=======================================================

type GetForderWaybillRequestList struct {
	// The order serial numbers. Make sure the order has trackingNo generated before calling this API.
	OrderSN string `json:"order_sn"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
}

//=======================================================
// GetForderWaybillResponse
//=======================================================

type GetForderWaybillResponseWaybill struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
	// The url of retrieving airway bill.
	Waybill string `json:"waybill"`
}

type GetForderWaybillResponseError struct {
	// The ordersn of orders which occurred error.
	OrderSN string `json:"order_sn"`
	// The forder_id of fulfillment orders which occurred error.
	ForderID string `json:"forder_id"`
	//
	ErrorCode
	// The detail information of this error.
	ErrorDescription string `json:"error_description"`
}

type GetForderWaybillResponseResult struct {
	// The number of ordersn to get airway bills in this call.
	TotalCount int `json:"total_count"`
	// This Object contains the airway bill to each order.
	Waybills []Waybill `json:"waybills"`
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
}

type GetForderWaybillResponseBatchResult struct {
	// The number of orderSN to get airway bills in this call.
	TotalCount int `json:"total_count"`
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
	// The url of retrieving airway bill.
	Waybills []string `json:"waybills"`
}

//=======================================================
// GetReturnListResponse
//=======================================================

type GetReturnListResponseUser struct {
	// Buyer's nickname.
	Username int `json:"username"`
	// Buyer's email.
	Email string `json:"email"`
	// Buyer's portrait.
	Protrait string `json:"protrait"`
}

type GetReturnListResponseItem struct {
	// Item id.
	ItemID int `json:"item_id"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// Name of item in local language.
	Name string `json:"name"`
	// Image URLs of item.
	Images []string `json:"images"`
	// Amount of this item.
	Amount int `json:"amount"`
	// The price of Item.
	ItemPrice float64 `json:"item_price"`
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool `json:"is_add_on_deal"`
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool `json:"is_main_item"`
	// The unique identity of an addon deal.
	AddOnDealID int `json:"add_on_deal_id"`
}

type GetReturnListResponseReturn struct {
	// Image URLs of return.
	Images []string `json:"images"`
	// Reason for return product. Applicable values: See Data Definition- ReturnReason.
	Reason string `json:"reason"`
	// Reason that buyer provide.
	TextReason string `json:"text_reason"`
	// The serial number of return.
	ReturnSN int `json:"return_sn"`
	// Amount of the refund.
	RefundAmount float64 `json:"refund_amount"`
	// Currency of the return.
	Currency string `json:"currency"`
	// The time of return create.
	CreateTime int `json:"create_time"`
	// The time of modify return.
	UpdateTime int `json:"update_time"`
	// Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
	Status string `json:"status"`
	// The last time seller deal with this return.
	DueDate int `json:"due_date"`
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNumber string `json:"tracking_number"`
	// The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
	DisputeReason string `json:"dispute_reason"`
	// The reason that seller provide. While the return has been disputed, this field is useful.
	DisputeTextReason string `json:"dispute_text_reason"`
	// Items to be sent back to seller. Can be either integrated/non-integrated.
	NeedsLogistics bool `json:"needs_logistics"`
	// Order price before discount.
	AmountBeforeDiscount floatt64 `json:"amount_before_discount"`
	//
	User User `json:"user"`
	//
	Item []Item `json:"item"`
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
}

//=======================================================
// GetReturnDetailResponse
//=======================================================

type GetReturnDetailResponseUser struct {
	// Buyer's nickname.
	Username int `json:"username"`
	// Buyer's email.
	Email string `json:"email"`
	// Buyer's portrait.
	Portrait string `json:"portrait"`
}

type GetReturnDetailResponseItem struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int `json:"variation_id"`
	// Name of item in local language.
	Name string `json:"name"`
	// Image URLs of item.
	Images []string `json:"images"`
	// Amount of this item.
	Amount int `json:"amount"`
	// The price of item.
	ItemPrice float64 `json:"item_price"`
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool `json:"is_add_on_deal"`
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool `json:"is_main_item"`
	// The unique identity of an addon deal.
	AddOnDealID int `json:"add_on_deal_id"`
}

//=======================================================
// GetShopsByPartnerResponse
//=======================================================

type GetShopsByPartnerResponseSIP struct {
	// Affiliate Shop's area
	Country string `json:"country"`
	// Affiliate shop's id
	AShopID int `json:"a_shop_id"`
}

type GetShopsByPartnerResponseShop struct {
	// The two-digit code representing the country where the order was made.
	Country string `json:"country"`
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// The timestamp when the shop was authorized to the partner.
	AuthTime int `json:"auth_time"`
	// SIP affiliate shops info list
	SIPAShops []SIP `json:"sip_a_shops"`
	// Use this field to indicate the expiration date for shop authorization.
	ExpireTime int `json:"expire_time"`
}

//=======================================================
// GetCategoriesByCountryResponse
//=======================================================

type GetCategoriesByCountryResponseCategory struct {
	// The Identify of the parent of the category
	ParentID int `json:"parent_id"`
	// This is to indicate whether the category has children.
	HasChildren bool `json:"has_children"`
	// The Identify of each category
	CategoryID int `json:"category_id"`
	// The name of each category
	CategoryName string `json:"category_name"`
	// To indicate if this category supports size chart
	IsSuppSizechart bool `json:"is_supp_sizechart"`
}

//=======================================================
// GetPaymentListResponse
//=======================================================

type GetPaymentListResponseMethod struct {
	// The payment method
	PaymentMethod string `json:"payment_method"`
	// The country for this payment method
	Country string `json:"country"`
}

//=======================================================
// GetTopPicksListResponse
//=======================================================

type GetTopPicksListResponseItem struct {
	// Item ID
	ItemID int `json:"item_id"`
	// Item name
	ItemName string `json:"item_name"`
	// Item discounted price(original price if no discount). Item level price will return if it has variation.
	ItemPrice float64 `json:"item_price"`
	// The sales of the item
	Sales int `json:"sales"`
}

type GetTopPicksListResponseCollection struct {
	// Collection name
	Name string `json:"name"`
	// Collection ID
	TopPicksID int `json:"top_picks_id"`
	// True or False
	IsActivated bool `json:"is_activated"`
	// Item list of the collection
	Items []Item `json:"items"`
}

//=======================================================
// AddTopPicksResponse
//=======================================================

type AddTopPicksResponseItem struct {
	// Item ID
	ItemID int `json:"item_id"`
	// Item name
	ItemName string `json:"item_name"`
	// Item discounted price(original price if no discount). Item level price will return if it has variation.
	ItemPrice float64 `json:"item_price"`
	// The sales of the item
	Sales int `json:"sales"`
}

//=======================================================
// UpdateTopPicksResponse
//=======================================================

type UpdateTopPicksResponseItem struct {
	// Item ID
	ItemID int `json:"item_id"`
	// Item name
	ItemName string `json:"item_name"`
	// Item discounted price(original price if no discount). Item level price will return if it has variation.
	ItemPrice float64 `json:"item_price"`
	// The sales of the item
	Sales int `json:"sales"`
}

//=======================================================
// GenerateFMTrackingNoRequest
//=======================================================

type GenerateFMTrackingNoRequestInfo struct {
	// The full address of the seller.
	Address string `json:"address"`
	// Seller's name for the address.
	Name string `json:"name"`
	// Seller's postal code.
	Zipcode string `json:"zipcode"`
	// Seller's location.
	Area string `json:"area"`
	// Seller's phone number.
	Phone string `json:"phone"`
}

//=======================================================
// GetShopFMTrackingNoResponse
//=======================================================

type GetShopFMTrackingNoResponseList struct {
	// The specified delivery date.
	DeclareDate string `json:"declare_date"`
	// The logistics status for bound orders.
	Status string `json:"status"`
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
}

//=======================================================
// FirstMileCodeBindOrderRequest
//=======================================================

type FirstMileCodeBindOrderRequestOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
}

//=======================================================
// FirstMileCodeBindOrderResponse
//=======================================================

type FirstMileCodeBindOrderResponseList struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
	// The reason why the order/fulfillment order cannot be bound.
	Reason string `json:"reason"`
}

//=======================================================
// GetFmTnDetailResponse
//=======================================================

type GetFmTnDetailResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
	// The tracking number of SLS for orders/forders.
	SLSTN string `json:"slstn"`
}

//=======================================================
// GetFMTrackingNoWaybillResponse
//=======================================================

type GetFMTrackingNoWaybillResponseError struct {
	//
	ErrorCode string `json:"error_code"`
	// The detail information of this error.
	ErrorDescription string `json:"error_description"`
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
}

type GetFMTrackingNoWaybillResponseWaybill struct {
	// The first-mile tracking number.
	FMTN string `json:"fmtn"`
	// The url of retrieving waybill.
	Waybill string `json:"waybill"`
}

type GetFMTrackingNoWaybillResponseResult struct {
	// This Object contains the waybill to each tracking number.
	Waybills []Waybill `json:"waybills"`
	// The number of Tracking Number to get waybills in this call.
	TotalCount int `json:"total_count"`
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
}

type GetFMTrackingNoWaybillResponseBatch struct {
	// The list contains urls of retrieving waybill in PDF format. Each url contains the airway bills which is in the same logistics channel.
	Waybills []string `json:"waybills"`
	// The number of Tracking Number to get waybills in this call.
	TotalCount int `json:"total_count"`
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
}

//=======================================================
// GetShopFirstMileChannelResponse
//=======================================================

type GetShopFirstMileChannelResponseLogistic struct {
	// The identity of logistic channel.
	LogisticID int `json:"logistic_id"`
	// The name of logistic.
	LogisticName string `json:"logistic_name"`
	// The shipment method for bound orders, should be pickup or dropoff.
	ShipmentMethod string `json:"shipment_method"`
}
