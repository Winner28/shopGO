package dao

import (
	"connection"
	"log"
	"model"
)

type ProductCategoryDAOImpl struct{}

type ProductCategoryDAO interface {
	GetProductCategory() (model.ProductCategory, error)
	GetProductsByCategoryID(ID int) ([]model.Product, error)
	DeleteProductCategory(ProductID int)
}

func GetProductCategoryDAO() ProductCategoryDAO {
	return &ProductCategoryDAOImpl{}
}

func (dao *ProductCategoryDAOImpl) GetProductCategory() (model.ProductCategory, error) {
	return model.ProductCategory{}, nil
}

func (dao *ProductCategoryDAOImpl) GetProductsByCategoryID(ID int) ([]model.Product, error) {
	var products []model.Product
	if err := connection.GetConnection().DB.
		Table("products").Select("products.id, products.name, products.price,  products.description").
		Joins("inner join product_categories on products.id = product_categories.product_id").
		Where("product_categories.category_id = ?", ID).
		Scan(&products).Error; err != nil {

		return nil, err
	}

	log.Println(products)

	return products, nil
}

func (dao *ProductCategoryDAOImpl) DeleteProductCategory(ProductID int) {
	panic("Not implemented")
}
