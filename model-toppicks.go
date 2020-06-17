package shopeego

type GetTopPicksListRequest struct {
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type GetTopPicksListResponse struct {
	// Collection list
	Collections []GetTopPicksListResponseCollection `json:"collections,omitempty"`
}

type AddTopPicksRequest struct {
	// Collection name. 1 to 24 characters.
	Name string `json:"name,omitempty"`
	// the list of item id. Allow 4 to 8 items in one collection.
	ItemIDs []int `json:"item_i_ds,omitempty"`
	// True or False
	IsActivated bool `json:"is_activated,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type AddTopPicksResponse struct {
	// Collection id
	TopPicksID int64 `json:"top_picks_id,omitempty"`
	// Whether it is activated or not.
	IsActivated bool `json:"is_activated,omitempty"`
	// Collection name
	Name string `json:"name,omitempty"`
	// Item list of the collection
	Items []AddTopPicksResponseItem `json:"items,omitempty"`
}

type UpdateTopPicksRequest struct {
	// Collection id
	TopPicksID int64 `json:"top_picks_id,omitempty"`
	// Collection name. 1 to 24 characters.
	Name string `json:"name,omitempty"`
	// The list of item id. Existed item_ids will overridden by the new_item_ids.
	ItemIDs []int `json:"item_i_ds,omitempty"`
	// True or False. If true, it will activate this collection and deactivate the original one.
	IsActivated bool `json:"is_activated,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type UpdateTopPicksResponse struct {
	// Collection id
	TopPicksID int64 `json:"top_picks_id,omitempty"`
	// Whether it is activated or not.
	IsActivated bool `json:"is_activated,omitempty"`
	// Collection name
	Name string `json:"name,omitempty"`
	// Item list of the collection
	Items []UpdateTopPicksResponseItem `json:"items,omitempty"`
}

type DeleteTopPicksRequest struct {
	// Collection id. Cannot delete an activated collection.
	TopPicksID int64 `json:"top_picks_id,omitempty"`
	// Partner ID is assigned upon registration is successful. Required for all requests.
	PartnerID int64 `json:"partner_id,omitempty"`
	// Shopee's unique identifier for a shop. Required for all requests.
	ShopID int64 `json:"shopid,omitempty"`
	// This is to indicate the timestamp of the request. Required for all requests.
	Timestamp int `json:"timestamp,omitempty"`
}

type DeleteTopPicksResponse struct {
	// Collection id
	TopPicksID int64 `json:"top_picks_id,omitempty"`
}
