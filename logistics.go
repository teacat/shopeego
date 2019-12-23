package shopeego

type GetLogisticsRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Size struct {
	// The identity of size.
	SizeID int
	// The name of size.
	Name int
	// The pre-defined shipping fee for the specific size.
	DefaultPrice int
}

type Limit struct {
	// The max weight for an item on this logistic channel.If the value is 0 or null, that means there is no limit.
	ItemMaxWeight
	// The min weight for an item on this logistic channel. If the value is 0 or null, that means there is no limit.
	ItemMinWeight
}

type Dimension struct {
	// The max height limit.
	Height float64
	// The max width limit.
	Width float64
	// The max length limit.
	Length float64
	// The unit for the limit.
	Unit string
}

type Logistic struct {
	// The identity of logistic channel
	LogisticID int
	// The name of logistic channel
	LogisticName string
	// This is to indicate whether this logistic channel supports COD
	HasCOD bool
	// Whether this logistic channel is enabled on shop level.
	Enabled bool
	// See Define FeeType, related to FeeType Value
	FeeType string
	// Only for fee_type is SIZE_SELECTION
	Sizes Size
	// The weight limit for this logistic channel.
	WeightLimits Limit
	// The dimension limit for this logistic channel.
	ItemMaxDimension Dimension
}

type GetLogisticsResponse struct {
	Logistics []Logistic
	// The identifier for an API request for error tracking
	RequestID string
}

type UpdateShopLogisticsRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// The identity of logistic channel
	LogisticID int
	// Whether to enable this logistic channel
	Enabled bool
	// Whether to make this logistic channel preferred. Indonestia logistics channel are not applicable.
	Preferred bool
	// Whether to enable COD for this logistic channel. Only COD supported channels are applicable.
	COD bool
}

type UpdateShopLogisticsResponse struct {
	// Shopee's unique identifier for a shop.
	ShopID int
	// The identity of logistic channel
	LogisticID int
	// Whether this logistic channel is enabled
	Enabled bool
	// Whether this logistic channel is preferred
	Preferred bool
	// WHether COD is enabled for this channel
	COD bool
	// The identifier for an API request for error tracking
	RequestID string
}

type GetParameterForInitRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// The serial numbers of the order needs to be initiated logistic.
	OrderSN string
}

type GetParameterForInitResponse struct {
	// Could contain 'address_id', 'pickup_time_id'. If it contains 'address_id', should call shopee.logistics.GetAddress to get address list, then choose one of the address to init logistic. If it contains pickup_time_id, should call shopee.logistics.GetTimeSlot to get timeslot list, then choose one of the time to init logistic. If it has empty value, developer should still include "pickup" field in Init API.
	Pickup []string
	// Could contain 'branch_id', 'sender_real_name' or 'tracking_no'. If it contains 'branch_id', should call shopee.logistics.GetBranch to get branch list. If it contains 'sender_real_name' or 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "dropoff" field in Init API.
	Dropoff []string
	// Could contain 'tracking_no'. If it contains 'tracking_no', should manually input these values in Init API. If it has empty value, developer should still include "non-integrated" field in Init API.
	NonIntegrated [] string
	// The identifier for an API request for error tracking
	RequestID string
}

type GetAddressRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Address struct {
	// The identity of address
	AddressID int
	// The country of specify address
	Country string
	// The state of specify address
	State string
	// The city of specify address
	City string
	// The address description of specify address
	Address string
	// The zipcode of specify address
	Zipcode string
	// The district of specify address
	District string
	// The town of specify address
	Town string
}

type GetAddressResponse struct {
	// All pickup address that you can choose.
	AddressList []Address
	// The identifier for an API request for error tracking
	RequestID string
}

