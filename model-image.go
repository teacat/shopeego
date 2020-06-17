package shopeego

type UploadImgRequest struct {
	// Image url. max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max number of image is 9.
	Images []string `json:"images,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UploadImgResponse struct {
	//
	Images []UploadImgResponseImage `json:"images,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}
