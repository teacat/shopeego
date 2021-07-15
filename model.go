package shopeego

//=======================================================
// GetShopInfoResponse
//=======================================================

type GetShopInfoResponseShop struct {
	// Affiliate shop's id.
	AShopID string `json:"a_shopid,omitempty"`
	// Affiliate Shop's area.
	Country string `json:"country,omitempty"`
}

//=======================================================
// PerformanceResponse
//=======================================================

type PerformanceResponsePerformance struct {
	// The threshold used to compare shop's actual performance to the target performance. It has four types: lt(less than), gt(greater than), lte(less than or equal), gte(greater than or equal).
	ThresholdType string `json:"threshold_type,omitempty"`
	// Null, not applicable.
	Unit string `json:"unit,omitempty"`
	// Your target performance index.
	Target float64 `json:"target,omitempty"`
	// Your actual performance index.
	My float64 `json:"my,omitempty"`
}

//=======================================================
// GetShopCategoriesResponse
//=======================================================

type GetShopCategoriesResponseCategory struct {
	// ShopCategory's unique identifier.
	ShopCategoryID int64 `json:"shop_category_id,omitempty"`
	// ShopCategory's status. Applicable values: NORMAL, INACTIVE, DELETED.
	Status string `json:"status,omitempty"`
	// ShopCategory's name.
	Name string `json:"name,omitempty"`
	// ShopCategory's sort weight.
	SortWeight int `json:"sort_weight,omitempty"`
}

//=======================================================
// GetItemsResponse
//=======================================================

type GetItemsResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
}

//=======================================================
// GetCategoriesResponse
//=======================================================

type GetCategoriesResponseCategoryDaysToShipLimits struct {
	// The maximum of dts，-1 means no dts.
	MaxLimit int `json:"max_limit,omitempty"`
	// The minimum of dts, -1 means no dts.
	MinLimit int `json:"min_limit,omitempty"`
}

type GetCategoriesResponseCategory struct {
	// The Identify of each category.
	CategoryID int64 `json:"category_id,omitempty"`
	// The Identify of the parent of the category.
	ParentID int64 `json:"parent_id,omitempty"`
	// The name of each category.
	CategoryName string `json:"category_name,omitempty"`
	// This is to indicate whether the category has children. Attributes can ONLY be fetched for the category_id WITHOUT children.
	HasChildren bool `json:"has_children,omitempty"`
	// The limits of pre-order items' days_to_ship based on per category.
	DaysToShipLimits GetCategoriesResponseCategoryDaysToShipLimits `json:"days_to_ship_limits,omitempty"`
}

//=======================================================
// GetAttributesResponse
//=======================================================

type GetAttributesResponseAttributeValue struct {
	// Value in original language. It's MANDATORY to use attributes in original_value to add items.
	OriginalValue string `json:"original_value,omitempty"`
	// Value in translated language. As referrence only, CANNOT be used to add item. If the selected language is not supported in certain shop location, this field will be empty.
	TranslateValue string `json:"translate_value,omitempty"`
}

type GetAttributesResponseAttribute struct {
	// The Identify of each category.
	AttributeID int64 `json:"attribute_id,omitempty"`
	// The name of each attribute.
	AttributeName string `json:"attribute_name,omitempty"`
	// This is to indicate whether this attribute is mandantory.
	IsMandatory bool `json:"is_mandatory,omitempty"`
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttribuiteType string `json:"attribuite_type,omitempty"`
	// Enumerated type that defines the input type of the attribute. Applicable values: See Data Definition- AttributeInputType.
	InputType string `json:"input_type,omitempty"`
	// All options that attribute contains.
	Options []string `json:"options,omitempty"`
	// The option values in different language.
	Values []GetAttributesResponseAttributeValue `json:"values,omitempty"`
}

//=======================================================
// AddRequest
//=======================================================

type AddRequestVariation struct {
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price. Max Length: 20 letters
	Name string `json:"name,omitempty"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock,omitempty"`
	// The current price of the variation in the listing currency.
	Price float64 `json:"price,omitempty,string"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSku string `json:"variation_sku,omitempty"`
}

type AddRequestImage struct {
	// Url of items' image.The system would synchronous download the image one by one.if one of those image can not fetch, would get a warning in result.But can continue the AddItem proccessing.
	// if all image failed to fetch, would return an error.
	URL string `json:"url,omitempty"`
}

type AddRequestWholesale struct {
	// The min count of this tier wholesale. If the wholesale is not the first one, the min count must equal to max count of last tier plus one.
	Min int `json:"min,omitempty"`
	// The max count of this tier wholesale.
	Max int `json:"max,omitempty"`
	// The current price of the wholesale in the listing currency. The price must be cheaper than original price. And if the wholesale is not the first one, the price must be cheaper than previous tier.
	UnitPrice float64 `json:"unit_price,omitempty,string"`
}

type AddRequestAttribute struct {
	// related to shopee.item.GetAttributes result.attributes.attribute_id
	AttributesID int64 `json:"attributes_id,omitempty"`
	// related to shopee.item.GetAttributes one of result.attributes.options. Max length is 40 letters.
	Value string `json:"value,omitempty"`
}

type AddRequestLogistic struct {
	// related to shopee.logistics.GetLogistics result.logistics.logistic_id
	LogisticID int64 `json:"logistic_id,omitempty"`
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool `json:"enabled,omitempty"`
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64 `json:"shipping_fee,omitempty,string"`
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID int64 `json:"size_id,omitempty"`
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool `json:"is_free,omitempty"`
}

//=======================================================
// AddResponse
//=======================================================

type AddResponseItemVariation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku,omitempty"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name,omitempty"`
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price,omitempty,string"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock,omitempty"`
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status string `json:"status,omitempty"`
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int `json:"create_time,omitempty"`
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int `json:"update_time,omitempty"`
	// The original price of the variation in the listing currency.
	OriginalPrice float64 `json:"original_price,omitempty,string"`
	// The ID of discount activity the variation is currently in. One variation can only have one discount at a time. discount_id will be 0 if the variation has no discount applied.
	DiscountID int64 `json:"discount_id,omitempty"`
}

type AddResponseItemAttribute struct {
	// The Identify of each category.
	AttributeID int64 `json:"attribute_id,omitempty"`
	// The name of each attribute.
	AttributeName string `json:"attribute_name,omitempty"`
	// This is to indicate whether this attribute is mandantory.
	IsMandatory bool `json:"is_mandatory,omitempty"`
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttributeType string `json:"attribute_type,omitempty"`
	// The value of this item attribute.
	AtributeValue string `json:"atribute_value,omitempty"`
}

type AddResponseItemLogistic struct {
	// The identity of logistic channel.
	LogisticID int64 `json:"logistic_id,omitempty"`
	// The name of logistic.
	LogisticName string `json:"logistic_name,omitempty"`
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool `json:"enabled,omitempty"`
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64 `json:"shipping_fee,omitempty,string"`
	// If specify logistic fee_type is SIZE_SELECTION size_id is required.
	SizeID int64 `json:"size_id,omitempty"`
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool `json:"is_free,omitempty"`
	// Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
	EstimatedShippingFee float64 `json:"estimated_shipping_fee,omitempty,string"`
}

type AddResponseItemWholesale struct {
	// The min count of this tier wholesale.
	Min int `json:"min,omitempty"`
	// The max count of this tier wholesale.
	Max int `json:"max,omitempty"`
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64 `json:"unit_price,omitempty,string"`
}

type AddResponseItem struct {
	// Shopee's unique identifier for a shop.
	ShopID int64 `json:"shopid,omitempty"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku,omitempty"`
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED, UNLIST and BANNED.
	Status string `json:"status,omitempty"`
	// Name of the item in local language.
	Name string `json:"name,omitempty"`
	// Description of the item in local language.
	Description string `json:"description,omitempty"`
	// Image URLs of the item. It contains at most 9 URLs.
	Images []string `json:"images,omitempty"`
	// The three-digit code representing the currency unit used for the item in Shopee Listings.
	Currency string `json:"currency,omitempty"`
	// This is to indicate whether the item has variation(s).
	HasVariation bool `json:"has_variation,omitempty"`
	// The current price of the item in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price,omitempty,string"`
	// The current stock quantity of the item.
	Stock int `json:"stock,omitempty"`
	// Timestamp that indicates the date and time that the item was created.
	CreateTime int `json:"create_time,omitempty"`
	// Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
	UpdateTime int `json:"update_time,omitempty"`
	// the net weight of this item, the unit is KG.
	Weight float64 `json:"weight,omitempty,string"`
	// Could call shopee.item.GetCategories to get category detail.Related to result.categories.category_id.
	CategoryID int64 `json:"category_id,omitempty"`
	// The original price of the item in the listing currency.
	OriginalPrice float64 `json:"original_price,omitempty,string"`
	// The variation list of item.
	Variations []AddResponseItemVariation `json:"variations,omitempty"`
	//
	Attributes []AddResponseItemAttribute `json:"attributes,omitempty"`
	// The logistics list.
	Logistics []AddResponseItemLogistic `json:"logistics,omitempty"`
	// The wholesales tier list.
	Wholesales []AddResponseItemWholesale `json:"wholesales,omitempty"`
	// The sales volume of item.
	Sales int `json:"sales,omitempty"`
	// The page view of item.
	Views int `json:"views,omitempty"`
	// The conllection number of item.
	Likes int `json:"likes,omitempty"`
	// The length of package for this single item, the unit is CM
	PackageLength int `json:"package_length,omitempty"`
	// The width of package for this single item, the unit is CM
	PackageWidth int `json:"package_width,omitempty"`
	// The height of package for this single item, the unit is CM
	PackageHeight int `json:"package_height,omitempty"`
	// The guaranteed days to ship orders. For pre-order, please input value from 7 to 30; for non pre-order, please exclude this field and it will default to the respective standard per your shop location.(e.g. 3 for CrossBorder)
	DaysToShip int `json:"days_to_ship,omitempty"`
	// The rating star scores of this item.
	RatingStar float64 `json:"rating_star,omitempty,string"`
	// Count of comments for the item.
	CMTCount int `json:"cmt_count,omitempty"`
	// This indicates whether the item is secondhand.
	Condition string `json:"condition,omitempty"`
	// The ID of discount activity the item is currently in. One item can only have one discount at a time. discount_id will be 0 if the item has no discount applied, or item has variation.
	DiscountID int64 `json:"discount_id,omitempty"`
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool `json:"is_pre_order,omitempty"`
}

//=======================================================
// UnlistItemRequest
//=======================================================

type UnlistItemRequestItem struct {
	// Item's unique identifier.
	ItemID int64 `json:"item_id,omitempty"`
	// True: unlist this item; False: list this item.
	Unlist bool `json:"unlist,omitempty"`
}

//=======================================================
// UnlistItemResponse
//=======================================================

type UnlistItemResponseFailed struct {
	// Item's unique identifier.
	ItemID int64 `json:"item_id,omitempty"`
	// Error message.
	ErrorDesciption string `json:"error_desciption,omitempty"`
}

type UnlistItemResponseSuccess struct {
	// Item's unique identifier.
	ItemID int64 `json:"item_id,omitempty"`
	// True: item is unlisted; False: item is listed.
	Unlist bool `json:"unlist,omitempty"`
}

//=======================================================
// AddVariationsRequest
//=======================================================

type AddVariationsRequestVariation struct {
	// Name of the variation that belongs to the same item.A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name,omitempty"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock,omitempty"`
	// The current price of the variation in the listing currency.
	Price float64 `json:"price,omitempty,string"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku,omitempty"`
}

//=======================================================
// AddVariationsResponse
//=======================================================

type AddVariationsResponseVariation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku,omitempty"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name,omitempty"`
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price,omitempty,string"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock,omitempty"`
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status string `json:"status,omitempty"`
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int `json:"create_time,omitempty"`
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int `json:"update_time,omitempty"`
	// The original price of the variation in the listing currency.
	OriginalPrice float64 `json:"original_price,omitempty,string"`
}

//=======================================================
// GetItemsListResponse
//=======================================================

type GetItemsListResponseItemVariation struct {
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
}

type GetItemsListResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Shopee's unique identifier for a shop.
	ShopID int64 `json:"shopid,omitempty"`
	// The latest update time of the item.
	UpdateTime int `json:"update_time,omitempty"`
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, BANNED and UNLIST.
	Status string `json:"status,omitempty"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku,omitempty"`
	// The variation list of item
	Variations []GetItemsListResponseItemVariation `json:"variations,omitempty"`
	// Whether 2-tier variation structure is activated for this item
	Is2TierItem bool `json:"is_2_tier_item,omitempty"`
	// Only for TW seller. List of installments
	Tenures []int `json:"tenures,omitempty"`
}

