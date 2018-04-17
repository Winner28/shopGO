package service

import "net/http"

type Auth interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type AuthService struct{}

func (auth *AuthService) Login(w http.ResponseWriter, r *http.Request) {

}

func (auth *AuthService) Logout(w http.ResponseWriter, r *http.Request) {

}

func (auth *AuthService) Register(w http.ResponseWriter, r *http.Request) {

}
