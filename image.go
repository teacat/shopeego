package shopeego

type UploadImgRequest struct {
	// Image url. max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max number of image is 9.
	Images []string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Image struct {
	// origin image url
	ImageURL string
	// Shopee image url
	ShopeeImageURL string
}

type UploadImgResponse struct {
	//
	Images []Image
	// The identifier for an API request for error tracking
	RequestID string
}
