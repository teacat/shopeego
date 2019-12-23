package shopeego

type GetTopPicksListRequest struct {
	//
	RequestBase
}

type GetTopPicksListResponse struct {
	// Collection list
	Collections []GetTopPicksListResponseCollection `json:"collections"`
}

type AddTopPicksRequest struct {
	// Collection name. 1 to 24 characters.
	Name string `json:"name"`
	// the list of item id. Allow 4 to 8 items in one collection.
	ItemIDs []int `json:"item_i_ds"`
	// True or False
	IsActivated bool `json:"is_activated"`
	//
	RequestBase
}

type AddTopPicksResponse struct {
	// Collection id
	TopPicksID int `json:"top_picks_id"`
	// Whether it is activated or not.
	IsActivated bool `json:"is_activated"`
	// Collection name
	Name string `json:"name"`
	// Item list of the collection
	Items []AddTopPicksResponseItem `json:"items"`
}

type UpdateTopPicksRequest struct {
	// Collection id
	TopPicksID int `json:"top_picks_id"`
	// Collection name. 1 to 24 characters.
	Name string `json:"name"`
	// The list of item id. Existed item_ids will overridden by the new_item_ids.
	ItemIDs []int `json:"item_i_ds"`
	// True or False. If true, it will activate this collection and deactivate the original one.
	IsActivated bool `json:"is_activated"`
	//
	RequestBase
}

type UpdateTopPicksResponse struct {
	// Collection id
	TopPicksID int `json:"top_picks_id"`
	// Whether it is activated or not.
	IsActivated bool `json:"is_activated"`
	// Collection name
	Name string `json:"name"`
	// Item list of the collection
	Items []UpdateTopPicksResponseItem `json:"items"`
}

type DeleteTopPicksRequest struct {
	// Collection id. Cannot delete an activated collection.
	TopPicksID int `json:"top_picks_id"`
	//
	RequestBase
}

type DeleteTopPicksResponse struct {
	// Collection id
	TopPicksID int `json:"top_picks_id"`
}