//=======================================================
// GetItemDetailResponse
//=======================================================

type GetItemDetailResponseItemVariation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku,omitempty"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name,omitempty"`
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price,omitempty,string"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock,omitempty"`
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status string `json:"status,omitempty"`
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int `json:"create_time,omitempty"`
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int `json:"update_time,omitempty"`
	// The original price of the variation in the listing currency.
	OriginalPrice float64 `json:"original_price,omitempty,string"`
	// The ID of discount activity the variation is currently in. One variation can only have one discount at a time. discount_id will be 0 if the variation has no discount applied.
	DiscountID int64 `json:"discount_id,omitempty"`
	// Use this field to get the locked stock of variation by promotions.
	ReservedStock int64 `json:"reserved_stock,omitempty"`
	// Use this field to indicate the after-tax price of variation.
	InflatedPrice float64 `json:"inflated_price,omitempty,string"`
	// Use this field to indicate the after-tax original price of variation.
	InflatedOriginalPrice float64 `json:"inflated_original_price,omitempty,string"`
	// The settlement price of SIP item.
	SIPItemPrice float64 `json:"sip_item_price,omitempty,string"`
	// The strategy of creating sip_item_price. auto: automatically created; manual: manually created.
	PriceSource string `json:"price_source,omitempty"`
}

type GetItemDetailResponseItemAttribute struct {
	// The Identify of each category
	AttributeID int64 `json:"attribute_id,omitempty"`
	// The name of each attribute
	AttributeName string `json:"attribute_name,omitempty"`
	// This is to indicate whether this attribute is mandantory
	IsMandatory bool `json:"is_mandatory,omitempty"`
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttributeType string `json:"attribute_type,omitempty"`
	// The value of this item attribute.
	AttributeValue string `json:"attribute_value,omitempty"`
}

type GetItemDetailResponseItemLogistic struct {
	// The identity of logistic channel
	LogisticID int64 `json:"logistic_id,omitempty"`
	// The name of logistic
	LogisticName string `json:"logistic_name,omitempty"`
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool `json:"enabled,omitempty"`
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64 `json:"shipping_fee,omitempty,string"`
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID int64 `json:"size_id,omitempty"`
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool `json:"is_free,omitempty"`
	// Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
	EstimatedShippingFee float64 `json:"estimated_shipping_fee,omitempty,string"`
}

type GetItemDetailResponseItemWholesale struct {
	// The min count of this tier wholesale.
	Min int `json:"min,omitempty"`
	// The max count of this tier wholesale.
	Max int `json:"max,omitempty"`
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64 `json:"unit_price,omitempty,string"`
}

type GetItemDetailResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Shopee's unique identifier for a shop.
	ShopID int64 `json:"shopid,omitempty"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku,omitempty"`
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED, BANNED and UNLIST.
	Status string `json:"status,omitempty"`
	// Name of the item in local language.
	Name string `json:"name,omitempty"`
	// Description of the item in local language.
	Description string `json:"description,omitempty"`
	// Image URLs of the item. It contains at most 9 URLs.
	Images []string `json:"images,omitempty"`
	// The three-digit code representing the currency unit used for the item in Shopee Listings.
	Currency string `json:"currency,omitempty"`
	// This is to indicate whether the item has variation(s).
	HasVariaion bool `json:"has_variaion,omitempty"`
	// The current price of the item in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price,omitempty,string"`
	// The current stock quantity of the item.
	Stock int `json:"stock,omitempty"`
	// Timestamp that indicates the date and time that the item was created.
	CreateTime int `json:"create_time,omitempty"`
	// Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
	UpdateTime int `json:"update_time,omitempty"`
	// the net weight of this item, the unit is KG.
	Weight float64 `json:"weight,omitempty,string"`
	// Could call shopee.item.GetCategories to get category detail.Related to result.categories.category_id
	CategoryID int64 `json:"category_id,omitempty"`
	// The original price of the item in the listing currency.
	OriginalPrice float64 `json:"original_price,omitempty,string"`
	// The variation list of item
	Variations []GetItemDetailResponseItemVariation `json:"variations,omitempty"`
	//
	Attributes []GetItemDetailResponseItemAttribute `json:"attributes,omitempty"`
	// The logistics list.
	Logistics []GetItemDetailResponseItemLogistic `json:"logistics,omitempty"`
	// The wholesales tier list.
	Wholesales []GetItemDetailResponseItemWholesale `json:"wholesales,omitempty"`
	// The rating star scores of this item.
	RatingStar float64 `json:"rating_star,omitempty,string"`
	// Count of comments for the item.
	CMTCount int `json:"cmt_count,omitempty"`
	// The sales volume of item.
	Sales int `json:"sales,omitempty"`
	// The page view of item.
	Views int `json:"views,omitempty"`
	// The conllection number of item.
	Likes int `json:"likes,omitempty"`
	// The length of package for this single item, the unit is CM
	PackageLength float64 `json:"package_length,omitempty,string"`
	// The width of package for this single item, the unit is CM
	PackageWidth float64 `json:"package_width,omitempty,string"`
	// The height of package for this single item, the unit is CM
	PackageHeight float64 `json:"package_height,omitempty,string"`
	// The days to ship.
	DaysToShip int `json:"days_to_ship,omitempty"`
	// url of size chart image. Only particular categories support it.
	SizeChart string `json:"size_chart,omitempty"`
	// This indicates whether the item is secondhand.
	Condition string `json:"condition,omitempty"`
	// The ID of discount activity the item is currently in. One item can only have one discount at a time. discount_id will be 0 if the item has no discount applied, or item has variation.
	DiscountID int64 `json:"discount_id,omitempty"`
	// Whether 2-tier variation structure is activated for this item
	Is2TierItem bool `json:"is_2_tier_item,omitempty"`
	// Only for TW seller. List of installments
	Tenures []int `json:"tenures,omitempty"`
	// Use this field to get the locked stock of item by promotions.
	ReservedStock int `json:"reserved_stock,omitempty"`
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool `json:"is_pre_order,omitempty"`
}

//=======================================================
// UpdateItemRequest
//=======================================================

type UpdateItemRequestVariation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name,omitempty"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku,omitempty"`
}

type UpdateItemRequestAttribute struct {
	// related to shopee.item.GetAttributes result.attributes.attribute_id
	AttributesID int64 `json:"attributtes_id,omitempty"`
	// related to shopee.item.GetAttributes one of result.attributes.options
	Value string `json:"value,omitempty"`
}

type UpdateItemRequestWholesale struct {
	// The min count of this tier wholesale. If the wholesale is not the first one, the min count must equal to max count of last tier plus one.
	Min int `json:"min,omitempty"`
	// The max count of this tier wholesale.
	Max int `json:"max,omitempty"`
	// The current price of the wholesale in the listing currency. The price must be cheaper than original price. And if the wholesale is not the first one, the price must be cheaper than previous tier.'
	UnitPrice float64 `json:"unit_price,omitempty,string"`
}

type UpdateItemRequestLogistic struct {
	// related to shopee.logistics.GetLogistics result.logistics.logistic_id
	LogisticID int64 `json:"logistic_id,omitempty"`
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool `json:"enabled,omitempty"`
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64 `json:"shipping_fee,omitempty,string"`
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID int64 `json:"size_id,omitempty"`
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool `json:"is_free,omitempty"`
}

//=======================================================
// UpdateItemResponse
//=======================================================

type UpdateItemResponseItemVariation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku,omitempty"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	Name string `json:"name,omitempty"`
	// The current price of the variation in the listing currency.If item is in promotion, this value is discount price.
	Price float64 `json:"price,omitempty,string"`
	// The current stock quantity of the variation in the listing currency.
	Stock int `json:"stock,omitempty"`
	// Enumerated type that defines the current status of the variation. Applicable values: MODEL_NORMAL and MODEL_DELETED.
	Status string `json:"status,omitempty"`
	// Timestamp that indicates the date and time that the variation was created.
	CreateTime int `json:"create_time,omitempty"`
	// Timestamp that indicates the last time that there was a change in value of the variation, such as price/stock change.
	UpdateTime int `json:"update_time,omitempty"`
	// The original price of the variation in the listing currency.
	OriginalPirce float64 `json:"original_pirce,omitempty,string"`
	// The ID of discount activity the variation is currently in. One variation can only have one discount at a time. discount_id will be 0 if the variation has no discount applied.
	DiscountID int64 `json:"discount_id,omitempty"`
}

type UpdateItemResponseItemAttribute struct {
	// The Identify of each category
	AttributeID int64 `json:"attribute_id,omitempty"`
	// The name of each attribute
	AttributeName string `json:"attribute_name,omitempty"`
	// This is to indicate whether this attribute is mandantory
	IsMandatory bool `json:"is_mandatory,omitempty"`
	// Enumerated type that defines the type of the attribute. Applicable values: See Data Definition- AttributeType.
	AttributeType string `json:"attribute_type,omitempty"`
	// The value of this item attribute.
	AttribueValue string `json:"attribue_value,omitempty"`
}

type UpdateItemResponseItemLogistic struct {
	// The identity of logistic channel
	LogisticID int64 `json:"logistic_id,omitempty"`
	// The name of logistic
	LogisticName string `json:"logistic_name,omitempty"`
	// related to shopee.logistics.GetLogistics result.logistics.enabled only affect current item
	Enabled bool `json:"enabled,omitempty"`
	// Only needed when logistics fee_type = CUSTOM_PRICE.
	ShippingFee float64 `json:"shipping_fee,omitempty,string"`
	// If specify logistic fee_type is SIZE_SELECTION size_id is required
	SizeID int64 `json:"size_id,omitempty"`
	// when seller chooses this option, the shipping fee of this channel on item will be set to 0. Default value is False.
	IsFree bool `json:"is_free,omitempty"`
	// Estimated shipping fee calculated by weight. Don't exist if channel is no-integrated.
	EstimatedShippingFee float64 `json:"estimated_shipping_fee,omitempty,string"`
}

type UpdateItemResponseItemWholesale struct {
	// The min count of this tier wholesale.
	Min int `json:"min,omitempty"`
	// The max count of this tier wholesale.
	Max int `json:"max,omitempty"`
	// The current price of the wholesale in the listing currency.If item is in promotion, this price is useless.
	UnitPrice float64 `json:"unit_price,omitempty,string"`
}

type UpdateItemResponseItem struct {
	// Shopee's unique identifier for a shop.
	ShopID int64 `json:"shopid,omitempty"`
	// An item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku,omitempty"`
	// Enumerated type that defines the current status of the item. Applicable values: NORMAL, DELETED and BANNED.
	Status string `json:"status,omitempty"`
	// Name of the item in local language.
	Name string `json:"name,omitempty"`
	// Description of the item in local language.
	Description string `json:"description,omitempty"`
	// Image URLs of the item. It contains at most 9 URLs.
	Images []string `json:"images,omitempty"`
	// The three-digit code representing the currency unit used for the item in Shopee Listings.
	Currency string `json:"currency,omitempty"`
	// This is to indicate whether the item has variation(s).
	HasVariation bool `json:"has_variation,omitempty"`
	// The current price of the item in the listing currency. If item is in promotion, this value is discount price.
	Price float64 `json:"price,omitempty,string"`
	// The current stock quantity of the item.
	Stock int `json:"stock,omitempty"`
	// Timestamp that indicates the date and time that the item was created.
	CreateTime int `json:"create_time,omitempty"`
	// Timestamp that indicates the last time that there was a change in value of the item, such as price/stock change.
	UpdateTime int `json:"update_time,omitempty"`
	// the net weight of this item, the unit is KG.
	Weight float64 `json:"weight,omitempty,string"`
	// Could call shopee.item.GetCategories to get category detail.Related to result.categories.category_id
	CategoryID int64 `json:"category_id,omitempty"`
	// The original price of the item in the listing currency.
	OriginalPrice float64 `json:"original_price,omitempty,string"`
	// The variation list of item
	Variations []UpdateItemResponseItemVariation `json:"variations,omitempty"`
	//
	Attritube []UpdateItemResponseItemAttribute `json:"attritube,omitempty"`
	// The logistics list.
	Logistics []UpdateItemResponseItemLogistic `json:"logistics,omitempty"`
	// The wholesales tier list.
	Wholesales []UpdateItemResponseItemWholesale `json:"wholesales,omitempty"`
	// The rating star scores of this item.
	RatingStar float64 `json:"rating_star,omitempty,string"`
	// Count of comments for the item.
	CMTCount int `json:"cmt_count,omitempty"`
	// The sales volume of item.
	Sales int `json:"sales,omitempty"`
	// The page view of item.
	Views int `json:"views,omitempty"`
	// The conllection number of item.
	Likes int `json:"likes,omitempty"`
	// The length of package for this single item, the unit is CM
	PackageLength int `json:"package_length,omitempty"`
	// The width of package for this single item, the unit is CM
	PackageWidth int `json:"package_width,omitempty"`
	// The height of package for this single item, the unit is CM
	PackageHeight int `json:"package_height,omitempty"`
	// The guaranteed days to ship orders. Update value to less than 7 will default the value to the respective standard per your shop location and make this item non pre-order.(e.g. 3 for CrossBorder)
	DaysToShip int `json:"days_to_ship,omitempty"`
	// This indicates whether the item is secondhand.
	Condition string `json:"condition,omitempty"`
	// The ID of discount activity the item is currently in. One item can only have one discount at a time. discount_id will be 0 if the item has no discount applied, or item has variation.
	DiscountID int64 `json:"discount_id,omitempty"`
	// Use this field to identify whether the item is pre-order. Applicable value: true/false.
	IsPreOrder bool `json:"is_pre_order,omitempty"`
}

