package dao

import (
	"connection"
	"model"
)

type RoleDAO struct{}

func GetRoleDAO() *RoleDAO {
	return &RoleDAO{}
}

func (dao *RoleDAO) Get(userID int) (model.Role, error) {
	var role model.Role
	if err := connection.GetConnection().DB.First(&role, userID).Error; err != nil {
		return emptyRole(), err
	}
	return role, nil
}

func (dao *RoleDAO) Create(role model.Role) (model.Role, error) {
	if err := connection.GetConnection().DB.Create(&role).Error; err != nil {
		return emptyRole(), err
	}
	return role, nil

}

func (dao *RoleDAO) Update(ID int, user model.User) (model.User, error) {
	return model.User{}, nil
}

func (dao *RoleDAO) Delete(userID int) error {
	if err := connection.GetConnection().DB.Delete(model.User{}, userID).Error; err != nil {
		return err
	}
	return nil
}

func (dao *RoleDAO) FindAll() ([]model.Role, error) {
	var roles []model.Role
	if err := connection.GetConnection().DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func emptyRole() model.Role {
	return model.Role{}
}
