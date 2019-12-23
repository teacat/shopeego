package shopeego

type GetLogisticsRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int `json:"partner_id"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int `json:"shop_id"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp"`
}

type Size struct {
	// The identity of size.
	SizeID int `json:"size_id"`
	// The name of size.
	Name int `json:"name"`
	// The pre-defined shipping fee for the specific size.
	DefaultPrice int `json:"default_price"`
}

type Limit struct {
	// The max weight for an item on this logistic channel.If the value is 0 or null, that means there is no limit.
	ItemMaxWeight
	// The min weight for an item on this logistic channel. If the value is 0 or null, that means there is no limit.
	ItemMinWeight
}

type Dimension struct {
	// The max height limit.
	Height float64 `json:"height"`
	// The max width limit.
	Width float64 `json:"width"`
	// The max length limit.
	Length float64 `json:"length"`
	// The unit for the limit.
	Unit string `json:"unit"`
}

type Logistic struct {
	// The identity of logistic channel
	LogisticID int `json:"logistic_id"`
	// The name of logistic channel
	LogisticName string `json:"logistic_name"`
	// This is to indicate whether this logistic channel supports COD
	HasCOD bool `json:"has_cod"`
	// Whether this logistic channel is enabled on shop level.
	Enabled bool `json:"enabled"`
	// See Define FeeType, related to FeeType Value
	FeeType string `json:"fee_type"`
	// Only for fee_type is SIZE_SELECTION
	Sizes Size `json:"sizes"`
	// The weight limit for this logistic channel.
	WeightLimits Limit `json:"weight_limits"`
	// The dimension limit for this logistic channel.
	ItemMaxDimension Dimension `json:"item_max_dimension"`
}