//=======================================================
// UpdatePriceResponse
//=======================================================

type UpdatePriceResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time,omitempty"`
	// Specify the revised price of the item.
	Price float64 `json:"price,omitempty,string"`
}

//=======================================================
// UpdateStockResponse
//=======================================================

type UpdateStockResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time,omitempty"`
	// Specify the updated stock quantity.
	Stock int `json:"stock,omitempty"`
}

//=======================================================
// UpdateVariationPriceResponse
//=======================================================

type UpdateVariationPriceResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// Specify the revised price of one variation of the item.
	Price float64 `json:"price,omitempty,string"`
}

//=======================================================
// UpdateVariationStockResponse
//=======================================================

type UpdateVariationStockResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// The time when price of the item is updated.
	ModifiedTime int `json:"modified_time,omitempty"`
	// Specify the updated stock quantity.
	Stock int `json:"stock,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
}

//=======================================================
// UpdatePriceBatchRequest
//=======================================================

type UpdatePriceBatchRequestItem struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
	// New price value for this item.
	Price int `json:"price,omitempty"`
}

//=======================================================
// UpdatePriceBatchResponse
//=======================================================

type UpdatePriceBatchResponseBatchResultFailure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int64 `json:"item_id,omitempty"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description,omitempty"`
}

type UpdatePriceBatchResponseBatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []string `json:"modifications,omitempty"`
	// Informations for failed stock updating.
	Failures []UpdatePriceBatchResponseBatchResultFailure `json:"failures,omitempty"`
}

//=======================================================
// UpdateStockBatchRequest
//=======================================================

type UpdateStockBatchRequestItem struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
	// New stock value for this item.
	Stock int `json:"stock,omitempty"`
}

//=======================================================
// UpdateStockBatchResponse
//=======================================================

type UpdateStockBatchResponseBatchResultFailure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int64 `json:"item_id,omitempty"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description,omitempty"`
}

type UpdateStockBatchResponseBatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []string `json:"modifications,omitempty"`
	// Informations for failed stock updating.
	Failures []UpdateStockBatchResponseBatchResultFailure `json:"failures,omitempty"`
}

//=======================================================
// UpdateVariationPriceBatchRequest
//=======================================================

type UpdateVariationPriceBatchRequestVariation struct {
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int64 `json:"variation_id,omitempty"`
	// New price value of this variation.
	Price int `json:"price,omitempty"`
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
}

//=======================================================
// UpdateVariationPriceBatchResponse
//=======================================================

type UpdateVariationPriceBatchResponseBatchResultFailure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int64 `json:"item_id,omitempty"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
}

type UpdateVariationPriceBatchResponseBatchResultModification struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
}

type UpdateVariationPriceBatchResponseBatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []UpdateVariationPriceBatchResponseBatchResultModification `json:"modifications,omitempty"`
	// Informations for failed stock updating.
	Failures []UpdateVariationPriceBatchResponseBatchResultFailure `json:"failures,omitempty"`
}

//=======================================================
// UpdateVariationStockBatchRequest
//=======================================================

type UpdateVariationStockBatchRequestVariation struct {
	// Shopee's unique identifier for a variation of an item. Please input the variation_id of a variation to be changed. The variation_id and item_id pair must be matched in order to perform the update.
	VariationID int64 `json:"variation_id,omitempty"`
	// New stock value of this variation.
	Stock int `json:"stock,omitempty"`
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
}

//=======================================================
// UpdateVariationStockBatchResponse
//=======================================================

type UpdateVariationStockBatchResponseBatchResultFailure struct {
	// Shopee's unique identifier for an item. Indicating items which failed to update.
	ItemID int64 `json:"item_id,omitempty"`
	// Detailed information for the failed updating.
	ErrorDescription string `json:"error_description,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
}

type UpdateVariationStockBatchResponseBatchResultModification struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
}

type UpdateVariationStockBatchResponseBatchResult struct {
	// List of item_id which have been updated successfully.
	Modifications []UpdateVariationStockBatchResponseBatchResultModification `json:"modifications,omitempty"`
	// Informations for failed stock updating.
	Failures []UpdateVariationStockBatchResponseBatchResultFailure `json:"failures,omitempty"`
}

//=======================================================
// InitTierVariationRequest
//=======================================================

type InitTierVariationRequestTierVariation struct {
	// Tier variation name.
	Name string `json:"name,omitempty"`
	// Tier variation value options list. Option length should be under 20. Quantity of combinations of all 2 tier options is up to 50.
	Options []string `json:"options,omitempty"`
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string `json:"images_url,omitempty"`
}

type InitTierVariationRequestVariation struct {
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex []int `json:"tier_index,omitempty"`
	// Stock value of this variation item. The original variation stock will be override when calling this API to initialize 2-tier structure for an existed item. 0 stock will make this variation a greyout option for buyer.
	Stock int `json:"stock,omitempty"`
	// Price value of this variation item. The original variation price will be override when calling this API to initialize 2-tier structure for an existed item.
	Price float64 `json:"price,omitempty,string"`
	// SKU string of this variation.SKU length should be under 100.
	VariationSKU string `json:"variation_sku,omitempty"`
}

//=======================================================
// InitTierVariationResponse
//=======================================================

type InitTierVariationResponseVariation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int `json:"tier_index,omitempty"`
	// The identity of the variation.
	VariationID int64 `json:"variation_id,omitempty"`
}

//=======================================================
// AddTierVariationRequest
//=======================================================

type AddTierVariationRequestVariation struct {
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex []int `json:"tier_index,omitempty"`
	// Stock value of this variation item. 0 stock will make this variation a greyout option for buyer.
	Stock int `json:"stock,omitempty"`
	// Price value of this variation item.
	Price float64 `json:"price,omitempty,string"`
	// SKU string of this variation item.
	VariationSKU string `json:"variation_sku,omitempty"`
}

//=======================================================
// AddTierVariationResponse
//=======================================================

type AddTierVariationResponseVariation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int `json:"tier_index,omitempty"`
	// The identity of the variation.
	VariationID int64 `json:"variation_id,omitempty"`
}

//=======================================================
// GetVariationResponse
//=======================================================

type GetVariationResponseTierVariation struct {
	// Tier variation name.
	Name string `json:"name,omitempty"`
	// Tier variation value options list.
	Options []string `json:"options,omitempty"`
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string `json:"images_url,omitempty"`
}

type GetVariationResponseVariation struct {
	// Unique identifier of the variation.
	VariationID string `json:"variation_id,omitempty"`
	// A list of tier variation combination index, which indicates variation's option position in tier_variation['options'] list. e.g. [0,1] means tier variation 1 option 1 and tier variation 2 option 2.
	TierIndex []int `json:"tier_index,omitempty"`
}

//=======================================================
// UpdateTierVariationListRequest
//=======================================================

type UpdateTierVariationListRequestTierVariation struct {
	// Tier variation name.
	Name string `json:"name,omitempty"`
	// Tier variation value options list. Lenght should be under 20. Combinations of 2 level options should be under 50.
	Options []string `json:"options,omitempty"`
	// Tier variation images. Can only be applied for the first level options. Urls sequence match the options sequence and urls number cannot exceed options number.
	ImagesURL []string `json:"images_url,omitempty"`
}

//=======================================================
// UpdateTierVariationIndexRequest
//=======================================================

type UpdateTierVariationIndexRequestVariation struct {
	// A list of tier variation indexes, which indicate variation's options in tier_variation['options'] list.
	TierIndex []int `json:"tier_index,omitempty"`
	// The identity of product item variation.
	VariationID []int `json:"variation_id,omitempty"`
}

//=======================================================
// BoostItemResponse
//=======================================================

type BoostItemResponseBatchResultFailure struct {
	// to indicate error type.
	ErrorCode string `json:"error_code,omitempty"`
	// Failed item id.
	ID int64 `json:"id,omitempty"`
	// error description
	Description string `json:"description,omitempty"`
}

type BoostItemResponseBatchResult struct {
	// A list of item ids which have been boosted successfully.
	Successes []int `json:"successes,omitempty"`
	// A list of failed-to-boost items, including error details.
	Failures []BoostItemResponseBatchResultFailure `json:"failures,omitempty"`
}

//=======================================================
// GetBoostedItemResponse
//=======================================================

type GetBoostedItemResponseItem struct {
	// boosted items' id.
	ItemID int64 `json:"item_id,omitempty"`
	// Cooldown_second time is four hours after boost. After four hours you can boost this item again.
	CooldownSecond int `json:"cooldown_second,omitempty"`
}

//=======================================================
// GetPromotionInfoResponse
//=======================================================

type GetPromotionInfoResponseItemPromotion struct {
	//
	PromotionType string `json:"promotion_type,omitempty"`
	// The ID of promotion.
	PromotionID int64 `json:"promotion_id,omitempty"`
	// ID of the variation that belongs to the same item.
	VariationID int64 `json:"variation_id,omitempty"`
	// Start timestamp of promotion.
	StartTime int `json:"start_time,omitempty"`
	// End timestamp of promotion.
	EndTime int `json:"end_time,omitempty"`
	// The promotion price of item.
	PromotionPrice float64 `json:"promotion_price,omitempty,string"`
	// The Locked stock of item by promotion.
	ReservedStock int `json:"reserved_stock,omitempty"`
	// The sold out timestamp of promotion stock.
	StockoutTime int `json:"stockout_time,omitempty"`
	// The stage at which the promotion goes. Available values: ongoing/upcoming.
	Staging string `json:"staging,omitempty"`
}

type GetPromotionInfoResponseItemError struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
	// Error Message.
	ErrorMsg string `json:"error_msg,omitempty"`
}

type GetPromotionInfoResponseItem struct {
	// Shopee's unique identifier for an item. Please input the item_id of an item to be changed.
	ItemID int64 `json:"item_id,omitempty"`
	// Promotion information list.
	Promotions []GetPromotionInfoResponseItemPromotion `json:"promotions,omitempty"`
	// The list of error items.
	Errors []GetPromotionInfoResponseItemError `json:"errors,omitempty"`
}

//=======================================================
// UploadImgResponse
//=======================================================

type UploadImgResponseImage struct {
	// origin image url
	ImageURL string `json:"image_url,omitempty"`
	// Shopee image url
	ShopeeImageURL string `json:"shopee_image_url,omitempty"`
}

//=======================================================
// AddDiscountRequest
//=======================================================

type AddDiscountRequestItemVariation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int64 `json:"variation_id,omitempty"`
	// The discount price of the item.
	VariationPromotionPrice float64 `json:"variation_promotion_price,omitempty,string"`
}

type AddDiscountRequestItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	//
	Variations []AddDiscountRequestItemVariation `json:"variations,omitempty"`
	// The discount price of the item. If the item has no variation, this param is necessary.
	ItemPromotionPrice float64 `json:"item_promotion_price,omitempty,string"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit,omitempty"`
}

//=======================================================
// AddDiscountItemRequest
//=======================================================

type AddDiscountItemRequestItemVariation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int64 `json:"variation_id,omitempty"`
	// The discount price of the item.
	VariationPromotionPrice float64 `json:"variation_promotion_price,omitempty,string"`
}

type AddDiscountItemRequestItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	//
	Variations []AddDiscountItemRequestItemVariation `json:"variations,omitempty"`
	// The discount price of the item. If the item has no variation, this param is necessary.
	ItemPromotionPrice float64 `json:"item_promotion_price,omitempty,string"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit,omitempty"`
}

//=======================================================
// GetDiscountDetailResponse
//=======================================================

