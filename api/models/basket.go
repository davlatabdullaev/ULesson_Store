package models

type Basket struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	TotalSum   uint   `json:"total_sum"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateBasket struct {
	CustomerID string `json:"customer_id"`
	TotalSum   uint   `json:"total_sum"`
}

type UpdateBasket struct {
	ID         string `json:"-"`
	CustomerID string `json:"customer_id"`
	TotalSum   uint   `json:"total_sum"`
}

type BasketResponse struct {
	Baskets []Basket `json:"baskets"`
	Count   int      `json:"count"`
}
