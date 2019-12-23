package shopeego

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"io"

	"github.com/parnurzeal/gorequest"
)

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
	//
	IsSandbox bool
}

// CONSTS

func NewClient(opts *ClientOptions) {

}

// Client 定義了一個蝦皮的客戶端該有什麼功能。
type Client interface {
	GetShopInfo(*GetShopInfoRequest) (*GetShopInfoResponse, error)
	UpdateShopInfo(*UpdateShopInfoRequest) (*UpdateShopInfoResponse, error)
	Performance(*PerformanceRequest) (*PerformanceResponse, error)
	SetShopInstallmentStatus(*SetShopInstallmentStatusRequest) (*SetShopInstallmentStatusResponse, error)

	AddShopCategory(*AddShopCategoryRequest) (*AddShopCategoryResponse, error)
	GetShopCategories(*GetShopCategoriesRequest) (*GetShopCategoriesResponse, error)
	DeleteShopCategory(*DeleteShopCategoryRequest) (*DeleteShopCategoryResponse, error)
	UpdateShopCategory(*UpdateShopCategoryRequest) (*UpdateShopCategoryResponse, error)
	AddItems(*AddItemsRequest) (*AddItemsResponse, error)
	GetItems(*GetItemsRequest) (*GetItemsResponse, error)
	DeleteItems(*DeleteItemsRequest) (*DeleteItemsResponse, error)

	GetCategories(*GetCategoriesRequest) (*GetCategoriesResponse, error)
	GetAttributes(*GetAttributesRequest) (*GetAttributesResponse, error)
	Add(*AddRequest) (*AddResponse, error)
	Delete(*DeleteRequest) (*DeleteResponse, error)
	UnlistItem(*UnlistItemRequest) (*UnlistItemResponse, error)
	AddVariations(*AddVariationsRequest) (*AddVariationsResponse, error)
	DeleteVariation(*DeleteVariationRequest) (*DeleteVariationResponse, error)
	GetItemsList(*GetItemsListRequest) (*GetItemsListResponse, error)
	GetItemDetail(*GetItemDetailRequest) (*GetItemDetailResponse, error)
	UpdateItem(*UpdateItemRequest) (*UpdateItemResponse, error)
	AddItemImg(*AddItemImgRequest) (*AddItemImgResponse, error)
	UpdateItemImg(*UpdateItemImgRequest) (*UpdateItemImgResponse, error)
	InsertItemImg(*InsertItemImgRequest) (*InsertItemImgResponse, error)
	DeleteItemImg(*DeleteItemImgRequest) (*DeleteItemImgResponse, error)
	UpdatePrice(*UpdatePriceRequest) (*UpdatePriceResponse, error)
	UpdateStock(*UpdateStockRequest) (*UpdateStockResponse, error)
	UpdateVariationPrice(*UpdateVariationPriceRequest) (*UpdateVariationPriceResponse, error)
	UpdateVariationStock(*UpdateVariationStockRequest) (*UpdateVariationStockResponse, error)
	UpdatePriceBatch(*UpdatePriceBatchRequest) (*UpdatePriceBatchResponse, error)
	UpdateStockBatch(*UpdateStockBatchRequest) (*UpdateStockBatchResponse, error)
	UpdateVariationPriceBatch(*UpdateVariationPriceBatchRequest) (*UpdateVariationPriceBatchResponse, error)
	UpdateVariationStockBatch(*UpdateVariationStockBatchRequest) (*UpdateVariationStockBatchResponse, error)
	InitTierVariation(*InitTierVariationRequest) (*InitTierVariationResponse, error)
	AddTierVariation(*AddTierVariationRequest) (*AddTierVariationResponse, error)
	GetVariation(*GetVariationRequest) (*GetVariationResponse, error)
	UpdateTierVariationList(*UpdateTierVariationListRequest) (*UpdateTierVariationListResponse, error)
	UpdateTierVariationIndex(*UpdateTierVariationIndexRequest) (*UpdateTierVariationIndexResponse, error)
	BoostItem(*BoostItemRequest) (*BoostItemResponse, error)
	GetBoostedItem(*GetBoostedItemRequest) (*GetBoostedItemResponse, error)
	SetItemInstallmentTenures(*SetItemInstallmentTenuresRequest) (*SetItemInstallmentTenuresResponse, error)
	GetPromotionInfo(*GetPromotionInfoRequest) (*GetPromotionInfoResponse, error)
	GetRecommendCats(*GetRecommendCatsRequest) (*GetRecommendCatsResponse, error)

	UploadImg(*UploadImgRequest) (*UploadImgResponse, error)

	AddDiscount(*AddDiscountRequest) (*AddDiscountResponse, error)
	AddDiscountItem(*AddDiscountItemRequest) (*AddDiscountItemResponse, error)
	DeleteDiscount(*DeleteDiscountRequest) (*DeleteDiscountResponse, error)
	DeleteDiscountItem(*DeleteDiscountItemRequest) (*DeleteDiscountItemResponse, error)
	GetDiscountDetail(*GetDiscountDetailRequest) (*GetDiscountDetailResponse, error)
	GetDiscountsList(*GetDiscountsListRequest) (*GetDiscountsListResponse, error)
	UpdateDiscount(*UpdateDiscountRequest) (*UpdateDiscountResponse, error)
	UpdateDiscountItems(*UpdateDiscountItemsRequest) (*UpdateDiscountItemsResponse, error)

	GetOrdersList(*GetOrdersListRequest) (*GetOrdersListResponse, error)
	GetOrdersByStatus(*GetOrdersByStatusRequest) (*GetOrdersByStatusResponse, error)
	GetOrderDetails(*GetOrderDetailsRequest) (*GetOrderDetailsResponse, error)
	GetEscrowDetails(*GetEscrowDetailsRequest) (*GetEscrowDetailsResponse, error)
	AddOrderNote(*AddOrderNoteRequest) (*AddOrderNoteResponse, error)
	CancelOrder(*CancelOrderRequest) (*CancelOrderResponse, error)
	AcceptBuyerCancellation(*AcceptBuyerCancellationRequest) (*AcceptBuyerCancellationResponse, error)
	RejectBuyerCancellation(*RejectBuyerCancellationRequest) (*RejectBuyerCancellationResponse, error)
	GetForderInfo(*GetForderInfoRequest) (*GetForderInfoResponse, error)
	GetEscrowReleasedOrders(*GetEscrowReleasedOrdersRequest) (*GetEscrowReleasedOrdersResponse, error)
	SplitOrder(*SplitOrderRequest) (*SplitOrderResponse, error)
	UndoSplitOrder(*UndoSplitOrderRequest) (*UndoSplitOrderResponse, error)
	GetUnbindOrderList(*GetUnbindOrderListRequest) (*GetUnbindOrderListResponse, error)

	GetLogistics(*GetLogisticsRequest) (*GetLogisticsResponse, error)
	UpdateShopLogistics(*UpdateShopLogisticsRequest) (*UpdateShopLogisticsResponse, error)
	GetParameterForInit(*GetParameterForInitRequest) (*GetParameterForInitResponse, error)
	GetAddress(*GetAddressRequest) (*GetAddressResponse, error)
	GetTimeSlot(*GetTimeSlotRequest) (*GetTimeSlotResponse, error)
	GetBranch(*GetBranchRequest) (*GetBranchResponse, error)
	GetLogisticInfo(*GetLogisticInfoRequest) (*GetLogisticInfoResponse, error)
	Init(*InitRequest) (*InitResponse, error)
	GetAirwayBill(*GetAirwayBillRequest) (*GetAirwayBillResponse, error)
	GetOrderLogistics(*GetOrderLogisticsRequest) (*GetOrderLogisticsResponse, error)
	GetLogisticsMessage(*GetLogisticsMessageRequest) (*GetLogisticsMessageResponse, error)
	GetForderWaybill(*GetForderWaybillRequest) (*GetForderWaybillResponse, error)

	ConfirmReturn(*ConfirmReturnRequest) (*ConfirmReturnResponse, error)
	DisputeReturn(*DisputeReturnRequest) (*DisputeReturnResponse, error)
	GetReturnList(*GetReturnListRequest) (*GetReturnListResponse, error)
	GetReturnDetail(*GetReturnDetailRequest) (*GetReturnDetailResponse, error)

	GetShopsByPartner(*GetShopsByPartnerRequest) (*GetShopsByPartnerResponse, error)
	GetCategoriesByCountry(*GetCategoriesByCountryRequest) (*GetCategoriesByCountryResponse, error)
	GetPaymentList(*GetPaymentListRequest) (*GetPaymentListResponse, error)

	GetTopPicksList(*GetTopPicksListRequest) (*GetTopPicksListResponse, error)
	AddTopPicks(*AddTopPicksRequest) (*AddTopPicksResponse, error)
	UpdateTopPicks(*UpdateTopPicksRequest) (*UpdateTopPicksResponse, error)
	DeleteTopPicks(*DeleteTopPicksRequest) (*DeleteTopPicksResponse, error)

	GenerateFMTrackingNo(*GenerateFMTrackingNoRequest) (*GenerateFMTrackingNoResponse, error)
	GetShopFMTrackingNo(*GetShopFMTrackingNoRequest) (*GetShopFMTrackingNoResponse, error)
	FirstMileCodeBindOrder(*FirstMileCodeBindOrderRequest) (*FirstMileCodeBindOrderResponse, error)
	GetFmTnDetail(*GetFmTnDetailRequest) (*GetFmTnDetailResponse, error)
	GetFMTrackingNoWaybill(*GetFMTrackingNoWaybillRequest) (*GetFMTrackingNoWaybillResponse, error)
	GetShopFirstMileChannel(*GetShopFirstMileChannelRequest) (*GetShopFirstMileChannelResponse, error)
}

type ShopeeClient struct {
	Secret    string
	IsSandbox bool
}

type ResponseError struct {
	RequestID string
	Msg       string
	Error     error
}

//
func (s *ShopeeClient) getPath(method string) {
	var host string
	if s.IsSandbox {
		host := "https://partner.shopeemobile.com/"
	} else {
		host := "https://partner.uat.shopeemobile.com/"
	}
	return fmt.Sprtinf("%s%s", host, availablePaths[method])
}

//
func (s *ShopeeClient) sign(url string, body []byte) string {
	h := hmac.New(sha256.New, []byte(s.Secret))
	io.WriteString(h, fmt.Sprintf("%s|%s", url, string(body)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//
func (s *ShopeeClient) post(method string, in interface{}) ([]byte, error) {
	body, err := json.Marshal(in)
	if err != nil {
		return []byte(``), err
	}
	url := s.getPath(method)
	req := gorequest.New()
	req.post(url)
	//
	req.Set("Authorization", s.sign(url, body))

	//
	// HANDLE ERRROR!

	resp, _, err := req.End()
	if err != nil {
		return []byte(``), err
	}
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
// Discountt
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
// FirsttMileTracking
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
