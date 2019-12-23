package shopeego

type UploadImgRequest struct {
	// Image url. max 2.0 MB each.Image format accepted: JPG, JPEG, PNG.Suggested dimension: 1024 x 1024 px. Max number of image is 9.
	Images []string `json:"images"`
	//
	RequestBase
}

type UploadImgResponse struct {
	//
	Images []UploadImgResponseImage `json:"images"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}
