package models

type PrimaryKey struct {
	ID string `json:"id"`
}

type GetListRequest struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Search   string `json:"search"`
	BasketID string `json:"basket_id"`
}
