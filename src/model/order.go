package model

import "time"

type Order struct {
	ID       int `json:"id"`
	UserID   int `json:"userID"`
	AmountID int `json:"amountID"`
	Date     time.Time
}
