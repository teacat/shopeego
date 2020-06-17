package shopeego

type GetShopsByPartnerRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetShopsByPartnerResponse struct {
	// A list of shops authorized to the partner.
	AuthedShops []GetShopsByPartnerResponseShop `json:"authed_shops,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetCategoriesByCountryRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Two-digit country code(capital letter).
	Country string `json:"country,omitempty"`
	// Is cross border or not. 1: cross border; 0: not cross border
	IsCB bool `json:"is_cb,omitempty"`
	// Indicate the translation language. Language values include: en(English), vi(Vietnamese), id(Indonesian), th(Thai), zh-Hant(Traditional Chinese), zh-Hans(Simplified Chinese), ms-my(Malaysian Malay). If the selected language is not supported in certain shop location, the response will be in default language.
	Language string `json:"language,omitempty"`
}

type GetCategoriesByCountryResponse struct {
	// List of categories info.
	Categories []GetCategoriesByCountryResponseCategory `json:"categories,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetPaymentListRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// ID/MY/SG/VN/PH/TH/TW
	Country string `json:"country,omitempty"`
}

type GetPaymentListResponse struct {
	//
	PaymentMethodList []GetPaymentListResponseMethod `json:"payment_method_list,omitempty"`
}
