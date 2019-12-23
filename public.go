package shopeego

type GetShopsByPartnerRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type SIP struct {
	// Affiliate Shop's area
	Country string
	// Affiliate shop's id
	AShopID int
}

type Shop struct {
	// The two-digit code representing the country where the order was made.
	Country string
	// Shopee's unique identifier for a shop.
	ShopID int
	// The timestamp when the shop was authorized to the partner.
	AuthTime int
	// SIP affiliate shops info list
	SIPAShops []SIP
	// Use this field to indicate the expiration date for shop authorization.
	ExpireTime int
}

type GetShopsByPartnerResponse struct {
	// A list of shops authorized to the partner.
	AuthedShops []Shop
	// The identifier for an API request for error tracking
	RequestID string
}

type GetCategoriesByCountryRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Two-digit country code(capital letter).
	Country string
	// Is cross border or not. 1: cross border; 0: not cross border
	IsCB bool
	// Indicate the translation language. Language values include: en(English), vi(Vietnamese), id(Indonesian), th(Thai), zh-Hant(Traditional Chinese), zh-Hans(Simplified Chinese), ms-my(Malaysian Malay). If the selected language is not supported in certain shop location, the response will be in default language.
	Language string
}

type Category struct {
	// The Identify of the parent of the category
	ParentID int
	// This is to indicate whether the category has children.
	HasChildren bool
	// The Identify of each category
	CategoryID int
	// The name of each category
	CategoryName string
	// To indicate if this category supports size chart
	IsSuppSizechart bool
}

type GetCategoriesByCountryResponse struct {
	// List of categories info.
	Categories []Category
	// The identifier for an API request for error tracking
	RequestID string
}

type GetPaymentListRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// ID/MY/SG/VN/PH/TH/TW
	Country string
}

type Method struct {
	// The payment method
	PaymentMethod string
	// The country for this payment method
	Country string
}

type GetPaymentListResponse struct {
	//
	PaymentMethodList []Method
}
