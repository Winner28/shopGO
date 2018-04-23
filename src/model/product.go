package model

type Product struct {
	ID          int
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Category    string
}
