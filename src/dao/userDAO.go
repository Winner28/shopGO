package dao

import (
	"connection"
	"model"
)

type UserDAO struct{}

func GetUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) Get(ID int) (model.User, error) {
	var user model.User
	if err := connection.GetConnection().DB.First(&user, ID).Error; err != nil {
		return emptyUser(), err
	}
	return user, nil
}

func (dao *UserDAO) Create(user model.User) (model.User, error) {
	if err := connection.GetConnection().DB.Create(&user).Error; err != nil {
		return emptyUser(), err
	}
	return user, nil

}

func (dao *UserDAO) Update(ID int, user model.User) (model.User, error) {
	return model.User{}, nil
}

func (dao *UserDAO) Delete(ID int) error {
	if err := connection.GetConnection().DB.Delete(model.User{}, ID).Error; err != nil {
		return err
	}
	return nil

}

func (dao *UserDAO) FindAll() ([]model.User, error) {
	var users []model.User
	if err := connection.GetConnection().DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func emptyUser() model.User {
	return model.User{}
}