type GetDiscountDetailResponseItemVariation struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// Name of the variation that belongs to the same item.
	VariationName string `json:"variation_name,omitempty"`
	// The original price before discount of the variation.
	VariationOriginalPrice float64 `json:"variation_original_price,omitempty,string"`
	// The discount price of the variation.
	VariationPromotionPrice float64 `json:"variation_promotion_price,omitempty,string"`
	// The current stock quantity of the variation.
	VariationStock int `json:"variation_stock,omitempty"`
}

type GetDiscountDetailResponseItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Name of the item in local language.
	ItemName string `json:"item_name,omitempty"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit,omitempty"`
	// The original price before discount of the item. If there is variation, this value is 0.
	ItemOriginalPrice float64 `json:"item_original_price,omitempty,string"`
	// The discount price of the item. If there is variation, this value is 0.
	ItemPromotionPrice float64 `json:"item_promotion_price,omitempty,string"`
	// The current stock quantity of the item.
	Stock int `json:"stock,omitempty"`
	//
	Variations []GetDiscountDetailResponseItemVariation `json:"variations,omitempty"`
}

//=======================================================
// GetDiscountsListResponse
//=======================================================

type GetDiscountsListResponseDiscount struct {
	// Shopee's unique identifier for a discount activity.
	DiscountID int64 `json:"discount_id,omitempty"`
	// Title of the discount.
	DiscountName string `json:"discount_name,omitempty"`
	// The time when discount activity start.
	StartTime int `json:"start_time,omitempty"`
	// The time when discount activity end.
	EndTime int `json:"end_time,omitempty"`
	// The status of discount, applicable values: expired, ongoing, upcoming.
	Status string `json:"status,omitempty"`
}

//=======================================================
// UpdateDiscountItemsRequest
//=======================================================

type UpdateDiscountItemsRequestItemVariation struct {
	// Shopee's unique identifier for a variation of an item. If there is no variation of this item, you don't need to input this param. Dafault is 0.
	VariationID int64 `json:"variation_id,omitempty"`
	// The discount price of the item.
	VariationPromotionPrice float64 `json:"variation_promotion_price,omitempty,string"`
}

type UpdateDiscountItemsRequestItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// The max number of this product in the promotion price.
	PurchaseLimit int `json:"purchase_limit,omitempty"`
	// The discount price of the item.
	ItemOriginalPrice float64 `json:"item_original_price,omitempty,string"`
	//
	Variations []UpdateDiscountItemsRequestItemVariation `json:"variations,omitempty"`
}

//=======================================================
// GetOrdersListResponse
//=======================================================

type GetOrdersListResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// Enumerated type that defines the current status of the order. Applicable values: See Data Definition- OrderStatus.
	OrderStatus string `json:"order_status,omitempty"`
	// Timestamp that indicates the last time that there was a change in value of order, such as order status changed from 'Paid' to 'Completed'.
	UpdateTime int `json:"update_time,omitempty"`
}

//=======================================================
// GetOrdersByStatusResponse
//=======================================================

type GetOrdersByStatusResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// Enumerated type that defines the current status of the order. Applicable values: See Data Definition- OrderStatus.
	OrderStatus string `json:"order_status,omitempty"`
	// Timestamp that indicates the last time that there was a change in value of order, such as order status changed from 'Paid' to 'Completed'.
	UpdateTime int `json:"update_time,omitempty"`
}

//=======================================================
// GetOrderDetailsResponse
//=======================================================

type GetOrderDetailsResponseOrderAddress struct {
	// Recipient's name for the address.
	Name string `json:"name,omitempty"`
	// Recipient's phone number input when order was placed.
	Phone string `json:"phone,omitempty"`
	// The town of the recipient's address. Whether there is a town will depend on the region and/or country.
	Town string `json:"town,omitempty"`
	// The district of the recipient's address. Whether there is a town will depend on the region and/or country.
	District string `json:"district,omitempty"`
	// The city of the recipient's address. Whether there is a town will depend on the region and/or country.
	City string `json:"city,omitempty"`
	// The state/province of the recipient's address. Whether there is a town will depend on the region and/or country.
	State string `json:"state,omitempty"`
	// The two-digit code representing the country of the Recipient.
	Country string `json:"country,omitempty"`
	// Recipient's postal code.
	Zipcode string `json:"zipcode,omitempty"`
	// The full address of the recipient, including country, state, even street, and etc.
	FullAddress string `json:"full_address,omitempty"`
}

type GetOrderDetailsResponseOrderItem struct {
	// ID of item
	ItemID int64 `json:"item_id,omitempty"`
	// Name of item
	ItemName string `json:"item_name,omitempty"`
	// A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku,omitempty"`
	// ID of the variation that belongs to the same item.
	VariationID int64 `json:"variation_id,omitempty"`
	// Name of the variation that belongs to the same item.
	// A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	VariationName string `json:"variation_name,omitempty"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku,omitempty"`
	// The number of identical items purchased at the same time by the same buyer from one listing/item.
	VariationQuantityPurchased int `json:"variation_quantity_purchased,omitempty"`
	// The original price of the item in the listing currency.
	VariationOriginalPrice float64 `json:"variation_original_price,omitempty,string"`
	// The after-discount price of the item in the listing currency. If there is no discount, this value will be same as that of variation_original_price.
	// In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/variation level. Due to technical restriction, the value will return the price before bundle deal if we don't configure it to 0. Please call GetEscrowDetails if you want to calculate item-level discounted price for bundle deal item.
	VariationDiscountedPrice float64 `json:"variation_discounted_price,omitempty,string"`
	// This value indicates whether buyer buy the order item in wholesale price.
	IsWholesale bool `json:"is_wholesale,omitempty"`
	// The weight of the item
	Weight float64 `json:"weight,omitempty"`
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool `json:"is_add_on_deal,omitempty"`
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool `json:"is_main_item,omitempty"`
	// The unique identity of an addon deal.
	AddOnDealID int64 `json:"add_on_deal_id,omitempty"`
	// The type of the promotion,
	PromotionType string `json:"promotion_type,omitempty"`
	// The ID of the promotion.
	PromotionID int64 `json:"promotion_id,omitempty"`
}

type GetOrderDetailsResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The two-digit code representing the country where the order was made.
	Country string `json:"country,omitempty"`
	// The three-digit code representing the currency unit for which the order was paid.
	Currency string `json:"currency,omitempty"`
	// This value indicates whether the order was a COD (cash on delivery) order.
	COD bool `json:"cod,omitempty"`
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no,omitempty"`
	// Shipping preparation time set by the seller when listing item on Shopee.
	DaysToShip int `json:"days_to_ship,omitempty"`
	// This object contains detailed breakdown for the recipient address.
	RecipientAddress GetOrderDetailsResponseOrderAddress `json:"recipient_address,omitempty"`
	// The estimated shipping fee is an estimation calculated by Shopee based on specific logistics courier's standard.
	EstimatedShippingFee float64 `json:"estimated_shipping_fee,omitempty,string"`
	// The actual shipping cost of the order if available from external logistics partners.
	ActualShippingCost float64 `json:"actual_shipping_cost,omitempty,string"`
	// The total amount paid by the buyer for the order. This amount includes the total sale price of items, shipping cost beared by buyer; and offset by Shopee promotions if applicable. This value will only return after the buyer has completed payment for the order.
	TotalAmount float64 `json:"total_amount,omitempty,string"`
	// The total amount that the seller is expected to receive for the order. This amount includes buyer paid order amount (total_amount), all forms of Shopee platform subsidy; and offset by any cost and commission incurred.
	EscrowAmount float64 `json:"escrow_amount,omitempty,string"`
	// Enumerated type that defines the current status of the order.
	OrderStatus string `json:"order_status,omitempty"`
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string `json:"shipping_carrier,omitempty"`
	// The payment method that the buyer selected to pay for the order.
	// Applicable values: See Data Definition- Payment Methods.
	PaymentMethod string `json:"payment_method,omitempty"`
	// Only work for cross-border order.This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	GoodsToDeclare bool `json:"goods_to_declare,omitempty"`
	// Message to seller.
	MessageToSeller string `json:"message_to_seller,omitempty"`
	// The note seller made for own reference.
	Note string `json:"note,omitempty"`
	// Update time for the note.
	NoteUpdateTime int `json:"note_update_time,omitempty"`
	// Timestamp that indicates the date and time that the order was created.
	CreateTime int `json:"create_time,omitempty"`
	// Timestamp that indicates the last time that there was a change in value of order, such as order status changed from 'Paid' to 'Completed'.
	UpdateTime int `json:"update_time,omitempty"`
	// This object contains the detailed breakdown for the result of this API call.
	Items []GetOrderDetailsResponseOrderItem `json:"items,omitempty"`
	// The time when the order status is updated from UNPAID to PAID. This value is NULL when order is not paid yet.
	PayTime int `json:"pay_time,omitempty"`
	// For Indonesia orders only. The name of the dropshipper.
	Dropshipper string `json:"dropshipper,omitempty"`
	// Last 4 digits of the credit card
	CreditCardNumber string `json:"credit_card_number,omitempty"`
	// The name of buyer
	BuyerUsername string `json:"buyer_username,omitempty"`
	// The phone number of dropshipper
	DropshipperPhone string `json:"dropshipper_phone,omitempty"`
	// The deadline to ship out the parcel.
	ShipByDate int `json:"ship_by_date,omitempty"`
	// To indicate whether this order is split to fullfil order(forder) level. Call GetForderInfo if it's "true".
	IsSplitUp bool `json:"is_split_up,omitempty"`
	// Cancel reason from buyer.
	BuyerCancelReason string `json:"buyer_cancel_reason,omitempty"`
	// Could be one of buyer, seller or system
	CancelBy string `json:"cancel_by,omitempty"`
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
	// Use this field to get reason for buyer, seller, and system cancellation.
	CancelReason string `json:"cancel_reason,omitempty"`
	// Cross-border tax imposed by the Indonesian government on sellers.
	EscrowTax float64 `json:"escrow_tax,omitempty,string"`
	// Use this filed to judge whether the actual_shipping_fee is confirmed.
	IsActualShippingFeeConfirmed bool `json:"is_actual_shipping_fee_confirmed,omitempty"`
	// Buyer's CPF number for taxation and invoice purposes. Only for Brazil order.
	BuyerCPFID string `json:"buyer_cpf_id,omitempty"`
	// Use this field to indicate the order is fulfilled by shopee or seller. Applicable values: fulfilled_by_shopee, fulfilled_by_cb_seller, fulfilled_by_local_seller.
	OrderFlag string `json:"order_flag,omitempty"`
	// The last-mile tracking number. Only for Cross Board BR seller.
	LMTN string `json:"lm_tn,omitempty"`
}

//=======================================================
// GetEscrowDetailsResponse
//=======================================================

