package models

type Product struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
	BranchID      string `json:"branch_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type CreateProduct struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
	BranchID      string `json:"branch_id"`
}

type UpdateProduct struct {
	ID            string `json:"-"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
}

type ProductResponse struct {
	Product []Product
	Count   int
}

type ProductSell struct {
	SelectedProducts       SellRequest    `json:"selected_products"`
	ProductPrices          map[string]int `json:"product_prices"`
	NotEnoughProducts      map[string]int `json:"not_enough_products"`
	NotEnoughProductPrices map[string]int `json:"prices"`
	ProductsBranchID       string         `json:"products_branch_id"`
}

type SellRequest struct {
	Products map[string]int `json:"products"`
	BasketID string         `json:"basket_id"`
	BranchID string         `json:"branch_id"`
}

type DeliverProducts struct {
	NotEnoughProducts map[string]int `json:"not_enough_products"`
	NewProducts       map[string]int `json:"new_products"`
	NewProductPrices  map[string]int `json:"new_product_prices"`
}
