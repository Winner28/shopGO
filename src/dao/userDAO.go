package dao

import (
	"model"
)

type UserDAO struct{}

func GetUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) Get(ID int) (model.User, error) {
	return model.User{}, nil
}

func (dao *UserDAO) Create(user model.User) (model.User, error) {
	return model.User{}, nil
}

func (dao *UserDAO) Update(ID int, user model.User) (model.User, error) {
	return model.User{}, nil
}

func (dao *UserDAO) Delete(ID int) error {
	return nil
}

func (dao *UserDAO) FindAll() ([]model.User, error) {
	return nil, nil
}