type GetEscrowDetailsResponseOrderIncomeDetail struct {
	// The three-digit code representing the currency unit used for all transactional amount under
	LocalCurrency string `json:"local_currency,omitempty"`
	// The total amount paid by the buyer for the order. This amount includes the total sale price of items, shipping cost beared by buyer; and offset by Shopee promotions if applicable.
	TotalAmount float64 `json:"total_amount,omitempty,string"`
	// Final value of coins used by seller for the order.
	Coin float64 `json:"coin,omitempty,string"`
	// Final value of voucher provided by Shopee for the order.
	Voucher float64 `json:"voucher,omitempty,string"`
	// Final value of voucher provided by Seller for the order.
	VoucherSeller float64 `json:"voucher_seller,omitempty,string"`
	// Final sum of each item Shopee discount of a specific order. This amount will rebate to seller.
	SellerRebate float64 `json:"seller_rebate,omitempty,string"`
	// The final shipping cost of order . For Non-integrated logistics channel is 0.
	ActualShippingCost float64 `json:"actual_shipping_cost,omitempty,string"`
	// The platform shipping subsidy to the seller
	ShippingFeeRebate float64 `json:"shipping_fee_rebate,omitempty,string"`
	// The commission fee charged by Shopee platform if applicable.
	CommissionFee float64 `json:"commission_fee,omitempty,string"`
	// The voucher code or promotion code the buyer used.
	VoucherCode float64 `json:"voucher_code,omitempty,string"`
	// The voucher name or promotion name the buyer used.
	VoucherName float64 `json:"voucher_name,omitempty,string"`
	// The total amount that the seller is expected to receive for the order and will change before order completed. escrow_amount=total_amount+voucher+credit_card_promotion+seller_rebate+coin-commission_fee-credit_card_transaction_fee-cross_border_tax-service_fee-buyer_shopee_kredit-seller_coin_cash_back+final_shipping_fee-seller_return_refund_amount.
	EscrowAmount float64 `json:"escrow_amount,omitempty,string"`
	// Amount incurred by Buyer for purchasing items outside of home country. Amount may change after Return Refund.
	CroossBorderTax float64 `json:"crooss_border_tax,omitempty,string"`
	// Include buyer transaction fee and seller transaction fee.
	CreditCardTransactionFee float64 `json:"credit_card_transaction_fee,omitempty,string"`
	// Amount charged by Shopee to seller for additional services.
	ServiceFee float64 `json:"service_fee,omitempty,string"`
	// Amount charged by Shopee to Buyer for using ShopeeKredit for the order. Currently only applicable in ID.
	BuyerShopeeKredit float64 `json:"buyer_shopee_kredit,omitempty,string"`
	// Value of coins provided by Seller for purchasing with his or her store for the order.
	SellerCoinCashBack float64 `json:"seller_coin_cash_back,omitempty,string"`
	// Final adjusted amount that seller has to bear as part of escrow. This amount could be negative or positive.
	FinalShippingFee float64 `json:"final_shipping_fee,omitempty,string"`
	// Amount returned to Seller in the event of partial return.
	SellerReturnRefundAmount float64 `json:"seller_return_refund_amount,omitempty,string"`
	// The amount offset via payment promotion. May include bank payment promotion and Shopee payment promotion.
	CreditCardPromotion float64 `json:"credit_card_promotion,omitempty,string"`
	// True means incoming won't change any more.
	IsCompleted bool `json:"is_completed,omitempty"`
	// Use this field to fetch the list of voucher codes.
	VoucherCodeList []string `json:"voucher_code_list,omitempty"`
	// The transaction_fee of seller.
	SellerTransactionFee float64 `json:"seller_transaction_fee,omitempty,string"`
	// The transaction_fee of buyer.
	BuyerTransactionFee float64 `json:"buyer_transaction_fee,omitempty,string"`
}

type GetEscrowDetailsResponseOrderBankAccount struct {
	// Name of the seller's receiving bank
	BankName string `json:"bank_name,omitempty"`
	// Account number of the seller's receiving bank
	BankAccountNumber string `json:"bank_account_number,omitempty"`
	// The two-digit code representing country of the seller's receiving bank account
	BankAccountCountry string `json:"bank_account_country,omitempty"`
}

type GetEscrowDetailsResponseOrderActivityItem struct {
	// ID of item.
	ItemID int64 `json:"item_id,omitempty"`
	// ID of the variation that belongs to the same item.
	VariationID int64 `json:"variation_id,omitempty"`
	// The number of identical items purchased at the same time by the same buyer from one listing/item.
	QuantityPurchased int `json:"quantity_purchased,omitempty"`
	// The price used to participate activity. E.g. itemA original price is $10, promo price is $9, and bundle deal is buy 2 get 20% off equals to $14.4. The original_price value will be $9 in this case.
	OriginalPrice float64 `json:"original_price,omitempty,string"`
}

type GetEscrowDetailsResponseOrderActivity struct {
	// ID of activity.
	ActivityID int64 `json:"activity_id,omitempty"`
	// Type of activity. Currently only one type: bundle_deal
	ActivityType string `json:"activity_type,omitempty"`
	// The original TOTAL price of ALL items in one activity(e.g. bundle deal. Define by activity_id) in the listing currency.
	OriginalPrice float64 `json:"original_price,omitempty,string"`
	// The after-discocunt TOTAL price of ALL items in one activity(e.g. bundle deal. Define by activity_id) in the listing currency.
	DiscountedPrice float64 `json:"discounted_price,omitempty,string"`
	// This object contains the items in this activity.
	Items []GetEscrowDetailsResponseOrderActivityItem `json:"items,omitempty"`
}

type GetEscrowDetailsResponseOrderItem struct {
	// ID of item
	ItemID int64 `json:"item_id,omitempty"`
	// Name of item
	ItemName string `json:"item_name,omitempty"`
	// A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.
	ItemSKU string `json:"item_sku,omitempty"`
	// ID of the variation that belongs to the same item.
	VariationID int64 `json:"variation_id,omitempty"`
	// Name of the variation that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.
	VariationName string `json:"variation_name,omitempty"`
	// A variation SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.
	VariationSKU string `json:"variation_sku,omitempty"`
	// This value indicates the number of identical items purchased at the same time by the same buyer from one listing/item.
	QuantityPurchased int `json:"quantity_purchased,omitempty"`
	// The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.
	OriginalPrice float64 `json:"original_price,omitempty,string"`
	// The after-discount price of the item in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1. If there is no discount, this value will be the same as that of original_price.
	// In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/variation level. Due to technical restriction, the value will return the price before bundle deal if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the bundle deal discount breakdown on item level.
	DiscountedPrice float64 `json:"discounted_price,omitempty,string"`
	// The offset of this item when the buyer consumed Shopee Coins upon checkout.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	DiscountFromCoin float64 `json:"discount_from_coin,omitempty,string"`
	// The offset of this item when the buyer use Shopee voucher.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	DiscountFromVoucher float64 `json:"discount_from_voucher,omitempty,string"`
	// The offset of this item when the buyer use seller-specific voucher.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	DiscountFromVoucherSeller float64 `json:"discount_from_voucher_seller,omitempty,string"`
	// Platform subsidy to the seller for this item.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	SellerRebate float64 `json:"seller_rebate,omitempty,string"`
	// This value indicates the actual price the buyer pay.
	// In case of bundle deal item, this value will return 0 as by design bundle deal discount will not be breakdown to item/variation level. Due to technical restriction, the value will return the price before bundle deal if we don't configure it to 0. Please use the value under "income_details" and "activity" to calculate the bundle deal discount breakdown on item level.
	DealPrice float64 `json:"deal_price,omitempty,string"`
	// This value indicate the offset via credit card promotion.
	// In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0. Please use the value under "income_details" and "activity" to calculate the breakdown on item level.
	CreditCardPromotion float64 `json:"credit_card_promotion,omitempty,string"`
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool `json:"is_add_on_deal,omitempty"`
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool `json:"is_main_item,omitempty"`
	// The unique identity of an addon deal.
	AddOnDealID int64 `json:"add_on_deal_id,omitempty"`
}

type GetEscrowDetailsResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The two-digit code representing the country where the order was made.
	Country string `json:"country,omitempty"`
	// This object contains detailed income breakdown for the order.
	IncomeDetails GetEscrowDetailsResponseOrderIncomeDetail `json:"income_details,omitempty"`
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string `json:"shipping_carrier,omitempty"`
	// The three-digit code representing the currency unit of total order amount (escorw_amount) at the point of payment to the seller.
	EscrowCurrency string `json:"escrow_currency,omitempty"`
	// The exchange rate used by Shopee to convert local_currency to escrow_currency.
	ExchangeRate string `json:"exchange_rate,omitempty"`
	// The payment channel that the seller selected to receive escrow for the order.
	EscrowChannel string `json:"escrow_channel,omitempty"`
	// The unique identifier for a payee by the 3rd party payment service provider selected in escrow_channel.
	PayeeID string `json:"payee_id,omitempty"`
	// This object contains detailed breakdown for bank account of the seller if selected escorw_channel is Bank Transfer.
	BankAccount GetEscrowDetailsResponseOrderBankAccount `json:"bank_account,omitempty"`
	// This object contains the detailed breakdown for all the items in this order, including regular items(non-activity) and activity items.
	Items []GetEscrowDetailsResponseOrderItem `json:"items,omitempty"`
	// This object contains the activity in this order.
	Activity []GetEscrowDetailsResponseOrderActivity `json:"activity,omitempty"`
}

//=======================================================
// GetForderInfoResponse
//=======================================================

type GetForderInfoResponseForderLog struct {
	// The time when logistics info has been updated.
	Ctime int `json:"ctime,omitempty"`
	// The order logistics tracking info.
	Description string `json:"description,omitempty"`
}

type GetForderInfoResponseForderItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// The number of identical items/variations purchased at the same time by the same buyer from one listing/item.
	Num int `json:"num,omitempty"`
	// The original price of the item in the listing currency.
	ItemPrice float64 `json:"item_price,omitempty,string"`
	// The original price of the variation in the listing currency.
	VariationPrice float64 `json:"variation_price,omitempty,string"`
}

type GetForderInfoResponseForderLogisticsInfo struct {
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string `json:"shipping_carrier,omitempty"`
	// Only work for cross-border order. This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	GoodsToDeclare bool `json:"goods_to_declare,omitempty"`
	// Only work for cross-border order. This code is required at some sorting hub. Please ensure the service_code is INCLUDED on your shipping label, otherwise the parcel cannot be processed by the warehouse. If you didn't retrieve service_code after you first called this API, please try few more times within 30 minutes.
	ServiceCode string `json:"service_code,omitempty"`
	// The tracking number of fullfill order assigned by the shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no,omitempty"`
}

type GetForderInfoResponseForder struct {
	// The unique identifier for a fulfill order.
	ForderID string `json:"forder_id,omitempty"`
	// The fulfill order logistics status. Applicable values: See Data Definition - LogisticsStatus.
	Status string `json:"status,omitempty"`
	// Logistics tracking info.
	TrackingLog []GetForderInfoResponseForderLog `json:"tracking_log,omitempty"`
	// The items included in this fulfill order.
	Items []GetForderInfoResponseForderItem `json:"items,omitempty"`
	//
	LogisticsInfo []GetForderInfoResponseForderLogisticsInfo `json:"logistics_info,omitempty"`
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
	// The last-mile tracking number. Only for Cross Board BR seller.
	LMTN string `json:"lm_tn,omitempty"`
}

//=======================================================
// GetEscrowReleasedOrdersResponse
//=======================================================

type GetEscrowReleasedOrdersResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// Order's escrow amount.
	PayoutAmount float64 `json:"payout_amount,omitempty,string"`
	// Timestamp of escrow amount transaction finished.
	EscrowReleaseTime int `json:"escrow_release_time,omitempty"`
}

//=======================================================
// SplitOrderRequest
//=======================================================

type SplitOrderRequestParcel struct {
	// Itemids that will be put into a fullfillment order.
	ItemID int64 `json:"item_id,omitempty"`
	// Variation_id that will be put into a fulfillment order, if no variation please input variation_id:0.
	VariationID int64 `json:"variation_id"`
}

//=======================================================
// SplitOrderResponse
//=======================================================

type SplitOrderResponseForderItem struct {
	// Shopee's unique identifier for an item.
	ItemID int64 `json:"item_id,omitempty"`
}

type SplitOrderResponseForder struct {
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
	// Item information contained in fulfillment orders.Number of items must be greater than or equal to 2. eg.[[{"item_id": 123}],[{"item_id": 456}]]
	Items []SplitOrderResponseForderItem `json:"items,omitempty"`
}

//=======================================================
// GetUnbindOrderListResponse
//=======================================================

type GetUnbindOrderListResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The Shopee logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.
	LogisticStatus string `json:"logistic_status,omitempty"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
}

//=======================================================
// GetLogisticsResponse
//=======================================================

type GetLogisticsResponseLogisticSize struct {
	// The identity of size.
	SizeID int64 `json:"size_id,omitempty"`
	// The name of size.
	Name int `json:"name,omitempty"`
	// The pre-defined shipping fee for the specific size.
	DefaultPrice int `json:"default_price,omitempty"`
}

type GetLogisticsResponseLogisticLimit struct {
	// The max weight for an item on this logistic channel.If the value is 0 or null, that means there is no limit.
	ItemMaxWeight float64 `json:"item_max_weight,omitempty,string"`
	// The min weight for an item on this logistic channel. If the value is 0 or null, that means there is no limit.
	ItemMinWeight float64 `json:"item_min_weight,omitempty,string"`
}

type GetLogisticsResponseLogisticDimension struct {
	// The max height limit.
	Height float64 `json:"height,omitempty,string"`
	// The max width limit.
	Width float64 `json:"width,omitempty,string"`
	// The max length limit.
	Length float64 `json:"length,omitempty,string"`
	// The unit for the limit.
	Unit string `json:"unit,omitempty"`
}

type GetLogisticsResponseLogistic struct {
	// The identity of logistic channel
	LogisticID int64 `json:"logistic_id,omitempty"`
	// The name of logistic channel
	LogisticName string `json:"logistic_name,omitempty"`
	// This is to indicate whether this logistic channel supports COD
	HasCOD bool `json:"has_cod,omitempty"`
	// Whether this logistic channel is enabled on shop level.
	Enabled bool `json:"enabled,omitempty"`
	// See Define FeeType, related to FeeType Value
	FeeType string `json:"fee_type,omitempty"`
	// Only for fee_type is SIZE_SELECTION
	Sizes []GetLogisticsResponseLogisticSize `json:"sizes,omitempty"`
	// The weight limit for this logistic channel.
	WeightLimits GetLogisticsResponseLogisticLimit `json:"weight_limits,omitempty"`
	// The dimension limit for this logistic channel.
	ItemMaxDimension GetLogisticsResponseLogisticDimension `json:"item_max_dimension,omitempty"`
}

//=======================================================
// GetAddressResponse
//=======================================================

type GetAddressResponseAddress struct {
	// The identity of address
	AddressID int64 `json:"address_id,omitempty"`
	// The country of specify address
	Country string `json:"country,omitempty"`
	// The state of specify address
	State string `json:"state,omitempty"`
	// The city of specify address
	City string `json:"city,omitempty"`
	// The address description of specify address
	Address string `json:"address,omitempty"`
	// The zipcode of specify address
	Zipcode string `json:"zipcode,omitempty"`
	// The district of specify address
	District string `json:"district,omitempty"`
	// The town of specify address
	Town string `json:"town,omitempty"`
	// The flag of shop address, applicable values: default_address, pickup_address, return_address
	AddressFlag []string `json:"address_flag,omitempty"`
}

//=======================================================
// GetTimeSlotResponse
//=======================================================

type GetTimeSlotResponseTime struct {
	// The identity of pickuptime.
	PickupTimeID string `json:"pickup_time_id,omitempty"`
	// The date of pickup time. In timestamp.
	Date int `json:"date,omitempty"`
	// The text description of pickup time. Only applicable for certain channels.
	TimeText string `json:"time_text,omitempty"`
}

//=======================================================
// GetBranchResponse
//=======================================================

type GetBranchResponseBranch struct {
	// The identity of branch.
	BranchID int64 `json:"branch_id,omitempty"`
	// The country of specify branch.
	Country string `json:"country,omitempty"`
	// The state of specify branch.
	State string `json:"state,omitempty"`
	// The city of specify branch.
	City string `json:"city,omitempty"`
	// The address description of specify branch.
	Address string `json:"address,omitempty"`
	// The zipcode of specify branch.
	Zipcode string `json:"zipcode,omitempty"`
	// The district of specify branch.
	District string `json:"district,omitempty"`
	// The town of specify branch.
	Town string `json:"town,omitempty"`
}

//=======================================================
// GetLogisticInfoResponse
//=======================================================

type GetLogisticInfoResponsePickupAddress struct {
	// The identity of address.
	AddressID int64 `json:"address_id,omitempty"`
	// The country of specify branch.
	Country string `json:"country,omitempty"`
	// The state of specify branch.
	State string `json:"state,omitempty"`
	// The city of specify branch.
	City string `json:"city,omitempty"`
	// The address description of specify branch.
	Address string `json:"address,omitempty"`
	// The zipcode of specify branch.
	Zipcode string `json:"zipcode,omitempty"`
	// The district of specify branch.
	District string `json:"district,omitempty"`
	// The town of specify branch.
	Town string `json:"town,omitempty"`
	// List of pickup_time information corresponding to the address_id.
	TimeSlotList []string `json:"time_slot_list,omitempty"`
	// The identity of pickuptime.
	PickupTimeID string `json:"pickup_time_id,omitempty"`
	// The date of pickup time. In timestamp.
	Date int `json:"date,omitempty"`
	// The text description of pickup time. Only applicable for certain channels.
	TimeText string `json:"time_text,omitempty"`
	// The flag of shop address, applicable values: default_address, pickup_address, return_address
	AddressFlag []string `json:"address_flag,omitempty"`
}

type GetLogisticInfoResponseDropoffBranch struct {
	// The identity of branch.
	BranchID int64 `json:"branch_id,omitempty"`
	// The country of specify branch.
	Country string `json:"country,omitempty"`
	// The state of specify branch.
	State string `json:"state,omitempty"`
	// The city of specify branch.
	City string `json:"city,omitempty"`
	// The address description of specify branch.
	Address string `json:"address,omitempty"`
	// The zipcode of specify branch.
	Zipcode string `json:"zipcode,omitempty"`
	// The district of specify branch.
	District string `json:"district,omitempty"`
	// The town of specify branch.
	Town string `json:"town,omitempty"`
}

type GetLogisticInfoResponsePickup struct {
	// List of available pickup address info.
	AddressList []GetLogisticInfoResponsePickupAddress `json:"address_list,omitempty"`
}

type GetLogisticInfoResponseDropoff struct {
	// List of available dropoff branches info.
	BranchList []GetLogisticInfoResponseDropoffBranch `json:"branch_list,omitempty"`
}

type GetLogisticInfoResponseInfo struct {
	// Logistics information for pickup mode order.
	Pickup []string `json:"pickup,omitempty"`
	// Logistics information for dropoff mode order.
	Dropoff []string `json:"dropoff,omitempty"`
	// The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init.
	NonIntegrated []string `json:"non_integrated,omitempty"`
}

//=======================================================
// InitRequest
//=======================================================

type InitRequestPickup struct {
	// The identity of address. Retrieved from shopee.logistics.GetAddress.
	AddressID int64 `json:"address_id,omitempty"`
	// The pickup time id. Retrieved from shopee.logistics.GetTimeSlot.
	PickupItemID string `json:"pickup_item_id,omitempty"`
}

type InitRequestDropoff struct {
	// The identity of branch. Retrieved from shopee.logistics.GetBranch branch.
	BranchID int64 `json:"branch_id,omitempty"`
	// The real name of sender.
	SenderRealName string `json:"sender_real_name,omitempty"`
	// Need input this field when "tracking_no" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no,omitempty"`
}

type InitRequestNonIntegrated struct {
	// Optional parameter for non-integrated channel order. The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no,omitempty"`
}

//=======================================================
// GetAirwayBillResponse
//=======================================================

type GetAirwayBillResponseResultAirwayBill struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The url of retrieving airway bill.
	AirwayBill string `json:"airway_bill,omitempty"`
}

