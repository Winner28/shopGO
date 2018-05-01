package dao

import (
	"connection"
	"model"
)

type OrderDAO interface {
	createOrder(order model.Order, productID int) (model.Order, error)
	getOrdersByUserID(userID int) ([]model.UserOrder, error)
}

type OrderDAOIMpl struct {
	productOrderDAO ProductOrderDAO
}

func GetOrderDAO() OrderDAO {
	return &OrderDAOIMpl{GetProductOrderDAO()}
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
func (dao *OrderDAOIMpl) getOrdersByUserID(userID int) ([]model.UserOrder, error) {
	return nil, nil
}

func emptyOrder() model.Order {
	return model.Order{}
}
