package dao

import (
	"connection"
	"model"
)

type (
	ProductOrderDAO interface {
		Create(orderID, productID int) error
	}

	ProductOrderDAOImpl struct{}
)

func GetProductOrderDAO() ProductOrderDAO {
	return &ProductOrderDAOImpl{}
}

func (productOder *ProductOrderDAOImpl) Create(orderID, productID int) error {
	if err := connection.GetConnection().DB.Create(getProductOrder(orderID, productID)).Error; err != nil {
		return err
	}
	return nil
}

func getProductOrder(orderID, productID int) model.ProductOrder {
	return model.ProductOrder{
		OrderID:   orderID,
		ProductID: productID,
	}
}
