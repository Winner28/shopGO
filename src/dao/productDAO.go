package dao

import (
	"connection"
	"log"
	"model"
)

type ProductDAO struct {
	categoryDAO ProductCategoryDAO
	orderDAO    OrderDAO
}

func GetProductDAO() *ProductDAO {
	return &ProductDAO{GetProductCategoryDAO(), GetOrderDAO()}
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
	pr := dao.compareAndReturnProductToUpdate(&product)
	log.Println("Updating product:", pr)

	if err := connection.GetConnection().DB.Save(&pr).Error; err != nil {
		return emptyProduct(), err
	}
	if !emptyCategory(category.Name) {
		if err := dao.categoryDAO.UpdateProductCategory(getProductCategory(pr.ID, category)); err != nil {
			return emptyProduct(), err
		}
	}
	return product, nil
}

func emptyCategory(category string) bool {
	return category == ""
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

func (dao *ProductDAO) compareAndReturnProductToUpdate(product *model.Product) Product {
	oldProduct, _ := dao.Get(product.ID)
	if product.Description == "" {
		product.Description = oldProduct.Description
	}
	if product.Name == "" {
		product.Name = oldProduct.Name
	}
	if product.Price == 0 {
		product.Price = oldProduct.Price
	}

	return getProduct(*product)
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

func (dao *ProductDAO) GetUserOrders(userID int) ([]model.UserOrder, error) {
	orders, err := dao.orderDAO.getOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}
	for _, value := range orders {
		if product, err := dao.Get(value.ProductID); err == nil {
			value.ProductName = product.Name
			value.ProductDescription = product.Description
			value.Price = product.Price * float64(value.Ammount)
		} else {
			log.Println("Problem occured")
			return nil, err
		}
	}
	return orders, err
}

func (dao *ProductDAO) CreateOrder(order model.Order, productID int) (model.Order, error) {
	order, err := dao.orderDAO.createOrder(order, productID)
	if err != nil {
		return emptyOrder(), err
	}
	return order, nil
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
		ID:          p.ID,
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
	}
}
