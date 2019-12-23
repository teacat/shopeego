package shopeego

type Client interface {
	GetShopInfo
	UpdateShopInfo
	Performance
	SetShopInstallmentStatus

	AddShopCategory
	GetShopCategories
	DeleteShopCategory
	UpdateShopCategory
	AddItems
	GetItems
	DeleteItems

	GetCategories
	GetAttributes
	Add
	Delete
	UnlistItem
	AddVariations
	DeleteVariation
	GetItemsList
	GetItemDetail
	UpdateItem
	AddItemImg
	UpdateItemImg
	InsertItemImg
	DeleteItemImg
	UpdatePrice
	UpdateStock
	UpdateVariationPrice
	UpdateVariationStock
	UpdatePriceBatch
	UpdateStockBatch
	UpdateVariationPriceBatch
	UpdateVariationStockBatch
	InitTierVariation
	AddTierVariation
	GetVariation
	UpdateTierVariationList
	UpdateTierVariationIndex
	BoostItem
	GetBoostedItem
	SetItemInstallmentTenures
	GetPromotionInfo
	GetRecommendCats

	UploadImg

	AddDiscount
	AddDiscountItem
	DeleteDiscount
	DeleteDiscountItem
	GetDiscountDetail
	GetDiscountsList
	UpdateDiscount
	UpdateDiscountItems

	GetOrdersList
	GetOrdersByStatus
	GetOrderDetails
	GetEscrowDetails
	AddOrderNote
	CancelOrder
	AcceptBuyerCancellation
	RejectBuyerCancellation
	GetForderInfo
	GetEscrowReleasedOrders
	SplitOrder
	UndoSplitOrder
	GetUnbindOrderList

	GetLogistics
	UpdateShopLogistics
	GetParameterForInit
	GetAddress
	GetTimeSlot
	GetBranch
	GetLogisticInfo
	Init
	GetAirwayBill
	GetOrderLogistics
	GetLogisticsMessage
	GetForderWaybill

	ConfirmReturn
	DisputeReturn
	GetReturnList
	GetReturnDetail

	GetShopsByPartner
	GetCategoriesByCountry
	GetPaymentList

	GetTopPicksList
	AddTopPicks
	UpdateTopPicks
	DeleteTopPicks

	GenerateFMTrackingNo
	GetShopFMTrackingNo
	FirstMileCodeBindOrder
	GetFmTnDetail
	GetFMTrackingNoWaybill
	GetShopFirstMileChannel
}

type GetShopInfoRequest struct {
	RequestBase
}

type GetShopInfoResponse struct {
	// Shopee's unique identifier for a shop.
	ShopID int
	// Name of the shop.
	ShopName string
	// The two-digit code representing the country where the order was made.
	Country string
	// Description of the shop.
	ShopDescripion string
	// List of videos URLs of the shop.
	Videos []string
	// List of images URLs of the shop.
	Images []string
	// Allow negotiations or not, 1: don't allow, 0: allow.
	DisableMakeOffer int
	// Display pickup address or not.
	EnableDisplayUnitNo bool
	// Listing limitation of items for the shop.
	ItemLimit int
	// The identifier for an API request for error tracking
	RequestID string
	// Applicable status: BANNED, FROZEN, NORMAL
	Status string
	// Only for TW seller. The status of whether shop support installment: 1 means true and 0 means false
	InstallmentStatus int
	// SIP affiliate shops info list
	SIPAShops []GetShopInfoResponseAffiliateShop
}

type UpdateShopInfoRequest struct {
	// Shop name of this shop.
	ShopName string
	// List of images url of shop banners.
	Images []string
	// List of videos url of shop banners.
	Videos []string
	// Allow negotiations or not, 1: don't allow, 0: allow.
	DisableMakeOffer int
	// Display pickup address or not.
	EnableDisplayUnitNo bool
	// Description of the shop.
	ShopDescription string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type UpdateShopInfoResponse struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// Shop name of this shop.
	ShopName string
	// List of images url of shop banners.
	Images []string
	// List of videos url of shop banners. Only accept youtube video urls.
	Videos []string
	// Allow negotiations or not, 1: don't allow, 0: allow.
	DisableMakeOffer int
	// Display pickup address or not.
	EnableDisplayUnitNo bool
	// Warning message if parts of image/video uploads failed.
	Warning string
	// Description of the shop.
	ShopDescription string
	// The two-digit code representing the country where the order was made.
	Country string
	// The identifier for an API request for error tracking
	RequestID string
}

type PerformanceRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Performance struct {
	// The threshold used to compare shop's actual performance to the target performance. It has four types: lt(less than), gt(greater than), lte(less than or equal), gte(greater than or equal).
	ThresholdType string
	// Null, not applicable.
	Unit string
	// Your target performance index.
	Target int
	// Your actual performance index.
	My float64
}

type PerformanceResponse struct {
	// To ensure that buyers can easily find what they are looking for, any attempts by sellers to manipulate search results to gain an unfair advantage will be penalized.
	SpamListingViolation []Performance
	// Overall review rating is the average of all order ratings submitted by your buyers.
	OverallReviewRating []Performance
	// Preparation time is the number of days it takes a seller to prepare and ship out an order.
	AveragePreparationTime []Performance
	// Late shipment rate is the percentage of orders (out of total orders) that were shipped late in the past 30 days. You should maintain your late shipment rate at a healthy level of <5%. If your late shipment rate exceeds 15%, you will receive a penalty point under the Seller Penalty Points system.
	LateShipmentRate []Performance
	// Return-refund rate is the percentage of orders (out of total orders) that were requested by buyers for a return-refund in the past 30 days.
	ReturnRefundRate []Performance
	// Chat response time is the average time it takes a seller to respond to a buyer's chat message.
	ResponseTime []Performance
	// It is the responsibility of sellers to ensure all items listed under their profiles are fully compliant with local laws, as well as Shopee’s terms and policies.
	ProhibitedListingViolation []Performance
	// Cancellation rate is the percentage of orders (out of total orders) cancelled by a seller in the past 30 days. Buyers initiatied cancellations are not included in the calculation.
	CancellationRate []Performance
	// Sellers should only list authentic products. Counterfeit items are products that were made in exact imitation of an existing brand with the intention to deceive or defraud, and may include, but are not limited to: - Products that are fake or replicas of an existing official product - Products that have never been produced by a specific brand - Products that bear such similarities with other products (e.g. a replica of a branded item with or without altered logos) without the authorization of the trademark owner.
	CounterfeitListingViolation []Performance
	// Your shop rating.
	ShopRating []Performance
	// Chat response rate is the percentage of new chats and offers (out of total) that a seller responds to within 12 hours of receiving them. Auto replies are not included in the chat response rate computation.
	ResponseRate []Performance
	// Non-fulfilment rate is the percentage of orders (out of total orders) that were either cancelled or returned in the past 30 days. Only orders cancelled by sellers are taken into consideration when computing non-fulfilment rate. Non-fulfilment rate is also the sum of your cancellation rate and return-refund rate.
	NonFullfillmentRate []Performance
	// The identifier for an API request for error tracking.
	RequestID string
}

type SetShopInstallmentStatusRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// The status of whether shop support installment: 1 means true and 0 means false.
	InstallmentStatus int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type SetShopInstallmentStatusResponse struct {
	// The status of whether shop support installment: 1 means true and 0 means false.
	InstallmentStatus int
}
