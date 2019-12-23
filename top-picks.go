package shopeego

type GetTopPicksListRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Item struct {
	// Item ID
	ItemID int
	// Item name
	ItemName string
	// Item discounted price(original price if no discount). Item level price will return if it has variation.
	ItemPrice float64
	// The sales of the item
	Sales int
}

type Collection struct {
	// Collection name
	Name string
	// Collection ID
	TopPicksID int
	// True or False
	IsActivated bool
	// Item list of the collection
	Items []Item
}

type GetTopPicksListResponse struct {
	// Collection list
	Collections []Collection
}

type AddTopPicksRequest struct {
	// Collection name. 1 to 24 characters.
	Name string
	// the list of item id. Allow 4 to 8 items in one collection.
	ItemIDs []int
	// True or False
	IsActivated bool
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Item struct {
	// Item ID
	ItemID int
	// Item name
	ItemName string
	// Item discounted price(original price if no discount). Item level price will return if it has variation.
	ItemPrice float64
	// The sales of the item
	Sales int
}

type AddTopPicksResponse struct {
	// Collection id
	TopPicksID int
	// Whether it is activated or not.
	IsActivated bool
	// Collection name
	Name string
	// Item list of the collection
	Items []Item
}

type UpdateTopPicksRequest struct {
	// Collection id
	TopPicksID int
	// Collection name. 1 to 24 characters.
	Name string
	// The list of item id. Existed item_ids will overridden by the new_item_ids.
	ItemIDs []int
	// True or False. If true, it will activate this collection and deactivate the original one.
	IsActivated bool
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type Item struct {
	// Item ID
	ItemID int
	// Item name
	ItemName string
	// Item discounted price(original price if no discount). Item level price will return if it has variation.
	ItemPrice float64
	// The sales of the item
	Sales int
}

type UpdateTopPicksResponse struct {
	// Collection id
	TopPicksID int
	// Whether it is activated or not.
	IsActivated bool
	// Collection name
	Name string
	// Item list of the collection
	Items []Item
}

type DeleteTopPicksRequest struct {
	// Collection id. Cannot delete an activated collection.
	TopPicksID int
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int
}

type DeleteTopPicksResponse struct {
	// Collection id
	TopPicksID int
}
