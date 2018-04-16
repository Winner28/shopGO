package model

type ProductOrders struct {
	ID        int `json:"id"`
	OrderID   int `json:"orderID"`
	ProductID int `json:"productID"`
}
