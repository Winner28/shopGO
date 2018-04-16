package service

import (
	"model"
)

type CRUDRepository interface {
	Get(ID string, value interface{}) (interface{}, error)
	Create(ID int, value interface{}) (interface{}, error)
	Update(ID string, value interface{}) (interface{}, error)
	Delete(ID string, value interface{}) error
	FindAll(values interface{}) ([]interface{}, error)
}

type CRUDOperations struct{}

var repository CRUDRepository

func init() {
	repository = CRUDOperations{}
}

func (oper CRUDOperations) Get(ID string, value interface{}) (interface{}, error) {
	if user, ok := value.(model.User); ok {
		if err := getConnection().DB.First(&user, ID).Error; err != nil {
			return nil, err
		}
		return user, nil
	}
	if product, ok := value.(model.Product); ok {
		if err := getConnection().DB.First(&product, ID).Error; err != nil {
			return nil, err
		}
		return product, nil
	}

	return nil, nil
}

func (oper CRUDOperations) Create(ID int, value interface{}) (interface{}, error) {
	if user, ok := value.(model.User); ok {
		user.ID = ID
		if err := getConnection().DB.Create(&user).Error; err != nil {
			return nil, err
		}
		return user, nil
	}
	if product, ok := value.(model.Product); ok {
		product.ID = ID
		if err := getConnection().DB.Create(&product).Error; err != nil {
			return nil, err
		}
		return product, nil
	}

	return nil, nil
}

func (oper CRUDOperations) Update(ID string, value interface{}) (interface{}, error) {
	return nil, nil

}

func (oper CRUDOperations) Delete(ID string, value interface{}) error {
	if _, ok := value.(model.User); ok {
		if err := getConnection().DB.Delete(model.User{}, ID).Error; err != nil {
			return err
		}
		return nil
	}
	if _, ok := value.(model.Product); ok {
		if err := getConnection().DB.Delete(model.Product{}, ID).Error; err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (oper CRUDOperations) FindAll(value interface{}) ([]interface{}, error) {
	if _, ok := value.(model.User); ok {
		var users []model.User
		if err := getConnection().DB.Find(&users).Error; err != nil {
			return nil, err
		}

		result := make([]interface{}, len(users))
		for i, value := range users {
			result[i] = value
		}

		return result, nil
	}

	if _, ok := value.(model.Product); ok {
		var products []model.Product
		if err := getConnection().DB.Find(&products).Error; err != nil {
			return nil, err
		}

		result := make([]interface{}, len(products))
		for i, value := range products {
			result[i] = value
		}

		return result, nil
	}

	return nil, nil
}
