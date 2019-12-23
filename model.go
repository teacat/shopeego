package shopeego

type RequestBase struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetShopInfoResponseAffiliateShop struct {
	// Affiliate shop's id.
	AShopID string `json:"a_shop_id"`
	// Affiliate Shop's area.
	Country string `json:"country"`
}