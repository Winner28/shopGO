package dao

import "model"

type AuthDAO struct {
	userDAO *UserDAO
}

func GetAuthDAO() *AuthDAO {
	return &AuthDAO{GetUserDAO()}
}

func (auth *AuthDAO) Login(email, passwordHash string) error {
	return nil
}

func (auth *AuthDAO) Register(model.User) error {
	return nil
}

func (auth *AuthDAO) checkUserCredo(email, passwordHash string) {

}
