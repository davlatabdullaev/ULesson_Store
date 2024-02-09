package models

type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateCategory struct {
	Name string `json:"name"`
}

type UpdateCategory struct {
	ID   string `json:"-"`
	Name string `json:"name"`
}

type CategoryResponse struct {
	Category []Category `json:"category"`
	Count    int        `json:"count"`
}
