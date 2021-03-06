package dao

import (
	"connection"
	"errors"
	"model"
)

type (
	ProductOrderDAO interface {
		Create(orderID, productID int) error
		GetProductIDByOrderID(orderID int) (productID int, err error)
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

func (dao *ProductOrderDAOImpl) GetProductIDByOrderID(orderID int) (productID int, err error) {
	rows, err := connection.GetConnection().DB.Table("product_orders").Select("product_id").Where("order_id = ?", orderID).Rows()
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if !rows.Next() {
		return -1, errors.New("Have no product with such order id!")
	}
	if err = rows.Scan(&productID); err != nil {
		return -1, err
	}
	return productID, nil
}

func emptyProductOrder() model.ProductOrder {
	return model.ProductOrder{}
}

func getProductOrder(orderID, productID int) model.ProductOrder {
	return model.ProductOrder{
		OrderID:   orderID,
		ProductID: productID,
	}
}
