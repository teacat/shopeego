package shopeego

type GetShopInfoRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetShopInfoResponse struct {
	// Shopee's unique identifier for a shop.
	ShopID int64 `json:"shopid,omitempty"`
	// Name of the shop.
	ShopName string `json:"shop_name,omitempty"`
	// The two-digit code representing the country where the order was made.
	Country string `json:"country,omitempty"`
	// Description of the shop.
	ShopDescription string `json:"shop_description,omitempty"`
	// List of videos URLs of the shop.
	Videos []string `json:"videos,omitempty"`
	// List of images URLs of the shop.
	Images []string `json:"images,omitempty"`
	// Allow negotiations or not, 1: don't allow, 0: allow.
	DisableMakeOffer int `json:"disable_make_offer,omitempty"`
	// Display pickup address or not.
	EnableDisplayUnitNo bool `json:"enable_display_unit_no,omitempty"`
	// Listing limitation of items for the shop.
	ItemLimit int `json:"item_limit,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
	// Applicable status: BANNED, FROZEN, NORMAL
	Status string `json:"status,omitempty"`
	// Only for TW seller. The status of whether shop support installment: 1 means true and 0 means false
	InstallmentStatus int `json:"installment_status,omitempty"`
	// SIP affiliate shops info list
	SIPAShops []GetShopInfoResponseShop `json:"sip_a_shops,omitempty"`
	// Use this filed to indicate whether the shop is a cross-border shop.
	IsCB bool `json:"is_cb,omitempty"`
	// The days-to-ship value for non-pre orders.
	NonPreOrderDTS int64 `json:"non_pre_order_dts,omitempty"`
	// The timestamp when the shop was authorized to the partner.
	AuthTime int64 `json:"auth_time,omitempty"`
	// Use this field to indicate the expiration date for shop authorization.
	ExpireTime int64 `json:"expire_time,omitempty"`
}

type UpdateShopInfoRequest struct {
	// Shop name of this shop.
	ShopName string `json:"shop_name,omitempty"`
	// List of images url of shop banners.
	Images []string `json:"images,omitempty"`
	// List of videos url of shop banners.
	Videos []string `json:"videos,omitempty"`
	// Allow negotiations or not, 1: don't allow, 0: allow.
	DisableMakeOffer int `json:"disable_make_offer,omitempty"`
	// Display pickup address or not.
	EnableDisplayUnitNo bool `json:"enable_display_unit_no,omitempty"`
	// Description of the shop.
	ShopDescription string `json:"shop_description,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdateShopInfoResponse struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// Shop name of this shop.
	ShopName string `json:"shop_name,omitempty"`
	// List of images url of shop banners.
	Images []string `json:"images,omitempty"`
	// List of videos url of shop banners. Only accept youtube video urls.
	Videos []string `json:"videos,omitempty"`
	// Allow negotiations or not, 1: don't allow, 0: allow.
	DisableMakeOffer int `json:"disable_make_offer,omitempty"`
	// Display pickup address or not.
	EnableDisplayUnitNo bool `json:"enable_display_unit_no,omitempty"`
	// Warning message if parts of image/video uploads failed.
	Warning string `json:"warning,omitempty"`
	// Description of the shop.
	ShopDescription string `json:"shop_description,omitempty"`
	// The two-digit code representing the country where the order was made.
	Country string `json:"country,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type PerformanceRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type PerformanceResponse struct {
	// To ensure that buyers can easily find what they are looking for, any attempts by sellers to manipulate search results to gain an unfair advantage will be penalized.
	SpamListingViolation PerformanceResponsePerformance `json:"spam_listing_violation,omitempty"`
	// Overall review rating is the average of all order ratings submitted by your buyers.
	OverallReviewRating PerformanceResponsePerformance `json:"overall_review_rating,omitempty"`
	// Preparation time is the number of days it takes a seller to prepare and ship out an order.
	AveragePreparationTime PerformanceResponsePerformance `json:"average_preparation_time,omitempty"`
	// Late shipment rate is the percentage of orders (out of total orders) that were shipped late in the past 30 days. You should maintain your late shipment rate at a healthy level of <5%. If your late shipment rate exceeds 15%, you will receive a penalty point under the Seller Penalty Points system.
	LateShipmentRate PerformanceResponsePerformance `json:"late_shipment_rate,omitempty"`
	// Return-refund rate is the percentage of orders (out of total orders) that were requested by buyers for a return-refund in the past 30 days.
	ReturnRefundRate PerformanceResponsePerformance `json:"return_refund_rate,omitempty"`
	// Chat response time is the average time it takes a seller to respond to a buyer's chat message.
	ResponseTime PerformanceResponsePerformance `json:"response_time,omitempty"`
	// It is the responsibility of sellers to ensure all items listed under their profiles are fully compliant with local laws, as well as Shopee’s terms and policies.
	ProhibitedListingViolation PerformanceResponsePerformance `json:"prohibited_listing_violation,omitempty"`
	// Cancellation rate is the percentage of orders (out of total orders) cancelled by a seller in the past 30 days. Buyers initiatied cancellations are not included in the calculation.
	CancellationRate PerformanceResponsePerformance `json:"cancellation_rate,omitempty"`
	// Sellers should only list authentic products. Counterfeit items are products that were made in exact imitation of an existing brand with the intention to deceive or defraud, and may include, but are not limited to: - Products that are fake or replicas of an existing official product - Products that have never been produced by a specific brand - Products that bear such similarities with other products (e.g. a replica of a branded item with or without altered logos) without the authorization of the trademark owner.
	CounterfeitListingViolation PerformanceResponsePerformance `json:"counterfeit_listing_violation,omitempty"`
	// Your shop rating.
	ShopRating PerformanceResponsePerformance `json:"shop_rating,omitempty"`
	// Chat response rate is the percentage of new chats and offers (out of total) that a seller responds to within 12 hours of receiving them. Auto replies are not included in the chat response rate computation.
	ResponseRate PerformanceResponsePerformance `json:"response_rate,omitempty"`
	// Non-fulfilment rate is the percentage of orders (out of total orders) that were either cancelled or returned in the past 30 days. Only orders cancelled by sellers are taken into consideration when computing non-fulfilment rate. Non-fulfilment rate is also the sum of your cancellation rate and return-refund rate.
	NonFullfillmentRate PerformanceResponsePerformance `json:"non_fullfillment_rate,omitempty"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
}

type SetShopInstallmentStatusRequest struct {
	// The status of whether shop support installment: 1 means true and 0 means false.
	InstallmentStatus int `json:"installment_status,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type SetShopInstallmentStatusResponse struct {
	// The status of whether shop support installment: 1 means true and 0 means false.
	InstallmentStatus int `json:"installment_status,omitempty"`
}

type AuthPartnerRequest struct {
	// The credential retrieved in the APP console
	PartnerID int64 `json:"partner_id,omitempty"`
	// An encrypted string generated in accordance with Shopee rule
	// Token string `json:"token,omitempty"`
	// Redirect url. The url to which the page directs after the authorization is completed. It can be any url(e.g. your company official website).
	Redirect string `json:"redirect,omitempty"`
	// Needs to be consistent with the timestamp in sign base string. Expires in 5 minutes.
	Timestamp int `json:"timestamp,omitempty"`
}
