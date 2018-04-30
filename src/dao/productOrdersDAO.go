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

func (dao *ProductOrderDAOImpl) Create(orderID, productID int) error {
	productOrder := getProductOrder(orderID, productID)
	if err := connection.GetConnection().DB.Create(&productOrder).Error; err != nil {
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
