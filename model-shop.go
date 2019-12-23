package shopeego

type GetShopInfoRequest struct {
	RequestBase
}

type GetShopInfoResponse struct {
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// Name of the shop.
	ShopName string `json:"shop_name"`
	// The two-digit code representing the country where the order was made.
	Country string `json:"country"`
	// Description of the shop.
	ShopDescripion string `json:"shop_descripion"`
	// List of videos URLs of the shop.
	Videos []string `json:"videos"`
	// List of images URLs of the shop.
	Images []string `json:"images"`
	// Allow negotiations or not, 1: don't allow, 0: allow.
	DisableMakeOffer int `json:"disable_make_offer"`
	// Display pickup address or not.
	EnableDisplayUnitNo bool `json:"enable_display_unit_no"`
	// Listing limitation of items for the shop.
	ItemLimit int `json:"item_limit"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
	// Applicable status: BANNED, FROZEN, NORMAL
	Status string `json:"status"`
	// Only for TW seller. The status of whether shop support installment: 1 means true and 0 means false
	InstallmentStatus int `json:"installment_status"`
	// SIP affiliate shops info list
	SIPAShops []AffiliateShop `json:"sipa_shops"`
}

type UpdateShopInfoRequest struct {
	// Shop name of this shop.
	ShopName string `json:"shop_name"`
	// List of images url of shop banners.
	Images []string `json:"images"`
	// List of videos url of shop banners.
	Videos []string `json:"videos"`
	// Allow negotiations or not, 1: don't allow, 0: allow.
	DisableMakeOffer int `json:"disable_make_offer"`
	// Display pickup address or not.
	EnableDisplayUnitNo bool `json:"enable_display_unit_no"`
	// Description of the shop.
	ShopDescription string `json:"shop_description"`
	//
	RequestBase
}

type UpdateShopInfoResponse struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// Shop name of this shop.
	ShopName string `json:"shop_name"`
	// List of images url of shop banners.
	Images []string `json:"images"`
	// List of videos url of shop banners. Only accept youtube video urls.
	Videos []string `json:"videos"`
	// Allow negotiations or not, 1: don't allow, 0: allow.
	DisableMakeOffer int `json:"disable_make_offer"`
	// Display pickup address or not.
	EnableDisplayUnitNo bool `json:"enable_display_unit_no"`
	// Warning message if parts of image/video uploads failed.
	Warning string `json:"warning"`
	// Description of the shop.
	ShopDescription string `json:"shop_description"`
	// The two-digit code representing the country where the order was made.
	Country string `json:"country"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type PerformanceRequest struct {
	//
	RequestBase
}

type PerformanceResponse struct {
	// To ensure that buyers can easily find what they are looking for, any attempts by sellers to manipulate search results to gain an unfair advantage will be penalized.
	SpamListingViolation []Performance `json:"spam_listing_violation"`
	// Overall review rating is the average of all order ratings submitted by your buyers.
	OverallReviewRating []Performance `json:"overall_review_rating"`
	// Preparation time is the number of days it takes a seller to prepare and ship out an order.
	AveragePreparationTime []Performance `json:"average_preparation_time"`
	// Late shipment rate is the percentage of orders (out of total orders) that were shipped late in the past 30 days. You should maintain your late shipment rate at a healthy level of <5%. If your late shipment rate exceeds 15%, you will receive a penalty point under the Seller Penalty Points system.
	LateShipmentRate []Performance `json:"late_shipment_rate"`
	// Return-refund rate is the percentage of orders (out of total orders) that were requested by buyers for a return-refund in the past 30 days.
	ReturnRefundRate []Performance `json:"return_refund_rate"`
	// Chat response time is the average time it takes a seller to respond to a buyer's chat message.
	ResponseTime []Performance `json:"response_time"`
	// It is the responsibility of sellers to ensure all items listed under their profiles are fully compliant with local laws, as well as Shopee’s terms and policies.
	ProhibitedListingViolation []Performance `json:"prohibited_listing_violation"`
	// Cancellation rate is the percentage of orders (out of total orders) cancelled by a seller in the past 30 days. Buyers initiatied cancellations are not included in the calculation.
	CancellationRate []Performance `json:"cancellation_rate"`
	// Sellers should only list authentic products. Counterfeit items are products that were made in exact imitation of an existing brand with the intention to deceive or defraud, and may include, but are not limited to: - Products that are fake or replicas of an existing official product - Products that have never been produced by a specific brand - Products that bear such similarities with other products (e.g. a replica of a branded item with or without altered logos) without the authorization of the trademark owner.
	CounterfeitListingViolation []Performance `json:"counterfeit_listing_violation"`
	// Your shop rating.
	ShopRating []Performance `json:"shop_rating"`
	// Chat response rate is the percentage of new chats and offers (out of total) that a seller responds to within 12 hours of receiving them. Auto replies are not included in the chat response rate computation.
	ResponseRate []Performance `json:"response_rate"`
	// Non-fulfilment rate is the percentage of orders (out of total orders) that were either cancelled or returned in the past 30 days. Only orders cancelled by sellers are taken into consideration when computing non-fulfilment rate. Non-fulfilment rate is also the sum of your cancellation rate and return-refund rate.
	NonFullfillmentRate []Performance `json:"non_fullfillment_rate"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id"`
}

type SetShopInstallmentStatusRequest struct {
	// The status of whether shop support installment: 1 means true and 0 means false.
	InstallmentStatus int `json:"installment_status"`
	//
	RequestBase
}

type SetShopInstallmentStatusResponse struct {
	// The status of whether shop support installment: 1 means true and 0 means false.
	InstallmentStatus int `json:"installment_status"`
}
