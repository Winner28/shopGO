package model

import "time"

type Order struct {
	ID     int `json:"id"`
	UserID int `json:"userID"`
	Amount int `json:"amount"`
	Date   time.Time
}
