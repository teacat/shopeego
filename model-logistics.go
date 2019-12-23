package shopeego

type GetLogisticsRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetLogisticsResponse struct {
	Logistics []GetLogisticsResponseLogistic `json:"logistics"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type UpdateShopLogisticsRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// The identity of logistic channel
	LogisticID int `json:"logistic_id"`
	// Whether to enable this logistic channel
	Enabled bool `json:"enabled"`
	// Whether to make this logistic channel preferred. Indonestia logistics channel are not applicable.
	Preferred bool `json:"preferred"`
	// Whether to enable COD for this logistic channel. Only COD supported channels are applicable.
	COD bool `json:"cod"`
}

type UpdateShopLogisticsResponse struct {
	// Shopee's unique identifier for a shop.
	ShopID int `json:"shop_id"`
	// The identity of logistic channel
	LogisticID int `json:"logistic_id"`
	// Whether this logistic channel is enabled
	Enabled bool `json:"enabled"`
	// Whether this logistic channel is preferred
	Preferred bool `json:"preferred"`
	// WHether COD is enabled for this channel
	COD bool `json:"cod"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetParameterForInitRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// The serial numbers of the order needs to be initiated logistic.
	OrderSN string `json:"order_sn"`
}

type GetParameterForInitResponse struct {
	// Could contain 'address_id', 'pickup_time_id'. If it contains 'address_id', should call shopee.logistics.GetAddress to get address list, then choose one of the address to init logistic. If it contains pickup_time_id, should call shopee.logistics.GetTimeSlot to get timeslot list, then choose one of the time to init logistic. If it has empty value, developer should still include "pickup" field in Init API.
	Pickup []string `json:"pickup"`
	// Could contain 'branch_id', 'sender_real_name' or 'tracking_no'. If it contains 'branch_id', should call shopee.logistics.GetBranch to get branch list. If it contains 'sender_real_name' or 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "dropoff" field in Init API.
	Dropoff []string `json:"dropoff"`
	// Could contain 'tracking_no'. If it contains 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "non-integrated" field in Init API.
	NonIntegrated []string `json:"non_integrated_[]"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetAddressRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetAddressResponse struct {
	// All pickup address that you can choose.
	AddressList []GetAddressResponseAddress `json:"address_list"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetTimeSlotRequest struct {
	// The order serial numbers.
	OrderSN string `json:"order_sn"`
	// The identify of address.
	AddressID int `json:"address_id"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetTimeSlotResponse struct {
	//
	PickupTime []GetTimeSlotResponseTime `json:"pickup_time"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetBranchRequest struct {
	// The order serial numbers.
	OrderSN string `json:"order_sn"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetBranchResponse struct {
	// All dropoff branches you can choose.
	Branch []GetBranchResponseBranch `json:"branch"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetLogisticInfoRequest struct {
	// The order serial numbers.
	OrderSN string `json:"order_sn"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetLogisticInfoResponse struct {
	// Logistics information for pickup mode order.
	Pickup GetLogisticInfoResponsePickup `json:"pickup"`
	// Logistics information for dropoff mode order.
	Dropoff GetLogisticInfoResponseDropoff `json:"dropoff"`
	// The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init.
	InfoNeeded GetLogisticInfoResponseInfo `json:"info_needed"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type InitRequest struct {
	// The order serial numbers.
	OrderSN string `json:"order_sn"`
	// Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.
	Pickup InitRequestPickup `json:"pickup"`
	// Required parameter ONLY if GetParameterForInit returns "dropoff" or if GetLogisticsInfo returns "dropoff" under "info_needed" for the same order. Developer should still include "dropoff" field in the call even if "dropoff" has empty value.
	Dropoff InitRequestDropoff `json:"dropoff"`
	// Optional parameter when GetParameterForInit returns "non-integrated" or GetLogisticsInfo returns "non-integrated" under "info_needed".
	NonIntegrated InitRequestNonIntegrated `json:"non_integrated"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
}

type InitResponse struct {
	// The tracking number of order
	TrackingNumber string `json:"tracking_number"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetAirwayBillRequest struct {
	// The set of order serial numbers. Up to 50 ordersn in one call.
	OrderSNList []string `json:"order_sn_list"`
	// Option to get batch airway bills in single file. Default value is false.
	IsBatch bool `json:"is_batch"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type GetAirwayBillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result GetAirwayBillResponseResult `json:"result"`
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult GetAirwayBillResponseBatchResult `json:"batch_result"`
	// This list contains the ordersn of all orders that failed to retrieve airway bill in this call. AirwayBill is no longer fetchable after the order status is updated to SHIPPED.
	Errors []string `json:"errors"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetOrderLogisticsRequest struct {
	// The order serial numbers. Make sure the order has trackingNo generated before calling this API.
	OrderSN string `json:"order_sn"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
}

type GetOrderLogisticsResponse struct {
	// The logistics of order.
	Logistics []GetOrderLogisticsResponseLogistic `json:"logistics"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type GetLogisticsMessageRequest struct {
	// The order serial numbers.
	OrderSN string `json:"order_sn"`
	// The tracking number assigned by the shipping carrier for item shipment. If there are more than one tracking number in one order, this param is necessary.
	TrackingNumber string `json:"tracking_number"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
}

type GetLogisticsMessageResponse struct {
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNumber string `json:"tracking_number"`
	//
	TrackingInfo []GetLogisticsMessageResponseInfo `json:"tracking_info"`
	// The Shopee logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.
	LogisticStatus string `json:"logisttic_status"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
	// The order serial numbers.
	OrderSN string `json:"order_sn"`
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
}

type GetForderWaybillRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
	// The set of order serial numbers. Up to 50 ordersn in one call.
	OrdersList []GetForderWaybillRequestOrder `json:"orders_list"`
	// Option to get batch airway bills in single file. Default value is false.
	IsBatch bool `json:"is_batch"`
}

type GetForderWaybillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result GetForderWaybillResponseResult `json:"result"`
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult GetForderWaybillResponseBatchResult `json:"batch_result"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}
