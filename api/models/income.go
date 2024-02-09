package models

type Income struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	TotalSum   int    `json:"total_sum"`
}

type IncomesResponse struct {
	Incomes []Income `json:"incomes"`
	Count   int      `json:"count"`
}
