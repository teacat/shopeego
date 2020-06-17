package shopeego

var availablePaths map[string]string = map[string]string{
	//=======================================================
	// Shop
	//=======================================================

	"GetShopInfo":              "api/v1/shop/get",
	"UpdateShopInfo":           "api/v1/shop/update",
	"Performance":              "api/v1/shop/performance",
	"SetShopInstallmentStatus": "api/v1/shop/set_installment_status",

	//=======================================================
	// ShopCategory
	//=======================================================

	"AddShopCategory":    "api/v1/shop_category/add",
	"GetShopCategories":  "api/v1/shop_categorys/get",
	"DeleteShopCategory": "api/v1/shop_category/delete",
	"UpdateShopCategory": "api/v1/shop_category/update",
	"AddItems":           "api/v1/shop_category/add/items",
	"GetItems":           "api/v1/shop_category/get/items",
	"DeleteItems":        "api/v1/shop_category/del/items",

	//=======================================================
	// Item
	//=======================================================

	"GetCategories":             "api/v1/item/categories/get",
	"GetAttributes":             "api/v1/item/attributes/get",
	"Add":                       "api/v1/item/add",
	"Delete":                    "api/v1/item/delete",
	"UnlistItem":                "api/v1/items/unlist",
	"AddVariations":             "api/v1/item/add_variations",
	"DeleteVariation":           "api/v1/item/delete_variation",
	"GetItemsList":              "api/v1/items/get",
	"GetItemDetail":             "api/v1/item/get",
	"UpdateItem":                "api/v1/item/update",
	"AddItemImg":                "api/v1/item/img/add",
	"UpdateItemImg":             "api/v1/item/img/update",
	"InsertItemImg":             "api/v1/item/img/insert",
	"DeleteItemImg":             "api/v1/item/img/delete",
	"UpdatePrice":               "api/v1/items/update_price",
	"UpdateStock":               "api/v1/items/update_stock",
	"UpdateVariationPrice":      "api/v1/items/update_variation_price",
	"UpdateVariationStock":      "api/v1/items/update_variation_stock",
	"UpdatePriceBatch":          "api/v1/items/update/items_price",
	"UpdateStockBatch":          "api/v1/items/update/items_stock",
	"UpdateVariationPriceBatch": "api/v1/items/update/vars_price",
	"UpdateVariationStockBatch": "api/v1/items/update/vars_stock",
	"InitTierVariation":         "api/v1/item/tier_var/init",
	"AddTierVariation":          "api/v1/item/tier_var/add",
	"GetVariation":              "api/v1/item/tier_var/get",
	"UpdateTierVariationList":   "api/v1/item/tier_var/update_list",
	"UpdateTierVariationIndex":  "api/v1/item/tier_var/update",
	"BoostItem":                 "api/v1/items/boost",
	"GetBoostedItem":            "api/v1/items/get_boosted",
	"SetItemInstallmentTenures": "api/v1/item/installment/set",
	"GetPromotionInfo":          "api/v1/item/promotion/get",
	"GetRecommendCats":          "api/v1/item/categories/get_recommend",
	"GetComment":                "api/v1/items/comments/get",
	"ReplyComments":             "api/v1/items/comments/reply",

	//=======================================================
	// Image
	//=======================================================

	"UploadImg": "api/v1/image/upload",

	//=======================================================
	// Discount
	//=======================================================

	"AddDiscount":         "api/v1/discount/add",
	"AddDiscountItem":     "api/v1/discount/items/add",
	"DeleteDiscount":      "api/v1/discount/delete",
	"DeleteDiscountItem":  "api/v1/discount/item/delete",
	"GetDiscountDetail":   "api/v1/discount/detail",
	"GetDiscountsList":    "api/v1/discounts/get",
	"UpdateDiscount":      "api/v1/discount/update",
	"UpdateDiscountItems": "api/v1/discount/items/update",

	//=======================================================
	// Orders
	//=======================================================

	"GetOrdersList":           "api/v1/orders/basics",
	"GetOrdersByStatus":       "api/v1/orders/get",
	"GetOrderDetails":         "api/v1/orders/detail",
	"GetEscrowDetails":        "api/v1/orders/my_income",
	"AddOrderNote":            "api/v1/orders/note/add",
	"CancelOrder":             "api/v1/orders/cancel",
	"AcceptBuyerCancellation": "api/v1/orders/buyers_cancellation/accept",
	"RejectBuyerCancellation": "api/v1/orders/buyers_cancellation/reject",
	"GetForderInfo":           "api/v1/orders/forder/get",
	"GetEscrowReleasedOrders": "api/v1/orders/get_escrow_detail",
	"SplitOrder":              "api/v1/orders/split",
	"UndoSplitOrder":          "api/v1/orders/unsplit",
	"GetUnbindOrderList":      "api/v1/orders/unbind/list",
	"MyIncome":                "api/v1/orders/income",

	//=======================================================
	// Logistics
	//=======================================================

	"GetLogistics":        "api/v1/logistics/channel/get",
	"UpdateShopLogistics": "api/v1/logistics/channels/update",
	"GetParameterForInit": "api/v1/logistics/init_parameter/get",
	"GetAddress":          "api/v1/logistics/address/get",
	"GetTimeSlot":         "api/v1/logistics/timeslot/get",
	"GetBranch":           "api/v1/logistics/branch/get",
	"GetLogisticInfo":     "api/v1/logistics/init_info/get",
	"Init":                "api/v1/logistics/init",
	"GetAirwayBill":       "api/v1/logistics/airway_bill/get_mass",
	"GetOrderLogistics":   "api/v1/logistics/order/get",
	"GetLogisticsMessage": "api/v1/logistics/tracking",
	"GetForderWaybill":    "api/v1/logistics/forder_waybill/get_mass",
	"SetAddress":          "api/v1/logistics/address/set",
	"DeleteAddress":       "api/v1/logistics/address/delete",

	//=======================================================
	// Returns
	//=======================================================

	"ConfirmReturn":   "api/v1/returns/confirm",
	"DisputeReturn":   "api/v1/returns/dispute",
	"GetReturnList":   "api/v1/returns/get",
	"GetReturnDetail": "api/v1/returns/detail",

	//=======================================================
	// Public
	//=======================================================

	"GetShopsByPartner":      "api/v1/shop/get_partner_shop",
	"GetCategoriesByCountry": "api/v1/item/categories/get_by_country",
	"GetPaymentList":         "api/v1/payment/list",

	//=======================================================
	// TopPicks
	//=======================================================

	"GetTopPicksList": "api/v1/top_picks/get",
	"AddTopPicks":     "api/v1/top_picks/add",
	"UpdateTopPicks":  "api/v1/top_picks/update",
	"DeleteTopPicks":  "api/v1/top_picks/delete",

	//=======================================================
	// FirstMileTracking
	//=======================================================

	"GenerateFMTrackingNo":    "api/v1/fm_tn/generate",
	"GetShopFMTrackingNo":     "api/v1/fm_tn/get",
	"FirstMileCodeBindOrder":  "api/v1/fm_tn/bind",
	"GetFmTnDetail":           "api/v1/fm_tn/detail",
	"GetFMTrackingNoWaybill":  "api/v1/fm_tn/waybill",
	"GetShopFirstMileChannel": "api/v1/fm_tn/channels",
	"FirstMileUnbind":         "api/v1/fm_tn/unbind",

	//=======================================================
	// Payment
	//=======================================================

	"GetTransactionList": "api/v1/wallet/transaction/list",

	//=======================================================
	// Push
	//=======================================================

	"GetPushConfig": "api/v1/push/get_config",
	"SetPushConfig": "api/v1/push/set_config",
}
