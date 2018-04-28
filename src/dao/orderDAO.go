package dao

import (
	"connection"
	"model"
)

type OrderDAO interface {
	createOrder(order model.Order) (model.Order, error)
	getOrdersByUserID(userID int)
}

type OrderDAOIMpl struct {
}

func GetOrderDAO() OrderDAO {
	return &OrderDAOIMpl{}
}

func (dao *OrderDAOIMpl) createOrder(order model.Order) (model.Order, error) {
	if err := connection.GetConnection().DB.Create(&order).Error; err != nil {
		return emptyOrder(), err
	}
	return model.Order{}, nil
}

// getting all user orders!
func (dao *OrderDAOIMpl) getOrdersByUserID(userID int) {

}

func emptyOrder() model.Order {
	return model.Order{}
}
