package dao

import (
	"connection"
	"model"
)

type UserDAO struct {
	roleDAO *RoleDAO
}

func GetUserDAO() *UserDAO {
	return &UserDAO{GetRoleDAO()}
}

func (dao *UserDAO) Get(ID int) (model.User, error) {
	var user model.User
	if err := connection.GetConnection().DB.First(&user, ID).Error; err != nil {
		return emptyUser(), err
	}
	return user, nil
}

func (dao *UserDAO) Create(user model.User, role model.Role) (model.User, error) {
	if err := connection.GetConnection().DB.Create(&user).Error; err != nil {
		return emptyUser(), err
	}
	role.UserID = user.ID
	if _, err := dao.roleDAO.Create(role); err != nil {
		return emptyUser(), err
	}
	return user, nil

}

func (dao *UserDAO) Update(user model.User, role model.Role) (model.User, error) {
	user = dao.compareAndReturnUserToUpdate(user)
	if err := connection.GetConnection().DB.Save(&user).Error; err != nil {
		return emptyUser(), err
	}
	if _, err := dao.roleDAO.Update(role); err != nil {
		return emptyUser(), err
	}
	return user, nil
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

func (dao *UserDAO) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	if err := connection.GetConnection().DB.Where("email = ?", email).First(&user).Error; err != nil {
		return emptyUser(), err
	}
	return user, nil
}

func emptyUser() model.User {
	return model.User{}
}

func (dao *UserDAO) compareAndReturnUserToUpdate(user model.User) model.User {
	oldUser, _ := dao.Get(user.ID)
	if user.FirstName == "" {
		user.FirstName = oldUser.FirstName
	}
	if user.LastName == "" {
		user.LastName = oldUser.LastName
	}
	if user.Email == "" {
		user.Email = oldUser.Email
	}
	if user.PasswordHash == "" {
		user.PasswordHash = oldUser.PasswordHash
	}
	return user
}
