package shopeego

type GetAccessTokenRequest struct {
	// The code in redirect url after shop authorization. Valid for one-time use, expires in 10 minutes.
	Code string `json:"code,omitempty"`
	// The shop_id of the shop that authorized the developer.
	ShopID int64 `json:"shop_id,omitempty"`
	// The credential retrieved in the APP console.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetAccessTokenResponse struct {
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
	// Error code. Empty when the api call succeeded.
	Error string `json:"error,omitempty"`
	// Use refresh_token to obtain new access_token. Valid for one-time use, expires in 30 days.
	RefreshToken string `json:"refresh_token,omitempty"`
	// Use access_token as a common request parameter for certian APIs. Valid for multiple use, expires in 4 hours.
	AccessToken string `json:"access_token,omitempty"`
	// Access_token expiration time, unit is second.
	ExpireIn int `json:"expire_in,omitempty"`
}

type RefreshAccessTokenRequest struct {
	// Use refresh_token to obtain new access_token. Valid for one-time use, expires in 30 days.
	RefreshToken string `json:"refresh_token,omitempty"`
	// The shop_id of the shop that authorized the developer.
	ShopID int64 `json:"shop_id,omitempty"`
	// The credential retrieved in the APP console.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type RefreshAccessTokenResponse struct {
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
	// Error code. Empty when the api call succeeded.
	Error string `json:"error,omitempty"`
	// Use refresh_token to obtain new access_token. Valid for one-time use, expires in 30 days.
	RefreshToken string `json:"refresh_token,omitempty"`
	// Use access_token as a common request parameter for certian APIs. Valid for multiple use, expires in 4 hours.
	AccessToken string `json:"access_token,omitempty"`
	// Access_token expiration time, unit is second.
	ExpireIn int `json:"expire_in,omitempty"`
	// The credential retrieved in the APP console.
	PartnerID int64 `json:"partner_id,omitempty"`
	// The shop_id of the shop that authorized the developer.
	ShopID int64 `json:"shop_id,omitempty"`
}
