package shopeego

type UploadImgRequest struct {
	// Image url. max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max number of image is 9.
	Images []string `json:"images"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type UploadImgResponse struct {
	//
	Images []UploadImgResponseImage `json:"images"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}
