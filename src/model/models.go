package model

import "time"

type User struct {
	ID           int
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

type Product struct {
	ID          int
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductCategory struct {
	ID         int `json:"id"`
	ProductID  int `json:"productID"`
	CategoryID int `json:"categoryID"`
}
type Order struct {
	ID       int `json:"id"`
	UserID   int `json:"userID"`
	AmountID int `json:"amountID"`
	Date     time.Time
}

type ProductOrders struct {
	ID        int `json:"id"`
	OrderID   int `json:"orderID"`
	ProductID int `json:"productID"`
}

type Role struct {
	ID     int    `json:"id"`
	UserID int    `json:"userID"`
	Name   string `json:"name"`
}
