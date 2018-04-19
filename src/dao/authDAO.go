package dao

import "model"

type AuthDAO struct {
	userDAO *UserDAO
}

func GetAuthDAO() *AuthDAO {
	return &AuthDAO{GetUserDAO()}
}

func (auth *AuthDAO) Login(email, passwordHash string) (bool, string) {
	user, err := auth.userDAO.Get(1)
	if err != nil {
		return false, "Got error"
	}
	if !auth.checkUserCredo(user, email, passwordHash) {
		return false, "Bad user email or password"
	}

	return true, "You logged in!"
}

func (auth *AuthDAO) Register(model.User) (bool, error) {
	return false, nil
}

func (auth *AuthDAO) checkUserCredo(user model.User, email, passwordHash string) bool {
	if user.Email == email && user.PasswordHash == passwordHash {
		return true
	}

	return false
}
