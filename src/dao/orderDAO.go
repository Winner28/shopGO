package dao

import (
	"model"
	"time"
)

type OrderDAO interface {
	createOrder(userID, amount int, date time.Time) (model.Order, error)
	getOrdersByUserID(userID int)
}

type OrderDAOIMpl struct {
}

func GetOrderDAO() OrderDAO {
	return &OrderDAOIMpl{}
}

func (dao *OrderDAOIMpl) createOrder(userID, amount int, date time.Time) (model.Order, error) {
	return model.Order{}, nil
}

func (dao *OrderDAOIMpl) getOrdersByUserID(userID int) {

}