type GetTimeSlotRequest struct {
	// The order serial numbers.
	OrderSN string
	// The identify of address.
	AddressID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Time struct {
	// The identity of pickuptime.
	PickupTimeID string
	// The date of pickup time. In timestamp.
	Date int
	// The text description of pickup time. Only applicable for certain channels.
	TimeText string
}

type GetTimeSlotResponse struct {
	//
	PickupTime []Time
	// The identifier for an API request for error tracking
	RequestID string
}

type GetBranchRequest struct {
	// The order serial numbers.
	OrderSN string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Branch struct {
	// The identity of branch.
	BranchID int
	// The country of specify branch.
	Country string
	// The state of specify branch.
	State string
	// The city of specify branch.
	City string
	// The address description of specify branch.
	Address string
	// The zipcode of specify branch.
	Zipcode string
	// The district of specify branch.
	District string
	// The town of specify branch.
	Town string
}

type GetBranchResponse struct {
	// All dropoff branches you can choose.
	Branch []Branch
	// The identifier for an API request for error tracking
	RequestID string
}

type GetLogisticInfoRequest struct {
	// The order serial numbers.
	OrderSN string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Address struct {
	// The identity of address.
	AddressID int
	// The country of specify branch.
	Country string
	// The state of specify branch.
	State string
	// The city of specify branch.
	City string
	// The address description of specify branch.
	Address string
	// The zipcode of specify branch.
	Zipcode string
	// The district of specify branch.
	District string
	// The town of specify branch.
	Town string
	// List of pickup_time information corresponding to the address_id.
	TimeSlotList []string
	// The identity of pickuptime.
	PickupTimeID string
	// The date of pickup time. In timestamp.
	Date int
	// The text description of pickup time. Only applicable for certain channels.
	TimeText string
}

type Branch struct {
	// The identity of branch.
	BranchID int
	// The country of specify branch.
	Country string
	// The state of specify branch.
	State string
	// The city of specify branch.
	City string
	// The address description of specify branch.
	Address string
	// The zipcode of specify branch.
	Zipcode string
	// The district of specify branch.
	District string
	// The town of specify branch.
	Town string
}

type Pickup struct {
	// List of available pickup address info.
	AddressList []Address
}

type Dropoff struct {
	// List of available dropoff branches info.
	BranchList []Branch
}

type Info structt {
	// Logistics information for pickup mode order.
	Pickup []string
	// Logistics information for dropoff mode order.
	Dropoff []string
	// The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init.
	NonIntegrated []string
}

type GetLogisticInfoResponse struct {
	// Logistics information for pickup mode order.
	Pickup Pickup
	// Logistics information for dropoff mode order.
	Dropoff Dropoff
	// The parameters required based on each specific order to Init. Must use the fields included under info_needed to call Init.
	InfoNeeded Info
	// The identifier for an API request for error tracking
	RequestID string
}

type Pickup struct {
	// The identity of address. Retrieved from shopee.logistics.GetAddress.
	AddressID int
	// The pickup time id. Retrieved from shopee.logistics.GetTimeSlot.
	PickupItemID string
}

type Dropoff struct {
	// The identity of branch. Retrieved from shopee.logistics.GetBranch branch.
	BranchID int
	// The real name of sender.
	SenderRealName string
	// Need input this field when "tracking_no" is returned from "info_need". Please note that this tracking number is assigned by third-party shipping carrier for item shipment.
	TrackingNo string
}

type NonIntegrated struct {
	// Optional parameter for non-integrated channel order. The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string
}

type InitRequest struct {
	// The order serial numbers.
	OrderSN string
	// Required parameter ONLY if GetParameterForInit returns "pickup" or if GetLogisticsInfo returns "pickup" under "info_needed" for the same order. Developer should still include "pickup" field in the call even if "pickup" has empty value.
	Pickup Pickup
	// Required parameter ONLY if GetParameterForInit returns "dropoff" or if GetLogisticsInfo returns "dropoff" under "info_needed" for the same order. Developer should still include "dropoff" field in the call even if "dropoff" has empty value.
	Dropoff Dropoff
	// Optional parameter when GetParameterForInit returns "non-integrated" or GetLogisticsInfo returns "non-integrated" under "info_needed".
	NonIntegrated NonIntegrated
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Shopee's unique identifier for a fulfillment order.
	ForderID string
}

type InitResponse struct {
	// The tracking number of order
	TrackingNumber string
	// The identifier for an API request for error tracking
	RequestID string
}

type GetAirwayBillRequest struct {
	// The set of order serial numbers. Up to 50 ordersn in one call.
	OrderSNList []string
	// Option to get batch airway bills in single file. Default value is false.
	IsBatch bool
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type AirwayBill struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The url of retrieving airway bill.
	AirwayBill string
}

type Error struct {
	// The ordersn of orders which occurred error.
	OrderSN string
	//
	ErrorCode string
	// The detail information of this error.
	ErrorDescription string
}

type Result struct {
	// The number of ordersn to get airway bills in this call.
	TotalCount int
	// This Object contains the airway bill to each order.
	AirwayBills []AirwayBill
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []Error
}

type BatchResult struct {
	// The number of orderSN to get airway bills in this call.
	TotalCount int
	// The list contains urls of retrieving airway bill in PDF format. Each url contains the airway bills which is in the same logistics channel.
	AirwayBills []string
}

type GetAirwayBillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result Result
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult BatchResult
	// This list contains the ordersn of all orders that failed to retrieve airway bill in this call. AirwayBill is no longer fetchable after the order status is updated to SHIPPED.
	Errors []string
	// The identifier for an API request for error tracking
	RequestID string
}

type GetOrderLogisticsRequest struct {
	// The order serial numbers. Make sure the order has trackingNo generated before calling this API.
	OrderSN string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Shopee's unique identifier for a fulfillment order.
	ForderID string
}

type Address struct {
	// Recipient's name for the address.
	Name string
	// Recipient's phone number input when order was placed.
	Phone string
	// The town of the recipient's address. Whether there is a town will depend on the region and/or country.
	Town string
	// The district of the recipient's address. Whether there is a town will depend on the region and/or country.
	District string
	// The city of the recipient's address. Whether there is a town will depend on the region and/or country.
	City string
	// The state/province of the recipient's address. Whether there is a town will depend on the region and/or country.
	State string
	// The two-digit code representing the country of the Recipient.
	Country string
	// Recipient's postal code.
	Zipcode string
	// The full address of the recipient, including country, state, even street, and etc.
	FullAddress string
}

type Logistic struct {
	// The logistics service provider that the buyer selected for the order to deliver items.
	ShippingCarrier string
	// The identity of logistic channel.
	LogisticID int
	// Only work for cross-border order. This code is required at some sorting hub. Please ensure the service_code is INCLUDED on your shipping label, otherwise the parcel cannot be processed by the warehouse. If you didn't retrieve service_code after you first called this API, please try few more times within 30 minutes.
	ServiceCode string
	// Only work for cross-border order.The name of the carrier ships cross countries.
	FirstMileName string
	// Only work for cross-border order.The name of the carrier delivers the parcels in local country.
	LastMileName string
	// Only work for cross-border order.This value indicates whether the order contains goods that are required to declare at customs. "T" means true and it will mark as "T" on the shipping label; "F" means false and it will mark as "P" on the shipping label. This value is accurate ONLY AFTER the order trackingNo is generated, please capture this value AFTER your retrieve the trackingNo.
	GoodsToDeclare bool
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNo string
	//
	Zone string
	// Only work for cross-border order. The string use for waybill printing. The format is "S - country_code and lane_number". For example, S-TH01, S-TH02
	LaneCode string
	// Only work for cross-border order in some special shop. The address info of the warehouse.
	WarehouseAddress string
	// Only work for cross-border order in some special shop. The ID of the warehouse.
	WarehouseID int
	// This object contains detailed breakdown for the recipient address.
	RecipientAddress Address
	// This value indicates whether the order was a COD (cash on delivery) order.
	COD bool
}

type GetOrderLogisticsResponse struct {
	// The logistics of order.
	Logistics []Logistic
	// The identifier for an API request for error tracking
	RequestID string
}

type GetLogisticsMessageRequest struct {
	// The order serial numbers.
	OrderSN string
	// The tracking number assigned by the shipping carrier for item shipment. If there are more than one tracking number in one order, this param is necessary.
	TrackingNumber string
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// Shopee's unique identifier for a fulfillment order.
	ForderID string
}

type Info struct {
	// The time when logistics info has been updated.
	CTime int
	// The order logistics tracking info.
	Description string
	// The 3PL logistics status for the order. Applicable values: See Data Definition - TrackingLogisticsStatus.
	Status string
}

type GetLogisticsMessageResponse struct {
	// The tracking number assigned by the shipping carrier for item shipment.
	TrackingNumber
	//
	TrackingInfo []Info
	// The Shopee logistics status for the order. Applicable values: See Data Definition- LogisticsStatus.
	LogisticStatus
	// The identifier for an API request for error tracking
	RequestID string
	// The order serial numbers.
	OrderSN string
	// Shopee's unique identifier for a fulfillment order.
	ForderID string
}

type List struct {
	// The order serial numbers. Make sure the order has trackingNo generated before calling this API.
	OrderSN string
	// The unique identifier for a fulfillment order.
	ForderID string
}

type GetForderWaybillRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
	// The set of order serial numbers. Up to 50 ordersn in one call.
	OrdersList []List
	// Option to get batch airway bills in single file. Default value is false.
	IsBatch bool
}

type Waybill struct {
	// Shopee's unique identifier for an order.
	OrderSN string
	// The unique identifier for a fulfillment order.
	ForderID string
	// The url of retrieving airway bill.
	Waybill string
}

type Error struct {
	// The ordersn of orders which occurred error.
	OrderSN string
	// The forder_id of fulfillment orders which occurred error.
	ForderID string
	//
	ErrorCode
	// The detail information of this error.
	ErrorDescription string
}

type Result struct {
	// The number of ordersn to get airway bills in this call.
	TotalCount int
	// This Object contains the airway bill to each order.
	Waybills []Waybill
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []Error
}

type BatchResult struct {
	// The number of orderSN to get airway bills in this call.
	TotalCount int
	// This list contains the ordersn and error descriptions of all orders that failed to retrieve airway bill in this call.
	Errors []Error
	// The url of retrieving airway bill.
	Waybills []string
}

type GetForderWaybillResponse struct {
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is false.
	Result Result
	// This object contains the detailed breakdown for the result of this API call if the param is_batch is true.
	BatchResult BatchResult
	// The identifier for an API request for error tracking
	RequestID string
}
