package models

type IncomeProduct struct {
	ID        string `json:"id"`
	IncomeID  string `json:"income_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
}

type CreateIncomeProduct struct {
	IncomeID  string `json:"income_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Price     int    `json:"price"`
}

type CreateIncomeProducts struct {
	IncomeProducts []CreateIncomeProduct `json:"income_products"`
}

type IncomeProductsResponse struct {
	IncomeProducts []IncomeProduct `json:"income_products"`
	Count          int             `json:"count"`
}

type IncomeProducts struct {
	IncomeProducts []IncomeProduct `json:"income_products"`
}

type DeleteIncomeProducts struct {
	IDs []PrimaryKey `json:"ids"`
}
