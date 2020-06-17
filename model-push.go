package shopeego

type GetPushConfigRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetPushConfigResponse struct {
	// The callback url of push mechanism.
	CallbackURL string `json:"callback_url,omitempty"`
	// The shutdown time caused by low successful rate of push mechanism.
	ShutTime int64 `json:"shut_time,omitempty"`
	//
	DeatiledConfig GetPushConfigResponseDeatiledConfig `json:"detailed_config,omitempty"`
	// Use this filed to set shops that need to be blocked.
	BlockedShopID []int64 `json:"blocked_shopid,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type SetPushConfigRequest struct {
	// The callback url of push mechanism.
	CallbackURL string `json:"callback_url,omitempty"`
	// The shutdown time caused by low successful rate of push mechanism.
	ShutTime int `json:"shut_time,omitempty"`
	//
	DeatiledConfig GetPushConfigResponseDeatiledConfig `json:"detailed_config,omitempty"`
	// Use this filed to set shops that need to be blocked.
	BlockedShopID []int64 `json:"blocked_shopid,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
}

type SetPushConfigResponse struct {
	// Use this field to indicate whether the configuration is set successfully.
	Status string `json:"status,omitempty"`
	// The identifier for an API request for error tracking.
	RequestID string `json:"request_id,omitempty"`
}