type GetAirwayBillResponseResultError struct {
	// The ordersn of orders which occurred error.
	OrderSN string `json:"ordersn,omitempty"`
	//
	ErrorCode string `json:"error_code,omitempty"`
	// The detail information of this error.
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetAirwayBillResponseResult struct {
	// The number of ordersn to get airway bills in this call.
	TotalCount int `json:"total_count,omitempty"`
	// This Object contains the airway bill to each order.
	AirwayBills []GetAirwayBillResponseResultAirwayBill `json:"airway_bills,omitempty"`
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []GetAirwayBillResponseResultError `json:"errors,omitempty"`
}

type GetAirwayBillResponseBatchResult struct {
	// The number of orderSN to get airway bills in this call.
	TotalCount int `json:"total_count,omitempty"`
	// The list contains urls of retrieving airway bill in PDF format. Each url contains the airway bills which is in the same logistics channel.
	AirwayBills []string `json:"airway_bills,omitempty"`
}

//=======================================================
// GetOrderLogisticsResponse
//=======================================================

type GetOrderLogisticsResponseLogisticRecipientAddress struct {
	// Recipient's name for the address.
	Name string `json:"name,omitempty"`
	// Recipient's phone number input when order was placed.
	Phone string `json:"phone,omitempty"`
	// The town of the recipient's address. Whether there is a town will depend on the region and/or country.
	Town string `json:"town,omitempty"`
	// The district of the recipient's address. Whether there is a town will depend on the region and/or country.
	District string `json:"district,omitempty"`
	// The city of the recipient's address. Whether there is a town will depend on the region and/or country.
	City string `json:"city,omitempty"`
	// The state/province of the recipient's address. Whether there is a town will depend on the region and/or country.
	State string `json:"state,omitempty"`
	// The two-digit code representing the country of the Recipient.
	Country string `json:"country,omitempty"`
	// Recipient's postal code.
	Zipcode string `json:"zipcode,omitempty"`
	// The full address of the recipient, including country, state, even street, and etc.
	FullAddress string `json:"full_address,omitempty"`
}

type GetOrderLogisticsResponseLogistic struct {
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string `json:"shipping_carrier,omitempty"`
	// The identity of logistic channel.
	LogisticID int64 `json:"logistic_id,omitempty"`
	// Only work for cross-border order. This code is required at some sorting hub. Please ensure the service_code is INCLUDED on your shipping label, otherwise the parcel cannot be processed by the warehouse. If you didn't retrieve service_code after you first called this API, please try few more times within 30 minutes.
	ServiceCode string `json:"service_code,omitempty"`
	// Only work for cross-border order.The name of the carrier ships cross countries.
	FirstMileName string `json:"first_mile_name,omitempty"`
	// Only work for cross-border order.The name of the carrier delivers the parcels in local country.
	LastMileName string `json:"last_mile_name,omitempty"`
	// Only work for cross-border order.This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	GoodsToDeclare bool `json:"goods_to_declare,omitempty"`
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no,omitempty"`
	//
	Zone string `json:"zone,omitempty"`
	// Only work for cross-border order. The string use for waybill printing. The format is "S - country_code and lane_number". For example, S-TH01, S-TH02
	LaneCode string `json:"lane_code,omitempty"`
	// Only work for cross-border order in some special shop. The address info of the warehouse.
	WarehouseAddress string `json:"warehouse_address,omitempty"`
	// Only work for cross-border order in some special shop. The ID of the warehouse.
	WarehouseID int64 `json:"warehouse_id,omitempty"`
	// This object contains detailed breakdown for the recipient address.
	RecipientAddress GetOrderLogisticsResponseLogisticRecipientAddress `json:"recipient_address,omitempty"`
	// This value indicates whether the order was a COD (cash on delivery) order.
	COD bool `json:"cod,omitempty"`
	// The sort_code of recipient.
	RecipientSortCode GetOrderLogisticsResponseLogisticRecipientSortCode `json:"recipient_sort_code,omitempty"`
	// The sort_code of sender.
	SenderSortCode GetOrderLogisticsResponseLogisticSenderSortCode `json:"sender_sort_code,omitempty"`
	// Only used for local TW sellers.
	ThirdPartyLogisticInfo GetOrderLogisticsResponseLogisticThirdPartyLogisticInfo `json:"third_party_logistic_info,omitempty"`
	// Buyer's CPF number for taxation and invoice purposes. Only for Brazil order.
	BuyerCPFID string `json:"buyer_cpf_id,omitempty"`
	// First mile tracking NO. for CrossBoard BR seller can be used to self-design CB Brazil AWB.
	ShopeeTrackingNo string `json:"shopee_tracking_no,omitempty"`
	// The last-mile tracking number. Only for Cross Board BR seller.
	LMTN string `json:"lm_tn,omitempty"`
}

type GetOrderLogisticsResponseLogisticRecipientSortCode struct {
	// The first-level sort_code of recipient.
	FirstRecipientSortCode string `json:"first_recipient_sort_code,omitempty"`
	// The second-level sort_code of recipient.
	SecondRecipientSortCode string `json:"second_recipient_sort_code,omitempty"`
	// The third-level sort_code of recipient.
	ThirdRecipientSortCode string `json:"third_recipient_sort_code,omitempty"`
}

type GetOrderLogisticsResponseLogisticSenderSortCode struct {
	// The first-level sort_code of sender.
	FirstSenderSortCode string `json:"first_sender_sort_code,omitempty"`
	// The second-level sort_code of sender.
	SecondSenderSortCode string `json:"second_sender_sort_code,omitempty"`
	// The third-level sort_code of sender.
	ThirdSenderSortCode string `json:"third_sender_sort_code,omitempty"`
}

type GetOrderLogisticsResponseLogisticThirdPartyLogisticInfo struct {
	// Use this field to indicate the order category.
	ServiceDescription string `json:"service_description,omitempty"`
	// The manufacturer barcode.
	Barcode string `json:"barcode,omitempty"`
	// The purchase_time of the store.
	PurchaseTime string `json:"purchase_time,omitempty"`
	// The return_time of the store.
	ReturnTime string `json:"return_time,omitempty"`
	// The name of manufacturers.
	ManufacturesName string `json:"manufacturers_name,omitempty"`
	// The website of manufacturers.
	ManufacturesWebsite string `json:"manufacturers_website,omitempty"`
	// The identification of recipient area.
	RecipientArea string `json:"recipient_area,omitempty"`
	// The route code of the waybill.
	RouteStep string `json:"route_step,omitempty"`
	// The tally code of the waybill.
	Suda5Code string `json:"suda5_code,omitempty"`
	// The code of large logistics.
	LargeLogisticsID string `json:"large_logistics_id,omitempty"`
	// The parent code of the waybill.
	ParentID string `json:"parent_id,omitempty"`
	// Use this field to indicate the return cycle.
	ReturnCycle string `json:"return_cycle,omitempty"`
	// Use this field to indicate the return mode.
	ReturnMode string `json:"return_mode,omitempty"`
	// The reminder of stork work.
	Prompt string `json:"prompt,omitempty"`
	// Shopee's unique identifier for an order.
	OrderNo string `json:"order_no,omitempty"`
	// The QR code of the waybill.
	QRCode string `json:"qrcode,omitempty"`
	// The supplier name of channel.
	ECSupplierName string `json:"ec_supplier_name,omitempty"`
	// Use this field to indicate the first barcode.
	ECBarCode16 string `json:"ec_bar_code16,omitempty"`
	// The device code.
	EquipmentID string `json:"equipment_id,omitempty"`
	// The child code for B2C Family-mart.
	EShopID string `json:"eshop_id,omitempty"`
	// Use this field to indicate the pick barcode.
	ECBarCode9 string `json:"ec_bar_code9,omitempty"`
}

//=======================================================
// GetLogisticsMessageResponse
//=======================================================

type GetLogisticsMessageResponseInfo struct {
	// The time when logistics info has been updated.
	CTime int `json:"c_time,omitempty"`
	// The order logistics tracking info.
	Description string `json:"description,omitempty"`
	// The 3PL logistics status for the order. Applicable values: See Data Definition - TrackingLogisticsStatus.
	Status string `json:"status,omitempty"`
}

//=======================================================
// GetForderWaybillRequest
//=======================================================

type GetForderWaybillRequestOrder struct {
	// The order serial numbers. Make sure the order has trackingNo generated before calling this API.
	OrderSN string `json:"ordersn,omitempty"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
}

//=======================================================
// GetForderWaybillResponse
//=======================================================

type GetForderWaybillResponseResultWaybill struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
	// The url of retrieving airway bill.
	Waybill string `json:"waybill,omitempty"`
}

type GetForderWaybillResponseError struct {
	// The ordersn of orders which occurred error.
	OrderSN string `json:"ordersn,omitempty"`
	// The forder_id of fulfillment orders which occurred error.
	ForderID string `json:"forder_id,omitempty"`
	//
	ErrorCode string `json:"error_code,omitempty"`
	// The detail information of this error.
	ErrorDescription string `json:"error_description,omitempty"`
}

type GetForderWaybillResponseResult struct {
	// The number of ordersn to get airway bills in this call.
	TotalCount int `json:"total_count,omitempty"`
	// This Object contains the airway bill to each order.
	Waybills []GetForderWaybillResponseResultWaybill `json:"waybills,omitempty"`
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []GetForderWaybillResponseError `json:"errors,omitempty"`
}

type GetForderWaybillResponseBatchResult struct {
	// The number of orderSN to get airway bills in this call.
	TotalCount int `json:"total_count,omitempty"`
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []GetForderWaybillResponseError `json:"errors,omitempty"`
	// The url of retrieving airway bill.
	Waybills []string `json:"waybills,omitempty"`
}

//=======================================================
// GetReturnListResponse
//=======================================================

type GetReturnListResponseReturnUser struct {
	// Buyer's nickname.
	Username int `json:"username,omitempty"`
	// Buyer's email.
	Email string `json:"email,omitempty"`
	// Buyer's portrait.
	Protrait string `json:"protrait,omitempty"`
}

type GetReturnListResponseReturnItem struct {
	// Item id.
	ItemID int64 `json:"item_id,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// Name of item in local language.
	Name string `json:"name,omitempty"`
	// Image URLs of item.
	Images []string `json:"images,omitempty"`
	// Amount of this item.
	Amount int `json:"amount,omitempty"`
	// The price of Item.
	ItemPrice float64 `json:"item_price,omitempty,string"`
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool `json:"is_add_on_deal,omitempty"`
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool `json:"is_main_item,omitempty"`
	// The unique identity of an addon deal.
	AddOnDealID int64 `json:"add_on_deal_id,omitempty"`
}

type GetReturnListResponseReturn struct {
	// Image URLs of return.
	Images []string `json:"images,omitempty"`
	// Reason for return product. Applicable values: See Data Definition- ReturnReason.
	Reason string `json:"reason,omitempty"`
	// Reason that buyer provide.
	TextReason string `json:"text_reason,omitempty"`
	// The serial number of return.
	ReturnSN int `json:"return_sn,omitempty"`
	// Amount of the refund.
	RefundAmount float64 `json:"refund_amount,omitempty,string"`
	// Currency of the return.
	Currency string `json:"currency,omitempty"`
	// The time of return create.
	CreateTime int `json:"create_time,omitempty"`
	// The time of modify return.
	UpdateTime int `json:"update_time,omitempty"`
	// Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
	Status string `json:"status,omitempty"`
	// The last time seller deal with this return.
	DueDate int `json:"due_date,omitempty"`
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNumber string `json:"tracking_number,omitempty"`
	// The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
	DisputeReason string `json:"dispute_reason,omitempty"`
	// The reason that seller provide. While the return has been disputed, this field is useful.
	DisputeTextReason string `json:"dispute_text_reason,omitempty"`
	// Items to be sent back to seller. Can be either integrated/non-integrated.
	NeedsLogistics bool `json:"needs_logistics,omitempty"`
	// Order price before discount.
	AmountBeforeDiscount float64 `json:"amount_before_discount,omitempty,string"`
	//
	User GetReturnListResponseReturnUser `json:"user,omitempty"`
	//
	Item []GetReturnListResponseReturnItem `json:"item,omitempty"`
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
}

//=======================================================
// GetReturnDetailResponse
//=======================================================

type GetReturnDetailResponseUser struct {
	// Buyer's nickname.
	Username int `json:"username,omitempty"`
	// Buyer's email.
	Email string `json:"email,omitempty"`
	// Buyer's portrait.
	Portrait string `json:"portrait,omitempty"`
}

type GetReturnDetailResponseItem struct {
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// Name of item in local language.
	Name string `json:"name,omitempty"`
	// Image URLs of item.
	Images []string `json:"images,omitempty"`
	// Amount of this item.
	Amount int `json:"amount,omitempty"`
	// The price of item.
	ItemPrice float64 `json:"item_price,omitempty,string"`
	// To indicate if this item belongs to an addon deal.
	IsAddOnDeal bool `json:"is_add_on_deal,omitempty"`
	// To indicate if this item is main item or sub item. True means main item, false means sub item.
	IsMainItem bool `json:"is_main_item,omitempty"`
	// The unique identity of an addon deal.
	AddOnDealID int64 `json:"add_on_deal_id,omitempty"`
}

//=======================================================
// GetShopsByPartnerResponse
//=======================================================

type GetShopsByPartnerResponseShopSIP struct {
	// Affiliate Shop's area
	Country string `json:"country,omitempty"`
	// Affiliate shop's id
	AShopID int64 `json:"a_shopid,omitempty"`
}

type GetShopsByPartnerResponseShop struct {
	// The two-digit code representing the country where the order was made.
	Country string `json:"country,omitempty"`
	// Shopee's unique identifier for a shop.
	ShopID int64 `json:"shopid,omitempty"`
	// The timestamp when the shop was authorized to the partner.
	AuthTime int `json:"auth_time,omitempty"`
	// SIP affiliate shops info list
	SIPAShops []GetShopsByPartnerResponseShopSIP `json:"sip_a_shops,omitempty"`
	// Use this field to indicate the expiration date for shop authorization.
	ExpireTime int `json:"expire_time,omitempty"`
}

//=======================================================
// GetCategoriesByCountryResponse
//=======================================================

type GetCategoriesByCountryResponseCategory struct {
	// The Identify of the parent of the category
	ParentID int64 `json:"parent_id,omitempty"`
	// This is to indicate whether the category has children.
	HasChildren bool `json:"has_children,omitempty"`
	// The Identify of each category
	CategoryID int64 `json:"category_id,omitempty"`
	// The name of each category
	CategoryName string `json:"category_name,omitempty"`
	// To indicate if this category supports size chart
	IsSuppSizechart bool `json:"is_supp_sizechart,omitempty"`
}

//=======================================================
// GetPaymentListResponse
//=======================================================

type GetPaymentListResponseMethod struct {
	// The payment method
	PaymentMethod string `json:"payment_method,omitempty"`
	// The country for this payment method
	Country string `json:"country,omitempty"`
}

//=======================================================
// GetTopPicksListResponse
//=======================================================

type GetTopPicksListResponseCollectionItem struct {
	// Item ID
	ItemID int64 `json:"item_id,omitempty"`
	// Item name
	ItemName string `json:"item_name,omitempty"`
	// Item discounted price(original price if no discount). Item level price will return if it has variation.
	ItemPrice float64 `json:"item_price,omitempty,string"`
	// The sales of the item
	Sales int `json:"sales,omitempty"`
}

type GetTopPicksListResponseCollection struct {
	// Collection name
	Name string `json:"name,omitempty"`
	// Collection ID
	TopPicksID int64 `json:"top_picks_id,omitempty"`
	// True or False
	IsActivated bool `json:"is_activated,omitempty"`
	// Item list of the collection
	Items []GetTopPicksListResponseCollectionItem `json:"items,omitempty"`
}

//=======================================================
// AddTopPicksResponse
//=======================================================

type AddTopPicksResponseItem struct {
	// Item ID
	ItemID int64 `json:"item_id,omitempty"`
	// Item name
	ItemName string `json:"item_name,omitempty"`
	// Item discounted price(original price if no discount). Item level price will return if it has variation.
	ItemPrice float64 `json:"item_price,omitempty,string"`
	// The sales of the item
	Sales int `json:"sales,omitempty"`
}

//=======================================================
// UpdateTopPicksResponse
//=======================================================

type UpdateTopPicksResponseItem struct {
	// Item ID
	ItemID int64 `json:"item_id,omitempty"`
	// Item name
	ItemName string `json:"item_name,omitempty"`
	// Item discounted price(original price if no discount). Item level price will return if it has variation.
	ItemPrice float64 `json:"item_price,omitempty,string"`
	// The sales of the item
	Sales int `json:"sales,omitempty"`
}

//=======================================================
// GenerateFMTrackingNoRequest
//=======================================================

type GenerateFMTrackingNoRequestSellerInfo struct {
	// The full address of the seller.
	Address string `json:"address,omitempty"`
	// Seller's name for the address.
	Name string `json:"name,omitempty"`
	// Seller's postal code.
	Zipcode string `json:"zipcode,omitempty"`
	// Seller's location.
	Area string `json:"area,omitempty"`
	// Seller's phone number.
	Phone string `json:"phone,omitempty"`
}

//=======================================================
// GetShopFMTrackingNoResponse
//=======================================================

type GetShopFMTrackingNoResponseFMTNList struct {
	// The specified delivery date.
	DeclareDate string `json:"declare_date,omitempty"`
	// The logistics status for bound orders.
	Status string `json:"status,omitempty"`
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
}

//=======================================================
// FirstMileCodeBindOrderRequest
//=======================================================

type FirstMileCodeBindOrderRequestOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
}

//=======================================================
// FirstMileCodeBindOrderResponse
//=======================================================

type FirstMileCodeBindOrderResponseFail struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
	// The reason why the order/fulfillment order cannot be bound.
	Reason string `json:"reason,omitempty"`
}

//=======================================================
// GetFmTnDetailResponse
//=======================================================

type GetFmTnDetailResponseOrder struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
	// The tracking number of SLS for orders/forders.
	SLSTN string `json:"slstn,omitempty"`
	// Use this filed to indicate whether the order has been picked up by carrier.
	IsPickUpDone bool `json:"is_pick_up_done,omitempty"`
	// Use this filed to indicate whether the order has arrived at transit warehouse.
	IsArrivedTWS bool `json:"is_arrived_tws,omitempty"`
}

//=======================================================
// GetFMTrackingNoWaybillResponse
//=======================================================

type GetFMTrackingNoWaybillResponseError struct {
	//
	ErrorCode string `json:"error_code,omitempty"`
	// The detail information of this error.
	ErrorDescription string `json:"error_description,omitempty"`
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
}

type GetFMTrackingNoWaybillResponseResultWaybill struct {
	// The first-mile tracking number.
	FMTN string `json:"fm_tn,omitempty"`
	// The url of retrieving waybill.
	Waybill string `json:"waybill,omitempty"`
}

type GetFMTrackingNoWaybillResponseResult struct {
	// This Object contains the waybill to each tracking number.
	Waybills []GetFMTrackingNoWaybillResponseResultWaybill `json:"waybills,omitempty"`
	// The number of Tracking Number to get waybills in this call.
	TotalCount int `json:"total_count,omitempty"`
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []GetFMTrackingNoWaybillResponseError `json:"errors,omitempty"`
}

