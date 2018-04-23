package dao

import (
	"connection"
	"model"
)

type ProductDAO struct {
	categoryDAO ProductCategoryDAO
}

func GetProductDAO() *ProductDAO {
	return &ProductDAO{GetProductCategoryDAO()}
}

func (dao *ProductDAO) Get(ID int) (model.Product, error) {
	var product model.Product
	if err := connection.GetConnection().DB.First(&product, ID).Error; err != nil {
		return emptyProduct(), err
	}
	return product, nil
}

func (dao *ProductDAO) Create(product model.Product) (model.Product, error) {
	if err := connection.GetConnection().DB.Create(&product).Error; err != nil {
		return emptyProduct(), err
	}
	return product, nil

}

func (dao *ProductDAO) Update(ID int, product model.Product) (model.Product, error) {
	return emptyProduct(), nil
}

func (dao *ProductDAO) Delete(ID int) error {
	if err := connection.GetConnection().DB.Delete(model.Product{}, ID).Error; err != nil {
		return err
	}
	if err := GetProductCategoryDAO().DeleteProductCategory(ID); err != nil {
		return err
	}
	return nil

}

func (dao *ProductDAO) FindAll() ([]model.Product, error) {
	var products []model.Product
	if err := connection.GetConnection().DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (dao *ProductDAO) GetClothesProducts() ([]model.Product, error) {
	if products, err := dao.categoryDAO.GetProductsByCategoryID(1); err == nil {
		return products, nil
	}

	return nil, nil
}

func (dao *ProductDAO) GetTechsProducts() ([]model.Product, error) {
	if products, err := dao.categoryDAO.GetProductsByCategoryID(2); err == nil {
		return products, nil
	}
	return nil, nil
}

func emptyProduct() model.Product {
	return model.Product{}
}
