package shopeego

import "errors"

var (
	// There are errors in the input parameters
	ErrParams = errors.New("shopeego: there are errors in the input parameters")
	// The request is not authenticated. Ex: signature is wrong
	ErrAuth = errors.New("shopeego: the request is not authenticated. Ex: signature is wrong")
	// An error has occurred
	ErrServer = errors.New("shopeego: an error has occurred")
	// Too many request.Exceed 1000 request per min will be banned for an hour
	ErrTooManyRequest = errors.New("shopeego: too many request.Exceed 1000 request per min will be banned for an hour")
	// Not support action
	ErrNotSupport = errors.New("shopeego: not support action")
	// An inner error has occurred
	ErrInnerError = errors.New("shopeego: an inner error has occurred")
	// No Permission
	ErrPermission = errors.New("shopeego: no permission")
	// Not exist
	ErrNotExists = errors.New("shopeego: not exists")
)

type ClientOptions struct {
	Secret string
	// 非必要
	PartnerID int
	// 非必要
	ShopID int
}

// CONSTS

func NewClient(opts *ClientOptions) {

}