type GetFMTrackingNoWaybillResponseBatchResult struct {
	// The list contains urls of retrieving waybill in PDF format. Each url contains the airway bills which is in the same logistics channel.
	Waybills []string `json:"waybills,omitempty"`
	// The number of Tracking Number to get waybills in this call.
	TotalCount int `json:"total_count,omitempty"`
	// This list contains the first-mile tracking number and error descriptions of all tracking numbers that failed to retrieve airway bill in this call.
	Errors []GetFMTrackingNoWaybillResponseError `json:"errors,omitempty"`
}

//=======================================================
// GetShopFirstMileChannelResponse
//=======================================================

type GetShopFirstMileChannelResponseLogistic struct {
	// The identity of logistic channel.
	LogisticID int64 `json:"logistic_id,omitempty"`
	// The name of logistic.
	LogisticName string `json:"logistic_name,omitempty"`
	// The shipment method for bound orders, should be pickup or dropoff.
	ShipmentMethod string `json:"shipment_method,omitempty"`
}

//=======================================================
// GetPushConfigResponse
//=======================================================

type GetPushConfigResponseDeatiledConfig struct {
	// 0 stands for off and 1 stands for on.
	OrderStatus int `json:"order_status,omitempty"`
	// 0 stands for off and 1 stands for on.
	OrderTrackingNo int `json:"order_trackingno,omitempty"`
	// 0 stands for off and 1 stands for on.
	ShopUpdate int `json:"shop_update,omitempty"`
	// 0 stands for off and 1 stands for on.
	BannedItem int `json:"banned_item,omitempty"`
	// 0 stands for off and 1 stands for on.
	ItemPromotion int `json:"item_promotion,omitempty"`
	// 0 stands for off and 1 stands for on.
	ReservedStockChange int `json:"reserved_stock_change,omitempty"`
}

//=======================================================
// SetPushConfigRequest
//=======================================================

type SetPushConfigRequestDeatiledConfig struct {
	// 0 stands for off and 1 stands for on.
	OrderStatus int `json:"order_status,omitempty"`
	// 0 stands for off and 1 stands for on.
	OrderTrackingNo int `json:"order_trackingno,omitempty"`
	// 0 stands for off and 1 stands for on.
	ShopUpdate int `json:"shop_update,omitempty"`
	// 0 stands for off and 1 stands for on.
	BannedItem int `json:"banned_item,omitempty"`
	// 0 stands for off and 1 stands for on.
	ItemPromotion int `json:"item_promotion,omitempty"`
	// 0 stands for off and 1 stands for on.
	ReservedStockChange int `json:"reserved_stock_change,omitempty"`
}

//=======================================================
// GetTransactionListResponse
//=======================================================

type GetTransactionListResponseTransactionListPayOrderList struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// Name of the shop.
	ShopName string `json:"shop_name,omitempty"`
}

type GetTransactionListResponseTransactionList struct {
	// The ID of transaction.
	TransactionID int64 `json:"transaction_id,omitempty"`
	// The status of the transaction，available values: FAILED,COMPLETED,PENDING,INITIAL.
	Status string `json:"status,omitempty"`
	// The type of wallet, available values: shopee pay, jko pay.
	WalletType string `json:"wallet_type,omitempty"`
	// The type of transaction.
	TransactionType string `json:"transaction_type,omitempty"`
	// The amount of transaction.
	Amount float64 `json:"amount,omitempty,string"`
	// The current balance of this account.
	CurrentBalance float64 `json:"current_balance,omitempty,string"`
	// The create time of the transaction.
	CreateTime int `json:"create_time,omitempty"`
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The serial number of return.
	RefundSN string `json:"refund_sn,omitempty"`
	// The type of withdrawal.
	WithdrawalType string `json:"withdrawal_type,omitempty"`
	// This field indicates the transaction fee.
	TransactionFee float64 `json:"transaction_fee,omitempty,string"`
	// The detailed description of TOPUP SUCCESS and TOPUP FAILED.
	Description string `json:"description,omitempty"`
	// The name of buyer.
	BuyerName string `json:"buyer_name,omitempty"`
	// List of ordersn included in the transaction.
	PayOrderList GetTransactionListResponseTransactionListPayOrderList `json:"pay_order_list,omitempty"`
	// Name of the shop.
	ShopName string `json:"shop_name,omitempty"`
	// Withdraw ID when transaction type is withdraw_created, withdrawal_completed, withdrawal_cancelled.
	WithdrawID float64 `json:"withdraw_id,omitempty,string"`
	// The reason for ADJUSTMENT_ADD and ADJUSTMENT_MINUS.
	Reason string `json:"reason,omitempty"`
	// Use this field to indicate the event where a withdrawal is split into several withdrawals due to the withdrawal limit.
	RootWithdrawID float64 `json:"root_withdrawal_id,omitempty"`
}

//=======================================================
// FirstMileUnbindResponse
//=======================================================

type FirstMileUnbindResponseFailList struct {
	// The reason why the order/fulfillment order cannot be unbound.
	Reason string `json:"reason,omitempty"`
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
}

//=======================================================
// FirstMileUnbindRequest
//=======================================================

type FirstMileUnbindRequestOrderList struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"ordersn,omitempty"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
}

//=======================================================
// MyIncomeResponse
//=======================================================

type MyIncomeResponseOrderIncome struct {
	// The total amount that the seller is expected to receive for the order and will change before order completed. escrow_amount=buyer_total_amount+shopee_discount+voucher_from_shopee+coins+payment_promotion-buyer_transaction_fee-cross_border_tax-commission_fee-service_fee-seller_transaction_fee-seller_coin_cash_back-escrow_tax-drc_adjustable_refund+final_shipping_fee（could be positive/negative)
	EscrowAmount float64 `json:"escrow_amount,omitempty,string"`
	// The total amount that paid by buyer.buyer_total_amount= original price -seller_discount -shopee_discount -voucher_from_seller -voucher_from_shopee -coin -payment_promotion +buyer_paid_shipping_fee +buyer_transaction_fee +cross_border_tax
	BuyerTotalAmount float64 `json:"buyer_total_amount,omitempty,string"`
	// The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.
	OriginalPrice float64 `json:"original_price,omitempty,string"`
	// Final sum of each item seller discount of a specific order.
	SellerDiscount float64 `json:"seller_discount,omitempty,string"`
	// Final sum of each item Shopee discount of a specific order. This amount will rebate to seller.
	ShopeeDiscount float64 `json:"shopee_discount,omitempty,string"`
	// Final value of voucher provided by Seller for the order.
	VoucherFromSeller float64 `json:"voucher_from_seller,omitempty,string"`
	// Final value of voucher provided by Shopee for the order.
	VoucherFromShopee float64 `json:"voucher_from_shopee,omitempty,string"`
	// Final value of coins used by seller for the order.
	Coins float64 `json:"coins,omitempty,string"`
	// The shipping fee paid by buyer.
	BuyerPaidShippingFee float64 `json:"buyer_paid_shipping_fee,omitempty,string"`
	// Tansaction fee paid by buyer for the order.
	BuyerTransactionFee float64 `json:"buyer_transaction_fee,omitempty,string"`
	// Amount incurred by Buyer for purchasing items outside of home country. Amount may change after Return Refund.
	CrossBorderTax float64 `json:"cross_border_tax,omitempty,string"`
	// The amount offset via payment promotion.
	PaymentPromotion float64 `json:"payment_promotion,omitempty,string"`
	// The commission fee charged by Shopee platform if applicable.
	CommissionFee float64 `json:"commission_fee,omitempty,string"`
	// Amount charged by Shopee to seller for additional services.
	ServiceFee float64 `json:"service_fee,omitempty,string"`
	// Tansaction fee paid by seller for the order.
	SellerTransactionFee float64 `json:"seller_transaction_fee,omitempty,string"`
	// Compensation to seller in case of lost parcel
	SellerLostCompensation float64 `json:"seller_lost_compensation,omitempty,string"`
	// Value of coins provided by Seller for purchasing with his or her store for the order.
	SellerCoinCashBack float64 `json:"seller_coin_cash_back,omitempty,string"`
	// Cross-border tax imposed by the Indonesian government on sellers.
	EscrowTax float64 `json:"escrow_tax,omitempty,string"`
	// Final adjusted amount that seller has to bear as part of escrow. This amount could be negative or positive. = min(actual_shipping_fee, shopee_shipping_rebate) + shipping_discount_from_3pl - actual_shipping_fee
	FinalShippingFee float64 `json:"final_shipping_fee,omitempty,string"`
	// The final shipping cost of order and it is negative. For Non-integrated logistics channel is 0.
	ActualShippingFee float64 `json:"actual_shipping_fee,omitempty,string"`
	// The platform shipping subsidy to the seller.
	ShopeeShippingRebate float64 `json:"shopee_shipping_rebate,omitempty,string"`
	// The discount of shipping fee from 3PL. Currently only applicable to ID
	ShippingFeeDiscountFrom3PL float64 `json:"shipping_fee_discount_from_3pl,omitempty,string"`
	// The shipping discount defined by seller.
	SellerShippinhDiscount float64 `json:"seller_shipping_discount,omitempty,string"`
	// The estimated shipping fee is an estimation calculated by Shopee based on specific logistics courier's standard.
	EstimatedShippingFee float64 `json:"estimated_shipping_fee,omitempty,string"`
	// The list of voucher code provided by seller.
	SellerVoucherCode []string `json:"seller_voucher_code,omitempty,string"`
	// The adjustable refund amount from Shopee Dispute Resolution Center.
	DRCAdjustableRefund float64 `json:"drc_adjustable_refund,omitempty,string"`
	// Final amount paid by the buyer for the items in a specific order.
	CostOfGoodsSold float64 `json:"cost_of_goods_sold,omitempty,string"`
	// Amount paid by the buyer for the items in a specific order.
	OriginalCostOfGoodsSold float64 `json:"original_cost_of_goods_sold,omitempty,string"`
	// Sum of each item Shopee discount of a specific order
	OriginalShopeeDiscount float64 `json:"original_shopee_discount,omitempty,string"`
	// Amount returned to Seller in the event of Partial Return.
	SellerReturnRefund float64 `json:"seller_return_refund,omitempty,string"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty,string"`
	// The list of the serial number of refund.
	RefundIDList []string `json:"refund_id_list,omitempty,string"`
}

//=======================================================
// GetCommentResponse
//=======================================================

type GetCommentResponseItemCMTListCMTReply struct {
	// The content of reply
	Reply string `json:"reply,omitempty"`
	// The status of comment, available values: DELETE/NORMAL/VALID/HIDDEN
	Status string `json:"status,omitempty"`
}

type GetCommentResponseItemCMTList struct {
	// The identity of comment.
	CMTID int64 `json:"cmt_id,omitempty"`
	// Content of the comment.
	Comment string `json:"comment,omitempty"`
	// Username of the buyer who posted the comment.
	BuyerUsername string `json:"buyer_username,omitempty"`
	// Commented ordersn
	OrderSN string `json:"ordersn,omitempty"`
	// Commented item's id
	ItemID int64 `json:"item_id,omitempty"`
	// Shopee's unique identifier for a variation of an item.
	VariationID int64 `json:"variation_id,omitempty"`
	// The create time of the comment
	CreateTime int `json:"create_time,omitempty"`
	// The status of comment, available values: DELETE/NORMAL/VALID/HIDDEN
	Status string `json:"status,omitempty"`
	// Buyer's rating for the item
	RatingStar int `json:"rating_star,omitempty"`
	// The edit status of comment, available values: EXPIRED/EDITABLE/HAVE_EDIT_ONCE
	Editable string `json:"editable,omitempty"`
	// The status of comment, available values: DELETE/NORMAL/VALID/HIDDEN
	CMTReply GetCommentResponseItemCMTListCMTReply `json:"cmt_reply,omitempty"`
}

//=======================================================
// ReplyCommentsRequest
//=======================================================

type ReplyCommentsRequestCMTList struct {
	// The identity of comment.
	CMTID int64 `json:"cmt_id,omitempty"`
	// Content of the comment.
	Comment string `json:"comment,omitempty"`
}

//=======================================================
// ReplyCommentsResponse
//=======================================================

type ReplyCommentsResponseSuccList struct {
	CMTID int64 `json:"cmt_id,omitempty"`
}

type ReplyCommentsResponseError struct {
	// The identity of comment.
	CMTID    int64  `json:"cmt_id,omitempty"`
	ErrorMsg string `json:"error_msg,omitempty"`
}
