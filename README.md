type Client interface {
	GetShopInfo
	UpdateShopInfo
	Performance
	SetShopInstallmentStatus

	AddShopCategory
	GetShopCategories
	DeleteShopCategory
	UpdateShopCategory
	AddItems
	GetItems
	DeleteItems

	GetCategories
	GetAttributes
	Add
	Delete
	UnlistItem
	AddVariations
	DeleteVariation
	GetItemsList
	GetItemDetail
	UpdateItem
	AddItemImg
	UpdateItemImg
	InsertItemImg
	DeleteItemImg
	UpdatePrice
	UpdateStock
	UpdateVariationPrice
	UpdateVariationStock
	UpdatePriceBatch
	UpdateStockBatch
	UpdateVariationPriceBatch
	UpdateVariationStockBatch
	InitTierVariation
	AddTierVariation
	GetVariation
	UpdateTierVariationList
	UpdateTierVariationIndex
	BoostItem
	GetBoostedItem
	SetItemInstallmentTenures
	GetPromotionInfo
	GetRecommendCats

	UploadImg

	AddDiscount
	AddDiscountItem
	DeleteDiscount
	DeleteDiscountItem
	GetDiscountDetail
	GetDiscountsList
	UpdateDiscount
	UpdateDiscountItems

	GetOrdersList
	GetOrdersByStatus
	GetOrderDetails
	GetEscrowDetails
	AddOrderNote
	CancelOrder
	AcceptBuyerCancellation
	RejectBuyerCancellation
	GetForderInfo
	GetEscrowReleasedOrders
	SplitOrder
	UndoSplitOrder
	GetUnbindOrderList

	GetLogistics
	UpdateShopLogistics
	GetParameterForInit
	GetAddress
	GetTimeSlot
	GetBranch
	GetLogisticInfo
	Init
	GetAirwayBill
	GetOrderLogistics
	GetLogisticsMessage
	GetForderWaybill

	ConfirmReturn
	DisputeReturn
	GetReturnList
	GetReturnDetail

	GetShopsByPartner
	GetCategoriesByCountry
	GetPaymentList

	GetTopPicksList
	AddTopPicks
	UpdateTopPicks
	DeleteTopPicks

	GenerateFMTrackingNo
	GetShopFMTrackingNo
	FirstMileCodeBindOrder
	GetFmTnDetail
	GetFMTrackingNoWaybill
	GetShopFirstMileChannel
}