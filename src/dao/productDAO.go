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

func (dao *ProductDAO) Create(product model.Product, category model.Category) (model.Product, error) {
	pr := getProduct(product)
	if err := connection.GetConnection().DB.Create(&pr).Error; err != nil {
		return emptyProduct(), err
	}
	if err := dao.categoryDAO.CreateProductCategory(getProductCategory(pr.ID, category)); err != nil {
		return emptyProduct(), err
	}
	return product, nil
}

func (dao *ProductDAO) Update(product model.Product, category model.Category) (model.Product, error) {
	pr := getProduct(product)
	if err := connection.GetConnection().DB.Save(&pr).Error; err != nil {
		return emptyProduct(), err
	}
	if err := dao.categoryDAO.UpdateProductCategory(getProductCategory(pr.ID, category)); err != nil {
		return emptyProduct(), err
	}
	return product, nil
}

func getProductCategory(ID int, cat model.Category) model.ProductCategory {
	var catID int
	switch cat.Name {
	case "Clothes":
		catID = 1
		break
	case "Technologies":
		catID = 2
		break
	}

	productCat := model.ProductCategory{
		ProductID:  ID,
		CategoryID: catID,
	}
	return productCat
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

type Product struct {
	ID          int
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func getProduct(p model.Product) Product {
	return Product{
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
	}
}
