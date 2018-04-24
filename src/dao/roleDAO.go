package dao

import (
	"connection"
	"errors"
	"model"
	"strings"
)

type RoleDAO struct{}

func GetRoleDAO() *RoleDAO {
	return &RoleDAO{}
}

func (dao *RoleDAO) Get(userID int) (model.Role, error) {
	var role model.Role
	if err := connection.GetConnection().DB.Where("user_id = ?", userID).First(&role).Error; err != nil {
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

func (dao *RoleDAO) createDefaultRole(userID int) bool {
	var role model.Role
	role.UserID = userID
	role.Name = "user"
	if err := connection.GetConnection().DB.Create(&role).Error; err != nil {
		return false
	}

	return true
}

func (dao *RoleDAO) Update(role model.Role) (model.Role, error) {
	if err := dao.compareAndReturnRoleToUpdate(&role); err != nil {
		return emptyRole(), errors.New("Error")
	}
	if err := connection.GetConnection().DB.Save(&role).Error; err != nil {
		return emptyRole(), err
	}
	return role, nil
}

func (dao *RoleDAO) compareAndReturnRoleToUpdate(role *model.Role) error {
	if !checkIfRoleExists(role.Name) {
		return errors.New("We dont have such role in system!")
	}
	oldRole, err := dao.Get(role.UserID)
	if err != nil {
		return err
	}
	if role.Name == "" {
		role.Name = oldRole.Name
	}
	return nil
}

func checkIfRoleExists(roleName string) bool {
	switch strings.ToLower(roleName) {
	case "admin":
		return true
	case "user":
		return true
	case "":
		return true
	default:
		return false
	}
}

func (dao *RoleDAO) Delete(userID int) error {
	if err := connection.GetConnection().DB.Where("user_id = ?", userID).Delete(model.Role{}).Error; err != nil {
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
