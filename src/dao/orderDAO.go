package dao

import (
	"connection"
	"log"
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
	var orders []model.Order
	var userOrders []model.UserOrder
	if err := connection.GetConnection().DB.Find(&orders).Where("user_id=?", userID).Error; err != nil {
		return nil, err
	}

	for _, value := range orders {
		log.Println(value)
		if productID, err := dao.productOrderDAO.GetProductIDByOrderID(value.ID); err == nil {
			userOrders = append(userOrders, model.UserOrder{
				ProductID: productID,
				Ammount:   value.Amount,
				Date:      value.Date,
			})
		} else {
			log.Println("Error when we try to get productID by order ID!")
			return nil, err
		}
	}

	return userOrders, nil
}

func emptyOrder() model.Order {
	return model.Order{}
}
