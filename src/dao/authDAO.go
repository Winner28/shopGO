package dao

import "model"

type AuthDAO struct {
	userDAO *UserDAO
	roleDAO *RoleDAO
}

func GetAuthDAO() *AuthDAO {
	return &AuthDAO{GetUserDAO(), GetRoleDAO()}
}

func (auth *AuthDAO) Login(email, passwordHash string) (bool, string) {
	user, err := auth.userDAO.GetUserByEmail(email)
	if err != nil {
		return false, "User not exists"
	}
	if !auth.checkUserCredo(user, email, passwordHash) {
		return false, "Bad user email or password"
	}

	return true, "You logged in!"
}

func (auth *AuthDAO) Register(user model.User) (bool, string) {
	if _, err := auth.userDAO.Create(user, getDefaultRole()); err != nil {
		return false, err.Error()
	}
	userID := auth.getUserID(user)
	auth.roleDAO.createDefaultRole(userID)
	return true, "Successfully signed up!"
}

func (auth *AuthDAO) checkUserCredo(user model.User, email, passwordHash string) bool {
	if user.Email == email && user.PasswordHash == passwordHash {
		return true
	}

	return false
}

func (auth *AuthDAO) getUserID(user model.User) int {
	user, _ = auth.userDAO.GetUserByEmail(user.Email)
	return user.ID
}

func getDefaultRole() model.Role {
	return model.Role{
		Name: "user",
	}
}
