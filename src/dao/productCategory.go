package dao

import (
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
	return nil, nil
}

func (dao *ProductCategoryDAOImpl) DeleteProductCategory(ProductID int) {

}
