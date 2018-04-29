package dao

import (
	"connection"
	"model"
)

type OrderDAO interface {
	createOrder(order model.Order, productID int) (model.Order, error)
	getOrdersByUserID(userID int)
}

type OrderDAOIMpl struct {
	productOrderDAO ProductOrderDAO
}

func GetOrderDAO() OrderDAO {
	return &OrderDAOIMpl{}
}

func (dao *OrderDAOIMpl) createOrder(order model.Order, productID int) (model.Order, error) {
	if err := connection.GetConnection().DB.Create(&order).Error; err != nil {
		return emptyOrder(), err
	}
	if err := dao.productOrderDAO.Create(order.ID, productID); err != nil {
		return emptyOrder(), err
	}

	return order, nil
}

// getting all user orders!
func (dao *OrderDAOIMpl) getOrdersByUserID(userID int) {

}

func emptyOrder() model.Order {
	return model.Order{}
}