type GetLogisticsResponse struct {
	Logistics []Logistic `json:"logistics"`
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
	NonIntegrated [] string `json:"non_integrated_[]"`
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

type Address struct {
	// The identity of address
	AddressID int `json:"address_id"`
	// The country of specify address
	Country string `json:"country"`
	// The state of specify address
	State string `json:"state"`
	// The city of specify address
	City string `json:"city"`
	// The address description of specify address
	Address string `json:"address"`
	// The zipcode of specify address
	Zipcode string `json:"zipcode"`
	// The district of specify address
	District string `json:"district"`
	// The town of specify address
	Town string `json:"town"`
}

type GetAddressResponse struct {
	// All pickup address that you can choose.
	AddressList []Address `json:"address_list"`
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

type Time struct {
	// The identity of pickuptime.
	PickupTimeID string `json:"pickup_time_id"`
	// The date of pickup time. In timestamp.
	Date int `json:"date"`
	// The text description of pickup time. Only applicable for certain channels.
	TimeText string `json:"time_text"`
}

type GetTimeSlotResponse struct {
	//
	PickupTime []Time `json:"pickup_time"`
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

type Branch struct {
	// The identity of branch.
	BranchID int `json:"branch_id"`
	// The country of specify branch.
	Country string `json:"country"`
	// The state of specify branch.
	State string `json:"state"`
	// The city of specify branch.
	City string `json:"city"`
	// The address description of specify branch.
	Address string `json:"address"`
	// The zipcode of specify branch.
	Zipcode string `json:"zipcode"`
	// The district of specify branch.
	District string `json:"district"`
	// The town of specify branch.
	Town string `json:"town"`
}

type GetBranchResponse struct {
	// All dropoff branches you can choose.
	Branch []Branch `json:"branch"`
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

type Address struct {
	// The identity of address.
	AddressID int `json:"address_id"`
	// The country of specify branch.
	Country string `json:"country"`
	// The state of specify branch.
	State string `json:"state"`
	// The city of specify branch.
	City string `json:"city"`
	// The address description of specify branch.
	Address string `json:"address"`
	// The zipcode of specify branch.
	Zipcode string `json:"zipcode"`
	// The district of specify branch.
	District string `json:"district"`
	// The town of specify branch.
	Town string `json:"town"`
	// List of pickup_time information corresponding to the address_id.
	TimeSlotList []string `json:"time_slot_list"`
	// The identity of pickuptime.
	PickupTimeID string `json:"pickup_time_id"`
	// The date of pickup time. In timestamp.
	Date int `json:"date"`
	// The text description of pickup time. Only applicable for certain channels.
	TimeText string `json:"time_text"`
}

type Branch struct {
	// The identity of branch.
	BranchID int `json:"branch_id"`
	// The country of specify branch.
	Country string `json:"country"`
	// The state of specify branch.
	State string `json:"state"`
	// The city of specify branch.
	City string `json:"city"`
	// The address description of specify branch.
	Address string `json:"address"`
	// The zipcode of specify branch.
	Zipcode string `json:"zipcode"`
	// The district of specify branch.
	District string `json:"district"`
	// The town of specify branch.
	Town string `json:"town"`
}

type Pickup struct {
	// List of available pickup address info.
	AddressList []Address `json:"address_list"`
}

type Dropoff struct {
	// List of available dropoff branches info.
	BranchList []Branch `json:"branch_list"`
}

type Info structt {
	// Logistics information for pickup mode order.
	Pickup []string `json:"pickup"`
	// Logistics information for dropoff mode order.
	Dropoff []string `json:"dropoff"`
	// The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init.
	NonIntegrated []string `json:"non_integrated"`
}

type GetLogisticInfoResponse struct {
	// Logistics information for pickup mode order.
	Pickup Pickup `json:"pickup"`
	// Logistics information for dropoff mode order.
	Dropoff Dropoff `json:"dropoff"`
	// The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init.
	InfoNeeded Info `json:"info_needed"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}

type Pickup struct {
	// The identity of address. Retrieved from shopee.logistics.GetAddress.
	AddressID int `json:"address_id"`
	// The pickup time id. Retrieved from shopee.logistics.GetTimeSlot.
	PickupItemID string `json:"pickup_item_id"`
}

type Dropoff struct {
	// The identity of branch. Retrieved from shopee.logistics.GetBranch branch.
	BranchID int `json:"branch_id"`
	// The real name of sender.
	SenderRealName string `json:"sender_real_name"`
	// Need input this field when "tracking_no" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no"`
}

type NonIntegrated struct {
	// Optional parameter for non-integrated channel order. The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no"`
}

type InitRequest struct {
	// The order serial numbers.
	OrderSN string `json:"order_sn"`
	// Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.
	Pickup Pickup `json:"pickup"`
	// Required parameter ONLY if GetParameterForInit returns "dropoff" or if GetLogisticsInfo returns "dropoff" under "info_needed" for the same order. Developer should still include "dropoff" field in the call even if "dropoff" has empty value.
	Dropoff Dropoff `json:"dropoff"`
	// Optional parameter when GetParameterForInit returns "non-integrated" or GetLogisticsInfo returns "non-integrated" under "info_needed".
	NonIntegrated NonIntegrated `json:"non_integrated"`
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

type AirwayBill struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The url of retrieving airway bill.
	AirwayBill string `json:"airway_bill"`
}

type Error struct {
	// The ordersn of orders which occurred error.
	OrderSN string `json:"order_sn"`
	//
	ErrorCode string `json:"error_code"`
	// The detail information of this error.
	ErrorDescription string `json:"error_description"`
}

type Result struct {
	// The number of ordersn to get airway bills in this call.
	TotalCount int `json:"total_count"`
	// This Object contains the airway bill to each order.
	AirwayBills []AirwayBill `json:"airway_bills"`
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
}

type BatchResult struct {
	// The number of orderSN to get airway bills in this call.
	TotalCount int `json:"total_count"`
	// The list contains urls of retrieving airway bill in PDF format. Each url contains the airway bills which is in the same logistics channel.
	AirwayBills []string `json:"airway_bills"`
}

type GetAirwayBillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result Result `json:"result"`
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult BatchResult `json:"batch_result"`
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

type Address struct {
	// Recipient's name for the address.
	Name string `json:"name"`
	// Recipient's phone number input when order was placed.
	Phone string `json:"phone"`
	// The town of the recipient's address. Whether there is a town will depend on the region and/or country.
	Town string `json:"town"`
	// The district of the recipient's address. Whether there is a town will depend on the region and/or country.
	District string `json:"district"`
	// The city of the recipient's address. Whether there is a town will depend on the region and/or country.
	City string `json:"city"`
	// The state/province of the recipient's address. Whether there is a town will depend on the region and/or country.
	State string `json:"state"`
	// The two-digit code representing the country of the Recipient.
	Country string `json:"country"`
	// Recipient's postal code.
	Zipcode string `json:"zipcode"`
	// The full address of the recipient, including country, state, even street, and etc.
	FullAddress string `json:"full_address"`
}

type Logistic struct {
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string `json:"shipping_carrier"`
	// The identity of logistic channel.
	LogisticID int `json:"logistic_id"`
	// Only work for cross-border order. This code is required at some sorting hub. Please ensure the service_code is INCLUDED on your shipping label, otherwise the parcel cannot be processed by the warehouse. If you didn't retrieve service_code after you first called this API, please try few more times within 30 minutes.
	ServiceCode string `json:"service_code"`
	// Only work for cross-border order.The name of the carrier ships cross countries.
	FirstMileName string `json:"first_mile_name"`
	// Only work for cross-border order.The name of the carrier delivers the parcels in local country.
	LastMileName string `json:"last_mile_name"`
	// Only work for cross-border order.This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	GoodsToDeclare bool `json:"goods_to_declare"`
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string `json:"tracking_no"`
	//
	Zone string `json:"zone"`
	// Only work for cross-border order. The string use for waybill printing. The format is "S - country_code and lane_number". For example, S-TH01, S-TH02
	LaneCode string `json:"lane_code"`
	// Only work for cross-border order in some special shop. The address info of the warehouse.
	WarehouseAddress string `json:"warehouse_address"`
	// Only work for cross-border order in some special shop. The ID of the warehouse.
	WarehouseID int `json:"warehouse_id"`
	// This object contains detailed breakdown for the recipient address.
	RecipientAddress Address `json:"recipient_address"`
	// This value indicates whether the order was a COD (cash on delivery) order.
	COD bool `json:"cod"`
}

type GetOrderLogisticsResponse struct {
	// The logistics of order.
	Logistics []Logistic `json:"logistics"`
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

type Info struct {
	// The time when logistics info has been updated.
	CTime int `json:"c_time"`
	// The order logistics tracking info.
	Description string `json:"description"`
	// The 3PL logistics status for the order. Applicable values: See Data Definition - TrackingLogisticsStatus.
	Status string `json:"status"`
}

type GetLogisticsMessageResponse struct {
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNumber
	//
	TrackingInfo []Info `json:"tracking_info"`
	// The Shopee logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.
	LogisticStatus
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
	// The order serial numbers.
	OrderSN string `json:"order_sn"`
	// Shopee's unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
}

type List struct {
	// The order serial numbers. Make sure the order has trackingNo generated before calling this API.
	OrderSN string `json:"order_sn"`
	// The unique identifier for a fulfillment order.
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
	OrdersList []List `json:"orders_list"`
	// Option to get batch airway bills in single file. Default value is false.
	IsBatch bool `json:"is_batch"`
}

type Waybill struct {
	// Shopee's unique identifier for an order.
	OrderSN string `json:"order_sn"`
	// The unique identifier for a fulfillment order.
	ForderID string `json:"forder_id"`
	// The url of retrieving airway bill.
	Waybill string `json:"waybill"`
}

type Error struct {
	// The ordersn of orders which occurred error.
	OrderSN string `json:"order_sn"`
	// The forder_id of fulfillment orders which occurred error.
	ForderID string `json:"forder_id"`
	//
	ErrorCode
	// The detail information of this error.
	ErrorDescription string `json:"error_description"`
}

type Result struct {
	// The number of ordersn to get airway bills in this call.
	TotalCount int `json:"total_count"`
	// This Object contains the airway bill to each order.
	Waybills []Waybill `json:"waybills"`
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
}

type BatchResult struct {
	// The number of orderSN to get airway bills in this call.
	TotalCount int `json:"total_count"`
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []Error `json:"errors"`
	// The url of retrieving airway bill.
	Waybills []string `json:"waybills"`
}

type GetForderWaybillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result Result `json:"result"`
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult BatchResult `json:"batch_result"`
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
}
