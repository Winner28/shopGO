package model

type Role struct {
	ID     int    `json:"id"`
	UserID int    `json:"userID"`
	Name   string `json:"name"`
}
