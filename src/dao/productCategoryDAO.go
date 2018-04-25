package dao

import (
	"connection"
	"log"
	"model"
)

type ProductCategoryDAOImpl struct{}

type ProductCategoryDAO interface {
	GetProductCategory(productID int) (model.ProductCategory, error)
	GetProductsByCategoryID(ID int) ([]model.Product, error)
	DeleteProductCategory(productID int) error
	CreateProductCategory(category model.ProductCategory) error
}

func GetProductCategoryDAO() ProductCategoryDAO {
	return &ProductCategoryDAOImpl{}
}

func (dao *ProductCategoryDAOImpl) CreateProductCategory(category model.ProductCategory) error {
	if err := connection.GetConnection().DB.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (dao *ProductCategoryDAOImpl) GetProductCategory(productID int) (model.ProductCategory, error) {
	var category model.ProductCategory
	if err := connection.GetConnection().DB.Table("product_categories").
		Select("product_categories.id, product_categories.product_id, product_categories.category_id").
		Where("product_categories.product_id = ?", productID).Scan(&category).Error; err != nil {
		return emptyProductCategory(), err
	}
	return category, nil
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

func (dao *ProductCategoryDAOImpl) DeleteProductCategory(productID int) error {
	if err := connection.GetConnection().DB.Where("product_id = ?", productID).Delete(model.ProductCategory{}).Error; err != nil {
		return err
	}
	return nil

}

func emptyProductCategory() model.ProductCategory {
	return model.ProductCategory{}
}
