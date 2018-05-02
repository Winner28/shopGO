package model

import "time"

type UserOrder struct {
	ProductID          int
	ProductName        string
	ProductDescription string
	Ammount            int
	Date               time.Time
	Price              float64
	TotalPrice         float64
}
