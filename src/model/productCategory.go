package model

type ProductCategory struct {
	ID         int `json:"id"`
	ProductID  int `json:"productID"`
	CategoryID int `json:"categoryID"`
}
