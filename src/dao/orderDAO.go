package dao

import (
	"model"
	"time"
)

type OrderDAO interface {
	createOrder(userID, productID, amount int, date time.Time) (model.Order, error)
	getOrdersByUserID(userID int)
}

type OrderDAOIMpl struct {
}

func GetOrderDAO() OrderDAO {
	return &OrderDAOIMpl{}
}

func (dao *OrderDAOIMpl) createOrder(userID, productID, amount int, date time.Time) (model.Order, error) {
	return model.Order{}, nil
}

// getting all user orders!
func (dao *OrderDAOIMpl) getOrdersByUserID(userID int) {

}
