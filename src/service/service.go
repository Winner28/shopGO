package service

import (
	"model"
)

type CRUDRepository interface {
	Get(ID string, value interface{}) (interface{}, error)
	Create(value interface{}) (interface{}, error)
	Update(ID string, value interface{}) (interface{}, error)
	Delete(ID string, value interface{}) error
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

	return nil, nil
}

func (oper CRUDOperations) Create(value interface{}) (interface{}, error) {
	if user, ok := value.(model.User); ok {
		if err := getConnection().DB.Create(&user).Error; err != nil {
			return nil, err
		}
		return user, nil
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
	return nil
}
