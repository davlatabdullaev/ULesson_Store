package models

type BasketProduct struct {
	ID        string `json:"id"`
	BasketID  string `json:"basket_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateBasketProduct struct {
	BasketID  string `json:"basket_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type UpdateBasketProduct struct {
	ID        string `json:"-"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type BasketProductResponse struct {
	BasketProducts []BasketProduct
	Count          int
}

type BasketProductSell struct {
	ProductName string `json:"product_name"`
	Quantity    uint   `json:"quantity"`
}
