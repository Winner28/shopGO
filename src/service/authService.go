package service

import (
	"model"
	"net/http"
)

type authDAO interface {
	Login(email, passwordHash string) error
	Register(model.User) error
}

type AuthService struct {
	DAO authDAO
}

func GetAuthService(dao authDAO) *AuthService {
	return &AuthService{dao}
}

func (auth *AuthService) Login(w http.ResponseWriter, r *http.Request) {

	auth.DAO.Login("", "")
}

func (auth *AuthService) Logout(w http.ResponseWriter, r *http.Request) {

}

func (auth *AuthService) Register(w http.ResponseWriter, r *http.Request) {
	auth.DAO.Register(model.User{})
}
