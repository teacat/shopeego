package shopeego

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

const (
	// ClientVersionV1 for Open API V1, see: https://open.shopee.com/documents?module=63&type=2&id=53.
	ClientVersionV1 ClientVersion = iota
	// ClientVersionV2 for Open API V2, see: https://open.shopee.com/documents?module=63&type=2&id=56.
	ClientVersionV2
)

// ClientVersion is the API version of the client should be using.
type ClientVersion int

const (
	urlTestSandbox string = "https://partner.uat.shopeemobile.com/"
	urlStandard    string = "https://partner.shopeemobile.com/"
)

type ClientOptions struct {
	Secret string
	// 非必要
	//PartnerID int64
	// 非必要
	//ShopID int64
	//
	IsSandbox bool
	//
	Version ClientVersion
}

// CONSTS

func NewClient(opts *ClientOptions) Client {
	return &ShopeeClient{
		Secret:    opts.Secret,
		IsSandbox: opts.IsSandbox,
		Version:   opts.Version,
	}
}

// Client 定義了一個蝦皮的客戶端該有什麼功能。
type Client interface {
	// SetAccessToken(t string) *ShopeeClient

	//=======================================================
	// Shop
	//=======================================================

	// GetShopInfo Use this call to get information of shop
	GetShopInfo(*GetShopInfoRequest) (*GetShopInfoResponse, error)
	// UpdateShopInfo Use this call to update information of shop
	UpdateShopInfo(*UpdateShopInfoRequest) (*UpdateShopInfoResponse, error)
	// Performance Shop performance includes the indexes from "My Performance" of Seller Center.
	Performance(*PerformanceRequest) (*PerformanceResponse, error)
	// SetShopInstallmentStatus Only for TW whitelisted shop.Use this API to set the installment status of shop.
	SetShopInstallmentStatus(*SetShopInstallmentStatusRequest) (*SetShopInstallmentStatusResponse, error)
	//
	AuthPartner(*AuthPartnerRequest) string

	//=======================================================
	// ShopCategory
	//=======================================================

	// AddShopCategory Use this call to add a new collecion
	AddShopCategory(*AddShopCategoryRequest) (*AddShopCategoryResponse, error)
	// GetShopCategories Use this call to get list of in-shop categories
	GetShopCategories(*GetShopCategoriesRequest) (*GetShopCategoriesResponse, error)
	// DeleteShopCategory Use this call to delete a existing collecion
	DeleteShopCategory(*DeleteShopCategoryRequest) (*DeleteShopCategoryResponse, error)
	// UpdateShopCategory Use this call to update a existing collecion
	UpdateShopCategory(*UpdateShopCategoryRequest) (*UpdateShopCategoryResponse, error)
	// AddItems Use this call to add items list to certain shop_category
	AddItems(*AddItemsRequest) (*AddItemsResponse, error)
	// GetItems Use this call to get items list of certain shop_category
	GetItems(*GetItemsRequest) (*GetItemsResponse, error)
	// DeleteItems Use this api to delete items from shop category
	DeleteItems(*DeleteItemsRequest) (*DeleteItemsResponse, error)

	//=======================================================
	// Item
	//=======================================================

	// GetCategories Use this call to get categories of product item
	GetCategories(*GetCategoriesRequest) (*GetCategoriesResponse, error)
	// GetAttributes Use this call to get attributes of product item
	GetAttributes(*GetAttributesRequest) (*GetAttributesResponse, error)
	// Add Use this call to add a product item. Should get dependency by calling below API first
	// shopee.item.GetCategories
	// shopee.item.GetAttributes
	// shopee.logistics.GetLogistics
	Add(*AddRequest) (*AddResponse, error)
	// Delete Use this call to delete a product item.
	Delete(*DeleteRequest) (*DeleteResponse, error)
	// UnlistItem Use this api to unlist or list items in batch.
	UnlistItem(*UnlistItemRequest) (*UnlistItemResponse, error)
	// AddVariations Use this call to add item variations, it only supports non-tier-variation items.
	AddVariations(*AddVariationsRequest) (*AddVariationsResponse, error)
	// DeleteVariation Use this call to delete item variation
	DeleteVariation(*DeleteVariationRequest) (*DeleteVariationResponse, error)
	// GetItemsList Use this call to get a list of items
	GetItemsList(*GetItemsListRequest) (*GetItemsListResponse, error)
	// GetItemDetail Use this call to get detail of item
	GetItemDetail(*GetItemDetailRequest) (*GetItemDetailResponse, error)
	// UpdateItem Use this call to update a product item. Should get dependency by calling below API first
	// shopee.item.GetItemDetail
	UpdateItem(*UpdateItemRequest) (*UpdateItemResponse, error)
	// AddItemImg Use this call to add product item images.
	AddItemImg(*AddItemImgRequest) (*AddItemImgResponse, error)
	// UpdateItemImg Override and update all the existing images of an item.
	UpdateItemImg(*UpdateItemImgRequest) (*UpdateItemImgResponse, error)
	// InsertItemImg Use this call to add one item image in assigned position.
	InsertItemImg(*InsertItemImgRequest) (*InsertItemImgResponse, error)
	// DeleteItemImg Use this call to delete a product item image.
	DeleteItemImg(*DeleteItemImgRequest) (*DeleteItemImgResponse, error)
	// UpdatePrice Use this call to update item price
	UpdatePrice(*UpdatePriceRequest) (*UpdatePriceResponse, error)
	// UpdateStock Use this call to update item stock
	UpdateStock(*UpdateStockRequest) (*UpdateStockResponse, error)
	// UpdateVariationPrice Use this call to update item variation price
	UpdateVariationPrice(*UpdateVariationPriceRequest) (*UpdateVariationPriceResponse, error)
	// UpdateVariationStock Use this call to update item variation stock
	UpdateVariationStock(*UpdateVariationStockRequest) (*UpdateVariationStockResponse, error)
	// UpdatePriceBatch Update items price in batch.
	UpdatePriceBatch(*UpdatePriceBatchRequest) (*UpdatePriceBatchResponse, error)
	// UpdateStockBatch Update items stock in batch.
	UpdateStockBatch(*UpdateStockBatchRequest) (*UpdateStockBatchResponse, error)
	// UpdateVariationPriceBatch Update variations price in batch.
	UpdateVariationPriceBatch(*UpdateVariationPriceBatchRequest) (*UpdateVariationPriceBatchResponse, error)
	// UpdateVariationStockBatch Update variations stock in batch.
	UpdateVariationStockBatch(*UpdateVariationStockBatchRequest) (*UpdateVariationStockBatchResponse, error)
	// InitTierVariation Initialize a non-tier-variation item to a tier-variation item, upload variation image and initialize stock and price for each variation. This API cannot edit existed tier_variation and variation price/stock.
	InitTierVariation(*InitTierVariationRequest) (*InitTierVariationResponse, error)
	// AddTierVariation Use this api to add new tier variations in batch. Tier variation index of variations in the same item must be unique.
	AddTierVariation(*AddTierVariationRequest) (*AddTierVariationResponse, error)
	// GetVariation Use this call to get tier-variation basic information under an item
	GetVariation(*GetVariationRequest) (*GetVariationResponse, error)
	// UpdateTierVariationList Use this api to update tier-variation list or upload variation image of a tier-variation item.
	UpdateTierVariationList(*UpdateTierVariationListRequest) (*UpdateTierVariationListResponse, error)
	// UpdateTierVariationIndex Use this api to update existing tier index under the same variation_id.
	UpdateTierVariationIndex(*UpdateTierVariationIndexRequest) (*UpdateTierVariationIndexResponse, error)
	// BoostItem Use this api to boost multiple items at once.
	BoostItem(*BoostItemRequest) (*BoostItemResponse, error)
	// GetBoostedItem Use this api to get all boosted items.
	GetBoostedItem(*GetBoostedItemRequest) (*GetBoostedItemResponse, error)
	// SetItemInstallmentTenures Only for TW whitelisted shop. Use this API to set the installment tenures of items.
	SetItemInstallmentTenures(*SetItemInstallmentTenuresRequest) (*SetItemInstallmentTenuresResponse, error)
	// GetPromotionInfo Use this api to get ongoing and upcoming promotions.
	GetPromotionInfo(*GetPromotionInfoRequest) (*GetPromotionInfoResponse, error)
	// GetRecommendCats Use this API to get recommended category ids according to item name.
	GetRecommendCats(*GetRecommendCatsRequest) (*GetRecommendCatsResponse, error)
	// GetComment Use this api to get comment by shopid/itemid/comment_id
	GetComment(*GetCommentRequest) (*GetCommentResponse, error)
	// ReplyComments Use this api to reply comments from buyers in batch
	ReplyComments(*ReplyCommentsRequest) (*ReplyCommentsResponse, error)

	//=======================================================
	// Image
	//=======================================================

	// UploadImg Use this optional API to pre-validate your image urls and convert them to Shopee image url to use in item upload APIs. This way your potential invalid urls will not block your item upload process.
	UploadImg(*UploadImgRequest) (*UploadImgResponse, error)

	//=======================================================
	// Discount
	//=======================================================

	// AddDiscount Use this api to add shop discount activity
	AddDiscount(*AddDiscountRequest) (*AddDiscountResponse, error)
	// AddDiscountItem Use this api to add shop discount item
	AddDiscountItem(*AddDiscountItemRequest) (*AddDiscountItemResponse, error)
	// DeleteDiscount Use this api to delete one discount activity BEFORE it starts.
	DeleteDiscount(*DeleteDiscountRequest) (*DeleteDiscountResponse, error)
	// DeleteDiscountItem Use this api to delete items of the discount activity
	DeleteDiscountItem(*DeleteDiscountItemRequest) (*DeleteDiscountItemResponse, error)
	// GetDiscountDetail Use this api to get one shop discount activity detail
	GetDiscountDetail(*GetDiscountDetailRequest) (*GetDiscountDetailResponse, error)
	// GetDiscountsList Use this api to get shop discount activity list
	GetDiscountsList(*GetDiscountsListRequest) (*GetDiscountsListResponse, error)
	// UpdateDiscount Use this api to update one discount information
	UpdateDiscount(*UpdateDiscountRequest) (*UpdateDiscountResponse, error)
	// UpdateDiscountItems Use this api to update items of the discount activity
	UpdateDiscountItems(*UpdateDiscountItemsRequest) (*UpdateDiscountItemsResponse, error)

	//=======================================================
	// Orders
	//=======================================================

	// GetOrdersList GetOrdersList is the recommended call to use for order management. Use this call to retrieve basic information of all orders which are updated within specific period of time. More details of each order can be retrieved from GetOrderDetails. [Only the recent one month orders can be fetch through this API. Please use GetOrderBySatus API to fetch more orders.]
	GetOrdersList(*GetOrdersListRequest) (*GetOrdersListResponse, error)
	// GetOrdersByStatus GetOrdersByStatus is the recommended call to use for order management. Use this call to retrieve basic information of all orders which are specific status. More details of each order can be retrieved from GetOrderDetails.
	GetOrdersByStatus(*GetOrdersByStatusRequest) (*GetOrdersByStatusResponse, error)
	// GetOrderDetails Use this call to retrieve detailed information about one or more orders based on OrderSN.
	GetOrderDetails(*GetOrderDetailsRequest) (*GetOrderDetailsResponse, error)
	// GetEscrowDetails Use this call to retrieve detailed escrow information about one order based on OrderSN.
	GetEscrowDetails(*GetEscrowDetailsRequest) (*GetEscrowDetailsResponse, error)
	// AddOrderNote Use this call to add note for an order
	AddOrderNote(*AddOrderNoteRequest) (*AddOrderNoteResponse, error)
	// CancelOrder Use this call to cancel an order from the seller side.
	CancelOrder(*CancelOrderRequest) (*CancelOrderResponse, error)
	// AcceptBuyerCancellation Use this call to accept buyer cancellation
	AcceptBuyerCancellation(*AcceptBuyerCancellationRequest) (*AcceptBuyerCancellationResponse, error)
	// RejectBuyerCancellation Use this call to reject buyer cancellation
	RejectBuyerCancellation(*RejectBuyerCancellationRequest) (*RejectBuyerCancellationResponse, error)
	// GetForderInfo Use this call to retrieve detailed information of all the fulfill orders(forder) under a single regular order based on ordersn.
	GetForderInfo(*GetForderInfoRequest) (*GetForderInfoResponse, error)
	// GetEscrowReleasedOrders Use this api to get orders' release time and escrow amount.
	GetEscrowReleasedOrders(*GetEscrowReleasedOrdersRequest) (*GetEscrowReleasedOrdersResponse, error)
	// SplitOrder Use this API to split order into fulfillment orders. This feature is only enabled for whitelisted shops.
	SplitOrder(*SplitOrderRequest) (*SplitOrderResponse, error)
	// UndoSplitOrder Use this API to cancel split order from the seller side.
	UndoSplitOrder(*UndoSplitOrderRequest) (*UndoSplitOrderResponse, error)
	// GetUnbindOrderList Use this call to get a list of unbind orders.
	GetUnbindOrderList(*GetUnbindOrderListRequest) (*GetUnbindOrderListResponse, error)
	// MyIncome Use this API to fetch the accounting detail of order.
	MyIncome(*MyIncomeRequest) (*MyIncomeResponse, error)

	//=======================================================
	// Logistics
	//=======================================================

	// GetLogistics Use this call to get all supported logistic channels.
	GetLogistics(*GetLogisticsRequest) (*GetLogisticsResponse, error)
	// UpdateShopLogistics Configure shop level logistics
	UpdateShopLogistics(*UpdateShopLogisticsRequest) (*UpdateShopLogisticsResponse, error)
	// GetParameterForInit Use this call to get all required param for logistic initiation.
	GetParameterForInit(*GetParameterForInitRequest) (*GetParameterForInitResponse, error)
	// GetAddress For integrated logistics channel, use this call to get pickup address for pickup mode order.
	GetAddress(*GetAddressRequest) (*GetAddressResponse, error)
	// GetTimeSlot For integrated logistics channel, use this call to get pickup timeslot for pickup mode order.
	GetTimeSlot(*GetTimeSlotRequest) (*GetTimeSlotResponse, error)
	// GetBranch For integrated logistics channel, use this call to get dropoff location for dropoff mode order.
	GetBranch(*GetBranchRequest) (*GetBranchResponse, error)
	// GetLogisticInfo Get all the logistics info of an order to Init.
	// This API consolidates the output of GetParameterForInit, GetAddresss, GetTimeSlot and GetBranch based on each order so that developer can get all the required parameters ready in this API for Init.This API is an alternative of GetParameterForInit, GetAddresss, GetTimeSlot and GetBranch as a set.
	GetLogisticInfo(*GetLogisticInfoRequest) (*GetLogisticInfoResponse, error)
	// Init Use this call to initiate logistics including arrange Pickup, Dropoff or shipment for non-integrated logistic channels. Should call shopee.logistics.GetLogisticInfo to fetch all required param first. It's recommended to initiate logistics one hour AFTER the orders were placed since there is one-hour window buyer can cancel any order without request to seller.
	Init(*InitRequest) (*InitResponse, error)
	// GetAirwayBill Use this API to get airway bill for orders. AirwayBill is only fetchable when the order status is under READY_TO_SHIP and RETRY_SHIP.
	GetAirwayBill(*GetAirwayBillRequest) (*GetAirwayBillResponse, error)
	// GetOrderLogistics  Use this call to fetch the logistics information of an order, these info can be used for airwaybill printing. Dedicated for crossborder SLS order airwaybill. May not be applicable for local channel airwaybill.
	GetOrderLogistics(*GetOrderLogisticsRequest) (*GetOrderLogisticsResponse, error)
	// GetLogisticsMessage Use this call to get the logistics tracking information of an order.
	GetLogisticsMessage(*GetLogisticsMessageRequest) (*GetLogisticsMessageResponse, error)
	// GetForderWaybill Use this API to get airwaybill for fulfillment orders.
	GetForderWaybill(*GetForderWaybillRequest) (*GetForderWaybillResponse, error)
	// SetAddress Use this API to set default_address/pick_up_address/return_address of shop. Please use GetAddress API to fetch the address_id.
	SetAddress(*SetAddressRequest) (*SetAddressResponse, error)
	// DeleteAddress Use this API to delete the default/pick_up/return address of shop by address_id. Please use GetAddress API to fetch the address_id.
	DeleteAddress(*DeleteAddressRequest) (*DeleteAddressResponse, error)

	//=======================================================
	// Returns
	//=======================================================

	// ConfirmReturn Confirm return
	ConfirmReturn(*ConfirmReturnRequest) (*ConfirmReturnResponse, error)
	// DisputeReturn Dispute return
	DisputeReturn(*DisputeReturnRequest) (*DisputeReturnResponse, error)
	// GetReturnList Get return list
	GetReturnList(*GetReturnListRequest) (*GetReturnListResponse, error)
	// GetReturnDetail Use this api to get detail information of a returned order
	GetReturnDetail(*GetReturnDetailRequest) (*GetReturnDetailResponse, error)

	//=======================================================
	// Public
	//=======================================================

	// GetShopsByPartner Use this call to get basic info of shops which have authorized to the partner.
	GetShopsByPartner(*GetShopsByPartnerRequest) (*GetShopsByPartnerResponse, error)
	// GetCategoriesByCountry Use this api to get categories list filtered by country and cross border without using shopID.
	GetCategoriesByCountry(*GetCategoriesByCountryRequest) (*GetCategoriesByCountryResponse, error)
	// GetPaymentList The supported payment method list by country
	GetPaymentList(*GetPaymentListRequest) (*GetPaymentListResponse, error)

	//=======================================================
	// TopPicks
	//=======================================================

	// GetTopPicksList Get the list of all collections.
	GetTopPicksList(*GetTopPicksListRequest) (*GetTopPicksListResponse, error)
	// AddTopPicks Add one collection. One shop can have up to 10 collections.
	AddTopPicks(*AddTopPicksRequest) (*AddTopPicksResponse, error)
	// UpdateTopPicks Use this API to update the collection name, the item list in a collection, or to activate a collection.
	UpdateTopPicks(*UpdateTopPicksRequest) (*UpdateTopPicksResponse, error)
	// DeleteTopPicks Delete a collection
	DeleteTopPicks(*DeleteTopPicksRequest) (*DeleteTopPicksResponse, error)

	//=======================================================
	// FirstMileTracking
	//=======================================================

	// GenerateFMTrackingNo Use this API to generate first-mile tracking number for the shipment method of pickup. Please note that the prerequisite for using this API is that the order status is ready_to_ship and the tracking number of order has been obtained. Only applicable to cross-border sellers in China.
	GenerateFMTrackingNo(*GenerateFMTrackingNoRequest) (*GenerateFMTrackingNoResponse, error)
	// GetShopFMTrackingNo Use this API to fetch first-mile tracking numbers of the shop. Only applicable to cross-border sellers in China.
	GetShopFMTrackingNo(*GetShopFMTrackingNoRequest) (*GetShopFMTrackingNoResponse, error)
	// FirstMileCodeBindOrder Use this API to bind orders with the first-mile tracking number. Only applicable to cross-border sellers in China.
	FirstMileCodeBindOrder(*FirstMileCodeBindOrderRequest) (*FirstMileCodeBindOrderResponse, error)
	// GetFmTnDetail Use this API to fetch the detailed information of first-mile tracking number. Only applicable to cross-border sellers in China.
	GetFmTnDetail(*GetFmTnDetailRequest) (*GetFmTnDetailResponse, error)
	// GetFMTrackingNoWaybill Use the API to get the waybill of first-mile tracking number.Please note that this API only used for the shipment method of pickup. Only applicable to cross-border sellers in China.
	GetFMTrackingNoWaybill(*GetFMTrackingNoWaybillRequest) (*GetFMTrackingNoWaybillResponse, error)
	// GetShopFirstMileChannel Use this call to get all supported logistic channels for first mile. Only applicable to cross-border sellers in China.
	GetShopFirstMileChannel(*GetShopFirstMileChannelRequest) (*GetShopFirstMileChannelResponse, error)
	// FirstMileUnbind Use this API to unbind orders with the first-mile tracking number. Only applicable to cross-border sellers in China.
	FirstMileUnbind(*FirstMileUnbindRequest) (*FirstMileUnbindResponse, error)

	//=======================================================
	// Payment
	//=======================================================

	// GetTransactionList Use this API to get the transaction records of wallet.
	GetTransactionList(*GetTransactionListRequest) (*GetTransactionListResponse, error)

	//=======================================================
	// Push
	//=======================================================

	// GetPushConfig Use this API to get the configuration information of push service.
	GetPushConfig(*GetPushConfigRequest) (*GetPushConfigResponse, error)
	// SetPushConfig Use this API to set the configuration information of push service.
	SetPushConfig(*SetPushConfigRequest) (*SetPushConfigResponse, error)

	//=======================================================
	// Auth (V2)
	//=======================================================

	// GetAccessToken Use this API and the code to obtain the access_token and refresh_token.
	GetAccessToken(*GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	// RefreshAccessToken Use this API to refresh the access_token after it expires.
	RefreshAccessToken(*RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error)
}

// ShopeeClient represents a client to Shopee
type ShopeeClient struct {
	Secret      string
	accessToken string
	IsSandbox   bool
	Version     ClientVersion
}

// ResponseError defines a error response
type ResponseError struct {
	RequestID string `json:"request_id,omitempty"`
	Msg       string `json:"msg,omitempty"`
	ErrorType string `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (e ResponseError) Error() string {
	// 如果 `msg` 是空的，就試試看 V2 的 `message`
	msg := e.Msg
	if msg == "" {
		msg = e.Message
	}
	return fmt.Sprintf("shopeego: %s - %s", e.ErrorType, msg)
}

//
func (s *ShopeeClient) getPath(method string) string {
	var host string
	if s.IsSandbox {
		host = urlTestSandbox
	} else {
		host = urlStandard
	}
	// 依照不同版本處理 API 前輟。
	switch s.Version {
	case ClientVersionV1:
		return fmt.Sprintf("%sapi/v1/%s", host, availablePaths[method])
	default:
		return fmt.Sprintf("%sapi/v2/%s", host, availablePaths[method])
	}
}

// signV1 會簽署 V1 API。
func (s *ShopeeClient) signV1(url string, body []byte) string {
	h := hmac.New(sha256.New, []byte(s.Secret))
	io.WriteString(h, fmt.Sprintf("%s|%s", url, string(body)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// signV2 會簽署 V2 API。
func (s *ShopeeClient) signV2(url string, b []byte) string {
	h := hmac.New(sha256.New, []byte(s.Secret))
	p := s.getBodyPart(b)
	if s.IsSandbox {
		url = "/api/" + strings.TrimLeft(url, urlTestSandbox)
	} else {
		url = "/api/" + strings.TrimLeft(url, urlStandard)
	}
	io.WriteString(h, fmt.Sprintf("%d%s%d%s%d", p.partnerID, url, p.timestamp, s.accessToken, p.shopID))
	return fmt.Sprintf("%x", h.Sum(nil))
}

type bodyPart struct {
	partnerID int64
	timestamp int
	shopID    int64
}

//
func (s *ShopeeClient) getBodyPart(b []byte) bodyPart {
	var p bodyPart
	var t map[string]interface{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		return p
	}
	if v, ok := t["partner_id"].(float64); ok {
		p.partnerID = int64(v)
	}
	if v, ok := t["timestamp"].(float64); ok {
		p.timestamp = int(v)
	}
	if v, ok := t["shopid"].(float64); ok {
		p.shopID = int64(v)
	}
	if v, ok := t["shop_id"].(float64); ok {
		p.shopID = int64(v)
	}
	return p
}

func (s *ShopeeClient) SetAccessToken(t string) *ShopeeClient {
	s.accessToken = t
	return s
}

//
func (s *ShopeeClient) makeV2Query(url string, b []byte) string {
	p := s.getBodyPart(b)
	q := fmt.Sprintf("?partner_id=%d&shop_id=%d&timestamp=%d&sign=%s", p.partnerID, p.shopID, p.timestamp, s.signV2(url, b))
	if s.accessToken != "" {
		q += fmt.Sprintf("&access_token=%s", s.accessToken)
	}
	return q
}

//
func (s *ShopeeClient) post(method string, in interface{}) ([]byte, error) {
	body, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	url := s.getPath(method)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	req.Header.Add("Content-Type", "application/json")

	switch s.Version {
	// 如果是 V1 就在 Header 安插 Sign。
	case ClientVersionV1:
		req.Header.Add("Authorization", s.signV1(url, body))
	// 如果是 V2 的 API，就在 Body 中自動安插 Sign。
	case ClientVersionV2:
		req, err = http.NewRequest("POST", fmt.Sprintf("%s%s", url, s.makeV2Query(url, body)), bytes.NewReader(body))
		req.Header.Add("Authorization", s.signV1(url, body))
	}

	//Do request by native lib
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var errResp ResponseError
	_ = json.Unmarshal(body, &errResp)
	if errResp.ErrorType != "" {
		return nil, errResp
	}

	return s.patchFloat(body), nil
}

// patchFloat 會修正無法將 JSON 的字串值轉換成 Float64 型態的錯誤，這主要是因為 Shopee 會在某些時候以 `""`（空字串）當作是 Float64 的零值，
// 但對於 Golang 來說，Float64 的零值必須是 `"0"`，所以這個函式會將 Raw JSON 裡的所有相關（請參閱 `replaces.go` 關鍵詞表） `""` 轉換為 `"0"` 以便正確解析。
//
// 這個函式修正了 `json: invalid use of ,string struct tag, trying to unmarshal unquoted value into float64` 錯誤。
// 相關 Issue： https://github.com/teacat/shopeego/issues/6
func (s *ShopeeClient) patchFloat(body []byte) []byte {
	replaceConcat := strings.Join(replaces, "|")
	for _, v := range replaces {
		body = []byte(strings.ReplaceAll(string(body), fmt.Sprintf(`"%s":""`, v), fmt.Sprintf(`"%s":"0"`, v)))
	}

	var r = regexp.MustCompile(fmt.Sprintf(`"(%s)":(.*?)(,|})`, replaceConcat))
	return []byte(r.ReplaceAllString(string(body), `"$1":"$2"$3`))
}

//=======================================================
// Shop
//=======================================================

// GetShopInfo Use this call to get information of shop
func (s *ShopeeClient) GetShopInfo(req *GetShopInfoRequest) (resp *GetShopInfoResponse, err error) {
	b, err := s.post("GetShopInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateShopInfo Use this call to update information of shop
func (s *ShopeeClient) UpdateShopInfo(req *UpdateShopInfoRequest) (resp *UpdateShopInfoResponse, err error) {
	b, err := s.post("UpdateShopInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// Performance Shop performance includes the indexes from "My Performance" of Seller Center.
func (s *ShopeeClient) Performance(req *PerformanceRequest) (resp *PerformanceResponse, err error) {
	b, err := s.post("Performance", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// SetShopInstallmentStatus Only for TW whitelisted shop.Use this API to set the installment status of shop.
func (s *ShopeeClient) SetShopInstallmentStatus(req *SetShopInstallmentStatusRequest) (resp *SetShopInstallmentStatusResponse, err error) {
	b, err := s.post("SetShopInstallmentStatus", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// AuthPartner for V2.
func (s *ShopeeClient) AuthPartner(req *AuthPartnerRequest) string {
	url := s.getPath("AuthPartner")
	switch s.Version {
	// V1
	case ClientVersionV1:
		h := sha256.Sum256([]byte(fmt.Sprintf("%s%s", s.Secret, req.Redirect)))
		token := fmt.Sprintf("%x", h[:])
		return fmt.Sprintf("%s?id=%d&token=%s&redirect=%s", url, req.PartnerID, token, req.Redirect)
	// V2
	default:
		h := hmac.New(sha256.New, []byte(s.Secret))
		io.WriteString(h, fmt.Sprintf("%d%s%d", req.PartnerID, "/api/v2/shop/auth_partner", req.Timestamp))
		token := fmt.Sprintf("%x", h.Sum(nil))
		return fmt.Sprintf("%s?partner_id=%d&redirect=%s&sign=%s&timestamp=%d", url, req.PartnerID, req.Redirect, token, req.Timestamp)
	}
}

//=======================================================
// ShopCategory
//=======================================================

// AddShopCategory Use this call to add a new collecion
func (s *ShopeeClient) AddShopCategory(req *AddShopCategoryRequest) (resp *AddShopCategoryResponse, err error) {
	b, err := s.post("AddShopCategory", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetShopCategories Use this call to get list of in-shop categories
func (s *ShopeeClient) GetShopCategories(req *GetShopCategoriesRequest) (resp *GetShopCategoriesResponse, err error) {
	b, err := s.post("GetShopCategories", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// DeleteShopCategory Use this call to delete a existing collecion
func (s *ShopeeClient) DeleteShopCategory(req *DeleteShopCategoryRequest) (resp *DeleteShopCategoryResponse, err error) {
	b, err := s.post("DeleteShopCategory", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateShopCategory Use this call to update a existing collecion
func (s *ShopeeClient) UpdateShopCategory(req *UpdateShopCategoryRequest) (resp *UpdateShopCategoryResponse, err error) {
	b, err := s.post("UpdateShopCategory", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// AddItems Use this call to add items list to certain shop_category
func (s *ShopeeClient) AddItems(req *AddItemsRequest) (resp *AddItemsResponse, err error) {
	b, err := s.post("AddItems", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetItems Use this call to get items list of certain shop_category
func (s *ShopeeClient) GetItems(req *GetItemsRequest) (resp *GetItemsResponse, err error) {
	b, err := s.post("GetItems", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// DeleteItems Use this api to delete items from shop category
func (s *ShopeeClient) DeleteItems(req *DeleteItemsRequest) (resp *DeleteItemsResponse, err error) {
	b, err := s.post("DeleteItems", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// Item
//=======================================================

// GetCategories Use this call to get categories of product item
func (s *ShopeeClient) GetCategories(req *GetCategoriesRequest) (resp *GetCategoriesResponse, err error) {
	b, err := s.post("GetCategories", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetAttributes Use this call to get attributes of product item
func (s *ShopeeClient) GetAttributes(req *GetAttributesRequest) (resp *GetAttributesResponse, err error) {
	b, err := s.post("GetAttributes", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// Add Use this call to add a product item. Should get dependency by calling below API first
// shopee.item.GetCategories
// shopee.item.GetAttributes
// shopee.logistics.GetLogistics
func (s *ShopeeClient) Add(req *AddRequest) (resp *AddResponse, err error) {
	b, err := s.post("Add", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// Delete Use this call to delete a product item.
func (s *ShopeeClient) Delete(req *DeleteRequest) (resp *DeleteResponse, err error) {
	b, err := s.post("Delete", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UnlistItem Use this api to unlist or list items in batch.
func (s *ShopeeClient) UnlistItem(req *UnlistItemRequest) (resp *UnlistItemResponse, err error) {
	b, err := s.post("UnlistItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// AddVariations Use this call to add item variations
func (s *ShopeeClient) AddVariations(req *AddVariationsRequest) (resp *AddVariationsResponse, err error) {
	b, err := s.post("AddVariations", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// DeleteVariation Use this call to delete item variation
func (s *ShopeeClient) DeleteVariation(req *DeleteVariationRequest) (resp *DeleteVariationResponse, err error) {
	b, err := s.post("DeleteVariation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetItemsList Use this call to get a list of items
func (s *ShopeeClient) GetItemsList(req *GetItemsListRequest) (resp *GetItemsListResponse, err error) {
	b, err := s.post("GetItemsList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetItemDetail Use this call to get detail of item
func (s *ShopeeClient) GetItemDetail(req *GetItemDetailRequest) (resp *GetItemDetailResponse, err error) {
	b, err := s.post("GetItemDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateItem Use this call to update a product item. Should get dependency by calling below API first
// shopee.item.GetItemDetail
func (s *ShopeeClient) UpdateItem(req *UpdateItemRequest) (resp *UpdateItemResponse, err error) {
	b, err := s.post("UpdateItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// AddItemImg Use this call to add product item images.
func (s *ShopeeClient) AddItemImg(req *AddItemImgRequest) (resp *AddItemImgResponse, err error) {
	b, err := s.post("AddItemImg", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateItemImg Override and update all the existing images of an item.
func (s *ShopeeClient) UpdateItemImg(req *UpdateItemImgRequest) (resp *UpdateItemImgResponse, err error) {
	b, err := s.post("UpdateItemImg", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// InsertItemImg Use this call to add one item image in assigned position.
func (s *ShopeeClient) InsertItemImg(req *InsertItemImgRequest) (resp *InsertItemImgResponse, err error) {
	b, err := s.post("InsertItemImg", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// DeleteItemImg Use this call to delete a product item image.
func (s *ShopeeClient) DeleteItemImg(req *DeleteItemImgRequest) (resp *DeleteItemImgResponse, err error) {
	b, err := s.post("DeleteItemImg", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdatePrice Use this call to update item price
func (s *ShopeeClient) UpdatePrice(req *UpdatePriceRequest) (resp *UpdatePriceResponse, err error) {
	b, err := s.post("UpdatePrice", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateStock Use this call to update item stock
func (s *ShopeeClient) UpdateStock(req *UpdateStockRequest) (resp *UpdateStockResponse, err error) {
	b, err := s.post("UpdateStock", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateVariationPrice Use this call to update item variation price
func (s *ShopeeClient) UpdateVariationPrice(req *UpdateVariationPriceRequest) (resp *UpdateVariationPriceResponse, err error) {
	b, err := s.post("UpdateVariationPrice", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateVariationStock Use this call to update item variation stock
func (s *ShopeeClient) UpdateVariationStock(req *UpdateVariationStockRequest) (resp *UpdateVariationStockResponse, err error) {
	b, err := s.post("UpdateVariationStock", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdatePriceBatch Update items price in batch.
func (s *ShopeeClient) UpdatePriceBatch(req *UpdatePriceBatchRequest) (resp *UpdatePriceBatchResponse, err error) {
	b, err := s.post("UpdatePriceBatch", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateStockBatch Update items stock in batch.
func (s *ShopeeClient) UpdateStockBatch(req *UpdateStockBatchRequest) (resp *UpdateStockBatchResponse, err error) {
	b, err := s.post("UpdateStockBatch", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateVariationPriceBatch Update variations price in batch.
func (s *ShopeeClient) UpdateVariationPriceBatch(req *UpdateVariationPriceBatchRequest) (resp *UpdateVariationPriceBatchResponse, err error) {
	b, err := s.post("UpdateVariationPriceBatch", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateVariationStockBatch Update variations stock in batch.
func (s *ShopeeClient) UpdateVariationStockBatch(req *UpdateVariationStockBatchRequest) (resp *UpdateVariationStockBatchResponse, err error) {
	b, err := s.post("UpdateVariationStockBatch", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// InitTierVariation Initialize a non-tier-variation item to a tier-variation item, upload variation image and initialize stock and price for each variation. This API cannot edit existed tier_variation and variation price/stock.
func (s *ShopeeClient) InitTierVariation(req *InitTierVariationRequest) (resp *InitTierVariationResponse, err error) {
	b, err := s.post("InitTierVariation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// AddTierVariation Use this api to add new tier variations in batch. Tier variation index of variations in the same item must be unique.
func (s *ShopeeClient) AddTierVariation(req *AddTierVariationRequest) (resp *AddTierVariationResponse, err error) {
	b, err := s.post("AddTierVariation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetVariation Use this call to get tier-variation basic information under an item
func (s *ShopeeClient) GetVariation(req *GetVariationRequest) (resp *GetVariationResponse, err error) {
	b, err := s.post("GetVariation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateTierVariationList Use this api to update tier-variation list or upload variation image of a tier-variation item.
func (s *ShopeeClient) UpdateTierVariationList(req *UpdateTierVariationListRequest) (resp *UpdateTierVariationListResponse, err error) {
	b, err := s.post("UpdateTierVariationList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateTierVariationIndex Use this api to update existing tier index under the same variation_id.
func (s *ShopeeClient) UpdateTierVariationIndex(req *UpdateTierVariationIndexRequest) (resp *UpdateTierVariationIndexResponse, err error) {
	b, err := s.post("UpdateTierVariationIndex", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// BoostItem Use this api to boost multiple items at once.
func (s *ShopeeClient) BoostItem(req *BoostItemRequest) (resp *BoostItemResponse, err error) {
	b, err := s.post("BoostItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetBoostedItem Use this api to get all boosted items.
func (s *ShopeeClient) GetBoostedItem(req *GetBoostedItemRequest) (resp *GetBoostedItemResponse, err error) {
	b, err := s.post("GetBoostedItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// SetItemInstallmentTenures Only for TW whitelisted shop. Use this API to set the installment tenures of items.
func (s *ShopeeClient) SetItemInstallmentTenures(req *SetItemInstallmentTenuresRequest) (resp *SetItemInstallmentTenuresResponse, err error) {
	b, err := s.post("SetItemInstallmentTenures", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetPromotionInfo Use this api to get ongoing and upcoming promotions.
func (s *ShopeeClient) GetPromotionInfo(req *GetPromotionInfoRequest) (resp *GetPromotionInfoResponse, err error) {
	b, err := s.post("GetPromotionInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetRecommendCats Use this API to get recommended category ids according to item name.
func (s *ShopeeClient) GetRecommendCats(req *GetRecommendCatsRequest) (resp *GetRecommendCatsResponse, err error) {
	b, err := s.post("GetRecommendCats", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetComment Use this api to get comment by shopid/itemid/comment_id
func (s *ShopeeClient) GetComment(req *GetCommentRequest) (resp *GetCommentResponse, err error) {
	b, err := s.post("GetComment", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// ReplyComments Use this api to reply comments from buyers in batch
func (s *ShopeeClient) ReplyComments(req *ReplyCommentsRequest) (resp *ReplyCommentsResponse, err error) {
	b, err := s.post("ReplyComments", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// Image
//=======================================================

// UploadImg Use this optional API to pre-validate your image urls and convert them to Shopee image url to use in item upload APIs. This way your potential invalid urls will not block your item upload process.
func (s *ShopeeClient) UploadImg(req *UploadImgRequest) (resp *UploadImgResponse, err error) {
	b, err := s.post("UploadImg", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// Discount
//=======================================================

// AddDiscount Use this api to add shop discount activity
func (s *ShopeeClient) AddDiscount(req *AddDiscountRequest) (resp *AddDiscountResponse, err error) {
	b, err := s.post("AddDiscount", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// AddDiscountItem Use this api to add shop discount item
func (s *ShopeeClient) AddDiscountItem(req *AddDiscountItemRequest) (resp *AddDiscountItemResponse, err error) {
	b, err := s.post("AddDiscountItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// DeleteDiscount Use this api to delete one discount activity BEFORE it starts.
func (s *ShopeeClient) DeleteDiscount(req *DeleteDiscountRequest) (resp *DeleteDiscountResponse, err error) {
	b, err := s.post("DeleteDiscount", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// DeleteDiscountItem Use this api to delete items of the discount activity
func (s *ShopeeClient) DeleteDiscountItem(req *DeleteDiscountItemRequest) (resp *DeleteDiscountItemResponse, err error) {
	b, err := s.post("DeleteDiscountItem", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetDiscountDetail Use this api to get one shop discount activity detail
func (s *ShopeeClient) GetDiscountDetail(req *GetDiscountDetailRequest) (resp *GetDiscountDetailResponse, err error) {
	b, err := s.post("GetDiscountDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetDiscountsList Use this api to get shop discount activity list
func (s *ShopeeClient) GetDiscountsList(req *GetDiscountsListRequest) (resp *GetDiscountsListResponse, err error) {
	b, err := s.post("GetDiscountsList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateDiscount Use this api to update one discount information
func (s *ShopeeClient) UpdateDiscount(req *UpdateDiscountRequest) (resp *UpdateDiscountResponse, err error) {
	b, err := s.post("UpdateDiscount", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateDiscountItems Use this api to update items of the discount activity
func (s *ShopeeClient) UpdateDiscountItems(req *UpdateDiscountItemsRequest) (resp *UpdateDiscountItemsResponse, err error) {
	b, err := s.post("UpdateDiscountItems", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// Orders
//=======================================================

// GetOrdersList GetOrdersList is the recommended call to use for order management. Use this call to retrieve basic information of all orders which are updated within specific period of time. More details of each order can be retrieved from GetOrderDetails. [Only the recent one month orders can be fetch through this API. Please use GetOrderBySatus API to fetch more orders.]
func (s *ShopeeClient) GetOrdersList(req *GetOrdersListRequest) (resp *GetOrdersListResponse, err error) {
	b, err := s.post("GetOrdersList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetOrdersByStatus GetOrdersByStatus is the recommended call to use for order management. Use this call to retrieve basic information of all orders which are specific status. More details of each order can be retrieved from GetOrderDetails.
func (s *ShopeeClient) GetOrdersByStatus(req *GetOrdersByStatusRequest) (resp *GetOrdersByStatusResponse, err error) {
	b, err := s.post("GetOrdersByStatus", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetOrderDetails Use this call to retrieve detailed information about one or more orders based on OrderSN.
func (s *ShopeeClient) GetOrderDetails(req *GetOrderDetailsRequest) (resp *GetOrderDetailsResponse, err error) {
	b, err := s.post("GetOrderDetails", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetEscrowDetails Use this call to retrieve detailed escrow information about one order based on OrderSN.
func (s *ShopeeClient) GetEscrowDetails(req *GetEscrowDetailsRequest) (resp *GetEscrowDetailsResponse, err error) {
	b, err := s.post("GetEscrowDetails", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// AddOrderNote Use this call to add note for an order
func (s *ShopeeClient) AddOrderNote(req *AddOrderNoteRequest) (resp *AddOrderNoteResponse, err error) {
	b, err := s.post("AddOrderNote", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// CancelOrder Use this call to cancel an order from the seller side.
func (s *ShopeeClient) CancelOrder(req *CancelOrderRequest) (resp *CancelOrderResponse, err error) {
	b, err := s.post("CancelOrder", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// AcceptBuyerCancellation Use this call to accept buyer cancellation
func (s *ShopeeClient) AcceptBuyerCancellation(req *AcceptBuyerCancellationRequest) (resp *AcceptBuyerCancellationResponse, err error) {
	b, err := s.post("AcceptBuyerCancellation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// RejectBuyerCancellation Use this call to reject buyer cancellation
func (s *ShopeeClient) RejectBuyerCancellation(req *RejectBuyerCancellationRequest) (resp *RejectBuyerCancellationResponse, err error) {
	b, err := s.post("RejectBuyerCancellation", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetForderInfo Use this call to retrieve detailed information of all the fulfill orders(forder) under a single regular order based on ordersn.
func (s *ShopeeClient) GetForderInfo(req *GetForderInfoRequest) (resp *GetForderInfoResponse, err error) {
	b, err := s.post("GetForderInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetEscrowReleasedOrders Use this api to get orders' release time and escrow amount.
func (s *ShopeeClient) GetEscrowReleasedOrders(req *GetEscrowReleasedOrdersRequest) (resp *GetEscrowReleasedOrdersResponse, err error) {
	b, err := s.post("GetEscrowReleasedOrders", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// SplitOrder Use this API to split order into fulfillment orders. This feature is only enabled for whitelisted shops.
func (s *ShopeeClient) SplitOrder(req *SplitOrderRequest) (resp *SplitOrderResponse, err error) {
	b, err := s.post("SplitOrder", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UndoSplitOrder Use this API to cancel split order from the seller side.
func (s *ShopeeClient) UndoSplitOrder(req *UndoSplitOrderRequest) (resp *UndoSplitOrderResponse, err error) {
	b, err := s.post("UndoSplitOrder", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetUnbindOrderList Use this call to get a list of unbind orders.
func (s *ShopeeClient) GetUnbindOrderList(req *GetUnbindOrderListRequest) (resp *GetUnbindOrderListResponse, err error) {
	b, err := s.post("GetUnbindOrderList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// MyIncome Use this API to fetch the accounting detail of order.
func (s *ShopeeClient) MyIncome(req *MyIncomeRequest) (resp *MyIncomeResponse, err error) {
	b, err := s.post("MyIncome", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// Logistics
//=======================================================

// GetLogistics Use this call to get all supported logistic channels.
func (s *ShopeeClient) GetLogistics(req *GetLogisticsRequest) (resp *GetLogisticsResponse, err error) {
	b, err := s.post("GetLogistics", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateShopLogistics Configure shop level logistics
func (s *ShopeeClient) UpdateShopLogistics(req *UpdateShopLogisticsRequest) (resp *UpdateShopLogisticsResponse, err error) {
	b, err := s.post("UpdateShopLogistics", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetParameterForInit Use this call to get all required param for logistic initiation.
func (s *ShopeeClient) GetParameterForInit(req *GetParameterForInitRequest) (resp *GetParameterForInitResponse, err error) {
	b, err := s.post("GetParameterForInit", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetAddress For integrated logistics channel, use this call to get pickup address for pickup mode order.
func (s *ShopeeClient) GetAddress(req *GetAddressRequest) (resp *GetAddressResponse, err error) {
	b, err := s.post("GetAddress", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetTimeSlot For integrated logistics channel, use this call to get pickup timeslot for pickup mode order.
func (s *ShopeeClient) GetTimeSlot(req *GetTimeSlotRequest) (resp *GetTimeSlotResponse, err error) {
	b, err := s.post("GetTimeSlot", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetBranch For integrated logistics channel, use this call to get dropoff location for dropoff mode order.
func (s *ShopeeClient) GetBranch(req *GetBranchRequest) (resp *GetBranchResponse, err error) {
	b, err := s.post("GetBranch", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetLogisticInfo Get all the logistics info of an order to Init.
// This API consolidates the output of GetParameterForInit, GetAddresss, GetTimeSlot and GetBranch based on each order so that developer can get all the required parameters ready in this API for Init.This API is an alternative of GetParameterForInit, GetAddresss, GetTimeSlot and GetBranch as a set.
func (s *ShopeeClient) GetLogisticInfo(req *GetLogisticInfoRequest) (resp *GetLogisticInfoResponse, err error) {
	b, err := s.post("GetLogisticInfo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// Init Use this call to initiate logistics including arrange Pickup, Dropoff or shipment for non-integrated logistic channels. Should call shopee.logistics.GetLogisticInfo to fetch all required param first. It's recommended to initiate logistics one hour AFTER the orders were placed since there is one-hour window buyer can cancel any order without request to seller.
func (s *ShopeeClient) Init(req *InitRequest) (resp *InitResponse, err error) {
	b, err := s.post("Init", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetAirwayBill Use this API to get airway bill for orders. AirwayBill is only fetchable when the order status is under READY_TO_SHIP and RETRY_SHIP.
func (s *ShopeeClient) GetAirwayBill(req *GetAirwayBillRequest) (resp *GetAirwayBillResponse, err error) {
	b, err := s.post("GetAirwayBill", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetOrderLogistics Use this call to fetch the logistics information of an order, these info can be used for waybill printing. Dedicated for crossborder SLS order airwaybill. May not be applicable for local channel airwaybill.
func (s *ShopeeClient) GetOrderLogistics(req *GetOrderLogisticsRequest) (resp *GetOrderLogisticsResponse, err error) {
	b, err := s.post("GetOrderLogistics", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetLogisticsMessage Use this call to get the logistics tracking information of an order.
func (s *ShopeeClient) GetLogisticsMessage(req *GetLogisticsMessageRequest) (resp *GetLogisticsMessageResponse, err error) {
	b, err := s.post("GetLogisticsMessage", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetForderWaybill Use this API to get airwaybill for fulfillment orders.
func (s *ShopeeClient) GetForderWaybill(req *GetForderWaybillRequest) (resp *GetForderWaybillResponse, err error) {
	b, err := s.post("GetForderWaybill", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// SetAddress Use this API to set default_address/pick_up_address/return_address of shop. Please use GetAddress API to fetch the address_id.
func (s *ShopeeClient) SetAddress(req *SetAddressRequest) (resp *SetAddressResponse, err error) {
	b, err := s.post("SetAddress", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// DeleteAddress Use this API to delete the default/pick_up/return address of shop by address_id. Please use GetAddress API to fetch the address_id.
func (s *ShopeeClient) DeleteAddress(req *DeleteAddressRequest) (resp *DeleteAddressResponse, err error) {
	b, err := s.post("DeleteAddress", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// Returns
//=======================================================

// ConfirmReturn Confirm return
func (s *ShopeeClient) ConfirmReturn(req *ConfirmReturnRequest) (resp *ConfirmReturnResponse, err error) {
	b, err := s.post("ConfirmReturn", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// DisputeReturn Dispute return
func (s *ShopeeClient) DisputeReturn(req *DisputeReturnRequest) (resp *DisputeReturnResponse, err error) {
	b, err := s.post("DisputeReturn", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetReturnList Get return list
func (s *ShopeeClient) GetReturnList(req *GetReturnListRequest) (resp *GetReturnListResponse, err error) {
	b, err := s.post("GetReturnList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetReturnDetail Use this api to get detail information of a returned order
func (s *ShopeeClient) GetReturnDetail(req *GetReturnDetailRequest) (resp *GetReturnDetailResponse, err error) {
	b, err := s.post("GetReturnDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// Public
//=======================================================

// GetShopsByPartner Use this call to get basic info of shops which have authorized to the partner.
func (s *ShopeeClient) GetShopsByPartner(req *GetShopsByPartnerRequest) (resp *GetShopsByPartnerResponse, err error) {
	b, err := s.post("GetShopsByPartner", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetCategoriesByCountry Use this api to get categories list filtered by country and cross border without using shopID.
func (s *ShopeeClient) GetCategoriesByCountry(req *GetCategoriesByCountryRequest) (resp *GetCategoriesByCountryResponse, err error) {
	b, err := s.post("GetCategoriesByCountry", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetPaymentList The supported payment method list by country
func (s *ShopeeClient) GetPaymentList(req *GetPaymentListRequest) (resp *GetPaymentListResponse, err error) {
	b, err := s.post("GetPaymentList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// TopPicks
//=======================================================

// GetTopPicksList Get the list of all collections.
func (s *ShopeeClient) GetTopPicksList(req *GetTopPicksListRequest) (resp *GetTopPicksListResponse, err error) {
	b, err := s.post("GetTopPicksList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// AddTopPicks Add one collection. One shop can have up to 10 collections.
func (s *ShopeeClient) AddTopPicks(req *AddTopPicksRequest) (resp *AddTopPicksResponse, err error) {
	b, err := s.post("AddTopPicks", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// UpdateTopPicks Use this API to update the collection name, the item list in a collection, or to activate a collection.
func (s *ShopeeClient) UpdateTopPicks(req *UpdateTopPicksRequest) (resp *UpdateTopPicksResponse, err error) {
	b, err := s.post("UpdateTopPicks", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// DeleteTopPicks Delete a collection
func (s *ShopeeClient) DeleteTopPicks(req *DeleteTopPicksRequest) (resp *DeleteTopPicksResponse, err error) {
	b, err := s.post("DeleteTopPicks", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// FirstMileTracking
//=======================================================

// GenerateFMTrackingNo Use this API to generate first-mile tracking number for the shipment method of pickup. Please note that the prerequisite for using this API is that the order status is ready_to_ship and the tracking number of order has been obtained.
func (s *ShopeeClient) GenerateFMTrackingNo(req *GenerateFMTrackingNoRequest) (resp *GenerateFMTrackingNoResponse, err error) {
	b, err := s.post("GenerateFMTrackingNo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetShopFMTrackingNo Use this API to fetch first-mile tracking numbers of the shop.
func (s *ShopeeClient) GetShopFMTrackingNo(req *GetShopFMTrackingNoRequest) (resp *GetShopFMTrackingNoResponse, err error) {
	b, err := s.post("GetShopFMTrackingNo", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// FirstMileCodeBindOrder Use this API to bind orders with the first-mile tracking number.
func (s *ShopeeClient) FirstMileCodeBindOrder(req *FirstMileCodeBindOrderRequest) (resp *FirstMileCodeBindOrderResponse, err error) {
	b, err := s.post("FirstMileCodeBindOrder", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetFmTnDetail Use this API to fetch the detailed information of first-mile tracking number.
func (s *ShopeeClient) GetFmTnDetail(req *GetFmTnDetailRequest) (resp *GetFmTnDetailResponse, err error) {
	b, err := s.post("GetFmTnDetail", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetFMTrackingNoWaybill Use the API to get the waybill of first-mile tracking number.Please note that this API only used for the shipment method of pickup.
func (s *ShopeeClient) GetFMTrackingNoWaybill(req *GetFMTrackingNoWaybillRequest) (resp *GetFMTrackingNoWaybillResponse, err error) {
	b, err := s.post("GetFMTrackingNoWaybill", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// GetShopFirstMileChannel Use this call to get all supported logistic channels for first mile.
func (s *ShopeeClient) GetShopFirstMileChannel(req *GetShopFirstMileChannelRequest) (resp *GetShopFirstMileChannelResponse, err error) {
	b, err := s.post("GetShopFirstMileChannel", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// FirstMileUnbind Use this API to unbind orders with the first-mile tracking number. Only applicable to cross-border sellers in China.
func (s *ShopeeClient) FirstMileUnbind(req *FirstMileUnbindRequest) (resp *FirstMileUnbindResponse, err error) {
	b, err := s.post("FirstMileUnbind", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// Payment
//=======================================================

// GetTransactionList Use this API to get the transaction records of wallet.
func (s *ShopeeClient) GetTransactionList(req *GetTransactionListRequest) (resp *GetTransactionListResponse, err error) {
	b, err := s.post("GetTransactionList", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// Push
//=======================================================

// GetPushConfig Use this API to get the configuration information of push service.
func (s *ShopeeClient) GetPushConfig(req *GetPushConfigRequest) (resp *GetPushConfigResponse, err error) {
	b, err := s.post("GetPushConfig", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// SetPushConfig Use this API to get the configuration information of push service.
func (s *ShopeeClient) SetPushConfig(req *SetPushConfigRequest) (resp *SetPushConfigResponse, err error) {
	b, err := s.post("SetPushConfig", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

//=======================================================
// Auth (V2)
//=======================================================

// GetAccessToken Use this API and the code to obtain the access_token and refresh_token.
func (s *ShopeeClient) GetAccessToken(req *GetAccessTokenRequest) (resp *GetAccessTokenResponse, err error) {
	b, err := s.post("GetAccessToken", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}

// RefreshAccessToken Use this API to refresh the access_token after it expires.
func (s *ShopeeClient) RefreshAccessToken(req *RefreshAccessTokenRequest) (resp *RefreshAccessTokenResponse, err error) {
	b, err := s.post("RefreshAccessToken", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}
	return
}
