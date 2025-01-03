package models

type Furniture struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	ImageURL    string   `json:"image_url"`
	Category    string   `json:"category"`
	Colors      []string `json:"colors"`
	Sizes       []string `json:"sizes"`
	Images      []string `json:"images"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Cart struct {
	ID         string     `json:"id"`
	UserID     string     `json:"user_id"`
	Items      []CartItem `json:"items"`
	TotalPrice float64    `json:"total_price"`
}

type CartItem struct {
	FurnitureID string `json:"furniture_id"`
	Quantity    int    `json:"quantity"`
}
