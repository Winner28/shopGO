package model

import "time"

type UserOrder struct {
	ProductName        string
	ProductDescription string
	Ammount            int
	Date               time.Time
	Price              float64
}
