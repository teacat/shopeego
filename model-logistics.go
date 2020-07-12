package shopeego

type GetLogisticsRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetLogisticsResponse struct {
	Logistics []GetLogisticsResponseLogistic `json:"logistics,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type UpdateShopLogisticsRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// The identity of logistic channel
	LogisticID int64 `json:"logistic_id,omitempty"`
	// Whether to enable this logistic channel
	Enabled bool `json:"enabled,omitempty"`
	// Whether to make this logistic channel preferred. Indonestia logistics channel are not applicable.
	Preferred bool `json:"preferred,omitempty"`
	// Whether to enable COD for this logistic channel. Only COD supported channels are applicable.
	COD bool `json:"cod,omitempty"`
}

type UpdateShopLogisticsResponse struct {
	// Shopee's unique identifier for a shop.
	ShopID int64 `json:"shopid,omitempty"`
	// The identity of logistic channel
	LogisticID int64 `json:"logistic_id,omitempty"`
	// Whether this logistic channel is enabled
	Enabled bool `json:"enabled,omitempty"`
	// Whether this logistic channel is preferred
	Preferred bool `json:"preferred,omitempty"`
	// WHether COD is enabled for this channel
	COD bool `json:"cod,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetParameterForInitRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// The serial numbers of the order needs to be initiated logistic.
	OrderSN string `json:"ordersn,omitempty"`
}

type GetParameterForInitResponse struct {
	// Could contain 'address_id', 'pickup_time_id'. If it contains 'address_id', should call shopee.logistics.GetAddress to get address list, then choose one of the address to init logistic. If it contains pickup_time_id, should call shopee.logistics.GetTimeSlot to get timeslot list, then choose one of the time to init logistic. If it has empty value, developer should still include "pickup" field in Init API.
	Pickup []string `json:"pickup,omitempty"`
	// Could contain 'branch_id', 'sender_real_name' or 'tracking_no'. If it contains 'branch_id', should call shopee.logistics.GetBranch to get branch list. If it contains 'sender_real_name' or 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "dropoff" field in Init API.
	Dropoff []string `json:"dropoff,omitempty"`
	// Could contain 'tracking_no'. If it contains 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "non-integrated" field in Init API.
	NonIntegrated []string `json:"non_integrated_[],omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetAddressRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetAddressResponse struct {
	// All pickup address that you can choose.
	AddressList []GetAddressResponseAddress `json:"address_list,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetTimeSlotRequest struct {
	// The order serial numbers.
	OrderSN string `json:"ordersn,omitempty"`
	// The identify of address.
	AddressID int64 `json:"address_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetTimeSlotResponse struct {
	//
	PickupTime []GetTimeSlotResponseTime `json:"pickup_time,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetBranchRequest struct {
	// The order serial numbers.
	OrderSN string `json:"ordersn,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetBranchResponse struct {
	// All dropoff branches you can choose.
	Branch []GetBranchResponseBranch `json:"branch,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetLogisticInfoRequest struct {
	// The order serial numbers.
	OrderSN string `json:"ordersn,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetLogisticInfoResponse struct {
	// Logistics information for pickup mode order.
	Pickup GetLogisticInfoResponsePickup `json:"pickup,omitempty"`
	// Logistics information for dropoff mode order.
	Dropoff GetLogisticInfoResponseDropoff `json:"dropoff,omitempty"`
	// The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init.
	InfoNeeded GetLogisticInfoResponseInfo `json:"info_needed,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type InitRequest struct {
	// The order serial numbers.
	OrderSN string `json:"ordersn,omitempty"`
	// Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.
	Pickup *InitRequestPickup `json:"pickup,omitempty"`
	// Required parameter ONLY if GetParameterForInit returns "dropoff" or if GetLogisticsInfo returns "dropoff" under "info_needed" for the same order. Developer should still include "dropoff" field in the call even if "dropoff" has empty value.
	Dropoff *InitRequestDropoff `json:"dropoff,omitempty"`
	// Optional parameter when GetParameterForInit returns "non-integrated" or GetLogisticsInfo returns "non-integrated" under "info_needed".
	NonIntegrated *InitRequestNonIntegrated `json:"non_integrated,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
}

type InitResponse struct {
	// The tracking number of order
	TrackingNumber string `json:"tracking_number,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetAirwayBillRequest struct {
	// The set of order serial numbers. Up to 50 ordersn in one call.
	OrderSNList []string `json:"ordersn_list,omitempty"`
	// Option to get batch airway bills in single file. Default value is false.
	IsBatch bool `json:"is_batch,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetAirwayBillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result GetAirwayBillResponseResult `json:"result,omitempty"`
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult GetAirwayBillResponseBatchResult `json:"batch_result,omitempty"`
	// This list contains the ordersn of all orders that failed to retrieve airway bill in this call. AirwayBill is no longer fetchable after the order status is updated to SHIPPED.
	Errors []string `json:"errors,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetOrderLogisticsRequest struct {
	// The order serial numbers. Make sure the order has trackingNo generated before calling this API.
	OrderSN string `json:"ordersn,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
}

type GetOrderLogisticsResponse struct {
	// The logistics of order.
	Logistics []GetOrderLogisticsResponseLogistic `json:"logistics,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type GetLogisticsMessageRequest struct {
	// The order serial numbers.
	OrderSN string `json:"ordersn,omitempty"`
	// The tracking number assigned by the shipping carrier for item shipment. If there are more than one tracking number in one order, this param is necessary.
	TrackingNumber string `json:"tracking_number,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
}

type GetLogisticsMessageResponse struct {
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNumber string `json:"tracking_number,omitempty"`
	//
	TrackingInfo []GetLogisticsMessageResponseInfo `json:"tracking_info,omitempty"`
	// The Shopee logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.
	LogisticStatus string `json:"logisttic_status,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
	// The order serial numbers.
	OrderSN string `json:"ordersn,omitempty"`
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id,omitempty"`
}

type GetForderWaybillRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// The set of order serial numbers. Up to 50 ordersn in one call.
	OrdersList []GetForderWaybillRequestOrder `json:"orders_list,omitempty"`
	// Option to get batch airway bills in single file. Default value is false.
	IsBatch bool `json:"is_batch,omitempty"`
}

type GetForderWaybillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result GetForderWaybillResponseResult `json:"result,omitempty"`
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult GetForderWaybillResponseBatchResult `json:"batch_result,omitempty"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type SetAddressRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// The identity of default address.
	DefaultAddressID int64 `json:"default_address_id,omitempty"`
	// The identity of pick_up address.
	PickUpAddressID int64 `json:"pick_up_address_id,omitempty"`
	// The identity of return address.
	ReturnAddressID int64 `json:"return_address_id,omitempty"`
}

type SetAddressResponse struct {
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}

type DeleteAddressRequest struct {
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
	// The identity of address
	AddressID int64 `json:"address_id,omitempty"`
}

type DeleteAddressResponse struct {
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id,omitempty"`
}
