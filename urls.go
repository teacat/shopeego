package shopeego

var availablePaths map[string]string = map[string]string{
	//=======================================================
	// Shop
	//=======================================================

	"GetShopInfo":              "shop/get",
	"UpdateShopInfo":           "shop/update",
	"Performance":              "shop/performance",
	"SetShopInstallmentStatus": "shop/set_installment_status",
	"AuthPartner":              "shop/auth_partner",

	//=======================================================
	// ShopCategory
	//=======================================================

	"AddShopCategory":    "shop_category/add",
	"GetShopCategories":  "shop_categorys/get",
	"DeleteShopCategory": "shop_category/delete",
	"UpdateShopCategory": "shop_category/update",
	"AddItems":           "shop_category/add/items",
	"GetItems":           "shop_category/get/items",
	"DeleteItems":        "shop_category/del/items",

	//=======================================================
	// Item
	//=======================================================

	"GetCategories":             "item/categories/get",
	"GetAttributes":             "item/attributes/get",
	"Add":                       "item/add",
	"Delete":                    "item/delete",
	"UnlistItem":                "items/unlist",
	"AddVariations":             "item/add_variations",
	"DeleteVariation":           "item/delete_variation",
	"GetItemsList":              "items/get",
	"GetItemDetail":             "item/get",
	"UpdateItem":                "item/update",
	"AddItemImg":                "item/img/add",
	"UpdateItemImg":             "item/img/update",
	"InsertItemImg":             "item/img/insert",
	"DeleteItemImg":             "item/img/delete",
	"UpdatePrice":               "items/update_price",
	"UpdateStock":               "items/update_stock",
	"UpdateVariationPrice":      "items/update_variation_price",
	"UpdateVariationStock":      "items/update_variation_stock",
	"UpdatePriceBatch":          "items/update/items_price",
	"UpdateStockBatch":          "items/update/items_stock",
	"UpdateVariationPriceBatch": "items/update/vars_price",
	"UpdateVariationStockBatch": "items/update/vars_stock",
	"InitTierVariation":         "item/tier_var/init",
	"AddTierVariation":          "item/tier_var/add",
	"GetVariation":              "item/tier_var/get",
	"UpdateTierVariationList":   "item/tier_var/update_list",
	"UpdateTierVariationIndex":  "item/tier_var/update",
	"BoostItem":                 "items/boost",
	"GetBoostedItem":            "items/get_boosted",
	"SetItemInstallmentTenures": "item/installment/set",
	"GetPromotionInfo":          "item/promotion/get",
	"GetRecommendCats":          "item/categories/get_recommend",
	"GetComment":                "items/comments/get",
	"ReplyComments":             "items/comments/reply",

	//=======================================================
	// Image
	//=======================================================

	"UploadImg": "image/upload",

	//=======================================================
	// Discount
	//=======================================================

	"AddDiscount":         "discount/add",
	"AddDiscountItem":     "discount/items/add",
	"DeleteDiscount":      "discount/delete",
	"DeleteDiscountItem":  "discount/item/delete",
	"GetDiscountDetail":   "discount/detail",
	"GetDiscountsList":    "discounts/get",
	"UpdateDiscount":      "discount/update",
	"UpdateDiscountItems": "discount/items/update",

	//=======================================================
	// Orders
	//=======================================================

	"GetOrdersList":           "orders/basics",
	"GetOrdersByStatus":       "orders/get",
	"GetOrderDetails":         "orders/detail",
	"GetEscrowDetails":        "orders/my_income",
	"AddOrderNote":            "orders/note/add",
	"CancelOrder":             "orders/cancel",
	"AcceptBuyerCancellation": "orders/buyers_cancellation/accept",
	"RejectBuyerCancellation": "orders/buyers_cancellation/reject",
	"GetForderInfo":           "orders/forder/get",
	"GetEscrowReleasedOrders": "orders/get_escrow_detail",
	"SplitOrder":              "orders/split",
	"UndoSplitOrder":          "orders/unsplit",
	"GetUnbindOrderList":      "orders/unbind/list",
	"MyIncome":                "orders/income",

	//=======================================================
	// Logistics
	//=======================================================

	"GetLogistics":        "logistics/channel/get",
	"UpdateShopLogistics": "logistics/channels/update",
	"GetParameterForInit": "logistics/init_parameter/get",
	"GetAddress":          "logistics/address/get",
	"GetTimeSlot":         "logistics/timeslot/get",
	"GetBranch":           "logistics/branch/get",
	"GetLogisticInfo":     "logistics/init_info/get",
	"Init":                "logistics/init",
	"GetAirwayBill":       "logistics/airway_bill/get_mass",
	"GetOrderLogistics":   "logistics/order/get",
	"GetLogisticsMessage": "logistics/tracking",
	"GetForderWaybill":    "logistics/forder_waybill/get_mass",
	"SetAddress":          "logistics/address/set",
	"DeleteAddress":       "logistics/address/delete",

	//=======================================================
	// Returns
	//=======================================================

	"ConfirmReturn":   "returns/confirm",
	"DisputeReturn":   "returns/dispute",
	"GetReturnList":   "returns/get",
	"GetReturnDetail": "returns/detail",

	//=======================================================
	// Public
	//=======================================================

	"GetShopsByPartner":      "shop/get_partner_shop",
	"GetCategoriesByCountry": "item/categories/get_by_country",
	"GetPaymentList":         "payment/list",

	//=======================================================
	// TopPicks
	//=======================================================

	"GetTopPicksList": "top_picks/get",
	"AddTopPicks":     "top_picks/add",
	"UpdateTopPicks":  "top_picks/update",
	"DeleteTopPicks":  "top_picks/delete",

	//=======================================================
	// FirstMileTracking
	//=======================================================

	"GenerateFMTrackingNo":    "fm_tn/generate",
	"GetShopFMTrackingNo":     "fm_tn/get",
	"FirstMileCodeBindOrder":  "fm_tn/bind",
	"GetFmTnDetail":           "fm_tn/detail",
	"GetFMTrackingNoWaybill":  "fm_tn/waybill",
	"GetShopFirstMileChannel": "fm_tn/channels",
	"FirstMileUnbind":         "fm_tn/unbind",

	//=======================================================
	// Payment
	//=======================================================

	"GetTransactionList": "wallet/transaction/list",

	//=======================================================
	// Push
	//=======================================================

	"GetPushConfig": "push/get_config",
	"SetPushConfig": "push/set_config",

	//=======================================================
	// Auth (V2)
	//=======================================================

	"GetAccessToken":     "auth/token/get",
	"RefreshAccessToken": "auth/access_token/get",
}
