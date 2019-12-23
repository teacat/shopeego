package shopeego

type GetShopsByPartnerRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type SIP struct {
	// Affiliate Shop's area
	Country string `json:"country"`
	// Affiliate shop's id
	AShopID int `json:"a_shop_id"`
}

type Shop struct {
	// The two-digit code representing the country where the order was made.
	Country string `json:"country"`
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// The timestamp when the shop was authorized to the partner.
	AuthTime int `json:"auth_time"`
	// SIP affiliate shops info list
	SIPAShops []SIP `json:"sipa_shops"`
	// Use this field to indicate the expiration date for shop authorization.
	ExpireTime int `json:"expire_time"`
}

type GetShopsByPartnerResponse struct {
	// A list of shops authorized to the partner.
	AuthedShops []Shop `json:"authed_shops"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetCategoriesByCountryRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// Two-digit country code(capital letter).
	Country string `json:"country"`
	// Is cross border or not. 1: cross border; 0: not cross border
	IsCB bool `json:"is_cb"`
	// Indicate the translation language. Language values include: en(English), vi(Vietnamese), id(Indonesian), th(Thai), zh-Hant(Traditional Chinese), zh-Hans(Simplified Chinese), ms-my(Malaysian Malay). If the selected language is not supported in certain shop location, the response will be in default language.
	Language string `json:"language"`
}

type Category struct {
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

type GetCategoriesByCountryResponse struct {
	// List of categories info.
	Categories []Category `json:"categories"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetPaymentListRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// ID/MY/SG/VN/PH/TH/TW
	Country string `json:"country"`
}

type Method struct {
	// The payment method
	PaymentMethod string `json:"payment_method"`
	// The country for this payment method
	Country string `json:"country"`
}

type GetPaymentListResponse struct {
	//
	PaymentMethodList []Method `json:"payment_method_list"`
}
